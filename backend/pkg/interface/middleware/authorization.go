package middleware

import (
	"fmt"
	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/cocoide/tech-guide/pkg/usecase/tknutils"
	"github.com/labstack/echo"
	"os"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (m *Middleware) AuthorizeAccountAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken, err := ctxutils.GetBearerToken(c)
		if err != nil {
			return c.JSON(403, fmt.Sprintf("Failed to get token from header: %v", err))
		}
		claims, err := tknutils.ParseJwt(bearerToken)
		if err != nil {
			return c.JSON(403, fmt.Sprintf("Failed to parse accessToken: %v", err))
		}
		ctxutils.SetInEchoCtx(c, key.CtxAccountID, claims.AccountID)
		return next(c)
	}
}

func (m *Middleware) AuthorizeSchedulerAccess(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		bearerToken, err := ctxutils.GetBearerToken(c)
		if err != nil {
			return c.JSON(403, fmt.Sprintf("Failed to get token from header: %v", err))
		}
		if bearerToken != os.Getenv("SCHEDULER_API_KEY") {
			return c.JSON(403, fmt.Sprintf("Unauthorized"))
		}
		return next(c)
	}
}
