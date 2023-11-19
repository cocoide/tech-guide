package usecase

import (
	"fmt"
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase/parser"
	"github.com/cocoide/tech-guide/pkg/usecase/tknutils"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
	"log"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
)

var (
	oauthGoogleScopes = []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}
)

const (
	CacheRefreshToken = "refresh_token"
)

var (
	accessTokenExpires  = 7 * 24 * time.Hour
	refreshTokenExpires = 30 * 7 * 24 * time.Hour
)

type AccountUsecase struct {
	tx            repository.Transaction
	repo          repository.AccountRepo
	cache         repository.CacheRepo
	signupSession service.SessionService[dto.SignupSession]
	validate      *validator.Validate
}

func NewAccountUsecase(repo repository.AccountRepo, cache repository.CacheRepo, signupSession service.SessionService[dto.SignupSession]) *AccountUsecase {
	return &AccountUsecase{repo: repo, cache: cache, signupSession: signupSession, validate: validator.New()}
}

type CompleteOnboardingRequest struct {
	SessionID       uuid.UUID
	FollowTopicIDs  []int
	FollowSourceIDs []int
}

func (u *AccountUsecase) CompleteOnboarding(req *CompleteOnboardingRequest) (*GenerateTokensResponse, error) {
	session, err := u.signupSession.Get(req.SessionID)
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
	if err := u.signupSession.Del(req.SessionID); err != nil {
		return nil, err
	}
	return u.generateTokens(session.AccountID)
}

func (u *AccountUsecase) GetSignupSession(sessionID uuid.UUID) (*dto.SignupSession, error) {
	res, err := u.signupSession.Get(sessionID)
	return &res, err
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
	session, err := u.signupSession.Get(req.SessionID)
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
	if err := u.signupSession.Set(req.SessionID, session, signupSessionDuration); err != nil {
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
	signupSessionDuration = time.Hour * 24
)

type CacheTemporarySignUpRequest struct {
	DisplayName string `validate:"required"`
	AvatarURL   string `validate:"url"`
	Email       string `validate:"required,email"`
}

func (u *AccountUsecase) CacheTemporarySignUp(req *CacheTemporarySignUpRequest) (*uuid.UUID, error) {
	if err := u.validate.Struct(req); err != nil {
		return nil, err
	}
	session := dto.SignupSession{
		DisplayName:     req.DisplayName,
		AvatarURL:       req.AvatarURL,
		OnboardingIndex: dto.RegisterProfile,
		Email:           req.Email,
	}
	signupResp, err := u.signupSession.New(session, signupSessionDuration)
	if err != nil {
		return nil, err
	}
	return &signupResp, nil
}

func (u *AccountUsecase) GenerateOAuthRedirectURL(oauth model.OAuthType) (string, error) {
	switch oauth {
	case model.Google:
		oauth2Conf := newGoogleOAuth2Conf()
		url := oauth2Conf.AuthCodeURL(tknutils.GenerateStrUUID(), oauth2.AccessTypeOffline)
		return url, nil
	default:
		return "", fmt.Errorf("Error selecting oauth method")
	}
}

type OAuthAuthenticateResponse struct {
	DisplayName string
	AvatarURL   string
	Email       string
}

func (u *AccountUsecase) AuthenticateWithGoogle(ctx context.Context, authCode string) (*OAuthAuthenticateResponse, error) {
	oauthConf := newGoogleOAuth2Conf()
	tok, err := oauthConf.Exchange(ctx, authCode)
	peopleService, err := people.NewService(context.Background(), option.WithTokenSource(oauthConf.TokenSource(ctx, tok)))
	if err != nil {
		return nil, err
	}
	person, err := peopleService.People.Get("people/me").PersonFields("names,emailAddresses,photos").Do()
	if err != nil {
		return nil, err
	}
	return &OAuthAuthenticateResponse{
		DisplayName: person.Names[0].DisplayName,
		AvatarURL:   person.Photos[0].Url,
		Email:       person.EmailAddresses[0].Value,
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

func newGoogleOAuth2Conf() *oauth2.Config {
	b, err := conf.TokenData.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Failed to read oauth credentials")
	}
	oauthConf, err := google.ConfigFromJSON(b, oauthGoogleScopes...)
	if err != nil {
		log.Fatalf("Failed to read oauth credentials")
	}
	return oauthConf
}
