package repository

import (
	"time"

	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type TopicRepo interface {
	CreateTopics(topics []model.Topic) error
	GetAllTopics() ([]model.Topic, error)
	GetTopicToArticleArrayByArticleID(articleID int) ([]model.TopicsToArticles, error)
	GetRecentPopularArticleIDs(duration time.Duration, limit int) ([]int, error)
}

type topicRepo struct {
	db *gorm.DB
}

func NewTopicRepo(db *gorm.DB) TopicRepo {
	return &topicRepo{db: db}
}
func (tr *topicRepo) CreateTopics(topics []model.Topic) error {
	return tr.db.Create(topics).Error
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

func (r *topicRepo) GetTopicToArticleArrayByArticleID(articleID int) ([]model.TopicsToArticles, error) {
	var TopicsToArticlesArray []model.TopicsToArticles
	if err := r.db.
		Select("article_id, topic_id, weight").
		Where("article_id = ?", articleID).
		Order("weight DESC").
		Find(&TopicsToArticlesArray).Error; err != nil {
		return nil, err
	}
	return TopicsToArticlesArray, nil
}

func (r *topicRepo) GetRecentPopularArticleIDs(duration time.Duration, limit int) ([]int, error) {
	var articleIDs []int
	err := r.db.Model(&model.TopicsToArticles{}).
		Select("article_id").
		Where("created_at > ?", time.Now().Add(-duration)).
		Group("article_id").
		Order("COUNT(*) DESC").
		Limit(limit).
		Pluck("article_id", &articleIDs).
		Error
	if err != nil {
		return nil, err
	}
	return articleIDs, nil
}
