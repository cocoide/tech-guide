package scheduler

import (
	"time"
)

type WorkerID int

const (
	personalizedArticlesWorker WorkerID = iota + 1
	trendingArticlesWorker
	qiitaTrendsWorker
	contributioinWorker
)

func NewAsyncJobRunner(sp *SchedulerPool, tw TimelineWorker) {
	go func() {
		sp.AddScheduler(trendingArticlesWorker, 24*time.Hour, tw.CacheTredingArticlesWorker)
		sp.AddScheduler(personalizedArticlesWorker, 24*time.Hour, tw.CachePersonalizedArticlesWorker)
		sp.AddScheduler(qiitaTrendsWorker, 24*time.Hour, tw.RegisterQiitaTendsWorker)
		sp.AddScheduler(contributioinWorker, 24*time.Hour, tw.ContributionWorker)
		sp.StartScheduler(trendingArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(personalizedArticlesWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(qiitaTrendsWorker)
		time.Sleep(60 * time.Second)
		sp.StartScheduler(contributioinWorker)
	}()
}
