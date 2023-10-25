package handler

import (
	"github.com/cocoide/tech-guide/pkg/usecase/tknutils"
	"github.com/labstack/echo"
	"strings"
)

func (h *Handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.JSON(403, "invalid token format")
		}
		claims, err := tknutils.ParseJwt(bearerToken[1])
		if err != nil {
			return c.JSON(403, "Failed to parse accessToken")
		}
		c.Set("account_id", claims.AccountID)
		return next(c)
	}
}
