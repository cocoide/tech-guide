package handler

import (
	"strconv"

	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/labstack/echo"
)

func (h *Handler) Login(c echo.Context) error {
	type REQ struct {
		Email    string
		Password string
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(403, err.Error())
	}
	if len(req.Password) < 5 {
		return c.JSON(403, "")
	}
	account, token, err := h.uu.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(403, "email or password is incorrect")
	}
	type RES struct {
		UID   int    `json:"uid"`
		Name  string `json:"name"`
		Image string `json:"image"`
		Token string `json:"token"`
	}
	res := &RES{
		UID:   account.ID,
		Name:  account.DisplayName,
		Image: account.AvatarURL,
		Token: token,
	}
	return c.JSON(200, res)
}

func (h *Handler) SignUp(c echo.Context) error {
	type REQ struct {
		Email    string
		Password string
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	if err := h.uu.SignUp(req.Email, req.Password); err != nil {
		return c.JSON(403, err.Error())
	}
	return c.JSON(200, "signup successful")
}

func (h *Handler) GenerateToken(c echo.Context) error {
	token, err := util.GenerateToken(2)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, token)
}
func (h *Handler) GetAccountProfile(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	account, err := h.ur.GetAccountProfile(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	account.Password = ""
	return c.JSON(200, account)
}
