package cronjob

import (
	"context"
	"errors"
	"fmt"

	"github.com/robfig/cron/v3"
)

// Scheduler defines the contract for the job scheduling engine.
type Scheduler interface {
	// Start begins the scheduler's operation.
	// It loads jobs from the registry and starts the cron engine.
	Start() error

	// Stop gracefully shuts down the scheduler.
	Stop()

	// Add schedules a new job definition immediately without restarting the scheduler.
	Add(def *JobDefinition) error
}

type scheduler struct {
	registry   Registry
	executor   Executor
	cronEngine *cron.Cron
	// jobIDs maps our job definition ID to the cron engine's internal EntryID.
	jobIDs     map[string]cron.EntryID
	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewScheduler creates a new scheduler instance.
func NewScheduler(registry Registry, executor Executor) Scheduler {
	ctx, cancel := context.WithCancel(context.Background())
	return &scheduler{
		registry:   registry,
		executor:   executor,
		cronEngine: cron.New(),
		jobIDs:     make(map[string]cron.EntryID),
		ctx:        ctx,
		cancelFunc: cancel,
	}
}

func (s *scheduler) Start() error {
	jobs, err := s.registry.FindAll()
	if err != nil {
		return fmt.Errorf("could not fetch jobs from registry: %w", err)
	}

	for _, job := range jobs {
		// Create a closure to capture the job's task.
		task := job.Task
		entryID, err := s.cronEngine.AddFunc(job.CronSpec, func() {
			s.executor.Execute(s.ctx, task)
		})
		if err != nil {
			// In a real app, you might want to log this but continue with other jobs.
			return fmt.Errorf("could not schedule job ID %s: %w", job.ID, err)
		}
		s.jobIDs[job.ID] = entryID
	}

	s.cronEngine.Start()
	return nil
}

func (s *scheduler) Stop() {
	s.cancelFunc() // Signal all running jobs to stop.
	s.cronEngine.Stop()
}

func (s *scheduler) Add(def *JobDefinition) error {
	if def == nil {
		return errors.New("job definition cannot be nil")
	}
	task := def.Task
	entryID, err := s.cronEngine.AddFunc(def.CronSpec, func() {
		s.executor.Execute(s.ctx, task)
	})
	if err != nil {
		return fmt.Errorf("could not schedule job ID %s: %w", def.ID, err)
	}
	s.jobIDs[def.ID] = entryID
	return nil
}
