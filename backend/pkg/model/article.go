package model

import "time"

type Article struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	OriginalURL  string       `json:"original_url"`
	ThumbnailURL string       `json:"thumbnail_url"`
	CreatedAt    time.Time    `json:"created_at"`
	Topics       []Topic      `json:"topics" gorm:"many2many:topics_to_articles"`
	Collections  []Collection `json:"collections" gorm:"many2many:bookmarks"`
}
