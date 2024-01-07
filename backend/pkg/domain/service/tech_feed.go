package service

import (
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"time"
)

type TechFeedService interface {
	GetQiitaTrendFeed(limit, save int, fromTime time.Time) (dto.QiitaArticles, error)
	GetZennTrendFeed() (*dto.RSSFeed, error)
	GetHatenaStockCount(url dto.OriginalURL) (int, error)
	GetQiitaStockCount(url dto.OriginalURL) (int, error)
	GetPocketStockCount(url dto.OriginalURL) (int, error)
}
