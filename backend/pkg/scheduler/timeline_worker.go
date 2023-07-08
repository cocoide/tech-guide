package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cocoide/tech-guide/key"
	"github.com/cocoide/tech-guide/pkg/gateway"
	"github.com/cocoide/tech-guide/pkg/model"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/util"
)

type TimelineWorker interface {
	CacheTredingArticlesWorker()
	CachePersonalizedArticlesWorker()
	RegisterQiitaTendsWorker()
	ContributionWorker()
}

type timelineWorker struct {
	ur repo.AccountRepo
	ar repo.ArticleRepo
	cr repo.CacheRepo
	tr repo.TopicRepo
	vr repo.ActivityRepo
	vs service.ActivityService
	tg gateway.TechFeedGateway
	ps service.PersonalizeService
}

func NewTimelineWorker(ur repo.AccountRepo, ar repo.ArticleRepo, vr repo.ActivityRepo, vs service.ActivityService, cr repo.CacheRepo, tr repo.TopicRepo, ps service.PersonalizeService, tg gateway.TechFeedGateway) TimelineWorker {
	return &timelineWorker{ur: ur, ar: ar, cr: cr, vr: vr, vs: vs, tr: tr, ps: ps, tg: tg}
}

func (w *timelineWorker) RegisterQiitaTendsWorker() {
	oneWeekAgo := time.Now().AddDate(0, 0, -7)
	oneWeekAgoString := oneWeekAgo.Format("2006-01-02")
	bookmarkThreshold := 50
	trends, err := w.tg.GetQiitaTrendFeed(5, bookmarkThreshold, oneWeekAgoString)
	if err != nil {
		log.Panicln(err)
	}
	var articles []*model.Article
	for _, v := range trends {
		exists, err := w.ar.CheckArticleExistsByURL(v.URL)
		if err != nil {
			log.Println(err)
		}
		if !exists {
			articles = append(articles,
				&model.Article{Title: v.Title, OriginalURL: v.URL})
		}
	}
	createdIDs, err := w.ar.BatchCreate(articles)
	strArticleIDs, err := w.cr.Get(key.PopularArticleIDs)
	if err != nil {
		log.Panicln(err)
	}
	articleIDs, err := util.Deserialize[[]int](strArticleIDs)
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
	strUniqueIDs, err := util.Serialize(uniqueIDs)
	if err := w.cr.Set(key.PopularArticleIDs, strUniqueIDs, 24*time.Hour); err != nil {
		log.Println(err)
	}
}

func (w *timelineWorker) CachePersonalizedArticlesWorker() {
	accountIDs, err := w.ur.GetAllAccountIDs()
	if err != nil {
		log.Println(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(accountIDs))
	for _, accountId := range accountIDs {
		go func(id int) {
			articleIDs, err := w.ps.GetRecommendArticleIDs(id)
			if err != nil {
				log.Println(err)
				return
			}
			if len(articleIDs) > 0 {
				strArticleIDs, err := util.Serialize(articleIDs)
				if err != nil {
					log.Println(err)
					return
				}
				if err := w.cr.Set(fmt.Sprintf(key.PersonalizedArticleIDs, id), strArticleIDs, 24*time.Hour); err != nil {
					log.Println(err)
				}
			}
		}(accountId)
	}
	wg.Wait()
}

func (w *timelineWorker) CacheTredingArticlesWorker() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cr := w.cr.WithCtx(ctx)

	selectArticleIDs, err := w.tr.GetRecentPopularArticleIDs(30*24*time.Hour, 50)
	if err != nil {
		log.Println(err)
	}
	serializedIDs, err := util.Serialize(selectArticleIDs)
	if err != nil {
		log.Println(err)
	}
	if err := cr.Set(key.PopularArticleIDs, serializedIDs, 24*time.Hour); err != nil {
		log.Println(err)
	}
}

func (w *timelineWorker) ContributionWorker() {
	contributions, err := w.vs.GetContributionsFromCache()
	if err != nil {
		log.Println(err)
	}
	if err := w.vr.BatchCreateContributions(contributions); err != nil {
		log.Println(err)
	}
}
