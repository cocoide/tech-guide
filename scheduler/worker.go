package scheduler

import (
	"fmt"
	"os"
)

type WorkerID int

const (
	CheckJobs WorkerID = iota + 1
	RegisterArticles
	UpdateTopics
)

//go:generate mockgen -source=worker.go -destination=./gen/worker.go
type WorkerService interface {
	ExecuteWorker(id WorkerID) error
}

type workerService struct {
	client *HttpClient
}

func NewWorkerService() WorkerService {
	client := NewHttpClient().
		WithBaseURL(os.Getenv("BACKEND_URL")).
		WithPath("scheduler").
		WithBearerToken(os.Getenv("SCHEDULER_API_KEY"))
	return &workerService{client: client}
}

func (w *workerService) ExecuteWorker(id WorkerID) error {
	switch id {
	case CheckJobs:
		return w.updateRating()
	case RegisterArticles:
		return w.registerArticles()
	case UpdateTopics:
		return w.updateTopics()
	default:
		return fmt.Errorf("unknown WorkerID: %v", id)
	}
}

func (w *workerService) registerArticles() error {
	_, err := w.client.
		WithPath("register").
		ExecuteRequest(POST)
	return err
}

func (w *workerService) updateTopics() error {
	_, err := w.client.
		WithPath("register").
		ExecuteRequest(PUT)
	return err
}

func (w workerService) updateRating() error {
	return nil
}
