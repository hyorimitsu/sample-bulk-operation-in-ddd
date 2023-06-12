// Package cmd provides interfaces and structs for implementing the Command pattern.
package cmd

import (
	"github.com/hyorimitsu/sample-bulk-operation-in-ddd/api/pkg/domain/spec"
)

// Command is an interface for commands on an entity.
type Command[E any] interface {
	// CanExecute determines if the command can be executed. For in-memory.
	CanExecute(ent *E) bool
	// Execute executes the command. For in-memory.
	Execute(ent *E) error

	// ToQueryWithParams returns the query and params of the command. For bulk operation.
	ToQueryWithParams() (string, []interface{})
	// Values returns the values of the command. For bulk operation.
	Values() map[string]interface{}
}

// BaseCommand is a basic structure for commands on an entity.
type BaseCommand[E any] struct {
	spec   spec.Specification[E]
	values map[string]interface{}
}

func (bc *BaseCommand[E]) CanExecute(ent *E) bool {
	return bc.spec.IsSatisfiedBy(ent)
}

func (bc *BaseCommand[E]) Execute(ent *E) error {
	return nil
}

func (bc *BaseCommand[E]) ToQueryWithParams() (string, []interface{}) {
	return bc.spec.ToQueryWithParams()
}

func (bc *BaseCommand[E]) Values() map[string]interface{} {
	return bc.values
}
