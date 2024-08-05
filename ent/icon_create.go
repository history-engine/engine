// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent/icon"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IconCreate is the builder for creating a Icon entity.
type IconCreate struct {
	config
	mutation *IconMutation
	hooks    []Hook
}

// SetHost sets the "host" field.
func (ic *IconCreate) SetHost(s string) *IconCreate {
	ic.mutation.SetHost(s)
	return ic
}

// SetPath sets the "path" field.
func (ic *IconCreate) SetPath(s string) *IconCreate {
	ic.mutation.SetPath(s)
	return ic
}

// SetCreatedAt sets the "created_at" field.
func (ic *IconCreate) SetCreatedAt(t time.Time) *IconCreate {
	ic.mutation.SetCreatedAt(t)
	return ic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ic *IconCreate) SetNillableCreatedAt(t *time.Time) *IconCreate {
	if t != nil {
		ic.SetCreatedAt(*t)
	}
	return ic
}

// SetUpdatedAt sets the "updated_at" field.
func (ic *IconCreate) SetUpdatedAt(t time.Time) *IconCreate {
	ic.mutation.SetUpdatedAt(t)
	return ic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ic *IconCreate) SetNillableUpdatedAt(t *time.Time) *IconCreate {
	if t != nil {
		ic.SetUpdatedAt(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *IconCreate) SetID(i int64) *IconCreate {
	ic.mutation.SetID(i)
	return ic
}

// Mutation returns the IconMutation object of the builder.
func (ic *IconCreate) Mutation() *IconMutation {
	return ic.mutation
}

// Save creates the Icon in the database.
func (ic *IconCreate) Save(ctx context.Context) (*Icon, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IconCreate) SaveX(ctx context.Context) *Icon {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IconCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IconCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IconCreate) defaults() {
	if _, ok := ic.mutation.CreatedAt(); !ok {
		v := icon.DefaultCreatedAt()
		ic.mutation.SetCreatedAt(v)
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		v := icon.DefaultUpdatedAt()
		ic.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IconCreate) check() error {
	if _, ok := ic.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "Icon.host"`)}
	}
	if v, ok := ic.mutation.Host(); ok {
		if err := icon.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Icon.host": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "Icon.path"`)}
	}
	if v, ok := ic.mutation.Path(); ok {
		if err := icon.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "Icon.path": %w`, err)}
		}
	}
	if _, ok := ic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Icon.created_at"`)}
	}
	if _, ok := ic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Icon.updated_at"`)}
	}
	return nil
}

func (ic *IconCreate) sqlSave(ctx context.Context) (*Icon, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IconCreate) createSpec() (*Icon, *sqlgraph.CreateSpec) {
	var (
		_node = &Icon{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(icon.Table, sqlgraph.NewFieldSpec(icon.FieldID, field.TypeInt64))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ic.mutation.Host(); ok {
		_spec.SetField(icon.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := ic.mutation.Path(); ok {
		_spec.SetField(icon.FieldPath, field.TypeString, value)
		_node.Path = value
	}
	if value, ok := ic.mutation.CreatedAt(); ok {
		_spec.SetField(icon.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ic.mutation.UpdatedAt(); ok {
		_spec.SetField(icon.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// IconCreateBulk is the builder for creating many Icon entities in bulk.
type IconCreateBulk struct {
	config
	err      error
	builders []*IconCreate
}

// Save creates the Icon entities in the database.
func (icb *IconCreateBulk) Save(ctx context.Context) ([]*Icon, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Icon, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IconMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IconCreateBulk) SaveX(ctx context.Context) []*Icon {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IconCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IconCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}