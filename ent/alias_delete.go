// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"history-engine/engine/ent/alias"
	"history-engine/engine/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AliasDelete is the builder for deleting a Alias entity.
type AliasDelete struct {
	config
	hooks    []Hook
	mutation *AliasMutation
}

// Where appends a list predicates to the AliasDelete builder.
func (ad *AliasDelete) Where(ps ...predicate.Alias) *AliasDelete {
	ad.mutation.Where(ps...)
	return ad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ad *AliasDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ad.sqlExec, ad.mutation, ad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ad *AliasDelete) ExecX(ctx context.Context) int {
	n, err := ad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ad *AliasDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(alias.Table, sqlgraph.NewFieldSpec(alias.FieldID, field.TypeInt64))
	if ps := ad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ad.mutation.done = true
	return affected, err
}

// AliasDeleteOne is the builder for deleting a single Alias entity.
type AliasDeleteOne struct {
	ad *AliasDelete
}

// Where appends a list predicates to the AliasDelete builder.
func (ado *AliasDeleteOne) Where(ps ...predicate.Alias) *AliasDeleteOne {
	ado.ad.mutation.Where(ps...)
	return ado
}

// Exec executes the deletion query.
func (ado *AliasDeleteOne) Exec(ctx context.Context) error {
	n, err := ado.ad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{alias.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ado *AliasDeleteOne) ExecX(ctx context.Context) {
	if err := ado.Exec(ctx); err != nil {
		panic(err)
	}
}