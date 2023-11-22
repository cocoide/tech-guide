package scheduler

import (
	"os"
)

//go:generate mockgen -source=worker.go -destination=./gen/worker.go
type WorkerService interface {
	RegisterArticles() error
	UpdateTopics() error
	UpdateRating() error
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
func (w *workerService) RegisterArticles() error {
	_, err := w.client.
		WithPath("register").
		ExecuteRequest(POST)
	return err
}

func (w *workerService) UpdateTopics() error {
	_, err := w.client.
		WithPath("register").
		ExecuteRequest(PUT)
	return err
}

func (w workerService) UpdateRating() error {
	return nil
}
