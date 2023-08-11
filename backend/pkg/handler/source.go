package handler

import (
	"strconv"

	"github.com/labstack/echo"
)

func (h *Handler) GetPopularSources(c echo.Context) error {
	topics, err := h.sr.GetPopularSources(5)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) GetAllSources(c echo.Context) error {
	topics, err := h.sr.GetAllSources()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) DoFollowSource(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := int(c.Get("account_id").(float64))
	if err := h.tr.UnfollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.sr.DoFollowSource(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) UnFollowSource(c echo.Context) error {
	topicId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := int(c.Get("account_id").(float64))
	if err := h.tr.UnfollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.sr.UnFollowSource(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) GetFollowingSources(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	souces, err := h.sr.GetFollowingSources(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, souces)
}
