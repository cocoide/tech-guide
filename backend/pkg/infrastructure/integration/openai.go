package integration

import "github.com/sashabaranov/go-openai"

func (s *openAIService) GetAnswerFromPrompt(prompt string, variability float32) (string, error) {
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
	res, err := s.client.CreateChatCompletion(s.ctx, req)
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
