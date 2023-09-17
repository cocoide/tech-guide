package scheduler

import (
	"time"

	"github.com/cocoide/tech-guide/key"
)

func NewAsyncJobRunner(sp *SchedulerPool, tw TimelineWorker){
	go func() {
		sp.AddScheduler(key.TrendingArticlesWorker, 24*time.Hour, tw.CacheTredingArticlesWorker)
		sp.AddScheduler(key.PersonalizedArticlesWorker, 24*time.Hour, tw.CachePersonalizedArticlesWorker)
		sp.AddScheduler(key.QiitaTrendsWorker, 24*time.Hour, tw.RegisterQiitaTendsWorker)
		sp.AddScheduler(key.ContributioinWorker, 24*time.Hour, tw.ContributionWorker)
		sp.StartScheduler(key.TrendingArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.PersonalizedArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.QiitaTrendsWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(key.ContributioinWorker)
	}()
}