package cronjob

import "context"

// jobExecutor is a concrete implementation of the Executor interface.
type jobExecutor struct {
	logger Logger
}

// NewJobExecutor creates a new job executor.
func NewJobExecutor(logger Logger) Executor {
	return &jobExecutor{
		logger: logger,
	}
}

// Execute runs the task in a new goroutine and recovers from any panics.
func (e *jobExecutor) Execute(ctx context.Context, task func(ctx context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				e.logger.Errorw("recovered from panic in job", "panic", r)
			}
		}()
		task(ctx)
	}()
}
