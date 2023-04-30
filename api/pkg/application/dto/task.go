package dto

import (
	"time"

	"github.com/google/uuid"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/value"
)

type Task struct {
	Id       uuid.UUID
	Title    string
	DueDate  time.Time
	Status   value.TaskStatus
	Progress int
}

type Tasks []*Task
