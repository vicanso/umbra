// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/schema"
)

// DNSDetectorResultUpdate is the builder for updating DNSDetectorResult entities.
type DNSDetectorResultUpdate struct {
	config
	hooks    []Hook
	mutation *DNSDetectorResultMutation
}

// Where adds a new predicate for the DNSDetectorResultUpdate builder.
func (ddru *DNSDetectorResultUpdate) Where(ps ...predicate.DNSDetectorResult) *DNSDetectorResultUpdate {
	ddru.mutation.predicates = append(ddru.mutation.predicates, ps...)
	return ddru
}

// SetTask sets the "task" field.
func (ddru *DNSDetectorResultUpdate) SetTask(i int) *DNSDetectorResultUpdate {
	ddru.mutation.ResetTask()
	ddru.mutation.SetTask(i)
	return ddru
}

// AddTask adds i to the "task" field.
func (ddru *DNSDetectorResultUpdate) AddTask(i int) *DNSDetectorResultUpdate {
	ddru.mutation.AddTask(i)
	return ddru
}

// SetResult sets the "result" field.
func (ddru *DNSDetectorResultUpdate) SetResult(i int8) *DNSDetectorResultUpdate {
	ddru.mutation.ResetResult()
	ddru.mutation.SetResult(i)
	return ddru
}

// AddResult adds i to the "result" field.
func (ddru *DNSDetectorResultUpdate) AddResult(i int8) *DNSDetectorResultUpdate {
	ddru.mutation.AddResult(i)
	return ddru
}

// SetMaxDuration sets the "maxDuration" field.
func (ddru *DNSDetectorResultUpdate) SetMaxDuration(i int) *DNSDetectorResultUpdate {
	ddru.mutation.ResetMaxDuration()
	ddru.mutation.SetMaxDuration(i)
	return ddru
}

// AddMaxDuration adds i to the "maxDuration" field.
func (ddru *DNSDetectorResultUpdate) AddMaxDuration(i int) *DNSDetectorResultUpdate {
	ddru.mutation.AddMaxDuration(i)
	return ddru
}

// SetMessages sets the "messages" field.
func (ddru *DNSDetectorResultUpdate) SetMessages(s []string) *DNSDetectorResultUpdate {
	ddru.mutation.SetMessages(s)
	return ddru
}

// SetHost sets the "host" field.
func (ddru *DNSDetectorResultUpdate) SetHost(s string) *DNSDetectorResultUpdate {
	ddru.mutation.SetHost(s)
	return ddru
}

// SetResults sets the "results" field.
func (ddru *DNSDetectorResultUpdate) SetResults(sdsr schema.DNSDetectorSubResults) *DNSDetectorResultUpdate {
	ddru.mutation.SetResults(sdsr)
	return ddru
}

// Mutation returns the DNSDetectorResultMutation object of the builder.
func (ddru *DNSDetectorResultUpdate) Mutation() *DNSDetectorResultMutation {
	return ddru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ddru *DNSDetectorResultUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ddru.defaults()
	if len(ddru.hooks) == 0 {
		if err = ddru.check(); err != nil {
			return 0, err
		}
		affected, err = ddru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ddru.check(); err != nil {
				return 0, err
			}
			ddru.mutation = mutation
			affected, err = ddru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ddru.hooks) - 1; i >= 0; i-- {
			mut = ddru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ddru *DNSDetectorResultUpdate) SaveX(ctx context.Context) int {
	affected, err := ddru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ddru *DNSDetectorResultUpdate) Exec(ctx context.Context) error {
	_, err := ddru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddru *DNSDetectorResultUpdate) ExecX(ctx context.Context) {
	if err := ddru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddru *DNSDetectorResultUpdate) defaults() {
	if _, ok := ddru.mutation.UpdatedAt(); !ok {
		v := dnsdetectorresult.UpdateDefaultUpdatedAt()
		ddru.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddru *DNSDetectorResultUpdate) check() error {
	if v, ok := ddru.mutation.Result(); ok {
		if err := dnsdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	return nil
}

func (ddru *DNSDetectorResultUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsdetectorresult.Table,
			Columns: dnsdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsdetectorresult.FieldID,
			},
		},
	}
	if ps := ddru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ddru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsdetectorresult.FieldUpdatedAt,
		})
	}
	if value, ok := ddru.mutation.Task(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldTask,
		})
	}
	if value, ok := ddru.mutation.AddedTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldTask,
		})
	}
	if value, ok := ddru.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetectorresult.FieldResult,
		})
	}
	if value, ok := ddru.mutation.AddedResult(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetectorresult.FieldResult,
		})
	}
	if value, ok := ddru.mutation.MaxDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := ddru.mutation.AddedMaxDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := ddru.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetectorresult.FieldMessages,
		})
	}
	if value, ok := ddru.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetectorresult.FieldHost,
		})
	}
	if value, ok := ddru.mutation.Results(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetectorresult.FieldResults,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ddru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsdetectorresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DNSDetectorResultUpdateOne is the builder for updating a single DNSDetectorResult entity.
type DNSDetectorResultUpdateOne struct {
	config
	hooks    []Hook
	mutation *DNSDetectorResultMutation
}

// SetTask sets the "task" field.
func (ddruo *DNSDetectorResultUpdateOne) SetTask(i int) *DNSDetectorResultUpdateOne {
	ddruo.mutation.ResetTask()
	ddruo.mutation.SetTask(i)
	return ddruo
}

// AddTask adds i to the "task" field.
func (ddruo *DNSDetectorResultUpdateOne) AddTask(i int) *DNSDetectorResultUpdateOne {
	ddruo.mutation.AddTask(i)
	return ddruo
}

// SetResult sets the "result" field.
func (ddruo *DNSDetectorResultUpdateOne) SetResult(i int8) *DNSDetectorResultUpdateOne {
	ddruo.mutation.ResetResult()
	ddruo.mutation.SetResult(i)
	return ddruo
}

// AddResult adds i to the "result" field.
func (ddruo *DNSDetectorResultUpdateOne) AddResult(i int8) *DNSDetectorResultUpdateOne {
	ddruo.mutation.AddResult(i)
	return ddruo
}

// SetMaxDuration sets the "maxDuration" field.
func (ddruo *DNSDetectorResultUpdateOne) SetMaxDuration(i int) *DNSDetectorResultUpdateOne {
	ddruo.mutation.ResetMaxDuration()
	ddruo.mutation.SetMaxDuration(i)
	return ddruo
}

// AddMaxDuration adds i to the "maxDuration" field.
func (ddruo *DNSDetectorResultUpdateOne) AddMaxDuration(i int) *DNSDetectorResultUpdateOne {
	ddruo.mutation.AddMaxDuration(i)
	return ddruo
}

// SetMessages sets the "messages" field.
func (ddruo *DNSDetectorResultUpdateOne) SetMessages(s []string) *DNSDetectorResultUpdateOne {
	ddruo.mutation.SetMessages(s)
	return ddruo
}

// SetHost sets the "host" field.
func (ddruo *DNSDetectorResultUpdateOne) SetHost(s string) *DNSDetectorResultUpdateOne {
	ddruo.mutation.SetHost(s)
	return ddruo
}

// SetResults sets the "results" field.
func (ddruo *DNSDetectorResultUpdateOne) SetResults(sdsr schema.DNSDetectorSubResults) *DNSDetectorResultUpdateOne {
	ddruo.mutation.SetResults(sdsr)
	return ddruo
}

// Mutation returns the DNSDetectorResultMutation object of the builder.
func (ddruo *DNSDetectorResultUpdateOne) Mutation() *DNSDetectorResultMutation {
	return ddruo.mutation
}

// Save executes the query and returns the updated DNSDetectorResult entity.
func (ddruo *DNSDetectorResultUpdateOne) Save(ctx context.Context) (*DNSDetectorResult, error) {
	var (
		err  error
		node *DNSDetectorResult
	)
	ddruo.defaults()
	if len(ddruo.hooks) == 0 {
		if err = ddruo.check(); err != nil {
			return nil, err
		}
		node, err = ddruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ddruo.check(); err != nil {
				return nil, err
			}
			ddruo.mutation = mutation
			node, err = ddruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ddruo.hooks) - 1; i >= 0; i-- {
			mut = ddruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ddruo *DNSDetectorResultUpdateOne) SaveX(ctx context.Context) *DNSDetectorResult {
	node, err := ddruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ddruo *DNSDetectorResultUpdateOne) Exec(ctx context.Context) error {
	_, err := ddruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ddruo *DNSDetectorResultUpdateOne) ExecX(ctx context.Context) {
	if err := ddruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ddruo *DNSDetectorResultUpdateOne) defaults() {
	if _, ok := ddruo.mutation.UpdatedAt(); !ok {
		v := dnsdetectorresult.UpdateDefaultUpdatedAt()
		ddruo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddruo *DNSDetectorResultUpdateOne) check() error {
	if v, ok := ddruo.mutation.Result(); ok {
		if err := dnsdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	return nil
}

func (ddruo *DNSDetectorResultUpdateOne) sqlSave(ctx context.Context) (_node *DNSDetectorResult, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   dnsdetectorresult.Table,
			Columns: dnsdetectorresult.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsdetectorresult.FieldID,
			},
		},
	}
	id, ok := ddruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DNSDetectorResult.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ddruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsdetectorresult.FieldUpdatedAt,
		})
	}
	if value, ok := ddruo.mutation.Task(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldTask,
		})
	}
	if value, ok := ddruo.mutation.AddedTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldTask,
		})
	}
	if value, ok := ddruo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetectorresult.FieldResult,
		})
	}
	if value, ok := ddruo.mutation.AddedResult(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetectorresult.FieldResult,
		})
	}
	if value, ok := ddruo.mutation.MaxDuration(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := ddruo.mutation.AddedMaxDuration(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldMaxDuration,
		})
	}
	if value, ok := ddruo.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetectorresult.FieldMessages,
		})
	}
	if value, ok := ddruo.mutation.Host(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetectorresult.FieldHost,
		})
	}
	if value, ok := ddruo.mutation.Results(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetectorresult.FieldResults,
		})
	}
	_node = &DNSDetectorResult{config: ddruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ddruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dnsdetectorresult.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}