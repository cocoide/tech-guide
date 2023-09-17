package rdb

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
)

func (r *Repository) GetCollectionByID(id int) (*model.Collection, error) {
	var collection model.Collection
	err := r.db.
		Preload("Articles").
		Where("id = ?", id).
		First(&collection).Error
	if err != nil {
		return nil, err
	}
	return &collection, nil
}
func (r *Repository) CreateCollection(collection *model.Collection) error {
	return r.db.Create(collection).Error
}
func (r *Repository) CreateCollectionWithBookmark(collection *model.Collection, articleId int) error {
	tx := r.db.Begin()
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
func (r *Repository) CreateBookmark(bookmark *model.Bookmark) error {
	return r.db.Create(bookmark).Error
}

func (r *Repository) GetCollectionsByAccountID(accountId int) ([]*model.Collection, error) {
	var collections []*model.Collection
	err := r.db.
		Where("account_id =?", accountId).
		Preload("Articles").
		Find(&collections).Error
	if err != nil {
		return nil, err
	}
	return collections, nil
}

func (r *Repository) GetCollectionAuthorID(collectionId int) (int, error) {
	var collection model.Collection
	err := r.db.
		Where("id =?", collectionId).
		Select("account_id").
		First(&collection).Error
	if err != nil {
		return 0, err
	}
	return collection.AccountID, nil
}
