package usecase

import (
	"fmt"

	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/utils"
)

type ScrapingUsecase struct {
	openai service.OpenAIService
	ogp    service.OGPService
}

func NewScrapingUsecase(openai service.OpenAIService, ogp service.OGPService) *ScrapingUsecase {
	return &ScrapingUsecase{openai: openai, ogp: ogp}
}

func (s ScrapingUsecase) SummarizeArticle(url string) (string, error) {
	if len(url) < 1 {
		return "", fmt.Errorf("url must be set")
	}
	markdown, err := utils.GetMarkdownByURL(url)
	if err != nil {
		return "", err
	}
	ogp, err := s.ogp.GetOGPByURL(url)
	if err != nil {
		return "", err
	}
	articleInfo := fmt.Sprintf("[title: %s][description%s][content: %s]", ogp.Title, ogp.Description, markdown)
	prompt := fmt.Sprintf("[%s] 以上の内容を200文字以内で要点を絞って日本語でまとめて", articleInfo)
	summary, err := s.openai.GetAnswerFromPrompt(prompt, 0.01)
	if err != nil {
		return "", err
	}
	return summary, nil
}
