package handler

import (
	"context"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/labstack/echo"
	"net/http"
	"os"
	"strconv"
)

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
	refreshToken := c.QueryParam("refresh_token")
	if len(refreshToken) != 28 {
		return c.JSON(403, "Unauthorized")
	}
	accountID := getAccountIDInCtx(c)
	accessToken, err := h.account.RefreshToken(accountID, refreshToken)
	if err != nil {
		return c.JSON(403, "Unauthorized")
	}
	return c.JSON(200, accessToken)
}

func (h *Handler) GetAccountProfile(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	account, err := h.repo.GetAccountProfile(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	account.Password = ""
	return c.JSON(200, account)
}
