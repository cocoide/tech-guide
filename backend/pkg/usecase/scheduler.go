package usecase

import (
	"fmt"
	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/dto"
	"github.com/cocoide/tech-guide/pkg/domain/model/object"
	repo "github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase/timeutils"
	"log"
)

type SchedulerUsecase struct {
	repo   schedulerUsecaseRepo
	feed   service.TechFeedService
	github service.GithubService
	ogp    service.OGPService
}
type schedulerUsecaseRepo interface {
	repo.SourceRepo
	repo.ArticleRepo
}

func NewSchedulerUsecase(repo schedulerUsecaseRepo, feed service.TechFeedService, github service.GithubService, ogp service.OGPService) *SchedulerUsecase {
	return &SchedulerUsecase{repo: repo, feed: feed, github: github, ogp: ogp}
}

func (u *SchedulerUsecase) RegisterArticles() error {
	articles := make(model.Articles, 0)
	qiitaFeed, err := u.feed.GetQiitaTrendFeed(5, 50, timeutils.OneDayAgo())
	if err != nil {
		return err
	}
	//zennFeed, err := u.feed.GetZennTrendFeed()
	if err != nil {
		return err
	}
	qiitaSourceID, err := u.repo.FindIDByDomain(object.QiitaDomain)
	if err != nil {
		return err
	}
	//zennSourceID, err := u.repo.FindIDByDomain(object.QiitaDomain)
	//if err != nil {
	//	return err
	//}
	articles = append(articles, qiitaFeed.Articles().WithSourceID(qiitaSourceID)...)
	//articles = append(articles, zennFeed.Articles().WithSourceID(zennSourceID)...)

	//for i, a := range articles {
	//	if len(a.ThumbnailURL) == 0 {
	//		ogp, err := u.ogp.GetOGPByURL(a.OriginalURL)
	//		if err != nil {
	//			log.Printf("failed to get ogp: %v", err)
	//		}
	//		if len(ogp.Thumbnail) <= object.ThumbnailMaxLength {
	//			articles[i].ThumbnailURL = ogp.Thumbnail
	//		}
	//	}
	//}
	//if _, err := u.repo.MultiCreateArticles(articles); err != nil {
	//	return err
	//}
	return nil
}

func (u *SchedulerUsecase) AssignRatings() error {
	params := &repo.ListArticlesParams{
		Preloads: []string{"Source"},
		RawQuery: &repo.RawQuery{WhereQuery: "rating_id IS NULL"},
	}
	articles, err := u.repo.ListArticles(params)
	if err != nil {
		return err
	}
	ratings, err := u.fetchArticleRatings(articles)
	if err != nil {
		log.Printf("error while fetching ratings: %v", err)
	}
	if len(ratings) > 0 {
		if _, err := u.repo.MultiCreateArticleRatings(ratings); err != nil {
			return fmt.Errorf("failed to create rating: %v", err)
		}
	}
	return nil
}

func (u *SchedulerUsecase) AssignTopics() {
}

func (u *SchedulerUsecase) ReassignTopics() {
}

func (u *SchedulerUsecase) UpdateActivity() {
}

func (u *SchedulerUsecase) fetchArticleRatings(articles model.Articles) (ratings model.ArticleRatings, err error) {
	for _, a := range articles {
		url := dto.OriginalURL(a.OriginalURL)
		var origin int
		switch a.Source.Domain {
		case object.QiitaDomain:
			origin, err = u.feed.GetQiitaStockCount(url)
			if err != nil {
				return nil, err
			}
		case object.GithubDomain:
			origin, err = u.github.GetStarCountByURL(url)
			if err != nil {
				return nil, err
			}
		}
		hatena, err := u.feed.GetHatenaStockCount(url)
		if err != nil {
			return nil, err
		}
		pocket, err := u.feed.GetPocketStockCount(url)
		if err != nil {
			return nil, err
		}
		rating := &model.ArticleRating{
			ArticleID:    a.ID,
			OriginStocks: origin,
			HatenaStocks: hatena,
			PocketStocks: pocket,
		}
		ratings = append(ratings, rating)
	}
	return ratings, nil
}
