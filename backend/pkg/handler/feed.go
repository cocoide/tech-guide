package handler

import (
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) GetFeeds(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	strPageIndex := c.QueryParam("page")
	var pageIndex int
	if strPageIndex == "" {
		pageIndex = 1
	} else {
		pageIndex, _ = strconv.Atoi(strPageIndex)
	}
	topicIDs, err := h.tr.GetFollowingTopicIDs(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	sourceIDs, err := h.sr.GetFollowingSourceIDs(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	pageSize := 10
	articles, err := h.ar.BatchGetArticlesByTopicIDsAndSourceID(topicIDs, sourceIDs, pageIndex, pageSize)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}
