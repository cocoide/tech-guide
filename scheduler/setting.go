package scheduler

import "time"

type JobSetting struct {
	ID              int       `json:"id"`
	WorkerID        WorkerID  `json:"worker_id"`
	IntervalSeconds int       `json:"interval_seconds"`
	ScheduledAt     time.Time `json:"scheduled_at"`
	Status          JobStatus `json:"status"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type JobSettings []*JobSetting

type JobStatus int

func (j JobStatus) Ptr() *JobStatus {
	p := j
	return &p
}

const (
	JobStopped JobStatus = iota
	JobRunning
	JobFailed
)
