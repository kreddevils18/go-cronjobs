package cronjob

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestScheduler_Start(t *testing.T) {
	t.Run("should fetch jobs from registry and start cron engine", func(t *testing.T) {
		// Arrange
		mockRegistry := new(MockRegistry)
		mockExecutor := new(MockExecutor)

		jobs := []*JobDefinition{
			{ID: "job-1", CronSpec: "*/5 * * * *", Task: func(ctx context.Context) {}},
			{ID: "job-2", CronSpec: "@hourly", Task: func(ctx context.Context) {}},
		}

		// --- Expectations for Mocks ---
		// Expect FindAll() to be called once and return our predefined jobs.
		mockRegistry.On("FindAll").Return(jobs, nil).Once()
		// Expect Execute to be called for any task.
		mockExecutor.On("Execute", mock.Anything, mock.AnythingOfType("func(context.Context)")).Return()
		// -----------------------------

		sched := NewScheduler(mockRegistry, mockExecutor)

		// Act
		err := sched.Start()

		// Assert
		assert.NoError(t, err)
		mockRegistry.AssertExpectations(t) // Verify that FindAll was called.
	})

	t.Run("should return error if registry fails to fetch jobs", func(t *testing.T) {
		// Arrange
		mockRegistry := new(MockRegistry)
		mockExecutor := new(MockExecutor)

		expectedErr := errors.New("database is down")

		// Expect FindAll() to be called and return an error.
		mockRegistry.On("FindAll").Return(nil, expectedErr).Once()

		scheduler := NewScheduler(mockRegistry, mockExecutor)

		// Act
		err := scheduler.Start()

		// Assert
		assert.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		mockRegistry.AssertExpectations(t)
	})
}

func TestScheduler_Stop(t *testing.T) {
	t.Run("should stop the cron engine without errors", func(t *testing.T) {
		// Arrange
		mockRegistry := new(MockRegistry)
		mockExecutor := new(MockExecutor)

		mockRegistry.On("FindAll").Return([]*JobDefinition{}, nil)

		scheduler := NewScheduler(mockRegistry, mockExecutor)
		err := scheduler.Start()
		assert.NoError(t, err)

		// Act & Assert
		assert.NotPanics(t, func() {
			scheduler.Stop()
		}, "scheduler.Stop() should not panic")
	})
}

func TestScheduler_Add(t *testing.T) {
	t.Run("should add a new job to a running scheduler", func(t *testing.T) {
		// Arrange
		mockRegistry := new(MockRegistry)
		mockExecutor := new(MockExecutor)

		// Scheduler starts with no jobs.
		mockRegistry.On("FindAll").Return([]*JobDefinition{}, nil)
		// Executor will be called by the new job.
		mockExecutor.On("Execute", mock.Anything, mock.AnythingOfType("func(context.Context)"))

		sched := NewScheduler(mockRegistry, mockExecutor)
		err := sched.Start()
		assert.NoError(t, err)

		s := sched.(*scheduler)
		assert.Len(t, s.cronEngine.Entries(), 0, "Scheduler should start with 0 jobs")

		// Act
		newJob := &JobDefinition{ID: "new-job", CronSpec: "@every 1h", Task: func(ctx context.Context) {}}
		err = sched.Add(newJob)

		// Assert
		assert.NoError(t, err)
		assert.Len(t, s.cronEngine.Entries(), 1, "Scheduler should have 1 job after Add")
		mockRegistry.AssertExpectations(t)
	})
}
