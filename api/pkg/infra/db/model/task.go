package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/dto"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/value"
)

type Task struct {
	Id        uuid.UUID `gorm:"primaryKey"`
	Title     string
	DueDate   time.Time
	Status    int
	Progress  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Tasks []*Task

func NewTaskFromEntity(ent *entity.Task) *Task {
	return &Task{
		Id:       ent.Id(),
		Title:    ent.Title(),
		DueDate:  ent.DueDate(),
		Status:   ent.Status().Value(),
		Progress: ent.Progress(),
	}
}

func (m *Task) ToDto() *dto.Task {
	status, _ := value.TaskStatusFromInt(m.Status)
	return &dto.Task{
		Id:       m.Id,
		Title:    m.Title,
		DueDate:  m.DueDate,
		Status:   status,
		Progress: m.Progress,
	}
}

func (ms Tasks) ToDtos() dto.Tasks {
	dtos := make(dto.Tasks, len(ms))
	for i, m := range ms {
		dtos[i] = m.ToDto()
	}
	return dtos
}

func (m *Task) ToEntity() *entity.Task {
	status, _ := value.TaskStatusFromInt(m.Status)

	ent := &entity.Task{}
	ent.SetValuesFromDB(
		m.Id,
		m.Title,
		m.DueDate,
		status,
		m.Progress,
	)

	return ent
}

func (m *Task) ToUpdateFieldMap() map[string]interface{} {
	return map[string]interface{}{
		"status":   m.Status,
		"progress": m.Progress,
	}
}
