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
	tx := repo.NewTxRepo(db)

	or := repo.NewCollectionRepo(db)
	ar := repo.NewArticleRepo(db)
	ur := repo.NewAccountRepo(db)
	tr := repo.NewTopicRepo(db)
	mr := repo.NewCommentRepo(db)
	vr := repo.NewActivityRepo(db)
	sr := repo.NewSourceRepo(db)
	fr := repo.NewFavoriteRepo(db)

	og := gateway.NewOGPGateway()
	tg := gateway.NewTechFeedGateway()
	ag := gateway.NewOpenAIGateway(ctx)
	uu := usecase.NewAccountUseCase(ur)
	ts := service.NewTopicAnalysisService(ag, tr, ar)
	ps := service.NewPersonalizeService(tr, cr)
	as := service.NewActivityService(cr)
	ss := service.NewScrapingService(ag, og)
	h := handler.NewHandler(tx, ur, vr, ar, mr, or, cr, sr, fr, og, uu, ts, ps, tr, tg, as, ss)

	private := e.Group("/account", h.AuthMiddleware)
	private.GET("/private/profile/:id", h.GetAccountProfile)
	e.GET("account/profile/:id", h.GetAccountProfile)
	e.GET("account/collection/:id", h.GetCollections)

	e.GET("/comment/:articleId", h.GetComments)
	private.POST("/comment", h.CreateComment)
	private.POST("/comment/:articleId", h.AddComment)
	private.POST("/bookmark", h.DoBookmark)
	private.POST("/collection", h.CreateCollection)
	private.GET("/collection", h.GetCollectionForBookmark)
	private.PUT("/article/read", h.SetReadArticle)
	private.GET("/article/read", h.GetReadArticle)
	private.GET("/session", h.Session)
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

	private.POST("/article/favorite/:id", h.DoFavoriteArticle)
	private.DELETE("/article/favorite/:id", h.UnFavoriteArticle)

	e.GET("/topic/:id", h.GetTopicData)
	private.GET("/topic/follow/:id", h.CheckTopicFollow)
	private.GET("/topic/follow", h.GetFollowingTopics)
	private.PUT("/topic/follow/:id", h.DoFollowTopic)
	private.DELETE("/topic/follow/:id", h.UnFollowTopic)

	e.GET("/source/:id", h.GetSourceData)
	private.GET("/source/follow/:id", h.CheckSourceFollow)
	private.GET("/source/follow", h.GetFollowingSources)
	private.PUT("/source/follow/:id", h.DoFollowSource)
	private.DELETE("/source/follow/:id", h.UnFollowSource)

	private.GET("/article/recommend", h.GetRecommendArticles)
	private.GET("/feeds", h.GetFeeds)

	e.GET("/trend/qiita", h.GetQiitaTrend)
	e.GET("/trend/zenn", h.GetZennTrend)

	sp := scheduler.NewSchedulerPool()
	tw := scheduler.NewTimelineWorker(ur, ar, sr, vr, as, cr, tr, ps, tg)
	go func() {
		sp.AddScheduler(key.TrendingArticlesWorker, 24*time.Hour, tw.CacheTredingArticlesWorker)
		sp.AddScheduler(key.PersonalizedArticlesWorker, 24*time.Hour, tw.CachePersonalizedArticlesWorker)
		sp.AddScheduler(key.QiitaTrendsWorker, 24*time.Hour, tw.RegisterQiitaTendsWorker)
		sp.AddScheduler(key.ContributioinWorker, 24*time.Hour, tw.ContributionWorker)
		sp.StartScheduler(key.TrendingArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.PersonalizedArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.QiitaTrendsWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.ContributioinWorker)
	}()
	e.Logger.Fatal(e.Start(":8080"))
}
