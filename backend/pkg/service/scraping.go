package service

import (
	"fmt"

	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/util"
)

type ScrapingService struct {
	og gateway.OpenAIGateway
	gg gateway.OGPGateway
}

func NewScrapingService(og gateway.OpenAIGateway, gg gateway.OGPGateway) ScrapingService {
	return ScrapingService{og: og, gg: gg}
}

func (s ScrapingService) SummarizeArticle(url string) (string, error) {
	if len(url) < 1 {
		return "", fmt.Errorf("url must be set")
	}
	markdown, err := util.GetMarkdownByURL(url)
	if err != nil {
		return "", err
	}
	ogp, err := s.gg.GetOGPByURL(url)
	if err != nil {
		return "", err
	}
	articleInfo := fmt.Sprintf("[title: %s][description%s][content: %s]", ogp.Title, ogp.Description, markdown)
	prompt := fmt.Sprintf("[%s] 以上の内容を200文字以内で要点を絞って日本語でまとめて", articleInfo)
	summary, err := s.og.GetAnswerFromPrompt(prompt, 0.01)
	if err != nil {
		return "", err
	}
	return summary, nil
}
