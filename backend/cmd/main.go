package main

import (
	"context"
	"os"
	"time"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/database"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/handler"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/scheduler"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	conf.NewEnv()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{os.Getenv("BACKEDN_URL"), os.Getenv("FRONTNED_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowCredentials: true,
	}))

	ctx := context.Background()
	db := database.NewDatabase()
	rd := database.NewRedisCilent(ctx)
	cr := repo.NewCacheRepo(rd)

	or := repo.NewCollectionRepo(db)
	ar := repo.NewArticleRepo(db)
	ur := repo.NewAccountRepo(db)
	tr := repo.NewTopicRepo(db)

	og := gateway.NewOGPGateway()
	ag := gateway.NewOpenAIGateway(ctx)
	uu := usecase.NewAccountUseCase(ur)
	ts := service.NewTopicAnalysisService(ag, tr, ar)
	ps := service.NewPersonalizeService(tr, cr)
	h := handler.NewHandler(ur, ar, or, cr, og, uu, ts, ps, tr)

	private := e.Group("/account", h.AuthMiddleware)
	private.GET("/private/profile/:id", h.GetAccountProfile)
	e.GET("account/profile/:id", h.GetAccountProfile)
	e.GET("account/collection/:id", h.GetCollections)

	private.POST("/bookmark", h.DoBookmark)
	private.POST("/collection", h.CreateCollection)
	private.GET("/collection", h.GetCollectionForBookmark)
	e.POST("/login", h.Login)
	e.POST("/signup", h.SignUp)
	e.POST("/refresh", h.RefreshToken)
	e.GET("/email", h.CheckEmail)

	e.GET("/rss", h.GetRSS)
	e.GET("/ogp", h.GetOGP)
	e.GET("/article", h.GetArticles)
	e.GET("/article/:id", h.GetArticleDetail)
	e.GET("/article/related/:id", h.GetRelatedArticles)
	e.GET("/token", h.GenerateToken)
	e.POST("/article", h.CreateArticle)
	e.POST("/topic", h.CreateTopics)
	e.GET("/topic", h.GetTopics)

	private.GET("/topic/follow", h.GetFollowingTopics)
	private.PUT("/topic/follow/:id", h.DoFollowTopic)
	private.DELETE("/topic/follow/:id", h.UnFollowTopic)

	private.GET("/article/recommend", h.GetRecommendArticles)

	sp := scheduler.NewSchedulerPool()
	tw := scheduler.NewTimelineWorker(ur, cr, tr, ps)
	go func() {
		sp.AddScheduler(key.TrendingArticlesWorker, 1*time.Hour, tw.CacheTredingArticlesWorker)
		sp.AddScheduler(key.PersonalizedArticlesWorker, 1*time.Hour, tw.CachePersonalizedArticlesWorker)
		sp.StartScheduler(key.TrendingArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.PersonalizedArticlesWorker)
	}()
	e.Logger.Fatal(e.Start(":8080"))
}
