package model

import "time"

type Article struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	OriginalURL  string       `json:"original_url"`
	ThumbnailURL string       `json:"thumbnail_url"`
	CreatedAt    time.Time    `json:"created_at"`
	Topics       []Topic      `json:"topics" gorm:"many2many:articles_to_topics"`
	Collections  []Collection `json:"collections" gorm:"many2many:bookmarks"`
}

type ArticlesToTopics struct {
	ArticleID int `json:"article_id"`
	TopicID   int `json:"topic_id"`
	// 1 ~ 5 (3: default)
	Weight    int       `json:"weight"`
	CreatedAt time.Time `json:"created_at"`
}
