package handler

import "github.com/labstack/echo"

func (h *Handler) GetTweets(c echo.Context) error {
	url := c.QueryParam("url")
	h.tweet.GetTweetIDsByQuoteURL(url)
	return c.JSON(200, "ok")
}
