package integration

import (
	"fmt"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/utils"
)

func (s *techFeedService) GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error) {
	url := fmt.Sprintf("https://qiita.com/api/v2/items?page=1&per_page=%d&query=created:>%s+stocks:>%d", limit, start, save)
	res, err := utils.FetchJSON[[]*dto.QiitaArticleAPI](url, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *techFeedService) GetZennTrendFeed() ([]model.Article, error) {
	url := "https://zenn.dev/feed"
	res, err := utils.FetchXML[dto.RSSFeed](url, 5*time.Second)
	if err != nil {
		return nil, err
	}
	artcles, err := res.GetArticles()
	if err != nil {
		return nil, err
	}
	return artcles, err
}
