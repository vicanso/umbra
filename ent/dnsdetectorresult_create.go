// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/dnsdetectorresult"
	"github.com/vicanso/cybertect/ent/schema"
)

// DNSDetectorResultCreate is the builder for creating a DNSDetectorResult entity.
type DNSDetectorResultCreate struct {
	config
	mutation *DNSDetectorResultMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ddrc *DNSDetectorResultCreate) SetCreatedAt(t time.Time) *DNSDetectorResultCreate {
	ddrc.mutation.SetCreatedAt(t)
	return ddrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ddrc *DNSDetectorResultCreate) SetNillableCreatedAt(t *time.Time) *DNSDetectorResultCreate {
	if t != nil {
		ddrc.SetCreatedAt(*t)
	}
	return ddrc
}

// SetUpdatedAt sets the "updated_at" field.
func (ddrc *DNSDetectorResultCreate) SetUpdatedAt(t time.Time) *DNSDetectorResultCreate {
	ddrc.mutation.SetUpdatedAt(t)
	return ddrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ddrc *DNSDetectorResultCreate) SetNillableUpdatedAt(t *time.Time) *DNSDetectorResultCreate {
	if t != nil {
		ddrc.SetUpdatedAt(*t)
	}
	return ddrc
}

// SetTask sets the "task" field.
func (ddrc *DNSDetectorResultCreate) SetTask(i int) *DNSDetectorResultCreate {
	ddrc.mutation.SetTask(i)
	return ddrc
}

// SetResult sets the "result" field.
func (ddrc *DNSDetectorResultCreate) SetResult(i int8) *DNSDetectorResultCreate {
	ddrc.mutation.SetResult(i)
	return ddrc
}

// SetMaxDuration sets the "maxDuration" field.
func (ddrc *DNSDetectorResultCreate) SetMaxDuration(i int) *DNSDetectorResultCreate {
	ddrc.mutation.SetMaxDuration(i)
	return ddrc
}

// SetMessages sets the "messages" field.
func (ddrc *DNSDetectorResultCreate) SetMessages(s []string) *DNSDetectorResultCreate {
	ddrc.mutation.SetMessages(s)
	return ddrc
}

// SetHost sets the "host" field.
func (ddrc *DNSDetectorResultCreate) SetHost(s string) *DNSDetectorResultCreate {
	ddrc.mutation.SetHost(s)
	return ddrc
}

// SetResults sets the "results" field.
func (ddrc *DNSDetectorResultCreate) SetResults(sdsr schema.DNSDetectorSubResults) *DNSDetectorResultCreate {
	ddrc.mutation.SetResults(sdsr)
	return ddrc
}

// Mutation returns the DNSDetectorResultMutation object of the builder.
func (ddrc *DNSDetectorResultCreate) Mutation() *DNSDetectorResultMutation {
	return ddrc.mutation
}

// Save creates the DNSDetectorResult in the database.
func (ddrc *DNSDetectorResultCreate) Save(ctx context.Context) (*DNSDetectorResult, error) {
	var (
		err  error
		node *DNSDetectorResult
	)
	ddrc.defaults()
	if len(ddrc.hooks) == 0 {
		if err = ddrc.check(); err != nil {
			return nil, err
		}
		node, err = ddrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DNSDetectorResultMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ddrc.check(); err != nil {
				return nil, err
			}
			ddrc.mutation = mutation
			node, err = ddrc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ddrc.hooks) - 1; i >= 0; i-- {
			mut = ddrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ddrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ddrc *DNSDetectorResultCreate) SaveX(ctx context.Context) *DNSDetectorResult {
	v, err := ddrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (ddrc *DNSDetectorResultCreate) defaults() {
	if _, ok := ddrc.mutation.CreatedAt(); !ok {
		v := dnsdetectorresult.DefaultCreatedAt()
		ddrc.mutation.SetCreatedAt(v)
	}
	if _, ok := ddrc.mutation.UpdatedAt(); !ok {
		v := dnsdetectorresult.DefaultUpdatedAt()
		ddrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ddrc *DNSDetectorResultCreate) check() error {
	if _, ok := ddrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New("ent: missing required field \"created_at\"")}
	}
	if _, ok := ddrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New("ent: missing required field \"updated_at\"")}
	}
	if _, ok := ddrc.mutation.Task(); !ok {
		return &ValidationError{Name: "task", err: errors.New("ent: missing required field \"task\"")}
	}
	if _, ok := ddrc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New("ent: missing required field \"result\"")}
	}
	if v, ok := ddrc.mutation.Result(); ok {
		if err := dnsdetectorresult.ResultValidator(v); err != nil {
			return &ValidationError{Name: "result", err: fmt.Errorf("ent: validator failed for field \"result\": %w", err)}
		}
	}
	if _, ok := ddrc.mutation.MaxDuration(); !ok {
		return &ValidationError{Name: "maxDuration", err: errors.New("ent: missing required field \"maxDuration\"")}
	}
	if _, ok := ddrc.mutation.Messages(); !ok {
		return &ValidationError{Name: "messages", err: errors.New("ent: missing required field \"messages\"")}
	}
	if _, ok := ddrc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New("ent: missing required field \"host\"")}
	}
	if _, ok := ddrc.mutation.Results(); !ok {
		return &ValidationError{Name: "results", err: errors.New("ent: missing required field \"results\"")}
	}
	return nil
}

func (ddrc *DNSDetectorResultCreate) sqlSave(ctx context.Context) (*DNSDetectorResult, error) {
	_node, _spec := ddrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ddrc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ddrc *DNSDetectorResultCreate) createSpec() (*DNSDetectorResult, *sqlgraph.CreateSpec) {
	var (
		_node = &DNSDetectorResult{config: ddrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: dnsdetectorresult.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dnsdetectorresult.FieldID,
			},
		}
	)
	if value, ok := ddrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsdetectorresult.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ddrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: dnsdetectorresult.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ddrc.mutation.Task(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldTask,
		})
		_node.Task = value
	}
	if value, ok := ddrc.mutation.Result(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: dnsdetectorresult.FieldResult,
		})
		_node.Result = value
	}
	if value, ok := ddrc.mutation.MaxDuration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: dnsdetectorresult.FieldMaxDuration,
		})
		_node.MaxDuration = value
	}
	if value, ok := ddrc.mutation.Messages(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetectorresult.FieldMessages,
		})
		_node.Messages = value
	}
	if value, ok := ddrc.mutation.Host(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: dnsdetectorresult.FieldHost,
		})
		_node.Host = value
	}
	if value, ok := ddrc.mutation.Results(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: dnsdetectorresult.FieldResults,
		})
		_node.Results = value
	}
	return _node, _spec
}

// DNSDetectorResultCreateBulk is the builder for creating many DNSDetectorResult entities in bulk.
type DNSDetectorResultCreateBulk struct {
	config
	builders []*DNSDetectorResultCreate
}

// Save creates the DNSDetectorResult entities in the database.
func (ddrcb *DNSDetectorResultCreateBulk) Save(ctx context.Context) ([]*DNSDetectorResult, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ddrcb.builders))
	nodes := make([]*DNSDetectorResult, len(ddrcb.builders))
	mutators := make([]Mutator, len(ddrcb.builders))
	for i := range ddrcb.builders {
		func(i int, root context.Context) {
			builder := ddrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DNSDetectorResultMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ddrcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ddrcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ddrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ddrcb *DNSDetectorResultCreateBulk) SaveX(ctx context.Context) []*DNSDetectorResult {
	v, err := ddrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
