package usecase

import (
	"context"
	"time"

	"github.com/your-org/your-app/internal/entity"
)

// UseCase implements the application rules (transaction orchestration).
// Depends only on entity and the Repository interface; independent of HTTP and DB.
type UseCase struct {
	repo Repository
}

// New creates the use case with the repository injected.
func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

// GetEntity returns an entity by ID.
func (u *UseCase) GetEntity(ctx context.Context, id string) (*entity.Entity, error) {
	return u.repo.GetByID(ctx, id)
}

// CreateEntity creates an entity (applies application rules and persists).
func (u *UseCase) CreateEntity(ctx context.Context, e *entity.Entity) error {
	now := time.Now()
	e.CreatedAt = now
	e.UpdatedAt = now
	return u.repo.Save(ctx, e)
}
