// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"history-engine/engine/ent/filetype"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	mutation *FileTypeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (ftc *FileTypeCreate) SetUserID(i int64) *FileTypeCreate {
	ftc.mutation.SetUserID(i)
	return ftc
}

// SetSuffix sets the "suffix" field.
func (ftc *FileTypeCreate) SetSuffix(s string) *FileTypeCreate {
	ftc.mutation.SetSuffix(s)
	return ftc
}

// SetType sets the "type" field.
func (ftc *FileTypeCreate) SetType(i int) *FileTypeCreate {
	ftc.mutation.SetType(i)
	return ftc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableType(i *int) *FileTypeCreate {
	if i != nil {
		ftc.SetType(*i)
	}
	return ftc
}

// SetCreatedAt sets the "created_at" field.
func (ftc *FileTypeCreate) SetCreatedAt(t time.Time) *FileTypeCreate {
	ftc.mutation.SetCreatedAt(t)
	return ftc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableCreatedAt(t *time.Time) *FileTypeCreate {
	if t != nil {
		ftc.SetCreatedAt(*t)
	}
	return ftc
}

// SetUpdatedAt sets the "updated_at" field.
func (ftc *FileTypeCreate) SetUpdatedAt(t time.Time) *FileTypeCreate {
	ftc.mutation.SetUpdatedAt(t)
	return ftc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ftc *FileTypeCreate) SetNillableUpdatedAt(t *time.Time) *FileTypeCreate {
	if t != nil {
		ftc.SetUpdatedAt(*t)
	}
	return ftc
}

// SetID sets the "id" field.
func (ftc *FileTypeCreate) SetID(i int64) *FileTypeCreate {
	ftc.mutation.SetID(i)
	return ftc
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftc *FileTypeCreate) Mutation() *FileTypeMutation {
	return ftc.mutation
}

// Save creates the FileType in the database.
func (ftc *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	ftc.defaults()
	return withHooks(ctx, ftc.sqlSave, ftc.mutation, ftc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftc *FileTypeCreate) Exec(ctx context.Context) error {
	_, err := ftc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftc *FileTypeCreate) ExecX(ctx context.Context) {
	if err := ftc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftc *FileTypeCreate) defaults() {
	if _, ok := ftc.mutation.GetType(); !ok {
		v := filetype.DefaultType
		ftc.mutation.SetType(v)
	}
	if _, ok := ftc.mutation.CreatedAt(); !ok {
		v := filetype.DefaultCreatedAt()
		ftc.mutation.SetCreatedAt(v)
	}
	if _, ok := ftc.mutation.UpdatedAt(); !ok {
		v := filetype.DefaultUpdatedAt()
		ftc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftc *FileTypeCreate) check() error {
	if _, ok := ftc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "FileType.user_id"`)}
	}
	if _, ok := ftc.mutation.Suffix(); !ok {
		return &ValidationError{Name: "suffix", err: errors.New(`ent: missing required field "FileType.suffix"`)}
	}
	if v, ok := ftc.mutation.Suffix(); ok {
		if err := filetype.SuffixValidator(v); err != nil {
			return &ValidationError{Name: "suffix", err: fmt.Errorf(`ent: validator failed for field "FileType.suffix": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "FileType.type"`)}
	}
	if v, ok := ftc.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if _, ok := ftc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FileType.created_at"`)}
	}
	if _, ok := ftc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FileType.updated_at"`)}
	}
	return nil
}

func (ftc *FileTypeCreate) sqlSave(ctx context.Context) (*FileType, error) {
	if err := ftc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ftc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ftc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ftc.mutation.id = &_node.ID
	ftc.mutation.done = true
	return _node, nil
}

func (ftc *FileTypeCreate) createSpec() (*FileType, *sqlgraph.CreateSpec) {
	var (
		_node = &FileType{config: ftc.config}
		_spec = sqlgraph.NewCreateSpec(filetype.Table, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt64))
	)
	if id, ok := ftc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ftc.mutation.UserID(); ok {
		_spec.SetField(filetype.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := ftc.mutation.Suffix(); ok {
		_spec.SetField(filetype.FieldSuffix, field.TypeString, value)
		_node.Suffix = value
	}
	if value, ok := ftc.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeInt, value)
		_node.Type = value
	}
	if value, ok := ftc.mutation.CreatedAt(); ok {
		_spec.SetField(filetype.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ftc.mutation.UpdatedAt(); ok {
		_spec.SetField(filetype.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// FileTypeCreateBulk is the builder for creating many FileType entities in bulk.
type FileTypeCreateBulk struct {
	config
	err      error
	builders []*FileTypeCreate
}

// Save creates the FileType entities in the database.
func (ftcb *FileTypeCreateBulk) Save(ctx context.Context) ([]*FileType, error) {
	if ftcb.err != nil {
		return nil, ftcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ftcb.builders))
	nodes := make([]*FileType, len(ftcb.builders))
	mutators := make([]Mutator, len(ftcb.builders))
	for i := range ftcb.builders {
		func(i int, root context.Context) {
			builder := ftcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileTypeMutation)
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
					_, err = mutators[i+1].Mutate(root, ftcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ftcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ftcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ftcb *FileTypeCreateBulk) SaveX(ctx context.Context) []*FileType {
	v, err := ftcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ftcb *FileTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ftcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftcb *FileTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ftcb.Exec(ctx); err != nil {
		panic(err)
	}
}
