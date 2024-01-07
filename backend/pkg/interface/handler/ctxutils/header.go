package ctxutils

import (
	"fmt"
	"github.com/labstack/echo"
	"strings"
)

func GetBearerToken(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	headers := strings.Split(authHeader, " ")
	if len(headers) != 2 {
		return "", fmt.Errorf("invalid token format")
	}
	return headers[1], nil
}
