// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"nevissGo/ent/pixel"
	"nevissGo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PixelDelete is the builder for deleting a Pixel entity.
type PixelDelete struct {
	config
	hooks    []Hook
	mutation *PixelMutation
}

// Where appends a list predicates to the PixelDelete builder.
func (pd *PixelDelete) Where(ps ...predicate.Pixel) *PixelDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PixelDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, pd.sqlExec, pd.mutation, pd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PixelDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PixelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(pixel.Table, sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt))
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	pd.mutation.done = true
	return affected, err
}

// PixelDeleteOne is the builder for deleting a single Pixel entity.
type PixelDeleteOne struct {
	pd *PixelDelete
}

// Where appends a list predicates to the PixelDelete builder.
func (pdo *PixelDeleteOne) Where(ps ...predicate.Pixel) *PixelDeleteOne {
	pdo.pd.mutation.Where(ps...)
	return pdo
}

// Exec executes the deletion query.
func (pdo *PixelDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pixel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PixelDeleteOne) ExecX(ctx context.Context) {
	if err := pdo.Exec(ctx); err != nil {
		panic(err)
	}
}