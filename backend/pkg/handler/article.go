package handler

import (
	"context"
	"log"
	"strconv"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/model/dto"
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

func (h *Handler) GetArticleDetail(c echo.Context) error {
	articleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	articles, err := h.ar.GetArticleByID(articleID)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}

func (h *Handler) GetRelatedArticles(c echo.Context) error {
	excludeArticleID := 0
	if c.QueryParam("exclude") != "" {
		excludeArticleID, _ = strconv.Atoi(c.QueryParam("exclude"))
	}
	originArticleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	originTopicToArticleArray, err := h.tr.GetTopicToArticleArrayByArticleID(originArticleID)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	var result []model.Article
	if len(originTopicToArticleArray) > 0 {
		result, err = h.ts.GetRelatedArticlesByOriginTopicToArticleArray(originTopicToArticleArray, excludeArticleID)
		if err != nil {
			return c.JSON(400, err.Error())
		}
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
	topicAssignErrCh := make(chan error)
	go func() {
		wholeTopicWeights, err := h.ts.ExtractTopicsWithWeightFromArticleTitle(article.Title)
		if err != nil {
			topicAssignErrCh <- err
			return
		}
		topicWeights := dto.GetTopWeightedTopics(wholeTopicWeights, 3)
		var topicToArticles []model.ArticlesToTopics
		for _, v := range topicWeights {
			topicToArticles = append(topicToArticles,
				model.ArticlesToTopics{ArticleID: article.ID, TopicID: v.ID, Weight: v.Weight})
		}
		if err := h.ar.CreateTopicToArticle(topicToArticles); err != nil {
			topicAssignErrCh <- err
			return
		}
	}()
	go func() {
		for err := range topicAssignErrCh {
			log.Print(err)
		}
	}()
	return c.JSON(200, article)
}
