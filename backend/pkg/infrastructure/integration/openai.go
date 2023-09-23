package integration

import (
	"context"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/sashabaranov/go-openai"
	"os"
)

type openAIService struct {
	client *openai.Client
}

func NewNLPService() service.NLPService {
	OPENAI_SECRET := os.Getenv("OPENAI_SECRET")
	client := openai.NewClient(OPENAI_SECRET)
	return &openAIService{client: client}
}

func (s *openAIService) GetAnswerFromPrompt(prompt string, variability float32) (string, error) {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
		Temperature: variability,
	}
	res, err := s.client.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}
	answer := res.Choices[0].Message.Content
	return answer, nil
}

func (s *openAIService) AsyncGetAnswerFromPrompt(prompt string, variability float32) <-chan string {
	responseCh := make(chan string, 1)

	go func() {
		answer, _ := s.GetAnswerFromPrompt(prompt, variability)
		responseCh <- answer
	}()

	return responseCh
}
