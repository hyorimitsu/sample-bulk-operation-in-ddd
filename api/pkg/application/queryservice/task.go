package queryservice

import (
	"context"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/dto"
)

type TaskQueryService interface {
	List(ctx context.Context) (dto.Tasks, error)
}
