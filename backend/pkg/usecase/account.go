package usecase

import (
	"fmt"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase/parser"
	"github.com/cocoide/tech-guide/pkg/usecase/tknutils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
)

const (
	CacheRefreshToken = "refresh_token"
)

var (
	accessTokenExpires  = 7 * 24 * time.Hour
	refreshTokenExpires = 30 * 7 * 24 * time.Hour
)

type AccountUsecase struct {
	tx       repository.Transaction
	repo     repository.AccountRepo
	cache    repository.CacheRepo
	signup   service.SessionService[dto.SignupSession]
	oauth    service.OAuthService
	validate *validator.Validate
}

func NewAccountUsecase(repo repository.AccountRepo, cache repository.CacheRepo, signup service.SessionService[dto.SignupSession], oauth service.OAuthService) *AccountUsecase {
	return &AccountUsecase{repo: repo, cache: cache, signup: signup, oauth: oauth, validate: validator.New()}
}

type CompleteOnboardingRequest struct {
	SessionID       uuid.UUID
	FollowTopicIDs  []int
	FollowSourceIDs []int
}

func (u *AccountUsecase) CompleteOnboarding(req *CompleteOnboardingRequest) (*GenerateTokensResponse, error) {
	session, err := u.signup.Get(req.SessionID)
	if err != nil {
		return nil, err
	}
	if err := u.validate.Struct(session); err != nil {
		return nil, err
	}
	if err := u.repo.CreateFollowTopics(session.AccountID, req.FollowTopicIDs); err != nil {
		return nil, err
	}
	if err := u.repo.CreateFollowSources(session.AccountID, req.FollowTopicIDs); err != nil {
		return nil, err
	}
	if err := u.signup.Del(req.SessionID); err != nil {
		return nil, err
	}
	return u.generateTokens(session.AccountID)
}

func (u *AccountUsecase) GetSignupSession(sessionID uuid.UUID) (*dto.SignupSession, error) {
	res, err := u.signup.Get(sessionID)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

type RegisterAccountRequest struct {
	SessionID   uuid.UUID `validate:"required"`
	AvatarURl   string    `validate:"url,required"`
	DisplayName string    `validate:"required"`
}

func (u *AccountUsecase) RegisterAccount(req *RegisterAccountRequest) error {
	if err := u.validate.Struct(req); err != nil {
		return err
	}
	session, err := u.signup.Get(req.SessionID)
	if err != nil {
		return err
	}
	if err = u.validate.Var(session.Email, "email"); err != nil {
		return err
	}
	account, err := u.repo.CreateAccount(&model.Account{
		AvatarURL: req.AvatarURl, DisplayName: req.DisplayName, Email: session.Email,
	})
	if err != nil {
		return err
	}
	// Signup Sessionの更新
	session.AccountID = account.ID
	session.OnboardingIndex = dto.FollowTopics
	if err := u.signup.Set(req.SessionID, session, SignupDuration); err != nil {
		return err
	}
	return nil
}

func (u *AccountUsecase) GetLoginSession(accountID int) (*dto.LoginSession, error) {
	var session dto.LoginSession
	sessionKey := fmt.Sprintf("session.%d", accountID)
	strSession, exist, err := u.cache.Get(sessionKey)
	if err != nil {
		return nil, err
	}
	if !exist {
		account, err := u.repo.GetAccountProfile(accountID)
		if err != nil {
			return nil, err
		}
		session = dto.LoginSession{
			AccountID:   account.ID,
			DisplayName: account.DisplayName,
			AvatarURL:   account.AvatarURL,
		}
	}
	if len(strSession) > 0 {
		session, err = parser.Deserialize[dto.LoginSession](strSession)
	}
	if err != nil {
		return nil, err
	}
	return &session, nil
}

type LoginResponse struct {
	Account *model.Account
	Tokens  *GenerateTokensResponse
}

func (u *AccountUsecase) Login(email string) (*LoginResponse, error) {
	account, err := u.repo.GetByEmail(email)
	if err != nil || account == nil {
		return nil, fmt.Errorf("Error account not found")
	}
	tokens, err := u.generateTokens(account.ID)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		Account: account,
		Tokens:  tokens,
	}, nil
}

const (
	SignupDuration = time.Hour * 24
)

func (u *AccountUsecase) GenerateOAuthRedirectURL(oauth model.OAuthType) (string, error) {
	return u.oauth.GenerateOAuthURL(oauth)
}

type ProcessOAuthCallbackResponse struct {
	SessionID uuid.UUID
	IsSignup  bool
	Tokens    *GenerateTokensResponse
}

func (u *AccountUsecase) ProcessOAuthCallback(oauth model.OAuthType, authCode string) (*ProcessOAuthCallbackResponse, error) {
	var profile *service.OAuthProfileResponse
	var err error
	switch oauth {
	case model.Google:
		profile, err = u.oauth.GetGoogleOAuthProfile(authCode)
	case model.Github:
		profile, err = u.oauth.GetGithubOAuthProfile(authCode)
	default:
		return nil, fmt.Errorf("oauth type error")
	}
	if err != nil {
		return nil, err
	}
	if err := u.validate.Struct(profile); err != nil {
		return nil, err
	}
	account, err := u.repo.GetByEmail(profile.Email)
	if err != nil {
		return nil, err
	}
	if account.NotFound() {
		session := dto.SignupSession{
			DisplayName:     profile.DisplayName,
			AvatarURL:       profile.AvatarURL,
			OnboardingIndex: dto.RegisterProfile,
			Email:           profile.Email,
		}
		sessionID, err := u.signup.New(session, SignupDuration)
		if err != nil {
			return nil, err
		}
		return &ProcessOAuthCallbackResponse{
			SessionID: sessionID,
			IsSignup:  true,
		}, nil
	}
	tokens, err := u.generateTokens(account.ID)
	if err != nil {
		return nil, err
	}
	return &ProcessOAuthCallbackResponse{
		Tokens: tokens,
	}, nil
}

type GenerateTokensResponse struct {
	AccessToken  string
	RefreshToken string
}

// Generate new AccessToken
func (u *AccountUsecase) RefreshToken(refreshToken string) (string, error) {
	claims, err := tknutils.ParseJwt(refreshToken)
	if err != nil {
		return "", fmt.Errorf("Unauthenticated")
	}
	accountID := claims.AccountID
	cacheRefreshTokenKey := fmt.Sprintf("%s.%d", CacheRefreshToken, accountID)
	token, exists, err := u.cache.Get(cacheRefreshTokenKey)
	if err != nil || !exists {
		return "", fmt.Errorf("Unauthenticated")
	}
	if token != refreshToken {
		return "", fmt.Errorf("Unauthenticated")
	}
	accessToken, err := tknutils.GenerateJwt(accountID, time.Now().Add(accessTokenExpires))
	return accessToken, nil
}

func (u *AccountUsecase) generateTokens(accountID int) (*GenerateTokensResponse, error) {
	accessToken, err := tknutils.GenerateJwt(accountID, time.Now().Add(accessTokenExpires))
	if err != nil {
		return nil, err
	}
	refreshToken, err := tknutils.GenerateJwt(accountID, time.Now().Add(refreshTokenExpires))
	cacheRefreshTokenKey := fmt.Sprintf("%s.%d", CacheRefreshToken, accountID)
	if err := u.cache.Set(cacheRefreshTokenKey, refreshToken, refreshTokenExpires); err != nil {
		return nil, fmt.Errorf("Failed to cache refresh token: %v", err)
	}
	return &GenerateTokensResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
