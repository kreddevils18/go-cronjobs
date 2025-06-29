package cronjob

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManager_Integration(t *testing.T) {
	t.Run("should register, start, and stop correctly", func(t *testing.T) {
		// Arrange
		manager := NewManager()

		// Act: Register jobs
		err1 := manager.Register("*/5 * * * *", func(ctx context.Context) {})
		err2 := manager.Register("@hourly", func(ctx context.Context) {})

		// Assert registration
		assert.NoError(t, err1)
		assert.NoError(t, err2)

		// Act: Start the manager
		err := manager.Start()
		assert.NoError(t, err)

		// Assert internal state after start
		// We need to cast to access the internal scheduler for this test
		internalScheduler := manager.scheduler.(*scheduler)
		assert.Len(t, internalScheduler.cronEngine.Entries(), 2, "Scheduler should have 2 jobs running")

		// Act: Stop the manager
		assert.NotPanics(t, func() {
			manager.Stop()
		})
	})

	t.Run("should return error for invalid cron spec during registration", func(t *testing.T) {
		// Arrange
		manager := NewManager()

		// Act
		err := manager.Register("invalid spec", func(ctx context.Context) {})

		// Assert
		assert.Error(t, err)
	})
}
