package scheduler

import (
	"fmt"
	"github.com/cocoide/tech-guide/scheduler/utils"
	"log"
	"time"
)

type JobRunner struct {
	repo   Repository
	pool   *SchedulerPool
	worker WorkerService
}

func NewJobRunner(repo Repository, pool *SchedulerPool, worker WorkerService) *JobRunner {
	return &JobRunner{repo: repo, pool: pool, worker: worker}
}

const (
	minScheduledInterval = time.Hour * 3
)

func (j *JobRunner) StartJobs() {
	j.pool.AddScheduler(ScheduleJobs, minScheduledInterval, j.scheduleLatestJobs)
	j.pool.StartScheduler(ScheduleJobs, time.Now())
}

func (j *JobRunner) scheduleLatestJobs() {
	scheduledAtBefore := time.Now().Add(minScheduledInterval)
	settings, err := j.repo.ListJobSettings(&ListJobSettingsParams{
		ScheduledAtBefore: &scheduledAtBefore,
		Status:            JobRunning.Ptr(),
	})
	if err != nil {
		log.Fatalf("Failed to get settings: %v", err)
	}
	for _, s := range settings {
		interval := utils.ParseIntSeconds(s.IntervalSeconds)
		worker := func() {
			if err := j.executeWorker(s.WorkerID, s.ScheduledAt, interval); err != nil {
				j.handleExecuteWorkerError(s.WorkerID, err)
			}
		}
		j.pool.AddScheduler(s.WorkerID, interval, worker)
		j.pool.StartScheduler(s.WorkerID, s.ScheduledAt)
	}
}

type WorkerID int

const (
	ScheduleJobs WorkerID = iota + 1
	RegisterArticles
	UpdateTopics
)

func (j *JobRunner) executeWorker(workerID WorkerID, scheduledAt time.Time, interval time.Duration) error {
	if scheduledAt.After(time.Now()) {
		log.Printf("worker(%d) is attempted to be executed before scheduled at(%v)", workerID, scheduledAt)
		return nil
	}
	var err error
	switch workerID {
	case RegisterArticles:
		err = j.worker.RegisterArticles()
	case UpdateTopics:
		err = j.worker.UpdateTopics()
	default:
		err = fmt.Errorf("unknown WorkerID: %v", workerID)
	}
	if err != nil {
		return fmt.Errorf("failed to execute worker: %v", err)
	}
	log.Printf("worker(%d) executed", workerID)
	nextScheduledAt := time.Now().Add(interval)
	if err := j.repo.SetScheduledAt(workerID, nextScheduledAt); err != nil {
		return err
	}
	log.Printf("worker(%d) scheduled at %v", workerID, scheduledAt)
	return nil
}

func (j *JobRunner) handleExecuteWorkerError(workerID WorkerID, err error) {
	if err != nil {
		log.Printf("Failed to execute worker(%d): %v", workerID, err)
		if _, err := j.repo.UpdateJobStatus(workerID, JobFailed); err != nil {
			log.Fatalf("Failed to update settings: %v", err)
		}
		j.pool.RemoveScheduler(workerID)
	}
}
