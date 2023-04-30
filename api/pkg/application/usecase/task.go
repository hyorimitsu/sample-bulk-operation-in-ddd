package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/dto"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/input"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/queryservice"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/domainservice"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
)

type TaskUseCaser interface {
	ListTodos(ctx context.Context) (dto.Tasks, error)
	CreateTodo(ctx context.Context, param *input.TaskCreateParam) error
	UpdateTodo(ctx context.Context, id string, param *input.TaskUpdateParam) error
	DeleteTodo(ctx context.Context, id string) error
}

type TaskUseCase struct {
	taskQueryService  queryservice.TaskQueryService
	taskDomainService domainservice.TaskDomainService
}

func NewTaskUseCase(
	taskQueryService queryservice.TaskQueryService,
	taskDomainService domainservice.TaskDomainService,
) *TaskUseCase {
	return &TaskUseCase{
		taskQueryService:  taskQueryService,
		taskDomainService: taskDomainService,
	}
}

func (u *TaskUseCase) ListTodos(ctx context.Context) (dto.Tasks, error) {
	return u.taskQueryService.List(ctx)
}

func (u *TaskUseCase) CreateTodo(ctx context.Context, param *input.TaskCreateParam) error {
	ent, err := entity.NewTask(param.Title, param.DueDate)
	if err != nil {
		return err
	}
	return u.taskDomainService.Create(ctx, ent)
}

func (u *TaskUseCase) UpdateTodo(ctx context.Context, id string, param *input.TaskUpdateParam) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	ent, err := u.taskDomainService.FindByID(ctx, parsedId)
	if err != nil {
		return err
	}

	if err := ent.UpdateStatus(param.Status); err != nil {
		return err
	}

	return u.taskDomainService.Update(ctx, ent)
}

func (u *TaskUseCase) DeleteTodo(ctx context.Context, id string) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return u.taskDomainService.Delete(ctx, parsedId)
}
