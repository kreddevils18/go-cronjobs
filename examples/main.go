package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kreddevils18/go-crobjobs/pkg/cronjob"
	gologger "github.com/kreddevils18/go-logger"
)

func main() {
	// Create a new logger from the external library
	logger, err := gologger.NewDefaultLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Info("Starting cron job manager...")

	// 1. Create a new manager with the custom logger.
	manager := cronjob.NewManager(cronjob.WithLogger(logger))

	// 2. Register jobs.
	logger.Info("Registering jobs...")
	// Job 1: A long-running job that respects context cancellation.
	err = manager.Register("@every 2s", func(ctx context.Context) {
		logger.Infow("-> Long-running job: starting work...", "job", "long-runner")

		// Simulate work that takes 10 seconds, but check for cancellation every second.
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				// The scheduler is stopping. Clean up and exit.
				logger.Warnw("-> Long-running job: context cancelled, stopping work.", "job", "long-runner")
				return
			case <-time.After(1 * time.Second):
				// Continue doing work.
				logger.Infow("-> Long-running job: ...working...", "progress", fmt.Sprintf("%d/10", i+1))
			}
		}
		logger.Infow("-> Long-running job: finished work.", "job", "long-runner")
	})
	if err != nil {
		logger.Fatal("Failed to register long-running job", "error", err)
	}

	// 3. Start the manager.
	logger.Info("Starting scheduler...")
	if err := manager.Start(); err != nil {
		logger.Fatal("Failed to start manager", "error", err)
	}
	logger.Info("Scheduler started. Press Ctrl+C to stop.")

	// 4. Wait for a shutdown signal.
	waitForShutdown()
	logger.Info("Shutdown signal received.")

	// 5. Stop the manager gracefully.
	// The manager will now signal the long-running job to stop.
	logger.Info("Stopping scheduler...")
	manager.Stop()
	logger.Info("Scheduler stopped gracefully.")
}

// waitForShutdown blocks until a SIGINT or SIGTERM signal is received.
func waitForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
