package model

import "time"

type Topic struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Icon_URL string    `json:"icon_url"`
	Articles []Article `json:"articles" gorm:"many2many:topics_to_articles"`
}

type Topics []*Topic

type FollowTopic struct {
	AccountID int `json:"account_id"`
	TopicID   int `json:"topic_id"`
}

type TopicsToArticles struct {
	TopicID   int     `json:"topic_id"`
	Topic     Topic   `json:"topic" gorm:"foreignKey:TopicID"`
	ArticleID int     `json:"article_id"`
	Article   Article `json:"article" gorm:"foreignKey:ArticleID"`
	// 1 ~ 5 (3: default)
	Weight    int       `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}

type Category struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	IconURL string  `json:"icon_url"`
	Topics  []Topic `json:"topics" gorm:"many2many:topic_to_categories"`
}

type TopicToCategory struct {
	TopicID    int   `json:"topic_id"`
	Topic      Topic `json:"topic" gorm:"foreignKey:TopicID"`
	CategoryID int   `json:"category_id"`
}
