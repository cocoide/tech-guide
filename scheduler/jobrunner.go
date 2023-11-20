package scheduler

import (
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

func (j *JobRunner) StartJobs() {
	params := &ListJobSettingsParams{
		Status: JobRunning.Ptr(),
	}
	settings, err := j.repo.ListJobSettings(params)
	if err != nil {
		panic(err)
	}
	for _, s := range settings {
		interval := utils.ParseIntSeconds(s.IntervalSeconds)
		worker := func() {
			if err := j.worker.ExecuteWorker(s.WorkerID); err != nil {
				log.Printf("Failed to execute worker(%d): %v", s.WorkerID, err)
				if _, err := j.repo.UpdateJobStatus(s.WorkerID, JobFailed); err != nil {
					panic(err)
				}
				j.pool.RemoveScheduler(s.WorkerID)
				return
			}
		}
		j.pool.AddScheduler(s.WorkerID, interval, worker)
		j.pool.StartScheduler(s.WorkerID, s.ScheduledAt)
		if err := j.repo.SetScheduledAt(s.WorkerID, s.ScheduledAt.Add(interval)); err != nil {
			if _, err := j.repo.UpdateJobStatus(s.WorkerID, JobFailed); err != nil {
				panic(err)
			}
			j.pool.RemoveScheduler(s.WorkerID)
		}
	}
	j.pool.AddScheduler(CheckJobs, time.Hour*12, j.checkJobs)
	j.pool.StartScheduler(CheckJobs, utils.NowInJST())
}

func (j *JobRunner) checkJobs() {
	settings, err := j.repo.ListJobSettings(
		&ListJobSettingsParams{
			Status: JobStopped.Ptr(),
		})
	if err != nil {
		panic(err)
	}
	for _, s := range settings {
		j.pool.RemoveScheduler(s.WorkerID)
	}
}
