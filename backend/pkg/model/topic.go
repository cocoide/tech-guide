package model

type Topic struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Icon_URL string `json:"icon_url"`
	// CategoryID int       `json:"category_id"`
	// Category   Category  `json:"category"`
	Articles []Article `json:"articles" gorm:"many2many:articles_to_topics"`
}

type FollowTopic struct {
	AccountID int `json:"account_id"`
	TopicID   int `json:"topic_id"`
}

// type Category struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// }
