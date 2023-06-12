// Package spec provides interfaces and structs for implementing the Specification pattern.
package spec

import "fmt"

// Specification is an interface for determining whether a given entity satisfies a specific specification.
type Specification[E any] interface {
	// IsSatisfiedBy determines whether the given entity satisfies the specification specified. For in-memory.
	IsSatisfiedBy(ent *E) bool

	// ToQueryWithParams returns the query representation and params of the specification specified. For bulk operation.
	ToQueryWithParams() (string, []interface{})

	// And returns a new specification that represents the logical AND of the current specification and the given specification.
	And(spec Specification[E]) Specification[E]
	// Or returns a new specification that represents the logical OR of the current specification and the given specification.
	Or(spec Specification[E]) Specification[E]
	// Not returns a new specification that represents the logical NOT of the current specification.
	Not() Specification[E]
	// Relate sets the current specification to the given specification.
	Relate(spec Specification[E])
}

// AndSpecification is a struct for determining whether two specifications are both satisfied.
type AndSpecification[E any] struct {
	Specification[E]
	compare Specification[E]
}

func (as *AndSpecification[E]) IsSatisfiedBy(ent *E) bool {
	return as.Specification.IsSatisfiedBy(ent) && as.compare.IsSatisfiedBy(ent)
}

func (as *AndSpecification[E]) ToQueryWithParams() (string, []interface{}) {
	query1, params1 := as.Specification.ToQueryWithParams()
	query2, params2 := as.compare.ToQueryWithParams()
	// TODO: SQL should depend on the repository layer and should not be described here (in the domain layer). Some kind of ingenuity, such as query abstraction, needs to be implemented.
	return fmt.Sprintf("(%s AND %s)", query1, query2), append(params1, params2...)
}

// OrSpecification is a struct for determining whether at least one of two specifications is satisfied.
type OrSpecification[E any] struct {
	Specification[E]
	compare Specification[E]
}

func (os *OrSpecification[E]) IsSatisfiedBy(ent *E) bool {
	return os.Specification.IsSatisfiedBy(ent) || os.compare.IsSatisfiedBy(ent)
}

func (os *OrSpecification[E]) ToQueryWithParams() (string, []interface{}) {
	query1, params1 := os.Specification.ToQueryWithParams()
	query2, params2 := os.compare.ToQueryWithParams()
	// TODO: SQL should depend on the repository layer and should not be described here (in the domain layer). Some kind of ingenuity, such as query abstraction, needs to be implemented.
	return fmt.Sprintf("(%s OR %s)", query1, query2), append(params1, params2...)
}

// NotSpecification is a struct for determining entities that do not satisfy a given specification.
type NotSpecification[E any] struct {
	Specification[E]
}

func (ns *NotSpecification[E]) IsSatisfiedBy(ent *E) bool {
	return !ns.Specification.IsSatisfiedBy(ent)
}

func (ns *NotSpecification[E]) ToQueryWithParams() (string, []interface{}) {
	query, params := ns.Specification.ToQueryWithParams()
	// TODO: SQL should depend on the repository layer and should not be described here (in the domain layer). Some kind of ingenuity, such as query abstraction, needs to be implemented.
	return fmt.Sprintf("NOT(%s)", query), params
}

// BaseSpecification is a basic struct for determining whether a given entity satisfies a specification.
type BaseSpecification[E any] struct {
	Specification[E]
}

func (bs *BaseSpecification[E]) IsSatisfiedBy(ent *E) bool {
	return false
}

func (bs *BaseSpecification[E]) ToQueryWithParams() (string, []interface{}) {
	return "", []interface{}{}
}

func (bs *BaseSpecification[E]) And(spec Specification[E]) Specification[E] {
	s := &AndSpecification[E]{
		Specification: bs.Specification,
		compare:       spec,
	}
	s.Relate(s)
	return s
}

func (bs *BaseSpecification[E]) Or(spec Specification[E]) Specification[E] {
	s := &OrSpecification[E]{
		Specification: bs.Specification,
		compare:       spec,
	}
	s.Relate(s)
	return s
}

func (bs *BaseSpecification[E]) Not() Specification[E] {
	s := &NotSpecification[E]{
		Specification: bs.Specification,
	}
	s.Relate(s)
	return s
}

func (bs *BaseSpecification[E]) Relate(spec Specification[E]) {
	bs.Specification = spec
}
