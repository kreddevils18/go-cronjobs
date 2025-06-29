package cronjob

import (
	"context"
	"log"
)

// Option defines a functional option for configuring the Manager.
type Option func(*Manager)

// Manager is the main entry point for the cronjob package.
// It acts as a facade, simplifying interaction with the underlying components.
type Manager struct {
	handler   Handler
	registry  Registry
	scheduler Scheduler
	logger    Logger
}

// NewManager creates and initializes a new cron job manager and all its components.
func NewManager(opts ...Option) *Manager {
	// Default components
	mgr := &Manager{
		handler:  NewJobDefinitionHandler(),
		registry: NewInMemoryRegistry(),
		logger:   &defaultLogger{}, // Use a silent logger by default
	}

	// Apply all custom options
	for _, opt := range opts {
		opt(mgr)
	}

	// Components that depend on other (potentially customized) components
	executor := NewJobExecutor(mgr.logger)
	mgr.scheduler = NewScheduler(mgr.registry, executor)

	return mgr
}

// WithRegistry provides an option to use a custom job registry.
func WithRegistry(r Registry) Option {
	return func(m *Manager) {
		m.registry = r
	}
}

// WithLogger provides an option to use a custom logger.
func WithLogger(l Logger) Option {
	return func(m *Manager) {
		m.logger = l
	}
}

// Register defines a new job and adds it to the registry.
// This should be called before starting the manager.
func (m *Manager) Register(spec string, task func(ctx context.Context)) error {
	def, err := m.handler.Handle(spec, task)
	if err != nil {
		return err
	}
	return m.registry.Add(def)
}

// Start loads all registered jobs and starts the scheduler.
func (m *Manager) Start() error {
	return m.scheduler.Start()
}

// Stop gracefully stops the scheduler.
func (m *Manager) Stop() {
	m.scheduler.Stop()
}

// --- Default Logger (internal) ---

// defaultLogger is a no-op logger used when no other logger is provided.
// It ensures that the library doesn't log anything by default.
type defaultLogger struct{}

func (l *defaultLogger) Debug(args ...interface{})                       {}
func (l *defaultLogger) Info(args ...interface{})                        {}
func (l *defaultLogger) Warn(args ...interface{})                        {}
func (l *defaultLogger) Error(args ...interface{})                       {}
func (l *defaultLogger) DPanic(args ...interface{})                      {}
func (l *defaultLogger) Panic(args ...interface{})                       {}
func (l *defaultLogger) Fatal(args ...interface{})                       {}
func (l *defaultLogger) Debugf(template string, args ...interface{})     {}
func (l *defaultLogger) Infof(template string, args ...interface{})      {}
func (l *defaultLogger) Warnf(template string, args ...interface{})      {}
func (l *defaultLogger) Errorf(template string, args ...interface{})     {}
func (l *defaultLogger) DPanicf(template string, args ...interface{})    {}
func (l *defaultLogger) Panicf(template string, args ...interface{})     {}
func (l *defaultLogger) Fatalf(template string, args ...interface{})     {}
func (l *defaultLogger) Debugw(msg string, keysAndValues ...interface{}) {}
func (l *defaultLogger) Infow(msg string, keysAndValues ...interface{})  {}
func (l *defaultLogger) Warnw(msg string, keysAndValues ...interface{})  {}
func (l *defaultLogger) Errorw(msg string, keysAndValues ...interface{}) {
	log.Printf(msg, keysAndValues...)
}                                                                         // log panics at least
func (l *defaultLogger) DPanicw(msg string, keysAndValues ...interface{}) {}
func (l *defaultLogger) Panicw(msg string, keysAndValues ...interface{})  {}
func (l *defaultLogger) Fatalw(msg string, keysAndValues ...interface{})  {}
func (l *defaultLogger) Sync() error                                      { return nil }
