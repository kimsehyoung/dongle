// Code generated by ent, DO NOT EDIT.

package videodbgen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/originalvideo"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/predicate"
)

// OriginalVideoDelete is the builder for deleting a OriginalVideo entity.
type OriginalVideoDelete struct {
	config
	hooks    []Hook
	mutation *OriginalVideoMutation
}

// Where appends a list predicates to the OriginalVideoDelete builder.
func (ovd *OriginalVideoDelete) Where(ps ...predicate.OriginalVideo) *OriginalVideoDelete {
	ovd.mutation.Where(ps...)
	return ovd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ovd *OriginalVideoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ovd.sqlExec, ovd.mutation, ovd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ovd *OriginalVideoDelete) ExecX(ctx context.Context) int {
	n, err := ovd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ovd *OriginalVideoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(originalvideo.Table, sqlgraph.NewFieldSpec(originalvideo.FieldID, field.TypeInt))
	if ps := ovd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ovd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ovd.mutation.done = true
	return affected, err
}

// OriginalVideoDeleteOne is the builder for deleting a single OriginalVideo entity.
type OriginalVideoDeleteOne struct {
	ovd *OriginalVideoDelete
}

// Where appends a list predicates to the OriginalVideoDelete builder.
func (ovdo *OriginalVideoDeleteOne) Where(ps ...predicate.OriginalVideo) *OriginalVideoDeleteOne {
	ovdo.ovd.mutation.Where(ps...)
	return ovdo
}

// Exec executes the deletion query.
func (ovdo *OriginalVideoDeleteOne) Exec(ctx context.Context) error {
	n, err := ovdo.ovd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{originalvideo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ovdo *OriginalVideoDeleteOne) ExecX(ctx context.Context) {
	if err := ovdo.Exec(ctx); err != nil {
		panic(err)
	}
}
