package presenter

import (
	"time"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/application/dto"
)

type Task struct {
	Id       string
	Title    string
	DueDate  string
	Status   string
	Progress int
}

type Tasks []*Task

func MapToTask(mdl *dto.Task) *Task {
	return &Task{
		Id:       mdl.Id.String(),
		Title:    mdl.Title,
		DueDate:  mdl.DueDate.Format(time.RFC3339),
		Status:   mdl.Status.String(),
		Progress: mdl.Progress,
	}
}

func MapToTasks(mdls dto.Tasks) Tasks {
	resp := make(Tasks, len(mdls))
	for i, mdl := range mdls {
		resp[i] = MapToTask(mdl)
	}
	return resp
}
