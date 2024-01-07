package handler

import "github.com/labstack/echo"

func (h *Handler) RegisterArticles(c echo.Context) error {
	if err := h.scheduler.RegisterArticles(); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) UpdateArticleTopics(c echo.Context) error {
	h.scheduler.AssignTopics()
	h.scheduler.ReassignTopics()
	return c.JSON(200, nil)
}

func (h *Handler) AssignArticleRating(c echo.Context) error {
	if err := h.scheduler.AssignRatings(); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, nil)
}

func (h *Handler) UpdateActivity(c echo.Context) error {
	h.scheduler.AssignTopics()
	h.scheduler.ReassignTopics()
	return c.JSON(200, nil)
}
