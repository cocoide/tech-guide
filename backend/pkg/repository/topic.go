package repository

import (
	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type TopicRepo interface {
	GetAllTopics() ([]model.Topic, error)
	GetTopicToArticleArrayByArticleID(articleID int) ([]model.ArticlesToTopics, error)
}

type topicRepo struct {
	db *gorm.DB
}

func NewTopicRepo(db *gorm.DB) TopicRepo {
	return &topicRepo{db: db}
}

func (tr *topicRepo) GetAllTopics() ([]model.Topic, error) {
	var topics []model.Topic
	result := tr.db.
		Find(&topics)
	if result.Error != nil {
		return nil, result.Error
	}
	return topics, nil
}

func (r *topicRepo) GetTopicToArticleArrayByArticleID(articleID int) ([]model.ArticlesToTopics, error) {
	var articlesToTopicsArray []model.ArticlesToTopics
	if err := r.db.
		Select("article_id, topic_id, weight").
		Where("article_id = ?", articleID).
		Order("weight DESC").
		Find(&articlesToTopicsArray).Error; err != nil {
		return nil, err
	}
	return articlesToTopicsArray, nil
}
