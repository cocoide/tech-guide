package ctxutils

import (
	"errors"
	"github.com/cocoide/tech-guide/key"
	"github.com/labstack/echo"
)

func GetFromEchoCtx[T any](ctx echo.Context, key key.EchoCtxKey) (T, error) {
	var result T
	val, ok := ctx.Get(string(key)).(T)
	if !ok {
		return result, errors.New("key not found or invalid type")
	}
	result = val
	return val, nil
}

func GetAccountID(ctx echo.Context) int {
	accountId, err := GetFromEchoCtx[int](ctx, key.CtxAccountID)
	if err != nil {
		panic(err)
	}
	return accountId
}
