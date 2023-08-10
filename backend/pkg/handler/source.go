package handler

import "github.com/labstack/echo"

func (h *Handler) GetPopularSources(c echo.Context) error {
	topics, err := h.sr.GetPopularSources(5)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, topics)
}
