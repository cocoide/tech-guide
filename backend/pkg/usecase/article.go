package usecase

import (
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase/keyutils"
	"github.com/cocoide/tech-guide/pkg/usecase/parser"
	"github.com/cocoide/tech-guide/pkg/usecase/timeutils"
	"sort"
	"strconv"
	"strings"
)

const (
	FeedsLimit        = 100
	FeedsPaginateSize = 6
)

type ArticleUsecase struct {
	nlp   service.NLPService
	cache repository.CacheRepo
	repo  repository.Repository
}

func NewArticleUsecase(nlp service.NLPService, cache repository.CacheRepo, repo repository.Repository) *ArticleUsecase {
	return &ArticleUsecase{nlp: nlp, cache: cache, repo: repo}
}

func (u *ArticleUsecase) GetTrendArticles() (model.Articles, error) {
	params := &repository.ListArticlesParams{
		OrderBy:  repository.Trend,
		Limit:    FeedsLimit,
		Preloads: []string{"Source", "Rating"},
	}
	return u.repo.ListArticles(params)
}

func (u *ArticleUsecase) GetFeedsWithCache(accountId, pageIndex int) (model.Articles, error) {
	var articles model.Articles
	cacheKey := keyutils.CacheFeedArticleIDs(accountId)
	// Feedは一日ごとにキャッシュ
	str, exist, err := u.cache.Get(cacheKey)
	if err != nil {
		return nil, err
	}
	// キャッシュが存在する場合
	if exist {
		articleIDs, err := parser.Deserialize[[]int](str)
		if err != nil {
			return nil, err
		}
		paginateIDs := u.paginateArticleIDs(articleIDs, pageIndex)

		articles, err = u.repo.GetArticlesByIDs(paginateIDs, []string{"Source", "Rating"})
		if err != nil {
			return nil, err
		}
		return articles, err
	}

	topicIDs, err := u.repo.GetFollowingTopicIDs(accountId)
	if err != nil {
		return nil, err
	}

	sourceIDs, err := u.repo.GetFollowingSourceIDs(accountId)
	if err != nil {
		return nil, err
	}
	params := &repository.ListArticlesParams{
		TopicIDs:  topicIDs,
		SourceIDs: sourceIDs,
		OrderBy:   repository.Latest,
		Limit:     FeedsLimit,
	}
	articleIDs, err := u.repo.ListArticleIDs(params)

	if err != nil {
		return nil, err
	}
	paginateIDs := u.paginateArticleIDs(articleIDs, pageIndex)
	articles, err = u.repo.GetArticlesByIDs(paginateIDs, []string{"Rating", "Source"})
	if err != nil {
		return nil, err
	}

	strArticleIDs, err := parser.Serialize(articles.IDs())
	if err != nil {
		return nil, err
	}
	if err := u.cache.Set(cacheKey, strArticleIDs, timeutils.OneDay()); err != nil {
		return nil, err
	}
	return articles, nil
}

func (u *ArticleUsecase) paginateArticleIDs(allArticleIDs []int, pageIndex int) []int {
	start := FeedsPaginateSize * (pageIndex - 1)

	if start >= len(allArticleIDs) {
		return []int{}
	}
	end := start + FeedsPaginateSize
	if end > len(allArticleIDs) {
		end = len(allArticleIDs)
	}
	return allArticleIDs[start:end]
}

func (u *ArticleUsecase) GetArticlesByTopicIDWithPagination() model.Articles {
	articles, _ := u.repo.GetArticlesByIDs([]int{1, 2, 3, 5}, []string{"Source", "Rating"})
	return articles
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
