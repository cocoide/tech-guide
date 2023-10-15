package model

import (
	"time"
)

type Article struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	OriginalURL  string        `json:"original_url"`
	ThumbnailURL string        `json:"thumbnail_url"`
	CreatedAt    time.Time     `json:"created_at"`
	Topics       []Topic       `json:"topics" gorm:"many2many:topics_to_articles"`
	Collections  []Collection  `json:"collections" gorm:"many2many:bookmarks"`
	SourceID     int           `json:"source_id"`
	Source       Source        `json:"source" gorm:"foreignKey:SourceID"`
	Summary      string        `json:"summary"`
	RatingID     int           `json:"rating_id"`
	Rating       ArticleRating `json:"rating" gorm:"foreignKey:RatingID"`
}

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	AccountID int       `json:"account_id"`
	Account   Account   `json:"account" gorm:"foreignKey:AccountID"`
	ArticleID int       `json:"article_id"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time `json:"created_at"`
}

type ArticleRating struct {
	ID           int       `json:"id"`
	OwnedStocks  int       `json:"owned_stocks"`
	OriginStocks int       `json:"origin_stocks"`
	PocketStocks int       `json:"pocket_stocks"`
	HatenaStocks int       `json:"hatena_stocks"`
	UpdatedAt    time.Time `json:"updated_at"`
	ArticleID    int       `json:"article_id"`
	Article      *Article  `json:"article" gorm:"foreignKey:ArticleID"`
}
