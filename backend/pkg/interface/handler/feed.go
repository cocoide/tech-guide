package handler

import (
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/labstack/echo"
)

func (h *Handler) GetFeedsWithPaginate(c echo.Context) error {
	accountId := ctxutils.GetAccountID(c)
	pageIndex, err := ctxutils.NewPaginateIndex(c)
	if err != nil {
		return c.JSON(500, err)
	}
	articles, err := h.article.GetFeedsWithCache(accountId, pageIndex)
	if err != nil {
		return c.JSON(500, err)
	}

	return c.JSON(200, articles)
}
