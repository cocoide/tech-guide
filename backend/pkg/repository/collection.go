package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type CollectionRepo interface {
	CreateCollection(collection *model.Collection) error
	CreateCollectionWithBookmark(collectino *model.Collection, articleId int) error
	CreateBookmark(bookmark *model.Bookmark) error
	GetCollectionAuthorID(collectionId int) (int, error)
	GetCollectionsByAccountID(accountId int) ([]*model.Collection, error)
	GetCollectionByID(id int) (*model.Collection, error)
}

type collectionRepo struct {
	db *gorm.DB
}

func NewCollectionRepo(db *gorm.DB) CollectionRepo {
	return &collectionRepo{db: db}
}

func (cr *collectionRepo) GetCollectionByID(id int) (*model.Collection, error) {
	var collection model.Collection
	err := cr.db.
		Preload("Articles").
		Where("id = ?", id).
		First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}
func (cr *collectionRepo) CreateCollection(collection *model.Collection) error {
	return cr.db.Create(collection).Error
}
func (cr *collectionRepo) CreateCollectionWithBookmark(collection *model.Collection, articleId int) error {
	tx := cr.db.Begin()
	if err := tx.Create(collection).Error; err != nil {
		tx.Rollback()
		return err
	}
	bookmark := &model.Bookmark{
		ArticleID:    articleId,
		CollectionID: collection.ID,
	}
	if err := tx.Create(bookmark).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
func (cr *collectionRepo) CreateBookmark(bookmark *model.Bookmark) error {
	return cr.db.Create(bookmark).Error
}

func (cr *collectionRepo) GetCollectionsByAccountID(accountId int) ([]*model.Collection, error) {
	var collections []*model.Collection
	err := cr.db.
		Where("account_id =?", accountId).
		Preload("Articles").
		Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (cr *collectionRepo) GetCollectionAuthorID(collectionId int) (int, error) {
	var collection model.Collection
	err := cr.db.
		Where("id =?", collectionId).
		Select("account_id").
		First(&collection).Error
	if err != nil {
		return 0, err
	}
	return collection.AccountID, nil
}
