package model

import "time"

type Source struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	IconURL   string    `json:"icon_url"`
	Domain    string    `json:"domain"`
	CreatedAt time.Time `json:"created_at"`
}

type FollowSource struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	SourceID  int       `json:"source_id"`
	CreatedAt time.Time `json:"created_at"`
}
