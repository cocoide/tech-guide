package integration

import (
	"fmt"
	"github.com/cocoide/tech-guide/conf"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase/tknutils"
	"github.com/google/go-github/v56/github"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
	"log"
	"os"
)

type oauthService struct {
}

func NewOAuthService() service.OAuthService {
	return &oauthService{}
}

var (
	oauthGoogleScopes = []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"}
)

func (s *oauthService) GenerateOAuthURL(oauthType model.OAuthType) (string, error) {
	switch oauthType {
	case model.Google:
		oauth2Conf := s.getGoogleOAuthConf()
		url := oauth2Conf.AuthCodeURL(tknutils.GenerateStrUUID(), oauth2.AccessTypeOffline)
		return url, nil
	case model.Github:
		oauth2Conf := s.getGithubOAuthConf()
		url := oauth2Conf.AuthCodeURL(tknutils.GenerateStrUUID())
		return url, nil
	default:
		return "", fmt.Errorf("Error selecting oauth method")
	}
}

func (s *oauthService) GetGoogleOAuthProfile(authCode string) (*service.OAuthProfileResponse, error) {
	ctx := context.Background()
	oauthConf := s.getGoogleOAuthConf()
	tok, err := oauthConf.Exchange(ctx, authCode)
	peopleService, err := people.NewService(ctx, option.WithTokenSource(oauthConf.TokenSource(ctx, tok)))
	if err != nil {
		return nil, err
	}
	person, err := peopleService.People.Get("people/me").PersonFields("names,emailAddresses,photos").Do()
	if err != nil {
		return nil, err
	}
	return &service.OAuthProfileResponse{
		DisplayName: person.Names[0].DisplayName,
		AvatarURL:   person.Photos[0].Url,
		Email:       person.EmailAddresses[0].Value,
	}, nil
}

func (s *oauthService) GetGithubOAuthProfile(authCode string) (*service.OAuthProfileResponse, error) {
	ctx := context.Background()
	oauthConf := s.getGithubOAuthConf()
	log.Printf("GithubConf: redirect: %s", oauthConf.RedirectURL)
	log.Print(authCode)
	tok, err := oauthConf.Exchange(ctx, authCode)
	if err != nil {
		log.Printf("Exchange Error: %v", err)
		return nil, err
	}
	client := github.NewClient(oauthConf.Client(ctx, tok))
	user, res, err := client.Users.Get(ctx, "")
	log.Printf("get user status: %s", res.Status)
	if err != nil {
		log.Printf("Get Error: %v", err)
		return nil, err
	}
	return &service.OAuthProfileResponse{
		DisplayName: *user.Name,
		AvatarURL:   *user.AvatarURL,
		Email:       *user.Email,
	}, nil
}

func (s *oauthService) getGoogleOAuthConf() *oauth2.Config {
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

func (s *oauthService) getGithubOAuthConf() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("BACKEND_URL") + "/oauth/callback?type=2",
		Scopes:       []string{"user:email"},
		Endpoint:     oauth2github.Endpoint,
	}
}
