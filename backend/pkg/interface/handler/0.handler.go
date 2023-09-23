package handler

import (
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase"
)

type Handler struct {
	repo        repository.Repository
	cache       repository.CacheRepo
	openai      service.NLPService
	ogp         service.OGPService
	feed        service.TechFeedService
	git         service.GithubService
	account     *usecase.AccountUsecase
	article     *usecase.ArticleUsecase
	personalize *usecase.PersonalizeUsecase
	activity    *usecase.ActivityUsecase
	scraping    *usecase.ScrapingUsecase
}

func NewHandler(
	repo repository.Repository,
	cache repository.CacheRepo,
	openai service.NLPService,
	ogp service.OGPService,
	feed service.TechFeedService,
	git service.GithubService,
	account *usecase.AccountUsecase,
	article *usecase.ArticleUsecase,
	personalize *usecase.PersonalizeUsecase,
	activity *usecase.ActivityUsecase,
	scraping *usecase.ScrapingUsecase,
) *Handler {
	return &Handler{
		repo:        repo,
		cache:       cache,
		openai:      openai,
		ogp:         ogp,
		feed:        feed,
		git:         git,
		account:     account,
		article:     article,
		personalize: personalize,
		activity:    activity,
		scraping:    scraping,
	}
}
