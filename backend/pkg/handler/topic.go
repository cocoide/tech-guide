package handler

import (
	"strconv"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/labstack/echo"
)

func (h *Handler) GetPopularTopics(c echo.Context) error {
	topics, err := h.tr.GetPopularTopics(5)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) GetFollowingTopics(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	topics, err := h.tr.GetFollowingTopics(accountId)
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
	accountId := int(c.Get("account_id").(float64))
	if err := h.tr.DoFollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Do Follow Topic")
}
func (h *Handler) UnFollowTopic(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := int(c.Get("account_id").(float64))
	if err := h.tr.UnfollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "Unfollow Topic")
}

func (h *Handler) GetTopics(c echo.Context) error {
	topics, err := h.tr.GetAllTopics()
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
	if err := h.tr.CreateTopics(topics); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "topics created")
}

func (h *Handler) GetTopicData(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	topic, err := h.tr.GetTopicData(topicId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topic)
}

func (h *Handler) CheckTopicFollow(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	topicId, err := strconv.Atoi(c.Param("id"))
	souce, err := h.tr.IsFollowingTopic(accountId, topicId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, souce)
}
