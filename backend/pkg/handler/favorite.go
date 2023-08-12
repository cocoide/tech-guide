package handler

import (
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) DoFavoriteArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := int(c.Get("account_id").(float64))
	if err := h.fr.DoFavoriteArticle(accountId, articleId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) UnFavoriteArticle(c echo.Context) error {
	articleId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := int(c.Get("account_id").(float64))
	if err := h.fr.UnFavoriteArticle(accountId, articleId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}
