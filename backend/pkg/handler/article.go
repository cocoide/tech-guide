package handler

import (
	"context"

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
	if err := h.ar.Create(article); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, article)
}
