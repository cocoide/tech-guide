package rdb

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
)

func (r *Repository) CountFavoritesByArticleID(articleID int) (int64, error) {
	var count int64
	if err := r.db.Model(&model.FavoriteArticle{}).
		Where("article_id = ?", articleID).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *Repository) DoFavoriteArticle(articleID, accountID int) error {
	favorite := model.FavoriteArticle{
		AccountID: accountID,
		ArticleID: articleID,
	}
	if err := r.db.Create(&favorite).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UnFavoriteArticle(articleID, accountID int) error {
	if err := r.db.
		Where("article_id = ? AND account_id = ?", articleID, accountID).
		Delete(model.FavoriteArticle{}).Error; err != nil {
		return err
	}
	return nil
}
