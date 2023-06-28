package service

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/model/dto"
	repo "github.com/cocoide/tech-guide/pkg/repository"
)

type TopicAnalysisService interface {
	ExtractTopicsWithWeightFromArticleTitle(title string) ([]dto.TopicWeight, error)
	GetRelatedArticlesByOriginTopicToArticleArray(origin []model.ArticlesToTopics) ([]model.Article, error)
}

type topicAnalysisService struct {
	og gateway.OpenAIGateway
	tr repo.TopicRepo
	ar repo.ArticleRepo
}

func NewTopicAnalysisService(og gateway.OpenAIGateway, tr repo.TopicRepo, ar repo.ArticleRepo) TopicAnalysisService {
	return &topicAnalysisService{og: og, tr: tr, ar: ar}
}

func (ts *topicAnalysisService) GetRelatedArticlesByOriginTopicToArticleArray(origin []model.ArticlesToTopics) ([]model.Article, error) {
	originArticleID := origin[0].ArticleID

	originTopicWeights := make(map[int]int)
	for _, v := range origin {
		originTopicWeights[v.TopicID] = v.Weight
	}
	var originTopicIDs []int
	for topicID := range originTopicWeights {
		originTopicIDs = append(originTopicIDs, topicID)
	}
	targetTopicToArticles, err := ts.ar.GetArticlesByTopicIDs(originTopicIDs, originArticleID)
	if err != nil {
		return nil, err
	}
	type TopicSimilarity struct {
		Article    model.Article
		Similarity float64
	}
	var similarArticles []TopicSimilarity
	visited := make(map[int]TopicSimilarity)

	for _, v := range targetTopicToArticles {
		similarArticle, ok := visited[v.Article.ID]
		if ok {
			similarArticle.Similarity += 0.1
		} else {
			similarity := 0.0
			if weight, ok := originTopicWeights[v.TopicID]; ok {
				similarity = +float64(weight/5) + float64(v.Weight)/5
			}
			similarArticle = TopicSimilarity{
				Article:    v.Article,
				Similarity: similarity,
			}
		}
		visited[v.Article.ID] = similarArticle
	}
	for _, v := range visited {
		similarArticles = append(similarArticles, v)
	}
	sort.Slice(similarArticles, func(i, j int) bool {
		return similarArticles[i].Similarity > similarArticles[j].Similarity
	})
	result := []model.Article{}
	for _, v := range similarArticles {
		result = append(result, v.Article)
	}
	return result, nil
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
