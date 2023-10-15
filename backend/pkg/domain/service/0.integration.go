package service

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
)

//go:generate mockgen -source=0.integration.go -destination=../../../mock/service/integration.go

type NLPService interface {
	GetAnswerFromPrompt(prompt string) (string, error)
	AsyncGetAnswerFromPrompt(prompt string) <-chan string
}

type OGPService interface {
	GetOGPByURL(url string) (*dto.OGPResponse, error)
}

type TechFeedService interface {
	GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error)
	GetZennTrendFeed() ([]model.Article, error)
}

type GithubService interface {
	GetStarCountByURL(url string) (int, error)
	FindTopStarRepo(star, limit int, language string) ([]*dto.GithubRepo, error)
}

type ScrapingService interface {
	GetAllContentByURL(url string) (string, error)
	GetHeaders(url string) ([]*dto.Header, error)
	GetHeaderContent(url string, headerContent string, headerLevel int) (string, error)
	GetHeaderContentWithHTMLElements(url string, headerContent string, headerLevel int) (string, error)
}

type TwitterService interface {
	GetTweetIDsByQuoteURL(url string)
}
