package cronjob

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/robfig/cron/v3"
)

// Handler defines the contract for handling job definitions.
type Handler interface {
	Handle(spec string, task func(ctx context.Context)) (*JobDefinition, error)
}

type jobDefinitionHandler struct {
	// dependencies like validators will go here
}

// NewJobDefinitionHandler creates a new handler for job definitions.
func NewJobDefinitionHandler() Handler {
	return &jobDefinitionHandler{}
}

func (h *jobDefinitionHandler) Handle(spec string, task func(ctx context.Context)) (*JobDefinition, error) {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)
	if _, err := p.Parse(spec); err != nil {
		return nil, fmt.Errorf("invalid cron spec: %w", err)
	}

	jobDef := &JobDefinition{
		ID:        uuid.NewString(),
		CronSpec:  spec,
		Task:      task,
		IsEnabled: true,
	}

	return jobDef, nil
}
