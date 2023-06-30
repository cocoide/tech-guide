package scheduler

import (
	"context"
	"log"
	"time"

	repo "github.com/cocoide/tech-guide/pkg/repository"
	"github.com/cocoide/tech-guide/pkg/util"
)

type TimelineWorker interface {
	CacheTredingArticlesWorker()
	CachePersonalizedArticlesWorker()
}

type timelineWorker struct {
	cr repo.CacheRepo
	tr repo.TopicRepo
}

func NewTimelineWorker(cr repo.CacheRepo, tr repo.TopicRepo) TimelineWorker {
	return &timelineWorker{cr: cr, tr: tr}
}

func (w *timelineWorker) CachePersonalizedArticlesWorker() {
}

func (w *timelineWorker) CacheTredingArticlesWorker() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cr := w.cr.WithCtx(ctx)

	selectArticleIDs, err := w.tr.GetRecentPopularArticleIDs(30*24*time.Hour, 50)
	if err != nil {
		log.Print(err)
	}
	serializedIDs, err := util.Serialize(selectArticleIDs)
	if err != nil {
		log.Print(err)
	}
	if err := cr.Set("popular_articles", serializedIDs, 24*time.Hour); err != nil {
		log.Print(err)
	}
}
