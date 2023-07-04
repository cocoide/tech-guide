package handler

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/labstack/echo"
)

func (h *Handler) CreateComment(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	type body struct {
		OriginalURL string `json:"original_url"`
		Content     string `json:"comment"`
	}
	req := new(body)
	if err := c.Bind(req); err != nil {
		return c.JSON(403, err.Error())
	}

	articleId, err := h.ar.GetArticleIDByURL(req.OriginalURL)
	if err != nil {
		return c.JSON(404, err.Error())
	}
	if articleId == 0 {
		ogp, err := h.og.GetOGPByURL(req.OriginalURL)
		if err != nil {
			return c.JSON(404, err.Error())
		}
		createdArticleID, err := h.ar.Create(&model.Article{Title: ogp.Title, OriginalURL: req.OriginalURL})
		if err != nil {
			return c.JSON(404, err.Error())
		}
		comment := &model.Comment{
			AccountID: accountId,
			ArticleID: createdArticleID,
			Content:   req.Content,
		}
		if err := h.mr.Create(comment); err != nil {
			return c.JSON(400, err.Error())
		}
	} else {
		comment := &model.Comment{AccountID: accountId, ArticleID: articleId, Content: req.Content}
		if err := h.mr.Create(comment); err != nil {
			return c.JSON(400, err.Error())
		}
	}
	return c.JSON(200, "comment created")
}
