package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	Create(article *model.Article) error
	GetLatestArticleByLimit(limit int) ([]*model.Article, error)
	GetArticleByCollectionID()
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}

func (r *articleRepo) Create(article *model.Article) error {
	return r.db.Create(article).Error
}
func (r *articleRepo) GetLatestArticleByLimit(limit int) ([]*model.Article, error) {
	var articles []*model.Article
	if err := r.db.
		Preload("Topics").
		Order("created_at desc").
		Limit(limit).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
func (r *articleRepo) GetArticleByCollectionID() {

}
