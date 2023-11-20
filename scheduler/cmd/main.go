package main

import (
	"github.com/cocoide/tech-guide/scheduler"
	"github.com/cocoide/tech-guide/scheduler/utils"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	utils.LoadConfigFromEnv()
	repo := scheduler.NewRepository()
	pool := scheduler.NewSchedulerPool()
	worker := scheduler.NewWorkerService()
	runner := scheduler.NewJobRunner(repo, pool, worker)
	runner.StartJobs()
	e.Logger.Fatal(e.Start(":8000"))
}
