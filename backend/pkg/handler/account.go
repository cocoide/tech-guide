package handler

import (
	"strconv"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/labstack/echo"
)

func (h *Handler) Login(c echo.Context) error {
	type REQ struct {
		Email string
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(403, err.Error())
	}
	account, token, err := h.uu.Login(req.Email)
	if err != nil {
		c.JSON(403, "email or password is incorrect")
	}
	res := &LoginRES{
		UID:   account.ID,
		Name:  account.DisplayName,
		Image: account.AvatarURL,
		Token: token,
	}
	return c.JSON(200, res)
}

type LoginRES struct {
	UID   int    `json:"uid"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Token string `json:"token"`
}

func (h *Handler) SignUp(c echo.Context) error {
	type REQ struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Image    string `json:"image"`
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	account := &model.Account{
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: req.Name,
		AvatarURL:   req.Image,
	}
	account, err := h.uu.SignUp(account)
	if err != nil {
		return c.JSON(403, err.Error())
	}
	token, err := util.GenerateToken(account.ID)
	if err != nil {
		return c.JSON(403, err.Error())
	}
	res := &LoginRES{
		UID:   account.ID,
		Name:  account.DisplayName,
		Image: account.AvatarURL,
		Token: token,
	}
	return c.JSON(200, res)
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

func (h *Handler) CheckEmail(c echo.Context) error {
	email := c.QueryParam("email")
	if len(email) < 1 {
		return c.JSON(400, "email is not set for search query")
	}
	isExist, err := h.ur.CheckExistByEmail(email)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	if isExist {
		return c.JSON(400, true)
	}
	return c.JSON(200, false)
}
