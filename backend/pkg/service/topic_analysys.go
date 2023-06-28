package service

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/model/dto"
	repo "github.com/cocoide/tech-guide/pkg/repository"
)

type TopicAnalysisService interface {
	ExtractTopicsWithWeightFromArticleTitle(title string) ([]dto.TopicWeight, error)
}

type topicAnalysisService struct {
	og gateway.OpenAIGateway
	tr repo.TopicRepo
}

func NewTopicAnalysisService(og gateway.OpenAIGateway, tr repo.TopicRepo) TopicAnalysisService {
	return &topicAnalysisService{og: og, tr: tr}
}
func (ts *topicAnalysisService) ExtractTopicsWithWeightFromArticleTitle(title string) ([]dto.TopicWeight, error) {
	existingTopics, err := ts.tr.GetAllTopics()
	existingTopicsName := ""
	for i, t := range existingTopics {
		if i > 0 {
			existingTopicsName += ", "
		}
		existingTopicsName += t.Name
	}
	prompt := fmt.Sprintf(conf.SelectTopicsPrompt, title, existingTopicsName)
	answer, err := ts.og.GetAnswerFromPrompt(prompt, 0)
	if err != nil {
		return nil, err
	}
	topicWeights := make(map[string]int)
	topicWeightPairs := strings.Split(answer, ",")
	for _, pair := range topicWeightPairs {
		parts := strings.Split(strings.TrimSpace(pair), ":")
		if len(parts) != 2 {
			continue
		}
		topic := strings.TrimSpace(parts[0])
		weight, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}
		topicWeights[topic] = weight
	}
	var result []dto.TopicWeight
	for _, existingTopic := range existingTopics {
		if weight, ok := topicWeights[existingTopic.Name]; ok {
			result = append(result, dto.TopicWeight{ID: existingTopic.ID, Name: existingTopic.Name, Weight: weight})
		}
	}
	return result, nil
}
