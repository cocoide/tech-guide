package handler

import (
	"github.com/labstack/echo"
)

func getAccountIDInCtx(c echo.Context) int {
	return int(c.Get("account_id").(float64))
}
