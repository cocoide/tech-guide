package model

import "time"

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"comment"`
	AccountID int       `json:"account_id"`
	Account   Account   `json:"account" gorm:"foreignKey:AccountID"`
	ArticleID int       `json:"article_id"`
	Article   Article   `json:"article" gorm:"foreignKey:ArticleID"`
	CreatedAt time.Time `json:"created_at"`
}
