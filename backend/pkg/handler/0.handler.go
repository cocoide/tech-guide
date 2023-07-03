package handler

import (
	"github.com/cocoide/tech-guide/pkg/gateway"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/usecase"
)

type Handler struct {
	ur repo.AccountRepo
	ar repo.ArticleRepo
	cr repo.CollectionRepo
	rr repo.CacheRepo
	tr repo.TopicRepo
	og gateway.OGPGateway
	tg gateway.TechFeedGateway
	uu usecase.AccountUseCase
	ts service.TopicAnalysisService
	ps service.PersonalizeService
}

func NewHandler(ur repo.AccountRepo, ar repo.ArticleRepo, cr repo.CollectionRepo, rr repo.CacheRepo, og gateway.OGPGateway, uu usecase.AccountUseCase, ts service.TopicAnalysisService, ps service.PersonalizeService, tr repo.TopicRepo, tg gateway.TechFeedGateway) *Handler {
	return &Handler{ur: ur, ar: ar, cr: cr, rr: rr, og: og, uu: uu, ts: ts, ps: ps, tr: tr, tg: tg}
}
