package usecase_test

import (
	"context"
	"testing"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/infrastructure/integration"
	"github.com/cocoide/tech-guide/pkg/usecase"
)

func TestSummarizeArticle(t *testing.T) {
	conf.NewEnv()
	ctx:=context.Background()
	ogp := integration.NewOGPService()
	openai :=integration.NewOpenAIService(ctx)
	ss := usecase.NewScrapingUsecase(openai, ogp)
	url := "https://www.catapultsuplex.com/entry/producthunt-knowhow"
	result, err := ss.SummarizeArticle(url)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}
