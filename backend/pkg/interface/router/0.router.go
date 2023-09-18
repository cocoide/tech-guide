package router

import (
	"os"

	"github.com/cocoide/tech-guide/pkg/interface/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRootRouter(e *echo.Echo, h *handler.Handler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("BACKEDN_URL"), os.Getenv("FRONTNED_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	account := e.Group("/account", h.AuthMiddleware)
	newSecureRouterGroup(account, h)
 
	e.GET("/account/profile/:id", h.GetAccountProfile)
	e.GET("/account/collection/:id", h.GetCollections)

	e.GET("/comment/:articleId", h.GetComments)

	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.POST("/refresh", h.RefreshToken)
	e.GET("/email", h.CheckEmail)
	e.GET("/contribution/:id", h.GetContributions)

	e.GET("/rss", h.GetRSS)
	e.GET("/ogp", h.GetOGP)
	e.GET("/domain", h.GetDomainID)
	e.GET("/overview", h.GetOverview)

	e.GET("/speakerdeck", h.GetSpeakerDeckID)

	e.GET("/article", h.GetArticles)
	e.GET("/article/source/:sourceId", h.GetArticlesBySourceID)
	e.GET("/article/topic/:topicId", h.GetArticlesByTopicID)
	e.GET("/article/:id", h.GetArticleDetail)
	e.GET("/article/related/:id", h.GetRelatedArticles)
	e.GET("/token", h.GenerateToken)
	e.POST("/article", h.CreateArticle)
	e.POST("/topic", h.CreateTopics)
	e.GET("/topic", h.GetTopics)
	e.GET("/topic/popular", h.GetPopularTopics)

	e.GET("/source", h.GetAllSources)
	e.GET("/source/popular", h.GetPopularSources)
	e.GET("/collection/:id", h.GetCollectionData)
	e.GET("/collection/:id/topics", h.GetTopicsForCollection)

	e.GET("/topic/:id", h.GetTopicData)

	e.GET("/source/:id", h.GetSourceData)

	e.GET("/trend/qiita", h.GetQiitaTrend)
	e.GET("/trend/zenn", h.GetZennTrend)
}
