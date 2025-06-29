package cronjob

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobDefinitionHandler_Handle(t *testing.T) {
	t.Run("should return error for invalid cron spec", func(t *testing.T) {
		// Arrange
		handler := NewJobDefinitionHandler()
		invalidSpec := "this is not a cron spec"
		task := func(ctx context.Context) {}

		// Act
		def, err := handler.Handle(invalidSpec, task)

		// Assert
		assert.Error(t, err)
		assert.Nil(t, def)
	})

	t.Run("should return a job definition for a valid cron spec", func(t *testing.T) {
		// Arrange
		handler := NewJobDefinitionHandler()
		validSpec := "*/5 * * * *"
		task := func(ctx context.Context) {}

		// Act
		def, err := handler.Handle(validSpec, task)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, def)

		// Verify the content of the definition
		assert.Equal(t, validSpec, def.CronSpec)
		assert.True(t, def.IsEnabled)
		assert.NotEmpty(t, def.ID)
		assert.NotNil(t, def.Task) // Check that the function pointer is assigned
	})
}
