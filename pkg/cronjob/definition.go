package cronjob

import "context"

// JobDefinition represents the metadata of a cron job.
type JobDefinition struct {
	ID        string
	Name      string
	CronSpec  string
	Task      func(ctx context.Context)
	IsEnabled bool
}
