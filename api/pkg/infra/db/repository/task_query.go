package repository

import (
	"context"

	"gorm.io/gorm"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/dto"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/infra/db/model"
)

type TaskQueryRepository struct {
	db *gorm.DB
}

func NewTaskQueryRepository(db *gorm.DB) *TaskQueryRepository {
	return &TaskQueryRepository{
		db: db,
	}
}

func (r *TaskQueryRepository) List(ctx context.Context) (dto.Tasks, error) {
	var mdls model.Tasks

	err := r.db.
		Order("due_date").
		Find(&mdls).
		Error

	return mdls.ToDtos(), err
}
