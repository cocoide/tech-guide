package service_test

import (
	"context"
	"testing"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/service"
)

func TestSummarizeArticle(t *testing.T) {
	conf.NewEnv()
	ctx := context.Background()
	gg := gateway.NewOGPGateway()
	og := gateway.NewOpenAIGateway(ctx)
	ss := service.NewScrapingService(og, gg)
	url := "https://www.catapultsuplex.com/entry/producthunt-knowhow"
	result, err := ss.SummarizeArticle(url)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(result)
}
