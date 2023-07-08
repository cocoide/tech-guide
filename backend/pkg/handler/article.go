package handler

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/model/dto"
	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/labstack/echo"
)

func (h *Handler) GetRecommendArticles(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	strArticleIDs, err := h.rr.Get(fmt.Sprintf(key.PersonalizedArticleIDs, accountId))
	var articleIDs []int
	if len(strArticleIDs) < 1 || err != nil {
		articleIDs, err = h.ps.GetRecommendArticleIDs(accountId)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		strArticleIDs, err = util.Serialize(articleIDs)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if err := h.rr.Set(fmt.Sprintf(key.PersonalizedArticleIDs, accountId), strArticleIDs, 24*time.Hour); err != nil {
			return c.JSON(400, err.Error())
		}
	} else {
		articleIDs, err = util.Deserialize[[]int](strArticleIDs)
		if err != nil {
			return c.JSON(400, err.Error())
		}
	}
	articles, err := h.ar.GetArticlesByIDs(articleIDs)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}

func (h *Handler) GetRSS(c echo.Context) error {
	result, err := util.FetchXML[dto.RSSFeed](c.QueryParam("url"), 5*time.Second)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, result)
}

func (h *Handler) GetOGP(c echo.Context) error {
	ctx := context.Background()
	var limiter = util.NewRateLimiter(5, 10)
	if err := limiter.Limit(ctx, func() error {
		URL := c.QueryParam("url")
		ogp, err := h.og.GetOGPByURL(URL)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if len(ogp.Thumbnail) > 500 {
			ogp.Thumbnail = ""
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
	thumbnailURL := ""
	if len(ogp.Thumbnail) < 500 {
		thumbnailURL = ogp.Thumbnail
	}
	article := &model.Article{
		Title:        ogp.Title,
		ThumbnailURL: thumbnailURL,
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
		var topicToArticles []model.TopicsToArticles
		for _, v := range topicWeights {
			topicToArticles = append(topicToArticles,
				model.TopicsToArticles{ArticleID: article.ID, TopicID: v.ID, Weight: v.Weight})
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
