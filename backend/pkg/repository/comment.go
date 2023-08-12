package repository

import (
	"errors"

	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	Create(comment *model.Comment) error
	GetComments(articleID int) ([]model.Comment, error)
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db: db}
}

func (r *commentRepo) Create(comment *model.Comment) error {
	return r.db.Create(&comment).Error
}
func (r *commentRepo) GetByArticleID(articleID int) error {
	return r.db.
		Preload("Account").
		Find(&model.Comment{ArticleID: articleID}).Error
}

func (r *commentRepo) GetComments(articleID int) ([]model.Comment, error) {
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
