package usecase

import (
	"context"

	"github.com/your-org/your-app/internal/entity"
)

// Repository is the interface of persistence required by use cases (Interface Adapter - gateway).
// Implementations are in internal/repository (e.g. memory, postgres) or in pkg (plugins).
// The dependency points inward: usecase depends on the interface, not the concrete adapter.
type Repository interface {
	GetByID(ctx context.Context, id string) (*entity.Entity, error)
	Save(ctx context.Context, e *entity.Entity) error
	Delete(ctx context.Context, id string) error
}
