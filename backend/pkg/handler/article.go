package handler

import (
	"context"
	"fmt"
	"sort"
	"strconv"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/labstack/echo"
)

func (h *Handler) GetOGP(c echo.Context) error {
	ctx := context.Background()
	var limiter = util.NewRateLimiter(5, 10)
	if err := limiter.Limit(ctx, func() error {
		URL := c.QueryParam("url")
		ogp, err := h.og.GetOGPByURL(URL)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		return c.JSON(200, ogp)
	}); err != nil {
		return c.JSON(401, err.Error())
	}
	return nil
}

func (h *Handler) GetArticles(c echo.Context) error {
	articles, err := h.ar.GetLatestArticleByLimit(50)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}

func (h *Handler) GetRelatedArticles(c echo.Context) error {
	articlelId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	articlesToTopics, err := h.ar.GetTagsAndWeightsByArticleID(articlelId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	var topicIDs []int
	for _, v := range articlesToTopics {
		topicIDs = append(topicIDs, v.TopicID)
	}
	// Topicはpermoanceのためにあえてpreloadしていない
	articles, err := h.ar.GetArticlesByTopicIDs(topicIDs, articlelId)
	type SimilarArticle struct {
		Article    model.Article
		Similarity float64
	}
	var similarArticles []SimilarArticle
	// Key: ArticleID
	visited := make(map[int]SimilarArticle)
	for _, v := range articles {
		similarArticle, ok := visited[v.Article.ID]
		if ok {
			// 複数のTopicを共通して持つ場合はSimilarityを0.1増す
			similarArticle.Similarity += 0.1
		} else {
			similarArticle = SimilarArticle{
				Article:    v.Article,
				Similarity: float64(v.Weight) / 5, // 正規化
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
	return c.JSON(200, result)
}

func (h *Handler) CreateArticle(c echo.Context) error {
	type REQ struct {
		OriginalURL string `json:"original_url"`
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	URL := req.OriginalURL
	ogp, err := h.og.GetOGPByURL(URL)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	article := &model.Article{
		Title:        ogp.Title,
		ThumbnailURL: ogp.Thumbnail,
		OriginalURL:  URL,
	}
	if _, err := h.ar.Create(article); err != nil {
		return c.JSON(400, err.Error())
	}
	topicWeights, err := h.ts.ExtractTopicsWithWeightFromArticleTitle(article.Title)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	var topicToArticles []model.ArticlesToTopics
	for _, v := range topicWeights {
		topicToArticles = append(topicToArticles,
			model.ArticlesToTopics{ArticleID: article.ID, TopicID: v.ID, Weight: v.Weight})
	}
	fmt.Print(topicWeights)
	if err := h.ar.CreateTopicToArticle(topicToArticles); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, article)
}
