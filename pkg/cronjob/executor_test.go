package cronjob

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestJobExecutor_Execute(t *testing.T) {
	t.Run("should execute a task without panicking the main process", func(t *testing.T) {
		// Arrange
		mockLogger := new(MockLogger)
		mockLogger.On("Errorw", mock.Anything, mock.Anything, mock.Anything).Return()

		executor := NewJobExecutor(mockLogger)
		panickingTask := func(ctx context.Context) {
			panic("this is a test panic")
		}

		// Act
		// We can't directly assert that a panic was recovered in another goroutine.
		// The success of this test is that the `go test` process itself does not crash.
		executor.Execute(context.Background(), panickingTask)

		// Give the goroutine a moment to run and potentially panic.
		time.Sleep(100 * time.Millisecond)
		mockLogger.AssertExpectations(t)
	})

	t.Run("should execute a normal task", func(t *testing.T) {
		// Arrange
		mockLogger := new(MockLogger)
		executor := NewJobExecutor(mockLogger)
		var executed bool
		var wg sync.WaitGroup

		wg.Add(1)
		normalTask := func(ctx context.Context) {
			defer wg.Done()
			executed = true
		}

		// Act
		executor.Execute(context.Background(), normalTask)
		wg.Wait()

		// Assert
		assert.True(t, executed, "expected the task to be executed, but it was not")
	})
}
