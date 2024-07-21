// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent/host"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostCreate is the builder for creating a Host entity.
type HostCreate struct {
	config
	mutation *HostMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (hc *HostCreate) SetUserID(i int64) *HostCreate {
	hc.mutation.SetUserID(i)
	return hc
}

// SetHost sets the "host" field.
func (hc *HostCreate) SetHost(s string) *HostCreate {
	hc.mutation.SetHost(s)
	return hc
}

// SetType sets the "type" field.
func (hc *HostCreate) SetType(i int) *HostCreate {
	hc.mutation.SetType(i)
	return hc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (hc *HostCreate) SetNillableType(i *int) *HostCreate {
	if i != nil {
		hc.SetType(*i)
	}
	return hc
}

// SetCreatedAt sets the "created_at" field.
func (hc *HostCreate) SetCreatedAt(t time.Time) *HostCreate {
	hc.mutation.SetCreatedAt(t)
	return hc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (hc *HostCreate) SetNillableCreatedAt(t *time.Time) *HostCreate {
	if t != nil {
		hc.SetCreatedAt(*t)
	}
	return hc
}

// SetUpdatedAt sets the "updated_at" field.
func (hc *HostCreate) SetUpdatedAt(t time.Time) *HostCreate {
	hc.mutation.SetUpdatedAt(t)
	return hc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (hc *HostCreate) SetNillableUpdatedAt(t *time.Time) *HostCreate {
	if t != nil {
		hc.SetUpdatedAt(*t)
	}
	return hc
}

// SetID sets the "id" field.
func (hc *HostCreate) SetID(i int64) *HostCreate {
	hc.mutation.SetID(i)
	return hc
}

// Mutation returns the HostMutation object of the builder.
func (hc *HostCreate) Mutation() *HostMutation {
	return hc.mutation
}

// Save creates the Host in the database.
func (hc *HostCreate) Save(ctx context.Context) (*Host, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HostCreate) SaveX(ctx context.Context) *Host {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HostCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HostCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HostCreate) defaults() {
	if _, ok := hc.mutation.GetType(); !ok {
		v := host.DefaultType
		hc.mutation.SetType(v)
	}
	if _, ok := hc.mutation.CreatedAt(); !ok {
		v := host.DefaultCreatedAt()
		hc.mutation.SetCreatedAt(v)
	}
	if _, ok := hc.mutation.UpdatedAt(); !ok {
		v := host.DefaultUpdatedAt()
		hc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HostCreate) check() error {
	if _, ok := hc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Host.user_id"`)}
	}
	if _, ok := hc.mutation.Host(); !ok {
		return &ValidationError{Name: "host", err: errors.New(`ent: missing required field "Host.host"`)}
	}
	if v, ok := hc.mutation.Host(); ok {
		if err := host.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Host.host": %w`, err)}
		}
	}
	if _, ok := hc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Host.type"`)}
	}
	if v, ok := hc.mutation.GetType(); ok {
		if err := host.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Host.type": %w`, err)}
		}
	}
	if _, ok := hc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Host.created_at"`)}
	}
	if _, ok := hc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Host.updated_at"`)}
	}
	return nil
}

func (hc *HostCreate) sqlSave(ctx context.Context) (*Host, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HostCreate) createSpec() (*Host, *sqlgraph.CreateSpec) {
	var (
		_node = &Host{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(host.Table, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt64))
	)
	if id, ok := hc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := hc.mutation.UserID(); ok {
		_spec.SetField(host.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := hc.mutation.Host(); ok {
		_spec.SetField(host.FieldHost, field.TypeString, value)
		_node.Host = value
	}
	if value, ok := hc.mutation.GetType(); ok {
		_spec.SetField(host.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := hc.mutation.CreatedAt(); ok {
		_spec.SetField(host.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := hc.mutation.UpdatedAt(); ok {
		_spec.SetField(host.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// HostCreateBulk is the builder for creating many Host entities in bulk.
type HostCreateBulk struct {
	config
	err      error
	builders []*HostCreate
}

// Save creates the Host entities in the database.
func (hcb *HostCreateBulk) Save(ctx context.Context) ([]*Host, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Host, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HostMutation)
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
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HostCreateBulk) SaveX(ctx context.Context) []*Host {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HostCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HostCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}
