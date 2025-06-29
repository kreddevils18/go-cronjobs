# Go Cronjob Package

A simple, lightweight, and reusable Go package for managing cron jobs with minimal configuration. This package allows you to define and manage asynchronous jobs through a single line of code.

## Features

- **Simple API**: Register and manage cron jobs with one line of code
- **Memory-first**: In-memory job storage and management
- **Thread-safe**: Concurrent job execution with proper synchronization
- **Context-aware**: Graceful shutdown with context cancellation
- **Flexible logging**: Pluggable logger interface with default silent operation
- **Clean Architecture**: Well-structured, testable, and maintainable code
- **High test coverage**: Comprehensive unit and integration tests

## Installation

```bash
go get github.com/kreddevils18/go-crobjobs
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/kreddevils18/go-crobjobs/pkg/cronjob"
    gologger "github.com/kreddevils18/go-logger"
)

func main() {
    // Create logger (optional)
    logger, err := gologger.NewDefaultLogger()
    if err != nil {
        log.Fatalf("Failed to initialize logger: %v", err)
    }
    defer logger.Sync()

    // Create manager with logger
    manager := cronjob.NewManager(cronjob.WithLogger(logger))

    // Register a job
    err = manager.Register("@every 5s", func(ctx context.Context) {
        logger.Info("Job executed!")
    })
    if err != nil {
        log.Fatal("Failed to register job:", err)
    }

    // Start the scheduler
    if err := manager.Start(); err != nil {
        log.Fatal("Failed to start manager:", err)
    }

    // Wait for shutdown signal
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    // Graceful shutdown
    manager.Stop()
}
```

## API Reference

### Manager

The `Manager` is the main entry point for the cronjob package.

#### Creating a Manager

```go
// Default manager (silent logging)
manager := cronjob.NewManager()

// Manager with custom logger
manager := cronjob.NewManager(cronjob.WithLogger(logger))

// Manager with custom registry
manager := cronjob.NewManager(cronjob.WithRegistry(customRegistry))
```

#### Manager Methods

- `Register(schedule string, task func(ctx context.Context)) error`: Register a new cron job
- `Start() error`: Start the job scheduler
- `Stop()`: Stop the scheduler and cancel all running jobs

### Cron Schedule Formats

The package supports standard cron expressions and descriptors:

```go
// Standard cron expressions
manager.Register("0 30 * * * *", task)    // Every hour at 30 minutes
manager.Register("0 0 12 * * *", task)    // Every day at noon

// Descriptors
manager.Register("@every 5s", task)       // Every 5 seconds
manager.Register("@every 1m", task)       // Every minute
manager.Register("@every 1h", task)       // Every hour
manager.Register("@daily", task)          // Daily at midnight
manager.Register("@weekly", task)         // Weekly on Sunday at midnight
manager.Register("@monthly", task)        // Monthly on the 1st at midnight
```

## Advanced Usage

### Context-Aware Jobs

Jobs receive a context that gets cancelled when the scheduler stops:

```go
manager.Register("@every 10s", func(ctx context.Context) {
    for i := 0; i < 10; i++ {
        select {
        case <-ctx.Done():
            log.Println("Job cancelled, cleaning up...")
            return
        case <-time.After(1 * time.Second):
            log.Printf("Working... %d/10", i+1)
        }
    }
    log.Println("Job completed")
})
```

### Custom Logger

Implement the `Logger` interface for custom logging:

```go
type Logger interface {
    Info(msg string, keysAndValues ...interface{})
    Infow(msg string, keysAndValues ...interface{})
    Warn(msg string, keysAndValues ...interface{})
    Warnw(msg string, keysAndValues ...interface{})
    Error(msg string, keysAndValues ...interface{})
    Errorw(msg string, keysAndValues ...interface{})
}
```

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## Examples

See the `examples/` directory for complete working examples:

- `examples/simple/main.go`: Basic usage with logging
- `examples/main.go`: Advanced usage with graceful shutdown

## Architecture

The package follows Clean Architecture principles:

- **Manager**: Facade pattern for simple API
- **Registry**: In-memory job storage with thread-safety
- **Scheduler**: Cron job scheduling using robfig/cron
- **Executor**: Safe job execution with panic recovery
- **Handler**: Job definition validation and processing

## Dependencies

- `github.com/robfig/cron/v3`: Cron expression parsing and scheduling
- `github.com/google/uuid`: Unique job ID generation
- `github.com/kreddevils18/go-logger`: Optional structured logging

## Contributing

1. Fork the repository
2. Create a feature branch
3. Write tests for your changes
4. Ensure all tests pass
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Changelog

### v1.0.0

- Initial release
- Basic cron job management
- Context-aware job execution
- Thread-safe operations
- Comprehensive test coverage
