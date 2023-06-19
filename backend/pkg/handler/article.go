package handler

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/labstack/echo"
)

func (h *Handler) GetArticles(c echo.Context) error {
	articles, err := h.ar.GetLatestArticleByLimit(50)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
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
	image := ""
	if len(ogp.Image) > 700 {
		image = ""
	} else {
		image = ogp.Image
	}
	article := &model.Article{
		Title:        ogp.Title,
		ThumbnailURL: image,
		OriginalURL:  URL,
	}
	if err := h.ar.Create(article); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, article)
}
