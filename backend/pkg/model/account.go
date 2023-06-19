package model

type Account struct {
	ID           int     `json:"id"`
	Email        string  `json:"email"`
	DisplayName  string  `json:"display_name"`
	AvatarURL    string  `json:"avatar_url"`
	Password     string  `json:"password"`
	FollowTopics []Topic `json:"follow_topics" gorm:"many2many:follow_topics;"`
}
