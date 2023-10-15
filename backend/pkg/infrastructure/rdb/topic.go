package rdb

import (
	"log"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"gorm.io/gorm"
)

func (r *Repository) GetCategoriesWithTopics() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.
		Preload("Topics").
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
func (r *Repository) GetCategories() ([]model.Category, error) {
	var categories []model.Category
	if err := r.db.
		Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *Repository) GetTopicsByCategoryIDs(categoryIds []int) ([]model.Topic, error) {
	var topics []model.Topic
	var topicToCategories []model.TopicToCategory
	if err := r.db.
		Preload("Topic").
		Where("category_id IN (?)", categoryIds).
		Find(&topicToCategories).Error; err != nil {
		return nil, err
	}
	topicMap := make(map[int]model.Topic) // topicID -> topic
	for _, category := range topicToCategories {
		if _, ok := topicMap[category.TopicID]; !ok {
			topicMap[category.TopicID] = category.Topic
		}
	}
	log.Print(topicMap)
	for _, topic := range topicMap {
		topics = append(topics, topic)
	}
	log.Print(topics)
	return topics, nil
}
func (r *Repository) GetPopularTopics(limit int) ([]model.Topic, error) {
	var popularTopics []model.Topic
	err := r.db.
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

func (r *Repository) GetFollowingTopicIDs(accountID int) ([]int, error) {
	var topicIDs []int
	err := r.db.Model(&model.FollowTopic{}).
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

func (r *Repository) DoFollowTopic(accountID, topicID int) error {
	followTopic := model.FollowTopic{
		AccountID: accountID,
		TopicID:   topicID,
	}
	return r.db.Create(&followTopic).Error
}

func (r *Repository) UnfollowTopic(accountID, topicID int) error {
	return r.db.Where("account_id = ? AND topic_id = ?", accountID, topicID).
		Delete(&model.FollowTopic{}).Error
}

func (r *Repository) GetFollowingTopics(accountId int) ([]model.Topic, error) {
	var topics []model.Topic

	err := r.db.
		Joins("JOIN follow_topics ON topics.id = follow_topics.topic_id").
		Where("follow_topics.account_id = ?", accountId).
		Find(&topics).Error

	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (r *Repository) CreateTopics(topics []model.Topic) error {
	return r.db.Create(topics).Error
}

func (r *Repository) GetAllTopics() ([]model.Topic, error) {
	var topics []model.Topic
	result := r.db.
		Find(&topics)
	if result.Error != nil {
		return nil, result.Error
	}
	return topics, nil
}

func (r *Repository) GetTopicsByIDs(IDs []int) ([]model.Topic, error) {
	var topics []model.Topic
	result := r.db.
		Where("id IN ?", IDs).Find(&topics)
	if result.Error != nil {
		return nil, result.Error
	}
	return topics, nil
}

func (r *Repository) UpdateTopicToArticle(topicToArticles []model.TopicsToArticles) error {
	return r.db.Updates(&topicToArticles).Error
}

func (r *Repository) GetTopicToArticleArrayByArticleID(articleID int) ([]model.TopicsToArticles, error) {
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
func (r *Repository) GetTopicToArticleArrayByArticleIDs(articleIDs []int) ([]model.TopicsToArticles, error) {
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

func (r *Repository) GetRecentPopularArticleIDs(duration time.Duration, limit int) ([]int, error) {
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

func (r *Repository) GetTopicsByCollectionID(collectionID int) ([]model.Topic, error) {
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

func (r *Repository) GetTopicData(topicID int) (*model.Topic, error) {
	var topic model.Topic
	if err := r.db.First(&topic, topicID).Error; err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *Repository) IsFollowingTopic(accountID, topicID int) (bool, error) {
	var count int64
	if err := r.db.Model(&model.FollowTopic{}).
		Where("account_id = ? AND topic_id = ?", accountID, topicID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
