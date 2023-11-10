package repository

import (
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"time"
)

//go:generate mockgen -source=article.go -destination=../../../mock/repository/article.go
type ArticleRepo interface {
	CreateArticle(article *model.Article) (int, error)
	CheckArticleExistsByURL(url string) (bool, error)
	BatchCreate(articles []*model.Article) ([]int, error)
	CreateTopicToArticle(topicToArticles []model.TopicsToArticles) error
	UpdateTopicToArticle(topicToArticles []model.TopicsToArticles) error
	// summaryカラムはomit
	GetLatestArticleByLimitWithSourceData(pageIndex, pageSize int) ([]*model.Article, error)
	GetArticlesByIDs(articleIDs []int, preloadColumns []string) (model.Articles, error)
	GetArticleByID(articleID int) (*model.Article, error)
	GetArticleWithRelatedDataByID(articleID int) (*model.Article, error)
	GetTopicsByID(articleId int) ([]model.Topic, error)
	GetTagsAndWeightsByArticleID(articleID int) ([]model.TopicsToArticles, error)
	// Limit by 50 counts
	GetArticlesByTopicIDs(topicIDs []int, omitArticleId int) ([]model.TopicsToArticles, error)
	UpdateSummaryByID(id int, summary string) error
	BatchGetArticlesByTopicIDsAndSourceID(topicIDs, sourceIDs []int, pageIndex, pageSize int) (model.Articles, error)
	ListArticles(params *ListArticlesParams) (model.Articles, error)
	GetArticlesBySourceID(sourceID, pageIndex, pageSize int) ([]model.Article, error)
	GetArticlesByTopicID(topicID, pageIndex, pageSize int) ([]model.Article, error)
	FindArticlesByTitle(title string) ([]model.Article, error)
	GetArticlesExcludingIDs(excludeIDs []int) (model.Articles, error)
}

type ListArticlesParams struct {
	OrderBy        ArticleOrderType
	Preloads       []string
	Duration       Duration
	FeedOption     FeedOption
	PaginateOption PaginateOption
}

type Duration struct {
	Start time.Time
	End   time.Time
}

type ArticleOrderType int

const (
	Latest ArticleOrderType = iota + 1
	Older
	Trend
	Discuss
)

type FeedOption struct {
	AccountID      int
	IsFollowTopic  bool
	IsFollowDomain bool
}

type PaginateOption struct {
	PageIndex int
	PageSize  int
}
