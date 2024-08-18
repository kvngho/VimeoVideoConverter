// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/predicate"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/productreview"
)

// ProductReviewDelete is the builder for deleting a ProductReview entity.
type ProductReviewDelete struct {
	config
	hooks    []Hook
	mutation *ProductReviewMutation
}

// Where appends a list predicates to the ProductReviewDelete builder.
func (prd *ProductReviewDelete) Where(ps ...predicate.ProductReview) *ProductReviewDelete {
	prd.mutation.Where(ps...)
	return prd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (prd *ProductReviewDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, prd.sqlExec, prd.mutation, prd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (prd *ProductReviewDelete) ExecX(ctx context.Context) int {
	n, err := prd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (prd *ProductReviewDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(productreview.Table, sqlgraph.NewFieldSpec(productreview.FieldID, field.TypeInt))
	if ps := prd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, prd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	prd.mutation.done = true
	return affected, err
}

// ProductReviewDeleteOne is the builder for deleting a single ProductReview entity.
type ProductReviewDeleteOne struct {
	prd *ProductReviewDelete
}

// Where appends a list predicates to the ProductReviewDelete builder.
func (prdo *ProductReviewDeleteOne) Where(ps ...predicate.ProductReview) *ProductReviewDeleteOne {
	prdo.prd.mutation.Where(ps...)
	return prdo
}

// Exec executes the deletion query.
func (prdo *ProductReviewDeleteOne) Exec(ctx context.Context) error {
	n, err := prdo.prd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{productreview.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (prdo *ProductReviewDeleteOne) ExecX(ctx context.Context) {
	if err := prdo.Exec(ctx); err != nil {
		panic(err)
	}
}
