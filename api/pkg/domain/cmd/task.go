package cmd

import (
	"errors"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/spec"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/value"
)

// UpdateExpiredTaskToDoneCommand represents a command to update the status of an expired task to "done".
type UpdateExpiredTaskToDoneCommand struct {
	Command[entity.Task]
}

// NewUpdateExpiredTaskToDoneCommand create a new instance of UpdateExpiredTaskToDoneCommand.
func NewUpdateExpiredTaskToDoneCommand() Command[entity.Task] {
	overDueDateSpec := spec.NewOverDueDateSpecification()
	completedSpec := spec.NewCompletedSpecification()

	specification := overDueDateSpec.And(completedSpec.Not())
	values := map[string]interface{}{
		"status": value.TaskStatusDone,
		// TODO: The behavior of `progress` when `status` is updated is defined in entity.Task. To avoid scattered definitions, it should be reused here.
		"progress": 100,
	}

	return &UpdateExpiredTaskToDoneCommand{
		Command: &BaseCommand[entity.Task]{
			spec:   specification,
			values: values,
		},
	}
}

// Execute updates the status of the given task to "done" if the command can be executed, otherwise it returns an error.
func (c *UpdateExpiredTaskToDoneCommand) Execute(ent *entity.Task) error {
	if !c.CanExecute(ent) {
		return errors.New("unable to execute update task state command")
	}
	return ent.UpdateStatus(value.TaskStatusDone.String())
}
