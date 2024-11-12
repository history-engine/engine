// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent/predicate"
	"history-engine/engine/ent/setting"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SettingUpdate is the builder for updating Setting entities.
type SettingUpdate struct {
	config
	hooks    []Hook
	mutation *SettingMutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (su *SettingUpdate) Where(ps ...predicate.Setting) *SettingUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetUserID sets the "user_id" field.
func (su *SettingUpdate) SetUserID(i int64) *SettingUpdate {
	su.mutation.ResetUserID()
	su.mutation.SetUserID(i)
	return su
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (su *SettingUpdate) SetNillableUserID(i *int64) *SettingUpdate {
	if i != nil {
		su.SetUserID(*i)
	}
	return su
}

// AddUserID adds i to the "user_id" field.
func (su *SettingUpdate) AddUserID(i int64) *SettingUpdate {
	su.mutation.AddUserID(i)
	return su
}

// SetMaxVersion sets the "max_version" field.
func (su *SettingUpdate) SetMaxVersion(i int) *SettingUpdate {
	su.mutation.ResetMaxVersion()
	su.mutation.SetMaxVersion(i)
	return su
}

// SetNillableMaxVersion sets the "max_version" field if the given value is not nil.
func (su *SettingUpdate) SetNillableMaxVersion(i *int) *SettingUpdate {
	if i != nil {
		su.SetMaxVersion(*i)
	}
	return su
}

// AddMaxVersion adds i to the "max_version" field.
func (su *SettingUpdate) AddMaxVersion(i int) *SettingUpdate {
	su.mutation.AddMaxVersion(i)
	return su
}

// SetMinVersionInterval sets the "min_version_interval" field.
func (su *SettingUpdate) SetMinVersionInterval(i int) *SettingUpdate {
	su.mutation.ResetMinVersionInterval()
	su.mutation.SetMinVersionInterval(i)
	return su
}

// SetNillableMinVersionInterval sets the "min_version_interval" field if the given value is not nil.
func (su *SettingUpdate) SetNillableMinVersionInterval(i *int) *SettingUpdate {
	if i != nil {
		su.SetMinVersionInterval(*i)
	}
	return su
}

// AddMinVersionInterval adds i to the "min_version_interval" field.
func (su *SettingUpdate) AddMinVersionInterval(i int) *SettingUpdate {
	su.mutation.AddMinVersionInterval(i)
	return su
}

// SetMinSize sets the "min_size" field.
func (su *SettingUpdate) SetMinSize(i int) *SettingUpdate {
	su.mutation.ResetMinSize()
	su.mutation.SetMinSize(i)
	return su
}

// SetNillableMinSize sets the "min_size" field if the given value is not nil.
func (su *SettingUpdate) SetNillableMinSize(i *int) *SettingUpdate {
	if i != nil {
		su.SetMinSize(*i)
	}
	return su
}

// AddMinSize adds i to the "min_size" field.
func (su *SettingUpdate) AddMinSize(i int) *SettingUpdate {
	su.mutation.AddMinSize(i)
	return su
}

// SetMaxSize sets the "max_size" field.
func (su *SettingUpdate) SetMaxSize(i int) *SettingUpdate {
	su.mutation.ResetMaxSize()
	su.mutation.SetMaxSize(i)
	return su
}

// SetNillableMaxSize sets the "max_size" field if the given value is not nil.
func (su *SettingUpdate) SetNillableMaxSize(i *int) *SettingUpdate {
	if i != nil {
		su.SetMaxSize(*i)
	}
	return su
}

// AddMaxSize adds i to the "max_size" field.
func (su *SettingUpdate) AddMaxSize(i int) *SettingUpdate {
	su.mutation.AddMaxSize(i)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SettingUpdate) SetUpdatedAt(t time.Time) *SettingUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// Mutation returns the SettingMutation object of the builder.
func (su *SettingUpdate) Mutation() *SettingMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SettingUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SettingUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SettingUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SettingUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SettingUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := setting.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

func (su *SettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.UserID(); ok {
		_spec.SetField(setting.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedUserID(); ok {
		_spec.AddField(setting.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := su.mutation.MaxVersion(); ok {
		_spec.SetField(setting.FieldMaxVersion, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedMaxVersion(); ok {
		_spec.AddField(setting.FieldMaxVersion, field.TypeInt, value)
	}
	if value, ok := su.mutation.MinVersionInterval(); ok {
		_spec.SetField(setting.FieldMinVersionInterval, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedMinVersionInterval(); ok {
		_spec.AddField(setting.FieldMinVersionInterval, field.TypeInt, value)
	}
	if value, ok := su.mutation.MinSize(); ok {
		_spec.SetField(setting.FieldMinSize, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedMinSize(); ok {
		_spec.AddField(setting.FieldMinSize, field.TypeInt, value)
	}
	if value, ok := su.mutation.MaxSize(); ok {
		_spec.SetField(setting.FieldMaxSize, field.TypeInt, value)
	}
	if value, ok := su.mutation.AddedMaxSize(); ok {
		_spec.AddField(setting.FieldMaxSize, field.TypeInt, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(setting.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SettingUpdateOne is the builder for updating a single Setting entity.
type SettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SettingMutation
}

// SetUserID sets the "user_id" field.
func (suo *SettingUpdateOne) SetUserID(i int64) *SettingUpdateOne {
	suo.mutation.ResetUserID()
	suo.mutation.SetUserID(i)
	return suo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableUserID(i *int64) *SettingUpdateOne {
	if i != nil {
		suo.SetUserID(*i)
	}
	return suo
}

// AddUserID adds i to the "user_id" field.
func (suo *SettingUpdateOne) AddUserID(i int64) *SettingUpdateOne {
	suo.mutation.AddUserID(i)
	return suo
}

// SetMaxVersion sets the "max_version" field.
func (suo *SettingUpdateOne) SetMaxVersion(i int) *SettingUpdateOne {
	suo.mutation.ResetMaxVersion()
	suo.mutation.SetMaxVersion(i)
	return suo
}

// SetNillableMaxVersion sets the "max_version" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableMaxVersion(i *int) *SettingUpdateOne {
	if i != nil {
		suo.SetMaxVersion(*i)
	}
	return suo
}

// AddMaxVersion adds i to the "max_version" field.
func (suo *SettingUpdateOne) AddMaxVersion(i int) *SettingUpdateOne {
	suo.mutation.AddMaxVersion(i)
	return suo
}

// SetMinVersionInterval sets the "min_version_interval" field.
func (suo *SettingUpdateOne) SetMinVersionInterval(i int) *SettingUpdateOne {
	suo.mutation.ResetMinVersionInterval()
	suo.mutation.SetMinVersionInterval(i)
	return suo
}

// SetNillableMinVersionInterval sets the "min_version_interval" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableMinVersionInterval(i *int) *SettingUpdateOne {
	if i != nil {
		suo.SetMinVersionInterval(*i)
	}
	return suo
}

// AddMinVersionInterval adds i to the "min_version_interval" field.
func (suo *SettingUpdateOne) AddMinVersionInterval(i int) *SettingUpdateOne {
	suo.mutation.AddMinVersionInterval(i)
	return suo
}

// SetMinSize sets the "min_size" field.
func (suo *SettingUpdateOne) SetMinSize(i int) *SettingUpdateOne {
	suo.mutation.ResetMinSize()
	suo.mutation.SetMinSize(i)
	return suo
}

// SetNillableMinSize sets the "min_size" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableMinSize(i *int) *SettingUpdateOne {
	if i != nil {
		suo.SetMinSize(*i)
	}
	return suo
}

// AddMinSize adds i to the "min_size" field.
func (suo *SettingUpdateOne) AddMinSize(i int) *SettingUpdateOne {
	suo.mutation.AddMinSize(i)
	return suo
}

// SetMaxSize sets the "max_size" field.
func (suo *SettingUpdateOne) SetMaxSize(i int) *SettingUpdateOne {
	suo.mutation.ResetMaxSize()
	suo.mutation.SetMaxSize(i)
	return suo
}

// SetNillableMaxSize sets the "max_size" field if the given value is not nil.
func (suo *SettingUpdateOne) SetNillableMaxSize(i *int) *SettingUpdateOne {
	if i != nil {
		suo.SetMaxSize(*i)
	}
	return suo
}

// AddMaxSize adds i to the "max_size" field.
func (suo *SettingUpdateOne) AddMaxSize(i int) *SettingUpdateOne {
	suo.mutation.AddMaxSize(i)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SettingUpdateOne) SetUpdatedAt(t time.Time) *SettingUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// Mutation returns the SettingMutation object of the builder.
func (suo *SettingUpdateOne) Mutation() *SettingMutation {
	return suo.mutation
}

// Where appends a list predicates to the SettingUpdate builder.
func (suo *SettingUpdateOne) Where(ps ...predicate.Setting) *SettingUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SettingUpdateOne) Select(field string, fields ...string) *SettingUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Setting entity.
func (suo *SettingUpdateOne) Save(ctx context.Context) (*Setting, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SettingUpdateOne) SaveX(ctx context.Context) *Setting {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SettingUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SettingUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SettingUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := setting.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

func (suo *SettingUpdateOne) sqlSave(ctx context.Context) (_node *Setting, err error) {
	_spec := sqlgraph.NewUpdateSpec(setting.Table, setting.Columns, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Setting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, setting.FieldID)
		for _, f := range fields {
			if !setting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != setting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.UserID(); ok {
		_spec.SetField(setting.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedUserID(); ok {
		_spec.AddField(setting.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.MaxVersion(); ok {
		_spec.SetField(setting.FieldMaxVersion, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedMaxVersion(); ok {
		_spec.AddField(setting.FieldMaxVersion, field.TypeInt, value)
	}
	if value, ok := suo.mutation.MinVersionInterval(); ok {
		_spec.SetField(setting.FieldMinVersionInterval, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedMinVersionInterval(); ok {
		_spec.AddField(setting.FieldMinVersionInterval, field.TypeInt, value)
	}
	if value, ok := suo.mutation.MinSize(); ok {
		_spec.SetField(setting.FieldMinSize, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedMinSize(); ok {
		_spec.AddField(setting.FieldMinSize, field.TypeInt, value)
	}
	if value, ok := suo.mutation.MaxSize(); ok {
		_spec.SetField(setting.FieldMaxSize, field.TypeInt, value)
	}
	if value, ok := suo.mutation.AddedMaxSize(); ok {
		_spec.AddField(setting.FieldMaxSize, field.TypeInt, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(setting.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Setting{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{setting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
