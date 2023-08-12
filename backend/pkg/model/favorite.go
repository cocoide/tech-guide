package model

import "time"

type FavoriteArticle struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	ArticleID int       `json:"article_id"`
	CreatedAt time.Time `json:"created_at"`
}
