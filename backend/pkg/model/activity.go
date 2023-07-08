package model

import "time"

type Achievement struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

type AccountsToAchievements struct {
	ID            int       `json:"id"`
	AccountID     int       `json:"account_id"`
	AchievementID int       `json:"achievement_id"`
	Points        int       `json:"points"`
	EarnedAt      time.Time `json:"earned_at"`
}

type Contribution struct {
	ID        int       `json:"id"`
	AccountID int       `json:"account_id"`
	Account   Account   `json:"account"`
	Points    int       `json:"points"`
	Date      time.Time `json:"date"`
}
