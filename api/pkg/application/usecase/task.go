package usecase

import (
	"context"

	"github.com/google/uuid"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/dto"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/input"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/queryservice"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/cmd"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/domainservice"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
)

type TaskUseCaser interface {
	ListTasks(ctx context.Context) (dto.Tasks, error)
	CreateTask(ctx context.Context, param *input.TaskCreateParam) error
	UpdateTask(ctx context.Context, id string, param *input.TaskUpdateParam) error
	UpdateExpiredTaskToDone(ctx context.Context, id string) error
	BulkUpdateExpiredTasksToDone(ctx context.Context) error
	DeleteTask(ctx context.Context, id string) error
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

func (u *TaskUseCase) ListTasks(ctx context.Context) (dto.Tasks, error) {
	return u.taskQueryService.List(ctx)
}

func (u *TaskUseCase) CreateTask(ctx context.Context, param *input.TaskCreateParam) error {
	ent, err := entity.NewTask(param.Title, param.DueDate)
	if err != nil {
		return err
	}
	return u.taskDomainService.Create(ctx, ent)
}

func (u *TaskUseCase) UpdateTask(ctx context.Context, id string, param *input.TaskUpdateParam) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	ent, err := u.taskDomainService.FindByID(ctx, parsedId)
	if err != nil {
		return err
	}

	// This is an unconditional update pattern without using `specification` and `command`.
	// If it is undesirable to have a mixture of patterns that use and do not use command, one option would be to implement something like UpdateTaskCommand.
	if err := ent.UpdateStatus(param.Status); err != nil {
		return err
	}

	return u.taskDomainService.Update(ctx, ent)
}

func (u *TaskUseCase) UpdateExpiredTaskToDone(ctx context.Context, id string) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	ent, err := u.taskDomainService.FindByID(ctx, parsedId)
	if err != nil {
		return err
	}

	// This is an in-memory update pattern using `specification` and `command`.
	command := cmd.NewUpdateExpiredTaskToDoneCommand()
	if err := command.Execute(ent); err != nil {
		return err
	}

	return u.taskDomainService.Update(ctx, ent)
}

func (u *TaskUseCase) BulkUpdateExpiredTasksToDone(ctx context.Context) error {
	// This is an bulk update pattern using `specification` and `command`.
	command := cmd.NewUpdateExpiredTaskToDoneCommand()
	return u.taskDomainService.BulkUpdate(ctx, command)
}

func (u *TaskUseCase) DeleteTask(ctx context.Context, id string) error {
	parsedId, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	return u.taskDomainService.Delete(ctx, parsedId)
}
