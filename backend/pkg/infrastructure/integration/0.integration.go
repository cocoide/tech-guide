package integration

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/sashabaranov/go-openai"
)

type openAIService struct {
	client *openai.Client
	ctx    context.Context
}

func NewOpenAIService(ctx context.Context) service.OpenAIService {
	OPENAI_SECRET := os.Getenv("OPENAI_SECRET")
	client := openai.NewClient(OPENAI_SECRET)
	return &openAIService{client: client, ctx: ctx}
}

type ogpService struct {
	client *http.Client
}

func NewOGPService() service.OGPService {
	return &ogpService{client: &http.Client{Timeout: 10 * time.Second}}
}

type techFeedService struct{}

func NewTechFeedService() service.TechFeedService {
	return &techFeedService{}
}
