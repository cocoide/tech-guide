package handler

import "github.com/labstack/echo"

func (h *Handler) GetQiitaTrend(c echo.Context) error {
	feed, err := h.feed.GetQiitaTrendFeed(5, 30, "2023-05-04")
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, feed)
}

func (h *Handler) GetZennTrend(c echo.Context) error {
	articles, err := h.feed.GetZennTrendFeed()
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, articles)
}
