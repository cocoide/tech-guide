package scheduler

import (
	"fmt"
	"github.com/cocoide/tech-guide/scheduler/utils"
	"log"
	"time"
)

type SchedulerPool struct {
	schedulers map[WorkerID]Scheduler
}

func NewSchedulerPool() *SchedulerPool {
	return &SchedulerPool{
		schedulers: make(map[WorkerID]Scheduler),
	}
}

func (sp *SchedulerPool) AddScheduler(id WorkerID, interval time.Duration, job func()) {
	if _, exists := sp.schedulers[id]; exists {
		fmt.Printf("Scheduler with id %s already exists\n", id)
		return
	}

	sp.schedulers[id] = newScheduler(interval, job)
}

func (sp *SchedulerPool) RemoveScheduler(id WorkerID) {
	if s, exists := sp.schedulers[id]; exists {
		s.Stop()
		delete(sp.schedulers, id)
		log.Printf("Worker(%d) is deleted", id)
	}
}

func (sp *SchedulerPool) StartScheduler(id WorkerID, scheduledAt time.Time) {
	if s, exists := sp.schedulers[id]; exists {
		delay := sp.calculateDelayTime(scheduledAt)
		s.Start(delay)
		log.Printf("Worker(%d) is started", id)
	}
}

func (sp *SchedulerPool) StopScheduler(id WorkerID) {
	if s, exists := sp.schedulers[id]; exists {
		s.Stop()
		log.Printf("Worker(%d) is stopped", id)
	}
}

func (sp *SchedulerPool) RescheduleScheduler(id WorkerID, interval time.Duration, scheduledAt time.Time) {
	if s, exists := sp.schedulers[id]; exists {
		delay := sp.calculateDelayTime(scheduledAt)
		s.Reschedule(interval, delay)
	}
}

func (sp *SchedulerPool) calculateDelayTime(scheduledAt time.Time) time.Duration {
	now := utils.NowInJST()
	if scheduledAt.After(now) {
		return scheduledAt.Sub(now)
	}
	// 過去の場合は即時実行
	return 0
}
