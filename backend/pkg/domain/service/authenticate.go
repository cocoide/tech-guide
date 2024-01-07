package service

import "github.com/cocoide/tech-guide/pkg/domain/model"

type OAuthProfileResponse struct {
	DisplayName string `validate:"required"`
	AvatarURL   string `validate:"url"`
	Email       string `validate:"required,email"`
}

type OAuthService interface {
	GetGoogleOAuthProfile(authCode string) (*OAuthProfileResponse, error)
	GetGithubOAuthProfile(authCode string) (*OAuthProfileResponse, error)
	GenerateOAuthURL(authType model.OAuthType) (string, error)
}
