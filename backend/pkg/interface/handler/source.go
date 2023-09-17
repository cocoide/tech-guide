package handler

import (
	"github.com/labstack/echo"
	"strconv"
)

func (h *Handler) GetPopularSources(c echo.Context) error {
	topics, err := h.repo.GetPopularSources(5)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}

func (h *Handler) GetAllSources(c echo.Context) error {
	topics, err := h.repo.GetAllSources()
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
	if err := h.repo.UnfollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.repo.DoFollowSource(accountId, topicId); err != nil {
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
	if err := h.repo.UnfollowTopic(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.repo.UnFollowSource(accountId, topicId); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) GetFollowingSources(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	souces, err := h.repo.GetFollowingSources(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, souces)
}
func (h *Handler) GetSourceData(c echo.Context) error {
	sourceId, err := strconv.Atoi(c.Param("id"))
	souce, err := h.repo.GetSourceData(sourceId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, souce)
}
func (h *Handler) CheckSourceFollow(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	sourceId, err := strconv.Atoi(c.Param("id"))
	souce, err := h.repo.IsFollowingSource(accountId, sourceId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, souce)
}
