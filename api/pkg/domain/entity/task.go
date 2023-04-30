package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/value"
)

type Task struct {
	id       uuid.UUID
	title    string
	dueDate  time.Time
	status   value.TaskStatus
	progress int
}

func NewTask(title string, dueDateStr string) (*Task, error) {
	dueDate, err := time.Parse(time.RFC3339, dueDateStr)
	if err != nil {
		return nil, err
	}

	return &Task{
		id:       uuid.New(),
		title:    title,
		dueDate:  dueDate,
		status:   value.TaskStatusNone,
		progress: 0,
	}, nil
}

func (e *Task) UpdateStatus(statusStr string) error {
	status, err := value.TaskStatusFromString(statusStr)
	if err != nil {
		return err
	}

	e.status = status
	e.updateProgressByStatus()

	return nil
}

func (e *Task) updateProgressByStatus() {
	switch e.status {
	case value.TaskStatusNone:
		e.progress = 0
	case value.TaskStatusInProgress:
		e.progress = 50
	case value.TaskStatusDone:
		e.progress = 100
	}
}

func (e *Task) Id() uuid.UUID {
	return e.id
}

func (e *Task) Title() string {
	return e.title
}

func (e *Task) DueDate() time.Time {
	return e.dueDate
}

func (e *Task) Status() value.TaskStatus {
	return e.status
}

func (e *Task) Progress() int {
	return e.progress
}

func (e *Task) SetValuesFromDB(
	id uuid.UUID,
	title string,
	dueDate time.Time,
	status value.TaskStatus,
	progress int,
) {
	e.id = id
	e.title = title
	e.dueDate = dueDate
	e.status = status
	e.progress = progress
}
