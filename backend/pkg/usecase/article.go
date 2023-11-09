package usecase

import (
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"sort"
	"strconv"
	"strings"
)

const (
	ArticlePaginateSize = 12
)

type ArticleUsecase struct {
	nlp   service.NLPService
	cache repository.CacheRepo
	repo  repository.Repository
}

func NewArticleUsecase(nlp service.NLPService, cache repository.CacheRepo, repo repository.Repository) *ArticleUsecase {
	return &ArticleUsecase{nlp: nlp, cache: cache, repo: repo}
}

func (u *ArticleUsecase) GetLatestArticlesWithPaginate(pageIndex int) (model.Articles, error) {
	params := &repository.ListArticlesParams{
		OrderBy:  repository.Latest,
		Preloads: []string{"Rating", "Source"},
		PaginateOption: repository.PaginateOption{
			PageIndex: pageIndex,
			PageSize:  ArticlePaginateSize,
		},
	}
	return u.repo.ListArticles(params)
}

func (u *ArticleUsecase) GetDiscussArticlesWithPaginate(pageIndex int) (model.Articles, error) {
	params := &repository.ListArticlesParams{
		OrderBy:  repository.Discuss,
		Preloads: []string{"Rating", "Source"},
		PaginateOption: repository.PaginateOption{
			PageIndex: pageIndex,
			PageSize:  ArticlePaginateSize,
		},
	}
	return u.repo.ListArticles(params)
}

func (u *ArticleUsecase) GetTrendArticlesWithPaginate(pageIndex int) (model.Articles, error) {
	params := &repository.ListArticlesParams{
		OrderBy:  repository.Trend,
		Preloads: []string{"Rating", "Source"},
		PaginateOption: repository.PaginateOption{
			PageIndex: pageIndex,
			PageSize:  ArticlePaginateSize,
		},
	}
	return u.repo.ListArticles(params)
}

func (u *ArticleUsecase) GetFeedArticlesWithPaginate(accountId, pageIndex int) (model.Articles, error) {
	params := &repository.ListArticlesParams{
		OrderBy:  repository.Latest,
		Preloads: []string{"Rating", "Source"},
		FeedOption: repository.FeedOption{
			AccountID:      accountId,
			IsFollowDomain: true,
			IsFollowTopic:  true,
		},
		PaginateOption: repository.PaginateOption{
			PageIndex: pageIndex,
			PageSize:  ArticlePaginateSize,
		},
	}
	return u.repo.ListArticles(params)
}

func (u *ArticleUsecase) paginateArticleIDs(allArticleIDs []int, pageIndex int) []int {
	start := ArticlePaginateSize * (pageIndex - 1)

	if start >= len(allArticleIDs) {
		return []int{}
	}
	end := start + ArticlePaginateSize
	if end > len(allArticleIDs) {
		end = len(allArticleIDs)
	}
	return allArticleIDs[start:end]
}

func (ts *ArticleUsecase) GetRelatedArticlesByOriginTopicToArticleArray(origin []model.TopicsToArticles, excludeID int) ([]model.Article, error) {
	originArticleID := origin[0].ArticleID

	originTopicWeights := make(map[int]int)
	for _, v := range origin {
		originTopicWeights[v.TopicID] = v.Weight
	}
	var originTopicIDs []int
	for topicID := range originTopicWeights {
		originTopicIDs = append(originTopicIDs, topicID)
	}
	targetTopicToArticles, err := ts.repo.GetArticlesByTopicIDs(originTopicIDs, originArticleID)
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
		if v.Article.ID == excludeID {
			continue
		}
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

func (ts *ArticleUsecase) ExtractTopicsWithWeightFromArticleTitle(title, summary string) ([]dto.TopicWeight, error) {
	existingTopics, err := ts.repo.GetAllTopics()
	existingTopicsName := ""
	for i, t := range existingTopics {
		if i > 0 {
			existingTopicsName += ", "
		}
		existingTopicsName += t.Name
	}
	prompt := conf.NewSelectTopicsPrompt(title, summary, existingTopics)
	answer, err := ts.nlp.GetAnswerFromPrompt(prompt)
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
