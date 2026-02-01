package memory

import (
	"context"
	"sync"

	"github.com/your-org/your-app/internal/entity"
	"github.com/your-org/your-app/internal/usecase"
)

// Repository is the memory implementation of the persistence gateway (Interface Adapter).
// Replace with internal/repository/postgres or pkg/database when using plugins.
type Repository struct {
	mu   sync.RWMutex
	data map[string]*entity.Entity
}

// NewRepository creates a memory repository.
func NewRepository() *Repository {
	return &Repository{data: make(map[string]*entity.Entity)}
}

func (r *Repository) GetByID(ctx context.Context, id string) (*entity.Entity, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	e, ok := r.data[id]
	if !ok {
		return nil, entity.ErrNotFound
	}
	return e, nil
}

func (r *Repository) Save(ctx context.Context, e *entity.Entity) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data[e.ID] = e
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, id)
	return nil
}

// Ensures that Repository implements usecase.Repository at compile time.
var _ usecase.Repository = (*Repository)(nil)
