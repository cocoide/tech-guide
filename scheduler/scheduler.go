package scheduler

import (
	"time"
)

type Scheduler interface {
	Start(delay time.Duration)
	Stop()
	IsRunning() bool
	Reschedule(newInterval time.Duration, delay time.Duration)
}

type scheduler struct {
	ticker    *time.Ticker
	done      chan bool
	job       func()
	startTime time.Time
	interval  time.Duration
	isRunning bool
}

func newScheduler(interval time.Duration, job func()) Scheduler {
	return &scheduler{
		interval:  interval,
		job:       job,
		isRunning: false,
	}
}

func (s *scheduler) Start(delay time.Duration) {
	if s.isRunning {
		return
	}

	s.done = make(chan bool)
	s.ticker = time.NewTicker(s.interval)

	go func() {
		time.Sleep(delay)
		s.job()
		for {
			select {
			case <-s.done:
				return
			case <-s.ticker.C:
				s.job()
			}
		}
	}()

	s.isRunning = true
}

func (s *scheduler) Stop() {
	if !s.isRunning {
		return
	}

	s.ticker.Stop()
	s.done <- true
	s.isRunning = false
}

func (s *scheduler) IsRunning() bool {
	return s.isRunning
}

func (s *scheduler) Reschedule(newInterval time.Duration, delay time.Duration) {
	if s.isRunning {
		s.Stop()
	}

	s.interval = newInterval

	if s.isRunning {
		s.Start(delay)
	}
}
