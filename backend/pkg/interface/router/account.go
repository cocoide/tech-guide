package router

import (
	"github.com/cocoide/tech-guide/pkg/interface/handler"
	"github.com/labstack/echo"
)

func newAccountRouterGroup(account *echo.Group, h *handler.Handler) {
	account.POST("/article/favorite/:id", h.DoFavoriteArticle)
	account.DELETE("/article/favorite/:id", h.UnFavoriteArticle)

	account.GET("/private/profile/:id", h.GetAccountProfile)

	account.POST("/comment", h.CreateComment)
	account.POST("/comment/:articleId", h.AddComment)
	account.POST("/bookmark", h.DoBookmark)

	account.POST("/collection", h.CreateCollection)
	account.GET("/collection", h.GetCollectionForBookmark)

	account.PUT("/article/read", h.SetReadArticle)
	account.GET("/article/read", h.GetReadArticle)

	account.GET("/topic/follow/:id", h.CheckTopicFollow)
	account.GET("/topic/follow", h.GetFollowingTopics)
	account.PUT("/topic/follow/:id", h.DoFollowTopic)
	account.DELETE("/topic/follow/:id", h.UnFollowTopic)

	account.GET("/source/follow/:id", h.CheckSourceFollow)
	account.GET("/source/follow", h.GetFollowingSources)
	account.PUT("/source/follow/:id", h.DoFollowSource)
	account.DELETE("/source/follow/:id", h.UnFollowSource)

	account.GET("/article/recommend", h.GetRecommendArticles)
	account.GET("/feeds", h.GetFeedsWithPaginate)
	account.GET("/session", h.GetAccountSession)
}
