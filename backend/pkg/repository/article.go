package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	Create(article *model.Article) (int, error)
	CreateTopicToArticle(topicToArticles []model.ArticlesToTopics) error
	GetLatestArticleByLimit(limit int) ([]*model.Article, error)
	GetTopicsByID(articleId int) ([]model.Topic, error)
	GetTagsAndWeightsByArticleID(articleID int) ([]model.ArticlesToTopics, error)
	// Limit by 50 counts
	GetArticlesByTopicIDs(topicIDs []int, excludeArticleId int) ([]model.ArticlesToTopics, error)
}

type articleRepo struct {
	db *gorm.DB
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}

func (r *articleRepo) CreateTopicToArticle(topicToArticles []model.ArticlesToTopics) error {
	return r.db.Create(&topicToArticles).Error
}

func (r *articleRepo) Create(article *model.Article) (int, error) {
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}
	return article.ID, nil
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
func (r *articleRepo) GetTopicsByID(articleId int) ([]model.Topic, error) {
	var article model.Article
	if err := r.db.
		Preload("Topics").
		Where("id = ?", articleId).
		First(&article).Error; err != nil {
		return nil, err
	}
	topics := article.Topics
	return topics, nil
}

func (r *articleRepo) GetTagsAndWeightsByArticleID(articleID int) ([]model.ArticlesToTopics, error) {
	var articlesToTopics []model.ArticlesToTopics
	if err := r.db.
		Preload("Topic").
		Where("article_id IN (?)", articleID).
		Find(&articlesToTopics).Error; err != nil {
		return nil, err
	}
	return articlesToTopics, nil
}

func (r *articleRepo) GetArticlesByTopicIDs(topicIDs []int, excludeArticleId int) ([]model.ArticlesToTopics, error) {
	var articlesToTopicsArray []model.ArticlesToTopics
	if err := r.db.
		Preload("Article").
		Where("topic_id IN (?)", topicIDs).
		Not("article_id = ?", excludeArticleId).
		Order("weight DESC").
		Limit(50).
		Find(&articlesToTopicsArray).Error; err != nil {
		return nil, err
	}
	return articlesToTopicsArray, nil
}
