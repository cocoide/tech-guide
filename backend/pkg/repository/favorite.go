package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type FavoriteRepo interface {
	DoFavoriteArticle(articleID, accountID int) error
	UnFavoriteArticle(articleID, accountID int) error
	CountFavoritesByArticleID(articleID int) (int64, error)
}
type favoriteRepo struct {
	db *gorm.DB
}

func NewFavoriteRepo(db *gorm.DB) FavoriteRepo {
	return &favoriteRepo{db: db}
}

func (r *favoriteRepo) CountFavoritesByArticleID(articleID int) (int64, error) {
	var count int64
	if err := r.db.Model(&model.FavoriteArticle{}).
		Where("article_id = ?", articleID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *favoriteRepo) DoFavoriteArticle(articleID, accountID int) error {
	favorite := model.FavoriteArticle{
		AccountID: accountID,
		ArticleID: articleID,
	}
	if err := r.db.Create(&favorite).Error; err != nil {
		return err
	}
	return nil
}

func (r *favoriteRepo) UnFavoriteArticle(articleID, accountID int) error {
	if err := r.db.
		Where("article_id = ? AND account_id = ?", articleID, accountID).
		Delete(model.FavoriteArticle{}).Error; err != nil {
		return err
	}
	return nil
}
