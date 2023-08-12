package gateway

import (
	"fmt"
	"time"

	"github.com/cocoide/tech-guide/pkg/model"
	"github.com/cocoide/tech-guide/pkg/model/dto"
	"github.com/cocoide/tech-guide/pkg/util"
)

type TechFeedGateway interface {
	GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error)
	GetZennTrendFeed() ([]model.Article, error)
}
type techFeedGateway struct{}

func NewTechFeedGateway() TechFeedGateway {
	return &techFeedGateway{}
}

func (g *techFeedGateway) GetQiitaTrendFeed(limit, save int, start string) ([]*dto.QiitaArticleAPI, error) {
	url := fmt.Sprintf("https://qiita.com/api/v2/items?page=1&per_page=%d&query=created:>%s+stocks:>%d", limit, start, save)
	res, err := util.FetchJSON[[]*dto.QiitaArticleAPI](url, nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (g *techFeedGateway) GetZennTrendFeed() ([]model.Article, error) {
	url := "https://zenn.dev/feed"
	res, err := util.FetchXML[dto.RSSFeed](url, 5*time.Second)
	if err != nil {
		return nil, err
	}
	artcles, err := res.GetArticles()
	if err != nil {
		return nil, err
	}
	return artcles, err
}
