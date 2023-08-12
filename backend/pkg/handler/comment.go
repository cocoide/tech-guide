package handler

import (
	"log"
	"strconv"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/model/dto"
	"github.com/cocoide/tech-guide/pkg/model/object"
	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/labstack/echo"
)

func (h *Handler) AddComment(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	articleID, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	type request struct {
		Content string `json:"content"`
	}
	req := new(request)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}

	comment := &model.Comment{
		AccountID: accountId,
		ArticleID: articleID,
		Content:   req.Content,
	}
	if err := h.mr.Create(comment); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) CreateComment(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	type body struct {
		OriginalURL string `json:"original_url"`
		Content     string `json:"content"`
	}
	req := new(body)
	if err := c.Bind(req); err != nil {
		return c.JSON(403, err.Error())
	}
	articleId, err := h.ar.GetArticleIDByURL(req.OriginalURL)
	if err != nil {
		return c.JSON(404, err.Error())
	}
	if articleId != 0 {
		domain := "tech-guide.jp"
		sourceID, err := h.sr.FindIDByDomain(domain)
		if err != nil {
			return c.JSON(400, err.Error())
		}
		ogp, err := h.og.GetOGPByURL(req.OriginalURL)
		if err != nil {
			return c.JSON(404, err.Error())
		}
		thumbnailURL := ""
		if len(ogp.Thumbnail) < 500 {
			thumbnailURL = ogp.Thumbnail
		}
		article := &model.Article{Title: ogp.Title, ThumbnailURL: thumbnailURL, OriginalURL: req.OriginalURL, SourceID: sourceID}
		createdArticleID, err := h.ar.Create(article)
		if len(req.Content) > 0 {
			comment := &model.Comment{
				AccountID: accountId,
				ArticleID: createdArticleID,
				Content:   req.Content,
			}
			if err := h.mr.Create(comment); err != nil {
				return c.JSON(400, err.Error())
			}
		}
		return c.JSON(200, "community post created")
	}

	ogp, err := h.og.GetOGPByURL(req.OriginalURL)
	if err != nil {
		return c.JSON(404, err.Error())
	}
	domainName, err := util.ExtractDomainNameFromURL(req.OriginalURL)
	if err != nil {
		return c.JSON(404, err.Error)
	}
	sourceID, err := h.sr.FindIDByDomain(domainName)
	// domainが設定されてない場合、Communityのsourceを設定
	if sourceID == 0 {
		sourceID, err = h.sr.FindIDByDomain(object.OriginDomain)
		if err != nil {
			return c.JSON(400, err.Error())
		}
	}
	thumbnailURL := ""
	if len(ogp.Thumbnail) < 500 {
		thumbnailURL = ogp.Thumbnail
	}
	createdArticleID, err := h.ar.Create(&model.Article{Title: ogp.Title, ThumbnailURL: thumbnailURL, OriginalURL: req.OriginalURL, SourceID: sourceID})
	if err != nil {
		return c.JSON(404, err.Error())
	}
	// commentがある場合のみ保存
	if len(req.Content) > 0 {
		comment := &model.Comment{
			AccountID: accountId,
			ArticleID: createdArticleID,
			Content:   req.Content,
		}
		if err := h.mr.Create(comment); err != nil {
			return c.JSON(400, err.Error())
		}
	}
	topicAssignErrCh := make(chan error)
	go func() {
		wholeTopicWeights, err := h.ts.ExtractTopicsWithWeightFromArticleTitle(ogp.Title)
		if err != nil {
			topicAssignErrCh <- err
			return
		}
		topicWeights := dto.GetTopWeightedTopics(wholeTopicWeights, 3)
		var topicToArticles []model.TopicsToArticles
		for _, v := range topicWeights {
			// 関連度5以上の記事だけを保存
			if v.Weight >= 5 {
				topicToArticles = append(topicToArticles,
					model.TopicsToArticles{ArticleID: createdArticleID, TopicID: v.ID, Weight: v.Weight})
			}
		}
		if len(topicToArticles) > 0 {
			if err := h.ar.CreateTopicToArticle(topicToArticles); err != nil {
				topicAssignErrCh <- err
				return
			}
		}
	}()
	summaryErrCh := make(chan error)
	go func() {
		summary, err := h.ss.SummarizeArticle(req.OriginalURL)
		summaryErrCh <- err
		if err := h.ar.UpdateSummaryByID(createdArticleID, summary); err != nil {
			summaryErrCh <- err
		}
	}()
	go func() {
		for err := range topicAssignErrCh {
			if err != nil {
				log.Print(err)
			}
		}
	}()
	go func() {
		for err := range summaryErrCh {
			if err != nil {
				log.Print(err)
			}
		}
	}()
	return c.JSON(200, "comment created")
}

func (h *Handler) GetComments(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("articleId"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	comments, err := h.mr.GetComments(articleId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, comments)
}
