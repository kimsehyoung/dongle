// Code generated by ent, DO NOT EDIT.

package videodbgen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/predicate"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/subtitledvideo"
)

// SubtitledVideoDelete is the builder for deleting a SubtitledVideo entity.
type SubtitledVideoDelete struct {
	config
	hooks    []Hook
	mutation *SubtitledVideoMutation
}

// Where appends a list predicates to the SubtitledVideoDelete builder.
func (svd *SubtitledVideoDelete) Where(ps ...predicate.SubtitledVideo) *SubtitledVideoDelete {
	svd.mutation.Where(ps...)
	return svd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (svd *SubtitledVideoDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, svd.sqlExec, svd.mutation, svd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (svd *SubtitledVideoDelete) ExecX(ctx context.Context) int {
	n, err := svd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (svd *SubtitledVideoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(subtitledvideo.Table, sqlgraph.NewFieldSpec(subtitledvideo.FieldID, field.TypeInt))
	if ps := svd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, svd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	svd.mutation.done = true
	return affected, err
}

// SubtitledVideoDeleteOne is the builder for deleting a single SubtitledVideo entity.
type SubtitledVideoDeleteOne struct {
	svd *SubtitledVideoDelete
}

// Where appends a list predicates to the SubtitledVideoDelete builder.
func (svdo *SubtitledVideoDeleteOne) Where(ps ...predicate.SubtitledVideo) *SubtitledVideoDeleteOne {
	svdo.svd.mutation.Where(ps...)
	return svdo
}

// Exec executes the deletion query.
func (svdo *SubtitledVideoDeleteOne) Exec(ctx context.Context) error {
	n, err := svdo.svd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{subtitledvideo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (svdo *SubtitledVideoDeleteOne) ExecX(ctx context.Context) {
	if err := svdo.Exec(ctx); err != nil {
		panic(err)
	}
}
