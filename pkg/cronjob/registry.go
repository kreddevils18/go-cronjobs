package cronjob

import (
	"fmt"
	"sync"
)

// Registry defines the contract for storing and retrieving job definitions.
type Registry interface {
	// Add adds a new job definition to the registry.
	// It returns an error if a job with the same ID already exists.
	Add(def *JobDefinition) error

	// Find retrieves a job definition by its ID.
	// It returns an error if the job is not found.
	Find(id string) (*JobDefinition, error)

	// FindAll returns all job definitions in the registry.
	FindAll() ([]*JobDefinition, error)
}

// inMemoryRegistry is an in-memory implementation of the Registry interface.
type inMemoryRegistry struct {
	jobs map[string]*JobDefinition
	mu   sync.RWMutex
}

// NewInMemoryRegistry creates a new in-memory job registry.
func NewInMemoryRegistry() Registry {
	return &inMemoryRegistry{
		jobs: make(map[string]*JobDefinition),
	}
}

func (r *inMemoryRegistry) Add(def *JobDefinition) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.jobs[def.ID]; exists {
		return fmt.Errorf("job with ID '%s' already exists", def.ID)
	}
	r.jobs[def.ID] = def
	return nil
}

func (r *inMemoryRegistry) Find(id string) (*JobDefinition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	def, found := r.jobs[id]
	if !found {
		return nil, fmt.Errorf("job with ID '%s' not found", id)
	}
	return def, nil
}

func (r *inMemoryRegistry) FindAll() ([]*JobDefinition, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if len(r.jobs) == 0 {
		return []*JobDefinition{}, nil
	}

	defs := make([]*JobDefinition, 0, len(r.jobs))
	for _, def := range r.jobs {
		defs = append(defs, def)
	}

	return defs, nil
}
