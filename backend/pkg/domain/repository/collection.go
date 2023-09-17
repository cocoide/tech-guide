package repository

import "github.com/cocoide/tech-guide/pkg/domain/model"

//go:generate mockgen -source=collection.go -destination=../../../mock/repository/collection.go
type CollectionRepo interface {
	CreateCollection(collection *model.Collection) error
	CreateCollectionWithBookmark(collectino *model.Collection, articleId int) error
	CreateBookmark(bookmark *model.Bookmark) error
	GetCollectionAuthorID(collectionId int) (int, error)
	GetCollectionsByAccountID(accountId int) ([]*model.Collection, error)
	GetCollectionByID(id int) (*model.Collection, error)
}
