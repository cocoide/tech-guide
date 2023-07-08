package handler

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/labstack/echo"
)

func (h *Handler) GetContributions(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(403, err.Error())
	}
	contributions, err := h.vr.GetContributionsByAccountID(accountId)
	if err != nil {
		return c.JSON(403, err.Error())
	}
	return c.JSON(200, contributions)
}

func (h *Handler) GetReadArticle(c echo.Context) error {
	var articleIDs []int
	accountId := int(c.Get("account_id").(float64))
	cacheKey := fmt.Sprintf(key.RecentReads, accountId)
	sortedSet, err := h.rr.GetAllSortedSet(cacheKey)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	for i := len(sortedSet) - 1; i >= 0; i-- {
		id, err := strconv.Atoi(sortedSet[i].Member.(string))
		if err != nil {
			return c.JSON(400, err.Error())
		}
		articleIDs = append(articleIDs, id)
	}
	articles, err := h.ar.GetArticlesByIDs(articleIDs)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	sortedArticles := make([]model.Article, 0)
	articleMap := make(map[int]model.Article)
	for _, article := range articles {
		articleMap[article.ID] = article
	}
	for _, id := range articleIDs {
		article, ok := articleMap[id]
		if ok {
			sortedArticles = append(sortedArticles, article)
		}
	}
	return c.JSON(200, sortedArticles)
}

func (h *Handler) SetReadArticle(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	if strArticleID := c.QueryParam("article_id"); strArticleID != "" {
		// 既読記事は一週間保持する
		timestamp := time.Now().Unix()
		cacheKey := fmt.Sprintf(key.RecentReads, accountId)
		if err := h.rr.AddSortedSet(cacheKey, strArticleID, float64(timestamp), 24*7*time.Hour); err != nil {
			return c.JSON(400, err.Error())
		}
	} else {
		return c.JSON(400, "account_id is not set for search param")
	}
	if err := h.as.AddContribution(accountId, conf.ReadPoint); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}
