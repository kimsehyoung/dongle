// Code generated by ent, DO NOT EDIT.

package subtitlegen

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimsehyoung/dongle/app/subtitle/ent/subtitlegen/subtitle"
	"github.com/kimsehyoung/dongle/app/subtitle/ent/subtitlegen/video"
)

// SubtitleCreate is the builder for creating a Subtitle entity.
type SubtitleCreate struct {
	config
	mutation *SubtitleMutation
	hooks    []Hook
}

// SetVideoID sets the "video_id" field.
func (sc *SubtitleCreate) SetVideoID(i int) *SubtitleCreate {
	sc.mutation.SetVideoID(i)
	return sc
}

// SetLanguage sets the "language" field.
func (sc *SubtitleCreate) SetLanguage(s string) *SubtitleCreate {
	sc.mutation.SetLanguage(s)
	return sc
}

// SetURL sets the "url" field.
func (sc *SubtitleCreate) SetURL(s string) *SubtitleCreate {
	sc.mutation.SetURL(s)
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SubtitleCreate) SetCreatedAt(t time.Time) *SubtitleCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SubtitleCreate) SetNillableCreatedAt(t *time.Time) *SubtitleCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetRoleID sets the "role" edge to the Video entity by ID.
func (sc *SubtitleCreate) SetRoleID(id int) *SubtitleCreate {
	sc.mutation.SetRoleID(id)
	return sc
}

// SetRole sets the "role" edge to the Video entity.
func (sc *SubtitleCreate) SetRole(v *Video) *SubtitleCreate {
	return sc.SetRoleID(v.ID)
}

// Mutation returns the SubtitleMutation object of the builder.
func (sc *SubtitleCreate) Mutation() *SubtitleMutation {
	return sc.mutation
}

// Save creates the Subtitle in the database.
func (sc *SubtitleCreate) Save(ctx context.Context) (*Subtitle, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SubtitleCreate) SaveX(ctx context.Context) *Subtitle {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SubtitleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SubtitleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SubtitleCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := subtitle.DefaultCreatedAt
		sc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SubtitleCreate) check() error {
	if _, ok := sc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`subtitlegen: missing required field "Subtitle.video_id"`)}
	}
	if _, ok := sc.mutation.Language(); !ok {
		return &ValidationError{Name: "language", err: errors.New(`subtitlegen: missing required field "Subtitle.language"`)}
	}
	if _, ok := sc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`subtitlegen: missing required field "Subtitle.url"`)}
	}
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`subtitlegen: missing required field "Subtitle.created_at"`)}
	}
	if _, ok := sc.mutation.RoleID(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`subtitlegen: missing required edge "Subtitle.role"`)}
	}
	return nil
}

func (sc *SubtitleCreate) sqlSave(ctx context.Context) (*Subtitle, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SubtitleCreate) createSpec() (*Subtitle, *sqlgraph.CreateSpec) {
	var (
		_node = &Subtitle{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(subtitle.Table, sqlgraph.NewFieldSpec(subtitle.FieldID, field.TypeInt))
	)
	if value, ok := sc.mutation.Language(); ok {
		_spec.SetField(subtitle.FieldLanguage, field.TypeString, value)
		_node.Language = value
	}
	if value, ok := sc.mutation.URL(); ok {
		_spec.SetField(subtitle.FieldURL, field.TypeString, value)
		_node.URL = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(subtitle.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := sc.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   subtitle.RoleTable,
			Columns: []string{subtitle.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(video.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.VideoID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SubtitleCreateBulk is the builder for creating many Subtitle entities in bulk.
type SubtitleCreateBulk struct {
	config
	builders []*SubtitleCreate
}

// Save creates the Subtitle entities in the database.
func (scb *SubtitleCreateBulk) Save(ctx context.Context) ([]*Subtitle, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Subtitle, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SubtitleMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SubtitleCreateBulk) SaveX(ctx context.Context) []*Subtitle {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SubtitleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SubtitleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
