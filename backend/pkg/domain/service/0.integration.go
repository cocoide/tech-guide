package service

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
)

//go:generate mockgen -source=0.integration.go -destination=../../../mock/service/integration.go
type OpenAIService interface {
	GetAnswerFromPrompt(prompt string, variability float32) (string, error)
	AsyncGetAnswerFromPrompt(prompt string, variability float32) <-chan string
}

//go:generate mockgen -source=0.integration.go -destination=../../../mock/service/integration.go
type OGPService interface {
	GetOGPByURL(url string) (*dto.OGPResponse, error)
}

//go:generate mockgen -source=0.integration.go -destination=../../../mock/service/integration.go
type TechFeedService interface {
	GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error)
	GetZennTrendFeed() ([]model.Article, error)
}
