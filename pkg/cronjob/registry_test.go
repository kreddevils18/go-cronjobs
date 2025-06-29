package cronjob

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInMemoryRegistry_AddAndFind(t *testing.T) {
	t.Run("should add and find a job successfully", func(t *testing.T) {
		// Arrange
		registry := NewInMemoryRegistry()
		jobDef := &JobDefinition{
			ID:       uuid.NewString(),
			CronSpec: "*/5 * * * *",
			Task:     func(ctx context.Context) {},
		}

		// Act: Add the job
		err := registry.Add(jobDef)
		assert.NoError(t, err)

		// Act: Find the job
		foundDef, err := registry.Find(jobDef.ID)

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, foundDef)
		assert.Equal(t, jobDef.ID, foundDef.ID)
		assert.Equal(t, jobDef.CronSpec, foundDef.CronSpec)
	})

	t.Run("should return error when adding a job with a duplicate ID", func(t *testing.T) {
		// Arrange
		registry := NewInMemoryRegistry()
		jobDef1 := &JobDefinition{ID: "job-1"}

		// Act
		err1 := registry.Add(jobDef1)
		err2 := registry.Add(jobDef1) // Add the same job again

		// Assert
		assert.NoError(t, err1)
		assert.Error(t, err2)
	})

	t.Run("should return error when finding a non-existent job", func(t *testing.T) {
		// Arrange
		registry := NewInMemoryRegistry()

		// Act
		foundDef, err := registry.Find("non-existent-id")

		// Assert
		assert.Error(t, err)
		assert.Nil(t, foundDef)
	})
}

func TestInMemoryRegistry_FindAll(t *testing.T) {
	t.Run("should return all added jobs", func(t *testing.T) {
		// Arrange
		registry := NewInMemoryRegistry()
		job1 := &JobDefinition{ID: "job-1"}
		job2 := &JobDefinition{ID: "job-2"}

		_ = registry.Add(job1)
		_ = registry.Add(job2)

		// Act
		allJobs, err := registry.FindAll()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, allJobs)
		assert.Len(t, allJobs, 2)
	})

	t.Run("should return an empty slice when no jobs are added", func(t *testing.T) {
		// Arrange
		registry := NewInMemoryRegistry()

		// Act
		allJobs, err := registry.FindAll()

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, allJobs)
		assert.Len(t, allJobs, 0)
	})
}
