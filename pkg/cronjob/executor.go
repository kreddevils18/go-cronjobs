package cronjob

import "context"

// Executor defines the contract for executing a job's task.
type Executor interface {
	// Execute runs the given task.
	// The implementation should handle running the task in a separate goroutine,
	// logging, and recovering from panics.
	Execute(ctx context.Context, task func(ctx context.Context))
}
