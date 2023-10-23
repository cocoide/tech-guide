package router

import (
	"net/http"
	"os"

	"github.com/cocoide/tech-guide/pkg/interface/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRootRouter(e *echo.Echo, h *handler.Handler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL"), "https://accounts.google.com"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodPut, http.MethodOptions},
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
	}))
	account := e.Group("/account", h.AuthMiddleware)
	newAccountRouterGroup(account, h)

	e.GET("/oauth/login", h.HandleOAuthLogin)
	e.GET("/oauth/callback", h.HandleOAuthCallback)
	e.POST("/oauth/refresh", h.RefreshToken)

	e.GET("/account/profile/:id", h.GetAccountProfile)
	e.GET("/account/collection/:id", h.GetCollections)

	e.GET("/comment/:articleId", h.GetComments)

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
	e.GET("/github/star", h.GetGithubStar)
	e.GET("/github/popular", h.GetTopGitRepo)

	e.GET("/scraper/headers", h.GetHeaders)
	e.GET("/scraper/content", h.GetPageContent)

	e.GET("/search/title", h.FindByTitle)
	e.GET("/summary", h.Summarize)
	e.GET("/scraper/detail", h.ScrapeDetail)

	e.GET("/tweet", h.GetTweets)

	e.GET("category", h.GetCategories)
}
