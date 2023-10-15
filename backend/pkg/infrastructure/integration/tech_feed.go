package integration

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/service"
)

type techFeedService struct {
	client *HttpClient
}

func NewTechFeedService() service.TechFeedService {
	client := NewHttpClient()
	return &techFeedService{client: client}
}

func (s *techFeedService) GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error) {
	var result []*dto.QiitaArticleAPI

	params := make(map[string]interface{})
	params["page"] = 1
	params["per_page"] = limit
	params["query"] = fmt.Sprintf("created:>%s+stocks:>%d", start, save)

	b, err := s.client.
		WithBaseURL("https://qiita.com/api/v2/items").
		WithRawParams(params).
		ExecuteRequest(GET)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(b, &result); err != nil {
		return nil, fmt.Errorf("Failted to unmarshal: %v", err)
	}
	return result, nil
}

func (s *techFeedService) GetZennTrendFeed() ([]model.Article, error) {
	b, err := s.client.
		WithBaseURL("https://zenn.dev/feed").
		ExecuteRequest(GET)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch :%v", err)
	}

	var feed dto.RSSFeed
	if err := xml.Unmarshal(b, &feed); err != nil {
		return nil, fmt.Errorf("Failted to unmarshal: %v", err)
	}
	artcles, err := feed.GetArticles()
	if err != nil {
		return nil, err
	}

	return artcles, err
}
