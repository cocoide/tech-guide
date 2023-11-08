package handler

import (
	"context"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/cocoide/tech-guide/pkg/usecase/parser"
	"github.com/labstack/echo"
	"net/http"
	"os"
)

func (h *Handler) GetAccountSession(c echo.Context) error {
	var session dto.AccountSession
	accountId := ctxutils.GetAccountID(c)
	sessionKey := fmt.Sprintf("session.%d", accountId)
	strSession, exist, err := h.cache.Get(sessionKey)
	if err != nil {
		return c.JSON(400, err)
	}
	if !exist {
		account, err := h.repo.GetAccountProfile(accountId)
		if err != nil {
			return c.JSON(400, err)
		}
		session = dto.AccountSession{
			AccountID:   account.ID,
			DisplayName: account.DisplayName,
			AvatarURL:   account.AvatarURL,
		}
	}
	if len(strSession) > 0 {
		session, err = parser.Deserialize[dto.AccountSession](strSession)
	}
	if err != nil {
		return c.JSON(400, fmt.Sprintf("Failed to deserialized: %v", err))
	}
	return c.JSON(200, &session)
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
	authenticateRes, err := h.account.AuthenticateWithGoogle(ctx, code)
	if err != nil {
		return c.JSON(400, "Failed to authenticate")
	}
	userEmail := authenticateRes.Email
	displayName := authenticateRes.DisplayName
	account, err := h.repo.GetByEmail(userEmail)
	if err != nil {
		return c.JSON(400, "Failed to authenticate")
	}
	if account == nil {
		newAccount := &model.Account{
			Email:       userEmail,
			DisplayName: displayName,
		}
		tokens, err := h.account.TemporarySignUp(newAccount)
		if err != nil {
			return c.JSON(400, err)
		}
		redirectURL := fmt.Sprintf("%s/api/oauth?access=%s&refresh=%s", os.Getenv("FRONTEND_URL"), tokens.AccessToken, tokens.RefreshToken)
		return c.Redirect(http.StatusFound, redirectURL)
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
	accountId := ctxutils.GetIntFromPath(c)
	account, err := h.repo.GetAccountProfile(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	account.Password = ""
	return c.JSON(200, account)
}
