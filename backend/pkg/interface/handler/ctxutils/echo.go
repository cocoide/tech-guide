package ctxutils

import (
	"errors"
	"fmt"
	"github.com/cocoide/tech-guide/key"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"log"
	"strconv"
)

var err error

func GetIDFromPath(ctx echo.Context) int {
	accountId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Print(err)
	}
	return accountId
}

func GetFromEchoCtx[T any](ctx echo.Context, key key.EchoCtxKey) (T, error) {
	var result T
	val, ok := ctx.Get(string(key)).(T)
	result = val
	if !ok {
		return result, errors.New("key not found or invalid type")
	}
	return val, nil
}

func GetAccountID(ctx echo.Context) int {
	accountId, err := GetFromEchoCtx[int](ctx, key.CtxAccountID)
	if err != nil {
		panic(err)
	}
	return accountId
}

func SetInEchoCtx(c echo.Context, key key.EchoCtxKey, value interface{}) {
	c.Set(string(key), value)
}

func NewPaginateIndex(c echo.Context) (int, error) {
	var result int
	pageIndex := c.QueryParam("page")
	if pageIndex == "" {
		return 1, nil
	}
	result, err := strconv.Atoi(pageIndex)
	if err != nil {
		return 0, fmt.Errorf("Parse error: %v", err)
	}
	return result, nil
}

func ParseParamInt(c echo.Context, name string) int {
	value, _ := strconv.Atoi(c.Param(name))
	return value
}

func GetSessionIDFromHeader(c echo.Context) (uuid.UUID, error) {
	var result uuid.UUID
	strSessionID := c.Request().Header.Get("X-Session-ID")
	result, err = uuid.Parse(strSessionID)
	if err != nil {
		return result, err
	}
	return result, nil
}
