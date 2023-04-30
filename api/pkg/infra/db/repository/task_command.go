package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/infra/db/model"
)

type TaskCommandRepository struct {
	db *gorm.DB
}

func NewTaskCommandRepository(db *gorm.DB) *TaskCommandRepository {
	return &TaskCommandRepository{
		db: db,
	}
}

func (r *TaskCommandRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Task, error) {
	var mdl model.Task

	err := r.db.
		Where("id = ?", id).
		First(&mdl).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return mdl.ToEntity(), nil
}

func (r *TaskCommandRepository) Create(ctx context.Context, ent *entity.Task) error {
	mdl := model.NewTaskFromEntity(ent)
	return r.db.
		WithContext(ctx).
		Create(mdl).
		Error
}

func (r *TaskCommandRepository) Update(ctx context.Context, ent *entity.Task) error {
	mdl := model.NewTaskFromEntity(ent)
	return r.db.
		WithContext(ctx).
		Model(mdl).
		Updates(mdl.ToUpdateFieldMap()).
		Error
}

func (r *TaskCommandRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.
		WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Task{}).
		Error
}
