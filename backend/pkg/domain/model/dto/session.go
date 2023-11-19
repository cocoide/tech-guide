package dto

type LoginSession struct {
	AccountID   int    `json:"account_id" validate:"gte=1"`
	DisplayName string `json:"display_name"`
	AvatarURL   string `json:"avatar_url"`
	Features    []int  `json:"features"`
}

type SignupSession struct {
	AccountID       int            `json:"account_id"`
	DisplayName     string         `json:"display_name"`
	AvatarURL       string         `json:"avatar_url"`
	Email           string         `json:"email"`
	FollowTopicIDs  []string       `json:"follow_topic_ids"`
	OnboardingIndex OnboardingType `json:"onboarding_index"`
}

type OnboardingType int

const (
	RegisterProfile OnboardingType = iota + 1
	FollowTopics
)

type AccountFeaturesType int
type AccountFeatures AccountFeaturesType

const (
	OnboardingYet AccountFeaturesType = iota + 1
)
