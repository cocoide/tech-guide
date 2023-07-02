package handler

import (
	"errors"
	"strconv"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func (h *Handler) GetCollectionData(c echo.Context) error {
	collectionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	collection, err := h.cr.GetCollectionByID(collectionId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, collection)
}

func (h *Handler) GetCollectionForBookmark(c echo.Context) error {
	accountId := int(c.Get("account_id").(float64))
	if accountId == 0 {
		return c.JSON(403, "unauthorized for viewing bookmark")
	}
	collections, err := h.cr.GetCollectionsByAccountID(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, collections)
}

func (h *Handler) GetCollections(c echo.Context) error {
	accountId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(400, err.Error())
	}
	collections, err := h.cr.GetCollectionsByAccountID(accountId)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, collections)
}
func (h *Handler) CreateCollection(c echo.Context) error {
	type REQ struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Visibility  int    `json:"visibility"`
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	accountId := int(c.Get("account_id").(float64))

	collection := model.Collection{
		AccountID:   accountId,
		Name:        req.Name,
		Description: req.Description,
		Visibility:  req.Visibility,
	}
	if err := collection.ValidateCollection(); err != nil {
		return c.JSON(400, err.Error())
	}

	if c.QueryParam("articleId") != "" {
		articleId, err := strconv.Atoi(c.QueryParam("articleId"))
		if err != nil {
			return c.JSON(400, err.Error())
		}
		if err := h.cr.CreateCollectionWithBookmark(&collection, articleId); err != nil {
			return c.JSON(400, err.Error())
		}
	} else {
		if err := h.cr.CreateCollection(&collection); err != nil {
			return c.JSON(400, err.Error())
		}
	}
	return c.JSON(200, collection)
}

func (h *Handler) DoBookmark(c echo.Context) error {
	type REQ struct {
		CollectionID int `json:"collection_id"`
		ArticleID    int `json:"article_id"`
	}
	req := new(REQ)
	if err := c.Bind(req); err != nil {
		return c.JSON(400, err.Error())
	}
	authorId, err := h.cr.GetCollectionAuthorID(req.CollectionID)
	if err != nil {
		return c.JSON(400, err.Error())
	}
	if authorId != int(c.Get("account_id").(float64)) {
		return c.JSON(403, "unauthorized for creating bookmark")
	}
	bookmark := model.Bookmark{
		ArticleID:    req.ArticleID,
		CollectionID: req.CollectionID,
	}
	if err := h.cr.CreateBookmark(&bookmark); err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return c.JSON(409, err.Error())
		}
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, bookmark)
}
