package scheduler

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

//go:generate mockgen -source=repository.go -destination=./gen/repository.go
type Repository interface {
	ListJobSettings(params *ListJobSettingsParams) (JobSettings, error)
	UpdateJobStatus(workerID WorkerID, status JobStatus) (*JobSetting, error)
	SetScheduledAt(workerID WorkerID, scheduledAt time.Time) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository() Repository {
	DSN := os.Getenv("DSN")
	db, err := gorm.Open(
		mysql.Open(DSN),
	)
	if err != nil {
		log.Fatalf("Failed to init db: %v", err)
	}
	return &repository{db: db}
}

type ListJobSettingsParams struct {
	Status            *JobStatus
	ScheduledAtBefore *time.Time
}

func (r *repository) SetScheduledAt(workerID WorkerID, scheduledAt time.Time) error {
	result := r.db.Model(&JobSetting{}).
		Where("worker_id = ?", workerID).
		Update("scheduled_at", scheduledAt)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *repository) ListJobSettings(params *ListJobSettingsParams) (JobSettings, error) {
	var setting JobSettings
	query := r.db.Model(&JobSetting{})

	if params.Status != nil {
		query = query.Where("status = ?", *params.Status)
	}

	if params.ScheduledAtBefore != nil {
		query = query.Where("scheduled_at < ?", *params.ScheduledAtBefore)
	}
	if err := query.Find(&setting).Error; err != nil {
		return nil, err
	}
	return setting, nil
}

func (r *repository) UpdateJobStatus(workerID WorkerID, status JobStatus) (*JobSetting, error) {
	var setting JobSetting
	if err := r.db.Model(&setting).
		Where("worker_id = ?", workerID).
		Update("status", status).Error; err != nil {
		return nil, err
	}
	if err := r.db.Where("worker_id = ?", workerID).First(&setting).Error; err != nil {
		return nil, err
	}
	return &setting, nil
}
