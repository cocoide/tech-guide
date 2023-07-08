package handler

import (
	"github.com/cocoide/tech-guide/pkg/gateway"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/usecase"
)

type Handler struct {
	tx repo.TxRepo
	ur repo.AccountRepo
	ar repo.ArticleRepo
	vr repo.ActivityRepo
	mr repo.CommentRepo
	cr repo.CollectionRepo
	rr repo.CacheRepo
	tr repo.TopicRepo
	og gateway.OGPGateway
	tg gateway.TechFeedGateway
	uu usecase.AccountUseCase
	ts service.TopicAnalysisService
	ps service.PersonalizeService
	as service.ActivityService
}

func NewHandler(tx repo.TxRepo, ur repo.AccountRepo, vr repo.ActivityRepo, ar repo.ArticleRepo, mr repo.CommentRepo, cr repo.CollectionRepo, rr repo.CacheRepo, og gateway.OGPGateway, uu usecase.AccountUseCase, ts service.TopicAnalysisService, ps service.PersonalizeService, tr repo.TopicRepo, tg gateway.TechFeedGateway, as service.ActivityService) *Handler {
	return &Handler{tx: tx, ur: ur, ar: ar, vr: vr, mr: mr, cr: cr, rr: rr, og: og, uu: uu, ts: ts, ps: ps, tr: tr, tg: tg, as: as}
}
