package handler

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/interface/handler/ctxutils"
	"github.com/labstack/echo"
	"strconv"
)

func (h *Handler) GetCategories(c echo.Context) error {
	categories, err := h.repo.GetCategoriesWithTopics()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, categories)
}

func (h *Handler) GetPopularTopics(c echo.Context) error {
	topics, err := h.repo.GetPopularTopics(5)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) GetFollowingTopics(c echo.Context) error {
	accountId := ctxutils.GetAccountID(c)
	topics, err := h.repo.GetFollowingTopics(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) DoFollowTopic(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := ctxutils.GetAccountID(c)
	if err := h.repo.DoFollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Do Follow Topic")
}
func (h *Handler) UnFollowTopic(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := ctxutils.GetAccountID(c)
	if err := h.repo.UnfollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Unfollow Topic")
}

func (h *Handler) GetTopics(c echo.Context) error {
	topics, err := h.repo.GetAllTopics()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) CreateTopics(c echo.Context) error {
	topics := []model.Topic{}
	if err := c.Bind(&topics); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.repo.CreateTopics(topics); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "topics created")
}

func (h *Handler) GetTopicData(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	topic, err := h.repo.GetTopicData(topicId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topic)
}

func (h *Handler) CheckTopicFollow(c echo.Context) error {
	accountId := ctxutils.GetAccountID(c)
	topicId, err := strconv.Atoi(c.Param("id"))
	souce, err := h.repo.IsFollowingTopic(accountId, topicId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, souce)
}
