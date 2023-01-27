// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/concealedkan"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ConcealedKanDelete is the builder for deleting a ConcealedKan entity.
type ConcealedKanDelete struct {
	config
	hooks    []Hook
	mutation *ConcealedKanMutation
}

// Where appends a list predicates to the ConcealedKanDelete builder.
func (ckd *ConcealedKanDelete) Where(ps ...predicate.ConcealedKan) *ConcealedKanDelete {
	ckd.mutation.Where(ps...)
	return ckd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ckd *ConcealedKanDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, ConcealedKanMutation](ctx, ckd.sqlExec, ckd.mutation, ckd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ckd *ConcealedKanDelete) ExecX(ctx context.Context) int {
	n, err := ckd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ckd *ConcealedKanDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: concealedkan.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: concealedkan.FieldID,
			},
		},
	}
	if ps := ckd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ckd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ckd.mutation.done = true
	return affected, err
}

// ConcealedKanDeleteOne is the builder for deleting a single ConcealedKan entity.
type ConcealedKanDeleteOne struct {
	ckd *ConcealedKanDelete
}

// Exec executes the deletion query.
func (ckdo *ConcealedKanDeleteOne) Exec(ctx context.Context) error {
	n, err := ckdo.ckd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{concealedkan.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ckdo *ConcealedKanDeleteOne) ExecX(ctx context.Context) {
	ckdo.ckd.ExecX(ctx)
}
