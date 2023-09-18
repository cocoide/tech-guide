package main

import (
	"context"

	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/interface/handler"
	"github.com/cocoide/tech-guide/pkg/interface/router"

	cacheRepo "github.com/cocoide/tech-guide/pkg/infrastructure/cache"
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	rdbRepo "github.com/cocoide/tech-guide/pkg/infrastructure/rdb"

	"github.com/cocoide/tech-guide/conf"

	"github.com/cocoide/tech-guide/pkg/scheduler"
	"github.com/cocoide/tech-guide/pkg/usecase"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	conf.NewEnv()
	ctx := context.Background()

	gorm := database.NewGormDatabase()
	redis := database.NewRedisCilent(ctx)

	cache := cacheRepo.NewRepository(redis, ctx)
	repo := rdbRepo.NewRepository(gorm)

	openai := integration.NewOpenAIService(ctx)
	ogp := integration.NewOGPService()
	feed := integration.NewTechFeedService()

	account := usecase.NewAccountUsecase(repo)
	activity := usecase.NewActivityUsecase(cache)
	article := usecase.NewArticleUsecase(openai, repo)
	personalize := usecase.NewPersonalizeUsecase(repo, cache)
	scraping := usecase.NewScrapingUsecase(openai, ogp)

	rootHandler := handler.NewHandler(repo, cache, openai, ogp, feed, account, article, personalize, activity, scraping)
	router.NewRootRouter(e, rootHandler)

	sp := scheduler.NewSchedulerPool()
	tw := scheduler.NewTimelineWorker(repo, cache, activity, feed, personalize)
	scheduler.NewAsyncJobRunner(sp, tw)

	e.Logger.Fatal(e.Start(":8080"))
}
