package handler

import (
	"github.com/cocoide/tech-guide/pkg/gateway"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/usecase"
)

type Handler struct {
	ur repo.AccountRepo
	ar repo.ArticleRepo
	cr repo.CollectionRepo
	og gateway.OGPGateway
	uu usecase.AccountUseCase
}

func NewHandler(ur repo.AccountRepo, ar repo.ArticleRepo, cr repo.CollectionRepo, og gateway.OGPGateway, uu usecase.AccountUseCase) *Handler {
	return &Handler{ur: ur, ar: ar, cr: cr, og: og, uu: uu}
}
