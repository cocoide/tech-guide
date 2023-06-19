package model

import (
	"errors"
	"time"
)

type Bookmark struct {
	ArticleID    int       `json:"article_id"`
	CollectionID int       `json:"collection_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type Collection struct {
	ID          int    `json:"id"`
	AccountID   int    `json:"account_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// 0: Public, 2: Follower, 3: Private
	Visibility int       `json:"visibility"`
	Articles   []Article `json:"articles" gorm:"many2many:bookmarks"`
}

func (c *Collection) ValidateCollection() error {
	if len(c.Name) < 1 {
		return errors.New("collection name is too short")
	}
	if c.Visibility != 0 && c.Visibility != 1 && c.Visibility != 2 {
		return errors.New("visibility must be 0, 1, or 2")
	}
	if c.AccountID == 0 {
		return errors.New("account_id must be set")
	}
	return nil
}
