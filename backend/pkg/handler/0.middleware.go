package handler

import (
	"strings"
	"time"

	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo"
)

func (h *Handler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return c.JSON(403, "invalid token format")
		}
		token, err := util.ParseToken(bearerToken[1])
		if err != nil {
			return c.JSON(403, "failed to parse token")
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			c.Set("account_id", claims["account_id"].(float64))
		} else {
			return c.JSON(403, "failed to get user data form claims")
		}
		// tokenのexpire timeを更新
		expireTime := time.Now().Add(60 * time.Minute)
		claims["exp"] = expireTime.Unix()
		return next(c)
	}
}
