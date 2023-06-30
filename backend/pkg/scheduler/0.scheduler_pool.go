package scheduler

import (
	"fmt"
	"time"
)

type SchedulerPool struct {
	schedulers map[string]Scheduler
}

func NewSchedulerPool() *SchedulerPool {
	return &SchedulerPool{
		schedulers: make(map[string]Scheduler),
	}
}

func (sp *SchedulerPool) AddScheduler(id string, interval time.Duration, job func()) {
	if _, exists := sp.schedulers[id]; exists {
		fmt.Printf("Scheduler with id %s already exists\n", id)
		return
	}

	sp.schedulers[id] = NewScheduler(interval, job)
}

func (sp *SchedulerPool) RemoveScheduler(id string) {
	if s, exists := sp.schedulers[id]; exists {
		s.Stop()
		delete(sp.schedulers, id)
	}
}

func (sp *SchedulerPool) StartScheduler(id string) {
	if s, exists := sp.schedulers[id]; exists {
		s.Start()
	}
}

func (sp *SchedulerPool) StopScheduler(id string) {
	if s, exists := sp.schedulers[id]; exists {
		s.Stop()
	}
}

func (sp *SchedulerPool) RescheduleScheduler(id string, newInterval time.Duration) {
	if s, exists := sp.schedulers[id]; exists {
		s.Reschedule(newInterval)
	}
}
