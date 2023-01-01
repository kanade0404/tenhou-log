// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kanade0404/tenhou-log/services/ent/hand"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// HandDelete is the builder for deleting a Hand entity.
type HandDelete struct {
	config
	hooks    []Hook
	mutation *HandMutation
}

// Where appends a list predicates to the HandDelete builder.
func (hd *HandDelete) Where(ps ...predicate.Hand) *HandDelete {
	hd.mutation.Where(ps...)
	return hd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (hd *HandDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(hd.hooks) == 0 {
		affected, err = hd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*HandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			hd.mutation = mutation
			affected, err = hd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(hd.hooks) - 1; i >= 0; i-- {
			if hd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = hd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, hd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (hd *HandDelete) ExecX(ctx context.Context) int {
	n, err := hd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (hd *HandDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: hand.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: hand.FieldID,
			},
		},
	}
	if ps := hd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, hd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// HandDeleteOne is the builder for deleting a single Hand entity.
type HandDeleteOne struct {
	hd *HandDelete
}

// Exec executes the deletion query.
func (hdo *HandDeleteOne) Exec(ctx context.Context) error {
	n, err := hdo.hd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{hand.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (hdo *HandDeleteOne) ExecX(ctx context.Context) {
	hdo.hd.ExecX(ctx)
}
