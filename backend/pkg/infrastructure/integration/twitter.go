package integration

import (
	"fmt"
	"log"
	"os"
)

type TwitterService struct {
	client *HttpClient
}

func NewTwitterService() *TwitterService {
	client := NewHttpClient()
	return &TwitterService{client: client}
}

func (s *TwitterService) GetTweetIDsByQuoteURL(url string) {
	b, err := s.client.
		WithHeader("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TWITTER_TOKEN"))).
		WithBaseURL("https://api.twitter.com/2/tweets/recent").
		//WithParam("max_results", 3).
		//WithParam("query", fmt.Sprintf("has:links url:%s lang:ja", url)).
		ExecuteRequest(GET)
	if err != nil {
		log.Print(err)
	}
	log.Println(string(b))
}
