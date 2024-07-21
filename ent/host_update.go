// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent/host"
	"history-engine/engine/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HostUpdate is the builder for updating Host entities.
type HostUpdate struct {
	config
	hooks    []Hook
	mutation *HostMutation
}

// Where appends a list predicates to the HostUpdate builder.
func (hu *HostUpdate) Where(ps ...predicate.Host) *HostUpdate {
	hu.mutation.Where(ps...)
	return hu
}

// SetUserID sets the "user_id" field.
func (hu *HostUpdate) SetUserID(i int64) *HostUpdate {
	hu.mutation.ResetUserID()
	hu.mutation.SetUserID(i)
	return hu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (hu *HostUpdate) SetNillableUserID(i *int64) *HostUpdate {
	if i != nil {
		hu.SetUserID(*i)
	}
	return hu
}

// AddUserID adds i to the "user_id" field.
func (hu *HostUpdate) AddUserID(i int64) *HostUpdate {
	hu.mutation.AddUserID(i)
	return hu
}

// SetHost sets the "host" field.
func (hu *HostUpdate) SetHost(s string) *HostUpdate {
	hu.mutation.SetHost(s)
	return hu
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (hu *HostUpdate) SetNillableHost(s *string) *HostUpdate {
	if s != nil {
		hu.SetHost(*s)
	}
	return hu
}

// SetType sets the "type" field.
func (hu *HostUpdate) SetType(i int) *HostUpdate {
	hu.mutation.ResetType()
	hu.mutation.SetType(i)
	return hu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (hu *HostUpdate) SetNillableType(i *int) *HostUpdate {
	if i != nil {
		hu.SetType(*i)
	}
	return hu
}

// AddType adds i to the "type" field.
func (hu *HostUpdate) AddType(i int) *HostUpdate {
	hu.mutation.AddType(i)
	return hu
}

// SetUpdatedAt sets the "updated_at" field.
func (hu *HostUpdate) SetUpdatedAt(t time.Time) *HostUpdate {
	hu.mutation.SetUpdatedAt(t)
	return hu
}

// Mutation returns the HostMutation object of the builder.
func (hu *HostUpdate) Mutation() *HostMutation {
	return hu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hu *HostUpdate) Save(ctx context.Context) (int, error) {
	hu.defaults()
	return withHooks(ctx, hu.sqlSave, hu.mutation, hu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hu *HostUpdate) SaveX(ctx context.Context) int {
	affected, err := hu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hu *HostUpdate) Exec(ctx context.Context) error {
	_, err := hu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hu *HostUpdate) ExecX(ctx context.Context) {
	if err := hu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hu *HostUpdate) defaults() {
	if _, ok := hu.mutation.UpdatedAt(); !ok {
		v := host.UpdateDefaultUpdatedAt()
		hu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hu *HostUpdate) check() error {
	if v, ok := hu.mutation.Host(); ok {
		if err := host.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Host.host": %w`, err)}
		}
	}
	if v, ok := hu.mutation.GetType(); ok {
		if err := host.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Host.type": %w`, err)}
		}
	}
	return nil
}

func (hu *HostUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := hu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt64))
	if ps := hu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hu.mutation.UserID(); ok {
		_spec.SetField(host.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := hu.mutation.AddedUserID(); ok {
		_spec.AddField(host.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := hu.mutation.Host(); ok {
		_spec.SetField(host.FieldHost, field.TypeString, value)
	}
	if value, ok := hu.mutation.GetType(); ok {
		_spec.SetField(host.FieldType, field.TypeInt, value)
	}
	if value, ok := hu.mutation.AddedType(); ok {
		_spec.AddField(host.FieldType, field.TypeInt, value)
	}
	if value, ok := hu.mutation.UpdatedAt(); ok {
		_spec.SetField(host.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hu.mutation.done = true
	return n, nil
}

// HostUpdateOne is the builder for updating a single Host entity.
type HostUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HostMutation
}

// SetUserID sets the "user_id" field.
func (huo *HostUpdateOne) SetUserID(i int64) *HostUpdateOne {
	huo.mutation.ResetUserID()
	huo.mutation.SetUserID(i)
	return huo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableUserID(i *int64) *HostUpdateOne {
	if i != nil {
		huo.SetUserID(*i)
	}
	return huo
}

// AddUserID adds i to the "user_id" field.
func (huo *HostUpdateOne) AddUserID(i int64) *HostUpdateOne {
	huo.mutation.AddUserID(i)
	return huo
}

// SetHost sets the "host" field.
func (huo *HostUpdateOne) SetHost(s string) *HostUpdateOne {
	huo.mutation.SetHost(s)
	return huo
}

// SetNillableHost sets the "host" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableHost(s *string) *HostUpdateOne {
	if s != nil {
		huo.SetHost(*s)
	}
	return huo
}

// SetType sets the "type" field.
func (huo *HostUpdateOne) SetType(i int) *HostUpdateOne {
	huo.mutation.ResetType()
	huo.mutation.SetType(i)
	return huo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (huo *HostUpdateOne) SetNillableType(i *int) *HostUpdateOne {
	if i != nil {
		huo.SetType(*i)
	}
	return huo
}

// AddType adds i to the "type" field.
func (huo *HostUpdateOne) AddType(i int) *HostUpdateOne {
	huo.mutation.AddType(i)
	return huo
}

// SetUpdatedAt sets the "updated_at" field.
func (huo *HostUpdateOne) SetUpdatedAt(t time.Time) *HostUpdateOne {
	huo.mutation.SetUpdatedAt(t)
	return huo
}

// Mutation returns the HostMutation object of the builder.
func (huo *HostUpdateOne) Mutation() *HostMutation {
	return huo.mutation
}

// Where appends a list predicates to the HostUpdate builder.
func (huo *HostUpdateOne) Where(ps ...predicate.Host) *HostUpdateOne {
	huo.mutation.Where(ps...)
	return huo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (huo *HostUpdateOne) Select(field string, fields ...string) *HostUpdateOne {
	huo.fields = append([]string{field}, fields...)
	return huo
}

// Save executes the query and returns the updated Host entity.
func (huo *HostUpdateOne) Save(ctx context.Context) (*Host, error) {
	huo.defaults()
	return withHooks(ctx, huo.sqlSave, huo.mutation, huo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (huo *HostUpdateOne) SaveX(ctx context.Context) *Host {
	node, err := huo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (huo *HostUpdateOne) Exec(ctx context.Context) error {
	_, err := huo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (huo *HostUpdateOne) ExecX(ctx context.Context) {
	if err := huo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (huo *HostUpdateOne) defaults() {
	if _, ok := huo.mutation.UpdatedAt(); !ok {
		v := host.UpdateDefaultUpdatedAt()
		huo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (huo *HostUpdateOne) check() error {
	if v, ok := huo.mutation.Host(); ok {
		if err := host.HostValidator(v); err != nil {
			return &ValidationError{Name: "host", err: fmt.Errorf(`ent: validator failed for field "Host.host": %w`, err)}
		}
	}
	if v, ok := huo.mutation.GetType(); ok {
		if err := host.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Host.type": %w`, err)}
		}
	}
	return nil
}

func (huo *HostUpdateOne) sqlSave(ctx context.Context) (_node *Host, err error) {
	if err := huo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(host.Table, host.Columns, sqlgraph.NewFieldSpec(host.FieldID, field.TypeInt64))
	id, ok := huo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Host.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := huo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, host.FieldID)
		for _, f := range fields {
			if !host.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != host.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := huo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := huo.mutation.UserID(); ok {
		_spec.SetField(host.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := huo.mutation.AddedUserID(); ok {
		_spec.AddField(host.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := huo.mutation.Host(); ok {
		_spec.SetField(host.FieldHost, field.TypeString, value)
	}
	if value, ok := huo.mutation.GetType(); ok {
		_spec.SetField(host.FieldType, field.TypeInt, value)
	}
	if value, ok := huo.mutation.AddedType(); ok {
		_spec.AddField(host.FieldType, field.TypeInt, value)
	}
	if value, ok := huo.mutation.UpdatedAt(); ok {
		_spec.SetField(host.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Host{config: huo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, huo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{host.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	huo.mutation.done = true
	return _node, nil
}
