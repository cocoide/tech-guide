package scheduler

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/cocoide/tech-guide/key"
	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/service"
	"github.com/cocoide/tech-guide/pkg/util"
)

type TimelineWorker interface {
	CacheTredingArticlesWorker()
	CachePersonalizedArticlesWorker()
}

type timelineWorker struct {
	ar repo.AccountRepo
	cr repo.CacheRepo
	tr repo.TopicRepo
	ps service.PersonalizeService
}

func NewTimelineWorker(ar repo.AccountRepo, cr repo.CacheRepo, tr repo.TopicRepo, ps service.PersonalizeService) TimelineWorker {
	return &timelineWorker{ar: ar, cr: cr, tr: tr, ps: ps}
}

func (w *timelineWorker) CachePersonalizedArticlesWorker() {
	accountIDs, err := w.ar.GetAllAccountIDs()
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
