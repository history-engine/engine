// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent/filetype"
	"history-engine/engine/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetUserID sets the "user_id" field.
func (ftu *FileTypeUpdate) SetUserID(i int64) *FileTypeUpdate {
	ftu.mutation.ResetUserID()
	ftu.mutation.SetUserID(i)
	return ftu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableUserID(i *int64) *FileTypeUpdate {
	if i != nil {
		ftu.SetUserID(*i)
	}
	return ftu
}

// AddUserID adds i to the "user_id" field.
func (ftu *FileTypeUpdate) AddUserID(i int64) *FileTypeUpdate {
	ftu.mutation.AddUserID(i)
	return ftu
}

// SetSuffix sets the "suffix" field.
func (ftu *FileTypeUpdate) SetSuffix(s string) *FileTypeUpdate {
	ftu.mutation.SetSuffix(s)
	return ftu
}

// SetNillableSuffix sets the "suffix" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableSuffix(s *string) *FileTypeUpdate {
	if s != nil {
		ftu.SetSuffix(*s)
	}
	return ftu
}

// SetType sets the "type" field.
func (ftu *FileTypeUpdate) SetType(i int) *FileTypeUpdate {
	ftu.mutation.ResetType()
	ftu.mutation.SetType(i)
	return ftu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableType(i *int) *FileTypeUpdate {
	if i != nil {
		ftu.SetType(*i)
	}
	return ftu
}

// AddType adds i to the "type" field.
func (ftu *FileTypeUpdate) AddType(i int) *FileTypeUpdate {
	ftu.mutation.AddType(i)
	return ftu
}

// SetUpdatedAt sets the "updated_at" field.
func (ftu *FileTypeUpdate) SetUpdatedAt(t time.Time) *FileTypeUpdate {
	ftu.mutation.SetUpdatedAt(t)
	return ftu
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftu *FileTypeUpdate) Mutation() *FileTypeMutation {
	return ftu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {
	ftu.defaults()
	return withHooks(ctx, ftu.sqlSave, ftu.mutation, ftu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftu *FileTypeUpdate) defaults() {
	if _, ok := ftu.mutation.UpdatedAt(); !ok {
		v := filetype.UpdateDefaultUpdatedAt()
		ftu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FileTypeUpdate) check() error {
	if v, ok := ftu.mutation.Suffix(); ok {
		if err := filetype.SuffixValidator(v); err != nil {
			return &ValidationError{Name: "suffix", err: fmt.Errorf(`ent: validator failed for field "FileType.suffix": %w`, err)}
		}
	}
	if v, ok := ftu.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	return nil
}

func (ftu *FileTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ftu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(filetype.Table, filetype.Columns, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt64))
	if ps := ftu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftu.mutation.UserID(); ok {
		_spec.SetField(filetype.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ftu.mutation.AddedUserID(); ok {
		_spec.AddField(filetype.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ftu.mutation.Suffix(); ok {
		_spec.SetField(filetype.FieldSuffix, field.TypeString, value)
	}
	if value, ok := ftu.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeInt, value)
	}
	if value, ok := ftu.mutation.AddedType(); ok {
		_spec.AddField(filetype.FieldType, field.TypeInt, value)
	}
	if value, ok := ftu.mutation.UpdatedAt(); ok {
		_spec.SetField(filetype.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ftu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ftu.mutation.done = true
	return n, nil
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileTypeMutation
}

// SetUserID sets the "user_id" field.
func (ftuo *FileTypeUpdateOne) SetUserID(i int64) *FileTypeUpdateOne {
	ftuo.mutation.ResetUserID()
	ftuo.mutation.SetUserID(i)
	return ftuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableUserID(i *int64) *FileTypeUpdateOne {
	if i != nil {
		ftuo.SetUserID(*i)
	}
	return ftuo
}

// AddUserID adds i to the "user_id" field.
func (ftuo *FileTypeUpdateOne) AddUserID(i int64) *FileTypeUpdateOne {
	ftuo.mutation.AddUserID(i)
	return ftuo
}

// SetSuffix sets the "suffix" field.
func (ftuo *FileTypeUpdateOne) SetSuffix(s string) *FileTypeUpdateOne {
	ftuo.mutation.SetSuffix(s)
	return ftuo
}

// SetNillableSuffix sets the "suffix" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableSuffix(s *string) *FileTypeUpdateOne {
	if s != nil {
		ftuo.SetSuffix(*s)
	}
	return ftuo
}

// SetType sets the "type" field.
func (ftuo *FileTypeUpdateOne) SetType(i int) *FileTypeUpdateOne {
	ftuo.mutation.ResetType()
	ftuo.mutation.SetType(i)
	return ftuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableType(i *int) *FileTypeUpdateOne {
	if i != nil {
		ftuo.SetType(*i)
	}
	return ftuo
}

// AddType adds i to the "type" field.
func (ftuo *FileTypeUpdateOne) AddType(i int) *FileTypeUpdateOne {
	ftuo.mutation.AddType(i)
	return ftuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ftuo *FileTypeUpdateOne) SetUpdatedAt(t time.Time) *FileTypeUpdateOne {
	ftuo.mutation.SetUpdatedAt(t)
	return ftuo
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftuo *FileTypeUpdateOne) Mutation() *FileTypeMutation {
	return ftuo.mutation
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (ftuo *FileTypeUpdateOne) Where(ps ...predicate.FileType) *FileTypeUpdateOne {
	ftuo.mutation.Where(ps...)
	return ftuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FileTypeUpdateOne) Select(field string, fields ...string) *FileTypeUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FileType entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {
	ftuo.defaults()
	return withHooks(ctx, ftuo.sqlSave, ftuo.mutation, ftuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftuo *FileTypeUpdateOne) defaults() {
	if _, ok := ftuo.mutation.UpdatedAt(); !ok {
		v := filetype.UpdateDefaultUpdatedAt()
		ftuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FileTypeUpdateOne) check() error {
	if v, ok := ftuo.mutation.Suffix(); ok {
		if err := filetype.SuffixValidator(v); err != nil {
			return &ValidationError{Name: "suffix", err: fmt.Errorf(`ent: validator failed for field "FileType.suffix": %w`, err)}
		}
	}
	if v, ok := ftuo.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	return nil
}

func (ftuo *FileTypeUpdateOne) sqlSave(ctx context.Context) (_node *FileType, err error) {
	if err := ftuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(filetype.Table, filetype.Columns, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt64))
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ftuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetype.FieldID)
		for _, f := range fields {
			if !filetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != filetype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftuo.mutation.UserID(); ok {
		_spec.SetField(filetype.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ftuo.mutation.AddedUserID(); ok {
		_spec.AddField(filetype.FieldUserID, field.TypeInt64, value)
	}
	if value, ok := ftuo.mutation.Suffix(); ok {
		_spec.SetField(filetype.FieldSuffix, field.TypeString, value)
	}
	if value, ok := ftuo.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeInt, value)
	}
	if value, ok := ftuo.mutation.AddedType(); ok {
		_spec.AddField(filetype.FieldType, field.TypeInt, value)
	}
	if value, ok := ftuo.mutation.UpdatedAt(); ok {
		_spec.SetField(filetype.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &FileType{config: ftuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ftuo.mutation.done = true
	return _node, nil
}
