package main

import (
	"flag"
	"github.com/cocoide/tech-guide/cmd/update_articles/helper"
	"github.com/cocoide/tech-guide/pkg/infrastructure/database"
	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/infrastructure/rdb"
	"log"
)

const (
	ApplyThumbnailURL = 1
	ApplyTopics       = 2
)

func main() {
	gorm := database.NewGormClient()
	repo := rdb.NewRepository(gorm)
	ogp := integration.NewOGPService()
	scraper := integration.NewScrapingService()
	usecase := helper.NewUpdateArticlesUsecase(gorm, repo, ogp, scraper)

	targetPtr := flag.Int("type", 0, "usecase type")
	flag.Parse()
	target := *targetPtr
	switch target {
	case ApplyThumbnailURL:
		usecase.ApplyThumbnailURL()
	case ApplyTopics:
		usecase.ApplyTopics()
	default:
		log.Print("usecase not selected")
	}
}
