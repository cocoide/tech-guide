package repository

import (
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model"
)

//go:generate mockgen -source=topic.go -destination=../../../mock/repository/topic.go
type TopicRepo interface {
	DoFollowTopic(accountID, topicID int) error
	UnfollowTopic(accountID, topicID int) error
	GetFollowingTopicIDs(accountID int) ([]int, error)
	GetFollowingTopics(accountId int) ([]model.Topic, error)
	CreateTopics(topics []model.Topic) error
	GetAllTopics() ([]model.Topic, error)
	GetTopicsByCategoryIDs(categoryIds []int) ([]model.Topic, error)
	GetTopicToArticleArrayByArticleID(articleID int) ([]model.TopicsToArticles, error)
	GetTopicsByIDs(IDs []int) ([]model.Topic, error)
	GetTopicToArticleArrayByArticleIDs(articleIDs []int) ([]model.TopicsToArticles, error)
	GetRecentPopularArticleIDs(duration time.Duration, limit int) ([]int, error)
	GetPopularTopics(limit int) ([]model.Topic, error)
	GetTopicsByCollectionID(collectionID int) ([]model.Topic, error)
	GetTopicData(topicID int) (*model.Topic, error)
	IsFollowingTopic(accountID, topicID int) (bool, error)
	GetCategoriesWithTopics() ([]model.Category, error)
	GetCategories() ([]model.Category, error)
}
