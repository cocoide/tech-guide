package handler

import (
	"context"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/interface/handler/pathutils"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"net/http"
	"os"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/labstack/echo"
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
	return c.Redirect(http.StatusFound, pathutils.CompleteSignupSessionURL(tokens))
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
	redirectURL, err := h.account.GenerateOAuthRedirectURL(model.Google)
	if err != nil {
		return c.JSON(500, err)
	}
	return c.Redirect(http.StatusFound, redirectURL)
}

func (h *Handler) HandleOAuthCallback(c echo.Context) error {
	ctx := context.Background()
	code := c.QueryParam("code")
	authenticateResp, err := h.account.AuthenticateWithGoogle(ctx, code)
	if err != nil {
		return c.JSON(400, "Failed to authenticate")
	}
	userEmail := authenticateResp.Email
	account, err := h.repo.GetByEmail(userEmail)
	if err != nil {
		return c.JSON(403, "Failed to authenticate")
	}
	if account == nil {
		temporarySignupReq := &usecase.CacheTemporarySignUpRequest{
			DisplayName: authenticateResp.DisplayName,
			AvatarURL:   authenticateResp.AvatarURL,
			Email:       authenticateResp.Email,
		}
		sessionID, err := h.account.CacheTemporarySignUp(temporarySignupReq)
		if err != nil {
			return c.JSON(400, err)
		}
		return c.Redirect(http.StatusFound, pathutils.StartSignupSessionURL(sessionID))
	}
	loginRes, err := h.account.Login(userEmail)
	tokens := loginRes.Tokens
	if err != nil {
		return c.JSON(400, err)
	}
	redirectURL := fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
	return c.Redirect(http.StatusFound, redirectURL)
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
