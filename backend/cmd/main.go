package main

import (
	"context"
	"time"

	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/interface/handler"
	"github.com/cocoide/tech-guide/pkg/interface/router"

	cacheRepo "github.com/cocoide/tech-guide/pkg/infrastructure/cache"
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	rdbRepo "github.com/cocoide/tech-guide/pkg/infrastructure/rdb"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/key"

	"github.com/cocoide/tech-guide/pkg/scheduler"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	conf.NewEnv()
	ctx := context.Background()

	gorm := database.NewGORMDatabase()
	redis := database.NewRedisCilent(ctx)

	cache := cacheRepo.NewRepository(redis, ctx)
	repo := rdbRepo.NewRepository(gorm)

	openai := integration.NewOpenAIService(ctx)
	ogp := integration.NewOGPService()
	feed := integration.NewTechFeedService()

	account := usecase.NewAccountUsecase(repo)
	activity :=usecase.NewActivityUsecase(cache)
	article:=usecase.NewArticleUsecase(openai, repo)
	personalize := usecase.NewPersonalizeUsecase(repo, cache)
	scraping:=usecase.NewScrapingUsecase(openai, ogp)

	handler := handler.NewHandler(repo, cache, openai, ogp, feed, account, article, personalize, activity, scraping)
	router.NewRouter(e, handler)
	
	sp := scheduler.NewSchedulerPool()
	tw := scheduler.NewTimelineWorker(repo, cache, activity, feed, personalize)
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
