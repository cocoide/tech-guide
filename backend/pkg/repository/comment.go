package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type CommentRepo interface {
	Create(quote *model.Comment) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentRepo{db: db}
}

func (r *commentRepo) Create(quote *model.Comment) error {
	return r.db.Create(&quote).Error
}
