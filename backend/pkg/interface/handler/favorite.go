package handler

import (
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/labstack/echo"
	"strconv"
)

func (h *Handler) DoFavoriteArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := ctxutils.GetAccountID(c)
	if err := h.repo.DoFavoriteArticle(accountId, articleId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) UnFavoriteArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := ctxutils.GetAccountID(c)
	if err := h.repo.UnFavoriteArticle(accountId, articleId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}
