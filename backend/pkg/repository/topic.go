package repository

import (
	"time"

	"github.com/cocoide/tech-guide/pkg/model"
	"gorm.io/gorm"
)

type TopicRepo interface {
	DoFollowTopic(accountID, topicID int) error
	UnfollowTopic(accountID, topicID int) error
	GetFollowingTopicIDs(accountID int) ([]int, error)
	GetFollowingTopics(accountId int) ([]model.Topic, error)
	CreateTopics(topics []model.Topic) error
	GetAllTopics() ([]model.Topic, error)
	GetTopicToArticleArrayByArticleID(articleID int) ([]model.TopicsToArticles, error)
	GetTopicToArticleArrayByArticleIDs(articleIDs []int) ([]model.TopicsToArticles, error)
	GetRecentPopularArticleIDs(duration time.Duration, limit int) ([]int, error)
	GetPopularTopics(limit int) ([]model.Topic, error)
	GetTopicsByCollectionID(collectionID int) ([]model.Topic, error)
	GetTopicData(topicID int) (*model.Topic, error)
	IsFollowingTopic(accountID, topicID int) (bool, error)
}

type topicRepo struct {
	db *gorm.DB
}

func NewTopicRepo(db *gorm.DB) TopicRepo {
	return &topicRepo{db: db}
}

func (tr *topicRepo) GetPopularTopics(limit int) ([]model.Topic, error) {
	var popularTopics []model.Topic
	err := tr.db.
		Table("follow_topics").
		Select("topics.*, COUNT(follow_topics.account_id) as follow_count").
		Joins("INNER JOIN topics ON topics.id = follow_topics.topic_id").
		Group("topics.id").
		Order("follow_count DESC").
		Limit(limit).
		Find(&popularTopics).
		Error
	if err != nil {
		return nil, err
	}
	return popularTopics, nil
}

func (tr *topicRepo) GetFollowingTopicIDs(accountID int) ([]int, error) {
	var topicIDs []int
	err := tr.db.Model(&model.FollowTopic{}).
		Where("account_id = ?", accountID).
		Pluck("topic_id", &topicIDs).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return topicIDs, nil
}

func (tr *topicRepo) DoFollowTopic(accountID, topicID int) error {
	followTopic := model.FollowTopic{
		AccountID: accountID,
		TopicID:   topicID,
	}
	return tr.db.Create(&followTopic).Error
}

func (tr *topicRepo) UnfollowTopic(accountID, topicID int) error {
	return tr.db.Where("account_id = ? AND topic_id = ?", accountID, topicID).
		Delete(&model.FollowTopic{}).Error
}

func (tr *topicRepo) GetFollowingTopics(accountId int) ([]model.Topic, error) {
	var topics []model.Topic

	err := tr.db.
		Joins("JOIN follow_topics ON topics.id = follow_topics.topic_id").
		Where("follow_topics.account_id = ?", accountId).
		Find(&topics).Error

	if err != nil {
		return nil, err
	}

	return topics, nil
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
func (r *topicRepo) GetTopicToArticleArrayByArticleIDs(articleIDs []int) ([]model.TopicsToArticles, error) {
	var TopicsToArticlesArray []model.TopicsToArticles
	if err := r.db.
		Select("article_id, topic_id, weight").
		Where("article_id IN (?)", articleIDs).
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

func (r *topicRepo) GetTopicsByCollectionID(collectionID int) ([]model.Topic, error) {
	var collection model.Collection
	err := r.db.Preload("Articles.Topics").First(&collection, collectionID).Error
	if err != nil {
		return nil, err
	}
	var topics []model.Topic
	for _, article := range collection.Articles {
		topics = append(topics, article.Topics...)
	}

	topicMap := make(map[int]model.Topic)
	for _, topic := range topics {
		topicMap[topic.ID] = topic
	}

	var uniqueTopics []model.Topic
	for _, topic := range topicMap {
		uniqueTopics = append(uniqueTopics, topic)
	}

	return uniqueTopics, nil
}

func (r *topicRepo) GetTopicData(topicID int) (*model.Topic, error) {
	var topic model.Topic
	if err := r.db.First(&topic, topicID).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *topicRepo) IsFollowingTopic(accountID, topicID int) (bool, error) {
	var count int64
	if err := r.db.Model(&model.FollowTopic{}).
		Where("account_id = ? AND topic_id = ?", accountID, topicID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
