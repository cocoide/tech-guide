package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cocoide/tech-guide/pkg/domain/model"
	"github.com/cocoide/tech-guide/pkg/domain/model/object"
	"github.com/cocoide/tech-guide/pkg/domain/repository"
	"github.com/cocoide/tech-guide/pkg/domain/service"
	"github.com/cocoide/tech-guide/pkg/usecase"

	"github.com/cocoide/tech-guide/key"

	"github.com/cocoide/tech-guide/pkg/utils"
)

type TimelineWorker struct {
	repo        repository.Repository
	cache       repository.CacheRepo
	activity    *usecase.ActivityUsecase
	feed        service.TechFeedService
	ogp         service.OGPService
	personalize *usecase.PersonalizeUsecase
}

func NewTimelineWorker(repo repository.Repository, cache repository.CacheRepo, activity *usecase.ActivityUsecase, feed service.TechFeedService, ogp service.OGPService, personalize *usecase.PersonalizeUsecase) *TimelineWorker {
	return &TimelineWorker{repo: repo, cache: cache, feed: feed, activity: activity, ogp: ogp, personalize: personalize}
}

func (w *TimelineWorker) RegisterQiitaTendsWorker() {
	oneWeekAgo := time.Now().AddDate(0, 0, -7)
	oneWeekAgoString := oneWeekAgo.Format("2006-01-02")
	bookmarkThreshold := 200
	trends, err := w.feed.GetQiitaTrendFeed(5, bookmarkThreshold, oneWeekAgoString)
	if err != nil {
		log.Panicln(err)
	}
	var articles []*model.Article
	for _, v := range trends {
		exists, err := w.repo.CheckArticleExistsByURL(v.URL)
		if err != nil {
			log.Println(err)
		}
		if !exists {
			sourceID, err := w.repo.FindIDByDomain(object.QiitaDomain)
			if err != nil {
				log.Panicln(err)
			}
			ogpResponse, err := w.ogp.GetOGPByURL(v.URL)
			if err != nil {
				log.Panicln(err)
			}
			articles = append(articles,
				&model.Article{Title: v.Title, OriginalURL: v.URL, SourceID: sourceID, ThumbnailURL: ogpResponse.Thumbnail})
		}
	}
	createdIDs, err := w.repo.BatchCreate(articles)
	strArticleIDs, _, err := w.cache.Get(key.PopularArticleIDs)
	if err != nil {
		log.Panicln(err)
	}
	articleIDs, err := utils.Deserialize[[]int](strArticleIDs)
	if err != nil {
		log.Panicln(err)
	}
	existingIDSet := make(map[int]struct{})
	for _, existingID := range articleIDs {
		existingIDSet[existingID] = struct{}{}
	}
	var uniqueIDs []int
	for _, createdID := range createdIDs {
		if _, ok := existingIDSet[createdID]; !ok {
			uniqueIDs = append(uniqueIDs, createdID)
		}
	}
	strUniqueIDs, err := utils.Serialize(uniqueIDs)
	if err := w.cache.Set(key.PopularArticleIDs, strUniqueIDs, 24*time.Hour); err != nil {
		log.Println(err)
	}
}

func (w *TimelineWorker) CachePersonalizedArticlesWorker() {
	accountIDs, err := w.repo.GetAllAccountIDs()
	if err != nil {
		log.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(accountIDs))
	for _, accountId := range accountIDs {
		go func(id int) {
			articleIDs, err := w.personalize.GetRecommendArticleIDs(id)
			if err != nil {
				log.Println(err)
				return
			}
			if len(articleIDs) > 0 {
				strArticleIDs, err := utils.Serialize(articleIDs)
				if err != nil {
					log.Println(err)
					return
				}
				if err := w.cache.Set(fmt.Sprintf(key.PersonalizedArticleIDs, id), strArticleIDs, 24*time.Hour); err != nil {
					log.Println(err)
				}
			}
		}(accountId)
	}
	wg.Wait()
}

func (w *TimelineWorker) CacheTredingArticlesWorker() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cr := w.cache.WithCtx(ctx)

	selectArticleIDs, err := w.repo.GetRecentPopularArticleIDs(30*24*time.Hour, 50)
	if err != nil {
		log.Println(err)
	}
	serializedIDs, err := utils.Serialize(selectArticleIDs)
	if err != nil {
		log.Println(err)
	}
	if err := cr.Set(key.PopularArticleIDs, serializedIDs, 24*time.Hour); err != nil {
		log.Println(err)
	}
}

func (w *TimelineWorker) ContributionWorker() {
	contributions, err := w.activity.GetContributionsFromCache()
	if err != nil {
		log.Println(err)
	}
	if err := w.repo.BatchCreateContributions(contributions); err != nil {
		log.Println(err)
	}
}
