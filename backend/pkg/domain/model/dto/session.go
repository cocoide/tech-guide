package dto

type AccountSession struct {
	AccountID   int    `json:"account_id"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	Features    []int  `json:"features"`
}

type AccountFeaturesType int
type AccountFeatures AccountFeaturesType

const (
	OnboardingYet AccountFeaturesType = iota + 1
)
