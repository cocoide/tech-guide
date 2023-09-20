package handler

import (
	"github.com/labstack/echo"
	"strconv"
)

func (h *Handler) GetGithubStar(c echo.Context) error {
	url := c.QueryParam("url")
	count, err := h.git.GetStarCountByURL(url)
	//log.Print(err)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, count)
}

func (h *Handler) GetTopGitRepo(c echo.Context) error {
	star, _ := strconv.Atoi(c.QueryParam("star"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	language := c.QueryParam("language")
	count, err := h.git.FindTopStarRepo(star, limit, language)
	//log.Print(err)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, count)
}
