package handler

import (
	"strconv"
	"time"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/util"
	"github.com/golang-jwt/jwt/v5"
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
		UID:          account.ID,
		Name:         account.DisplayName,
		Image:        account.AvatarURL,
		Token:        token,
		TokenExpires: time.Now().Add(1 * time.Hour).Unix(),
	}
	return c.JSON(200, res)
}

type LoginRES struct {
	UID          int    `json:"uid"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Token        string `json:"token"`
	TokenExpires int64  `json:"token_expires"`
}

func (h *Handler) RefreshToken(c echo.Context) error {
	type body struct {
		RefreshToken string `json:"refresh_token"`
	}
	req := new(body)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	token, err := util.ParseToken(req.RefreshToken)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.JSON(400, "failed to parse claims")
	}
	accountId := int(claims["account_id"].(float64))
	newToken, err := util.GenerateToken(accountId)
	if err != nil {
		return c.JSON(403, err.Error())
	}
	type response struct {
		Token        string `json:"token"`
		TokenExpires int64
	}
	return c.JSON(200, &response{
		Token:        newToken,
		TokenExpires: time.Now().Add(1 * time.Hour).Unix(),
	})
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
	if err := h.cr.CreateCollection(&model.Collection{
		AccountID:  account.ID,
		Name:       "後で読む",
		Visibility: 0,
	}); err != nil {
		return c.JSON(403, err.Error())
	}
	token, err := util.GenerateToken(account.ID)
	if err != nil {
		return c.JSON(403, err.Error())
	}
	res := &LoginRES{
		UID:          account.ID,
		Name:         account.DisplayName,
		Image:        account.AvatarURL,
		Token:        token,
		TokenExpires: time.Now().Add(1 * time.Hour).Unix(),
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
