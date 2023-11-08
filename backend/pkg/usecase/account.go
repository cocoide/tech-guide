package usecase

import (
	"errors"
	"fmt"
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/usecase/tknutils"
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
	CacheTemporaryAccount = "temporary_account"
	CacheRefreshToken     = "refresh_token"
)

var (
	accessTokenExpires  = 7 * 24 * time.Hour
	refreshTokenExpires = 30 * 7 * 24 * time.Hour
)

type AccountUsecase struct {
	repo       repository.AccountRepo
	cache      repository.CacheRepo
	oauth2Conf *oauth2.Config
}

func NewAccountUsecase(repo repository.AccountRepo, cache repository.CacheRepo) *AccountUsecase {
	b, err := conf.TokenData.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Failed to read oauth credentials")
	}
	oauth2Conf, err := google.ConfigFromJSON(b, oauthGoogleScopes...)
	return &AccountUsecase{repo: repo, cache: cache, oauth2Conf: oauth2Conf}
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
	log.Print(tokens)
	if err != nil {
		return nil, err
	}
	return &LoginResponse{
		Account: account,
		Tokens:  tokens,
	}, nil
}

// Cache temporay account data
func (u *AccountUsecase) TemporarySignUp(account *model.Account) (*GenerateTokensResponse, error) {
	isEmailUsed, err := u.repo.CheckExistByEmail(account.Email)
	if err != nil {
		return nil, err
	}
	if isEmailUsed {
		return nil, errors.New("email is already used")
	}
	// ※ Refactor later
	account, err = u.repo.CreateAccount(account)
	if err != nil {
		return nil, err
	}
	tokens, err := u.generateTokens(account.ID)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func (u *AccountUsecase) GenerateOAuthRedirectURL(oauth model.OAuthType) (string, error) {
	switch oauth {
	case model.Google:
		url := u.oauth2Conf.AuthCodeURL(tknutils.GenerateStrUUID(), oauth2.AccessTypeOffline)
		return url, nil
	default:
		return "", fmt.Errorf("Error selecting oauth method")
	}
}

type OAuthAuthenticateResponse struct {
	DisplayName string
	Email       string
}

func (u *AccountUsecase) AuthenticateWithGoogle(ctx context.Context, authCode string) (*OAuthAuthenticateResponse, error) {
	b, err := conf.TokenData.ReadFile("credentials.json")
	if err != nil {
		return nil, err
	}
	oauthConf, err := google.ConfigFromJSON(b, oauthGoogleScopes...)
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
