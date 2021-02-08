// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/pingdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/schema"
)

// PingDetectorResultUpdate is the builder for updating PingDetectorResult entities.
type PingDetectorResultUpdate struct {
	config
	hooks    []Hook
	mutation *PingDetectorResultMutation
}

// Where adds a new predicate for the PingDetectorResultUpdate builder.
func (pdru *PingDetectorResultUpdate) Where(ps ...predicate.PingDetectorResult) *PingDetectorResultUpdate {
	pdru.mutation.predicates = append(pdru.mutation.predicates, ps...)
	return pdru
}

// SetTask sets the "task" field.
func (pdru *PingDetectorResultUpdate) SetTask(i int) *PingDetectorResultUpdate {
	pdru.mutation.ResetTask()
	pdru.mutation.SetTask(i)
	return pdru
}

// AddTask adds i to the "task" field.
func (pdru *PingDetectorResultUpdate) AddTask(i int) *PingDetectorResultUpdate {
	pdru.mutation.AddTask(i)
	return pdru
}

// SetResult sets the "result" field.
func (pdru *PingDetectorResultUpdate) SetResult(i int8) *PingDetectorResultUpdate {
	pdru.mutation.ResetResult()
	pdru.mutation.SetResult(i)
	return pdru
}

// AddResult adds i to the "result" field.
func (pdru *PingDetectorResultUpdate) AddResult(i int8) *PingDetectorResultUpdate {
	pdru.mutation.AddResult(i)
	return pdru
}

// SetMaxDuration sets the "maxDuration" field.
func (pdru *PingDetectorResultUpdate) SetMaxDuration(i int) *PingDetectorResultUpdate {
	pdru.mutation.ResetMaxDuration()
	pdru.mutation.SetMaxDuration(i)
	return pdru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (pdru *PingDetectorResultUpdate) AddMaxDuration(i int) *PingDetectorResultUpdate {
	pdru.mutation.AddMaxDuration(i)
	return pdru
}

// SetMessages sets the "messages" field.
func (pdru *PingDetectorResultUpdate) SetMessages(s []string) *PingDetectorResultUpdate {
	pdru.mutation.SetMessages(s)
	return pdru
}

// SetIps sets the "ips" field.
func (pdru *PingDetectorResultUpdate) SetIps(s string) *PingDetectorResultUpdate {
	pdru.mutation.SetIps(s)
	return pdru
}

// SetResults sets the "results" field.
func (pdru *PingDetectorResultUpdate) SetResults(sdsr schema.PingDetectorSubResults) *PingDetectorResultUpdate {
	pdru.mutation.SetResults(sdsr)
	return pdru
}

// Mutation returns the PingDetectorResultMutation object of the builder.
func (pdru *PingDetectorResultUpdate) Mutation() *PingDetectorResultMutation {
	return pdru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pdru *PingDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pdru.defaults()
	if len(pdru.hooks) == 0 {
		if err = pdru.check(); err != nil {
			return 0, err
		}
		affected, err = pdru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PingDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pdru.check(); err != nil {
				return 0, err
			}
			pdru.mutation = mutation
			affected, err = pdru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pdru.hooks) - 1; i >= 0; i-- {
			mut = pdru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pdru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pdru *PingDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := pdru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pdru *PingDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := pdru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdru *PingDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := pdru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdru *PingDetectorResultUpdate) defaults() {
	if _, ok := pdru.mutation.UpdatedAt(); !ok {
		v := pingdetectorresult.UpdateDefaultUpdatedAt()
		pdru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdru *PingDetectorResultUpdate) check() error {
	if v, ok := pdru.mutation.Result(); ok {
		if err := pingdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	return nil
}

func (pdru *PingDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pingdetectorresult.Table,
			Columns: pingdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pingdetectorresult.FieldID,
			},
		},
	}
	if ps := pdru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pdru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pingdetectorresult.FieldUpdatedAt,
		})
	}
	if value, ok := pdru.mutation.Task(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldTask,
		})
	}
	if value, ok := pdru.mutation.AddedTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldTask,
		})
	}
	if value, ok := pdru.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: pingdetectorresult.FieldResult,
		})
	}
	if value, ok := pdru.mutation.AddedResult(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: pingdetectorresult.FieldResult,
		})
	}
	if value, ok := pdru.mutation.MaxDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := pdru.mutation.AddedMaxDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := pdru.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: pingdetectorresult.FieldMessages,
		})
	}
	if value, ok := pdru.mutation.Ips(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pingdetectorresult.FieldIps,
		})
	}
	if value, ok := pdru.mutation.Results(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: pingdetectorresult.FieldResults,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pdru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pingdetectorresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PingDetectorResultUpdateOne is the builder for updating a single PingDetectorResult entity.
type PingDetectorResultUpdateOne struct {
	config
	hooks    []Hook
	mutation *PingDetectorResultMutation
}

// SetTask sets the "task" field.
func (pdruo *PingDetectorResultUpdateOne) SetTask(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.ResetTask()
	pdruo.mutation.SetTask(i)
	return pdruo
}

// AddTask adds i to the "task" field.
func (pdruo *PingDetectorResultUpdateOne) AddTask(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.AddTask(i)
	return pdruo
}

// SetResult sets the "result" field.
func (pdruo *PingDetectorResultUpdateOne) SetResult(i int8) *PingDetectorResultUpdateOne {
	pdruo.mutation.ResetResult()
	pdruo.mutation.SetResult(i)
	return pdruo
}

// AddResult adds i to the "result" field.
func (pdruo *PingDetectorResultUpdateOne) AddResult(i int8) *PingDetectorResultUpdateOne {
	pdruo.mutation.AddResult(i)
	return pdruo
}

// SetMaxDuration sets the "maxDuration" field.
func (pdruo *PingDetectorResultUpdateOne) SetMaxDuration(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.ResetMaxDuration()
	pdruo.mutation.SetMaxDuration(i)
	return pdruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (pdruo *PingDetectorResultUpdateOne) AddMaxDuration(i int) *PingDetectorResultUpdateOne {
	pdruo.mutation.AddMaxDuration(i)
	return pdruo
}

// SetMessages sets the "messages" field.
func (pdruo *PingDetectorResultUpdateOne) SetMessages(s []string) *PingDetectorResultUpdateOne {
	pdruo.mutation.SetMessages(s)
	return pdruo
}

// SetIps sets the "ips" field.
func (pdruo *PingDetectorResultUpdateOne) SetIps(s string) *PingDetectorResultUpdateOne {
	pdruo.mutation.SetIps(s)
	return pdruo
}

// SetResults sets the "results" field.
func (pdruo *PingDetectorResultUpdateOne) SetResults(sdsr schema.PingDetectorSubResults) *PingDetectorResultUpdateOne {
	pdruo.mutation.SetResults(sdsr)
	return pdruo
}

// Mutation returns the PingDetectorResultMutation object of the builder.
func (pdruo *PingDetectorResultUpdateOne) Mutation() *PingDetectorResultMutation {
	return pdruo.mutation
}

// Save executes the query and returns the updated PingDetectorResult entity.
func (pdruo *PingDetectorResultUpdateOne) Save(ctx context.Context) (*PingDetectorResult, error) {
	var (
		err  error
		node *PingDetectorResult
	)
	pdruo.defaults()
	if len(pdruo.hooks) == 0 {
		if err = pdruo.check(); err != nil {
			return nil, err
		}
		node, err = pdruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PingDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pdruo.check(); err != nil {
				return nil, err
			}
			pdruo.mutation = mutation
			node, err = pdruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pdruo.hooks) - 1; i >= 0; i-- {
			mut = pdruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pdruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (pdruo *PingDetectorResultUpdateOne) SaveX(ctx context.Context) *PingDetectorResult {
	node, err := pdruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pdruo *PingDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := pdruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pdruo *PingDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := pdruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pdruo *PingDetectorResultUpdateOne) defaults() {
	if _, ok := pdruo.mutation.UpdatedAt(); !ok {
		v := pingdetectorresult.UpdateDefaultUpdatedAt()
		pdruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pdruo *PingDetectorResultUpdateOne) check() error {
	if v, ok := pdruo.mutation.Result(); ok {
		if err := pingdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	return nil
}

func (pdruo *PingDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *PingDetectorResult, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pingdetectorresult.Table,
			Columns: pingdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pingdetectorresult.FieldID,
			},
		},
	}
	id, ok := pdruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing PingDetectorResult.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := pdruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pingdetectorresult.FieldUpdatedAt,
		})
	}
	if value, ok := pdruo.mutation.Task(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldTask,
		})
	}
	if value, ok := pdruo.mutation.AddedTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldTask,
		})
	}
	if value, ok := pdruo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: pingdetectorresult.FieldResult,
		})
	}
	if value, ok := pdruo.mutation.AddedResult(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: pingdetectorresult.FieldResult,
		})
	}
	if value, ok := pdruo.mutation.MaxDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := pdruo.mutation.AddedMaxDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pingdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := pdruo.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: pingdetectorresult.FieldMessages,
		})
	}
	if value, ok := pdruo.mutation.Ips(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pingdetectorresult.FieldIps,
		})
	}
	if value, ok := pdruo.mutation.Results(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: pingdetectorresult.FieldResults,
		})
	}
	_node = &PingDetectorResult{config: pdruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pdruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pingdetectorresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}