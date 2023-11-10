package usecase

import (
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"log"
	"strconv"
	"strings"
)

type topicUsecaseRepo interface {
	repository.TopicRepo
	repository.ArticleRepo
}

type TopicUsecase struct {
	repo    topicUsecaseRepo
	nlp     service.NLPService
	scraper service.ScrapingService
}

func NewTopicUsecase(repo topicUsecaseRepo, nlp service.NLPService, scraper service.ScrapingService) *TopicUsecase {
	return &TopicUsecase{repo: repo, nlp: nlp, scraper: scraper}
}

func (u *TopicUsecase) ListPopularTopics() (model.Topics, error) {
	params := &repository.ListTopicsParams{
		MinHasArticlesCount: 1,
		OrderBy:             repository.Follow,
	}
	return u.repo.ListTopics(params)
}

func (u *TopicUsecase) extractTopicToArticlesByArticleID(articleID int) ([]model.TopicsToArticles, error) {
	article, err := u.repo.GetArticleByID(articleID)
	if err != nil {
		return nil, err
	}
	categories, err := u.repo.GetCategories()
	if err != nil {
		return nil, err
	}

	categoryIDs, err := u.selectCategoryIDsFromArticle(categories, article)
	if err != nil {
		return nil, err
	}
	topics, _ := u.repo.GetTopicsByCategoryIDs(categoryIDs)
	topicToArticles, err := u.assignTopicsForArticle(article, topics)
	if err != nil {
		return nil, err
	}
	return topicToArticles, nil
}

func (u *TopicUsecase) UpsertTopicToArticlesByArticleID(articleID int) error {
	newTopicToArticles, err := u.extractTopicToArticlesByArticleID(articleID)
	if err != nil {
		return err
	}
	existTopicToArticle, err := u.repo.GetTopicToArticleArrayByArticleID(articleID)
	if err != nil {
		return err
	}

	weightMap := make(map[int]int) // key: topicID, value: weight
	for _, t := range existTopicToArticle {
		weightMap[t.TopicID] = t.Weight
	}

	var toInsert, toUpdate []model.TopicsToArticles
	for _, t := range newTopicToArticles {
		if weight, exists := weightMap[t.TopicID]; exists {
			t.Weight = (weight + t.Weight) / 2
			toUpdate = append(toUpdate, t)
		} else {
			toInsert = append(toInsert, t)
		}
	}
	if len(toInsert) > 0 {
		if err := u.repo.CreateTopicToArticle(toInsert); err != nil {
			return err
		}
	}
	if len(toUpdate) > 0 {
		if err := u.repo.UpdateTopicToArticle(toUpdate); err != nil {
			return err
		}
	}
	return nil
}

// 1. 割り振るTopicをCategoryで分類する
func (u *TopicUsecase) selectCategoryIDsFromArticle(categories []model.Category, article *model.Article) ([]int, error) {
	categoryWeightSet := make(map[int]int) // categoryID -> relatedWeight
	prompt := conf.NewSelectCategoriesPrompt(article.Title, article.Summary, categories)
	answer, err := u.nlp.GetAnswerFromPrompt(prompt)
	if err != nil {
		return nil, err
	}
	log.Print(answer)
	categoryMap := u.trimSelectCategoriesNLPResult(answer)
	for _, category := range categories {
		if weight, ok := categoryMap[category.Name]; ok {
			categoryWeightSet[category.ID] = weight
		}
	}
	categoryIDs, _ := extractCategoryIDs(categoryWeightSet)
	return categoryIDs, nil
}
func extractCategoryIDs(categoryWeightSet map[int]int) ([]int, error) {
	var relatedCategoryIDs []int
	var normalCategoryIDs []int

	// カテゴリを関連カテゴリと通常カテゴリに分類
	for categoryID, weight := range categoryWeightSet {
		if weight > 4 {
			relatedCategoryIDs = append(relatedCategoryIDs, categoryID)
		} else if weight > 3 {
			normalCategoryIDs = append(normalCategoryIDs, categoryID)
		}
	}

	// categoryIDsを設定
	var categoryIDs []int
	if len(relatedCategoryIDs) > 2 {
		categoryIDs = relatedCategoryIDs
	} else {
		// 調整必要
		categoryIDs = append(categoryIDs, relatedCategoryIDs...)
		remaining := 3 - len(relatedCategoryIDs)
		if len(normalCategoryIDs) > remaining {
			categoryIDs = append(categoryIDs, normalCategoryIDs[:remaining]...)
		} else {
			categoryIDs = append(categoryIDs, normalCategoryIDs...)
		}
	}

	return categoryIDs, nil
}

// 2. Categoryの割り振りが完了した後、にTopicsを割り振り保存する
func (u *TopicUsecase) assignTopicsForArticle(article *model.Article, topics []model.Topic) ([]model.TopicsToArticles, error) {
	var topicToArticles []model.TopicsToArticles
	summary := article.Summary
	if summary == "" {
		headers, _ := u.scraper.GetHeaders(article.OriginalURL)
		for _, header := range headers {
			summary += header.ToMarkdown()
		}
	}
	prompt := conf.NewSelectTopicsPrompt(article.Title, article.Summary, topics)
	answer, err := u.nlp.GetAnswerFromPrompt(prompt)
	log.Printf(answer)
	if err != nil {
		log.Fatal(err)
	}
	allTopics, _ := u.repo.GetAllTopics()
	topicWeightSet := u.trimSelectTopicsNLPResult(answer)
	for _, t := range allTopics {
		// 提示されたTopicNameが既存のDBに存在するものかをCheck
		if weight, ok := topicWeightSet[t.Name]; ok && weight > 4 {
			topicToArticles = append(topicToArticles,
				model.TopicsToArticles{ArticleID: article.ID, TopicID: t.ID, Weight: weight})
		}
	}
	return topicToArticles, nil
}

// Response: Map[categoryName -> relatedWeight]
func (u *TopicUsecase) trimSelectCategoriesNLPResult(result string) map[string]int {
	// (Result: EX): [categoryA: weight, categoryB: weight, ...]
	categoryWeightSet := make(map[string]int)
	categoryWeightPairs := strings.Split(result, ",")
	for _, pair := range categoryWeightPairs {
		parts := strings.Split(strings.TrimSpace(pair), ":")
		if len(parts) != 2 {
			continue
		}
		category := strings.TrimSpace(parts[0])
		weight, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}
		categoryWeightSet[category] = weight
	}
	return categoryWeightSet
}

func (u *TopicUsecase) trimSelectTopicsNLPResult(result string) map[string]int {
	// (Result: EX): [topicA: weight, topicB: weight, ...]
	topicWeightSet := make(map[string]int)
	topicWeightPairs := strings.Split(result, ",")
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
		topicWeightSet[topic] = weight
	}
	return topicWeightSet
}
