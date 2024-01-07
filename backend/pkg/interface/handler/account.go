package handler

import (
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/cocoide/tech-guide/pkg/interface/handler/pathutils"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type CompleteOnboardingRequest struct {
	FollowTopicIDs  []int `json:"follow_topic_ids"`
	FollowSourceIDs []int `json:"follow_source_ids"`
}

func (h *Handler) CompleteOnboarding(c echo.Context) error {
	sessionID, err := ctxutils.GetSessionIDFromHeader(c)
	if err != nil {
		return c.JSON(400, err)
	}
	var req CompleteOnboardingRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, err)
	}
	completeOnboardingRequest := &usecase.CompleteOnboardingRequest{
		SessionID:       sessionID,
		FollowTopicIDs:  req.FollowTopicIDs,
		FollowSourceIDs: req.FollowSourceIDs,
	}
	tokens, err := h.account.CompleteOnboarding(completeOnboardingRequest)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.Redirect(http.StatusFound, pathutils.CompleteSignupURL(tokens))
}

func (h *Handler) GetSignupSession(c echo.Context) error {
	sessionID, err := ctxutils.GetSessionIDFromHeader(c)
	if err != nil {
		return c.JSON(400, err)
	}
	session, err := h.account.GetSignupSession(sessionID)
	if session == nil {
	}
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, session)
}

func (h *Handler) GetLoginSession(c echo.Context) error {
	accountID := ctxutils.GetAccountID(c)
	session, err := h.account.GetLoginSession(accountID)
	if err != nil {
		return c.JSON(400, err)
	}
	return c.JSON(200, session)
}

type RegisterAccountRequest struct {
	AvatarURL   string `json:"avatar_url"`
	DisplayName string `json:"display_name"`
}

func (h *Handler) RegisterAccount(c echo.Context) error {
	sessionID, err := ctxutils.GetSessionIDFromHeader(c)
	if err != nil {
		return c.JSON(500, err)
	}
	var req RegisterAccountRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(500, err)
	}
	registerAccountReq := &usecase.RegisterAccountRequest{
		SessionID:   sessionID,
		AvatarURl:   req.AvatarURL,
		DisplayName: req.DisplayName,
	}
	if err := h.account.RegisterAccount(registerAccountReq); err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(200, "register account success")
}

func (h *Handler) HandleOAuthLogin(c echo.Context) error {
	oauthType, err := ctxutils.GetOAuthType(c)
	if err != nil {
		return c.JSON(400, err)
	}
	redirectURL, err := h.account.GenerateOAuthRedirectURL(oauthType)
	if err != nil {
		return c.JSON(500, err)
	}
	log.Printf("Authorize URL: %s", redirectURL)
	return c.Redirect(http.StatusFound, redirectURL)
}

func (h *Handler) HandleOAuthCallback(c echo.Context) error {
	oauthType, err := ctxutils.GetOAuthType(c)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	authCode := c.QueryParam("code")
	res, err := h.account.ProcessOAuthCallback(oauthType, authCode)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	if res.IsSignup {
		return c.Redirect(http.StatusFound, pathutils.StartSignupURL(res.SessionID))
	}
	return c.Redirect(http.StatusFound, pathutils.LoginURL(res.Tokens))
}

func (h *Handler) RefreshToken(c echo.Context) error {
	refreshToken := c.QueryParam("token")
	accessToken, err := h.account.RefreshToken(refreshToken)
	if err != nil {
		return c.JSON(403, "Unauthorized")
	}
	return c.JSON(200, accessToken)
}

func (h *Handler) GetAccountProfile(c echo.Context) error {
	accountId := ctxutils.GetIDFromPath(c)
	account, err := h.repo.GetAccountProfile(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	account.Password = ""
	return c.JSON(200, account)
}
