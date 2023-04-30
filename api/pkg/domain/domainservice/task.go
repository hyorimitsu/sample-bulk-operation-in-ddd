package domainservice

import (
	"context"

	"github.com/google/uuid"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
)

type TaskDomainService interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Task, error)
	Create(ctx context.Context, ent *entity.Task) error
	Update(ctx context.Context, ent *entity.Task) error
	Delete(ctx context.Context, id uuid.UUID) error
}
