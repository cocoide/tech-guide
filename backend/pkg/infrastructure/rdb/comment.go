package rdb

import (
	"errors"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"gorm.io/gorm"
)

func (r *Repository) CreateComment(comment *model.Comment) error {
	return r.db.Create(&comment).Error
}
func (r *Repository) GetByArticleID(articleID int) error {
	return r.db.
		Preload("Account").
		Find(&model.Comment{ArticleID: articleID}).Error
}

func (r *Repository) GetComments(articleID int) ([]model.Comment, error) {
	var comments []model.Comment
	if err := r.db.
		Preload("Account").
		Where("article_id = ?", articleID).
		Order("created_at desc").
		Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // or return a custom error
		}
		return nil, err
	}
	return comments, nil
}
