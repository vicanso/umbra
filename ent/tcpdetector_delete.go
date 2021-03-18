// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/tcpdetector"
)

// TCPDetectorDelete is the builder for deleting a TCPDetector entity.
type TCPDetectorDelete struct {
	config
	hooks    []Hook
	mutation *TCPDetectorMutation
}

// Where adds a new predicate to the TCPDetectorDelete builder.
func (tdd *TCPDetectorDelete) Where(ps ...predicate.TCPDetector) *TCPDetectorDelete {
	tdd.mutation.predicates = append(tdd.mutation.predicates, ps...)
	return tdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tdd *TCPDetectorDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tdd.hooks) == 0 {
		affected, err = tdd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TCPDetectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tdd.mutation = mutation
			affected, err = tdd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tdd.hooks) - 1; i >= 0; i-- {
			mut = tdd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tdd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdd *TCPDetectorDelete) ExecX(ctx context.Context) int {
	n, err := tdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tdd *TCPDetectorDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tcpdetector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tcpdetector.FieldID,
			},
		},
	}
	if ps := tdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, tdd.driver, _spec)
}

// TCPDetectorDeleteOne is the builder for deleting a single TCPDetector entity.
type TCPDetectorDeleteOne struct {
	tdd *TCPDetectorDelete
}

// Exec executes the deletion query.
func (tddo *TCPDetectorDeleteOne) Exec(ctx context.Context) error {
	n, err := tddo.tdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tcpdetector.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tddo *TCPDetectorDeleteOne) ExecX(ctx context.Context) {
	tddo.tdd.ExecX(ctx)
}
