package spec

import (
	"time"

	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/entity"
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/value"
)

// OverDueDateSpecification represents a specification for tasks that have passed their due date.
type OverDueDateSpecification struct {
	Specification[entity.Task]
}

// NewOverDueDateSpecification creates a new instance of OverDueDateSpecification.
func NewOverDueDateSpecification() Specification[entity.Task] {
	s := &OverDueDateSpecification{
		Specification: &BaseSpecification[entity.Task]{},
	}
	s.Relate(s)
	return s
}

// IsSatisfiedBy checks if a task has passed its due date.
func (s *OverDueDateSpecification) IsSatisfiedBy(task *entity.Task) bool {
	return task.DueDate().Before(time.Now())
}

// ToQueryWithParams returns the query representation and params of the specification specified by OverDueDateSpecification.
func (s *OverDueDateSpecification) ToQueryWithParams() (string, []interface{}) {
	// TODO: SQL should depend on the repository layer and should not be described here (in the domain layer). Some kind of ingenuity, such as query abstraction, needs to be implemented.
	return "due_date < ?", []interface{}{time.Now()}
}

// CompletedSpecification represents a specification for completed tasks.
type CompletedSpecification struct {
	Specification[entity.Task]
}

// NewCompletedSpecification creates a new instance of CompletedSpecification.
func NewCompletedSpecification() Specification[entity.Task] {
	s := &CompletedSpecification{
		Specification: &BaseSpecification[entity.Task]{},
	}
	s.Relate(s)
	return s
}

// IsSatisfiedBy checks if a task is completed.
func (s *CompletedSpecification) IsSatisfiedBy(task *entity.Task) bool {
	return task.Status() == value.TaskStatusDone
}

// ToQueryWithParams returns the query representation and params of the specification specified by CompletedSpecification.
func (s *CompletedSpecification) ToQueryWithParams() (string, []interface{}) {
	// TODO: SQL should depend on the repository layer and should not be described here (in the domain layer). Some kind of ingenuity, such as query abstraction, needs to be implemented.
	return "status = ?", []interface{}{value.TaskStatusDone}
}
