package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type CollectionRepo interface {
	CreateCollection(collectino *model.Collection) error
	CreateBookmark(bookmark *model.Bookmark) error
	GetCollectionAuthorID(collectionId int) (int, error)
	GetCollectionsByAccountID(accountId int) ([]*model.Collection, error)
}

type collectionRepo struct {
	db *gorm.DB
}

func NewCollectionRepo(db *gorm.DB) CollectionRepo {
	return &collectionRepo{db: db}
}
func (cr *collectionRepo) DeleteBookmark() {}

func (cr *collectionRepo) CreateCollection(collection *model.Collection) error {
	return cr.db.Create(collection).Error
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
