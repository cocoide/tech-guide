package router

import (
	"github.com/cocoide/tech-guide/pkg/interface/handler"
	custom_middleware "github.com/cocoide/tech-guide/pkg/interface/middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"os"
)

func NewSchedulerRouter(e *echo.Echo, h *handler.Handler, m *custom_middleware.Middleware) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("SCHEDULER_URL")},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPut, http.MethodOptions},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))
	scheduler := e.Group("/scheduler", m.AuthorizeSchedulerAccess)
	scheduler.POST("/register", h.RegisterArticles)
	scheduler.PUT("/topic", h.UpdateArticleTopics)
	scheduler.POST("/rating", h.AssignArticleRating)

	scheduler.PUT("/activity", h.UpdateActivity)
}
