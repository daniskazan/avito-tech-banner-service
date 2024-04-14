// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/banner"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/tag"
)

// TagCreate is the builder for creating a Tag entity.
type TagCreate struct {
	config
	mutation *TagMutation
	hooks    []Hook
}

// SetTagName sets the "tag_name" field.
func (tc *TagCreate) SetTagName(s string) *TagCreate {
	tc.mutation.SetTagName(s)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TagCreate) SetCreatedAt(t time.Time) *TagCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TagCreate) SetNillableCreatedAt(t *time.Time) *TagCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TagCreate) SetUpdatedAt(t time.Time) *TagCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TagCreate) SetNillableUpdatedAt(t *time.Time) *TagCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// AddBannerIDs adds the "banners" edge to the Banner entity by IDs.
func (tc *TagCreate) AddBannerIDs(ids ...int) *TagCreate {
	tc.mutation.AddBannerIDs(ids...)
	return tc
}

// AddBanners adds the "banners" edges to the Banner entity.
func (tc *TagCreate) AddBanners(b ...*Banner) *TagCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tc.AddBannerIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tc *TagCreate) Mutation() *TagMutation {
	return tc.mutation
}

// Save creates the Tag in the database.
func (tc *TagCreate) Save(ctx context.Context) (*Tag, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TagCreate) SaveX(ctx context.Context) *Tag {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TagCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TagCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TagCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := tag.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := tag.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TagCreate) check() error {
	if _, ok := tc.mutation.TagName(); !ok {
		return &ValidationError{Name: "tag_name", err: errors.New(`ent: missing required field "Tag.tag_name"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Tag.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Tag.updated_at"`)}
	}
	return nil
}

func (tc *TagCreate) sqlSave(ctx context.Context) (*Tag, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TagCreate) createSpec() (*Tag, *sqlgraph.CreateSpec) {
	var (
		_node = &Tag{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(tag.Table, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	)
	if value, ok := tc.mutation.TagName(); ok {
		_spec.SetField(tag.FieldTagName, field.TypeString, value)
		_node.TagName = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(tag.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(tag.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.BannersTable,
			Columns: tag.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TagCreateBulk is the builder for creating many Tag entities in bulk.
type TagCreateBulk struct {
	config
	err      error
	builders []*TagCreate
}

// Save creates the Tag entities in the database.
func (tcb *TagCreateBulk) Save(ctx context.Context) ([]*Tag, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Tag, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TagMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TagCreateBulk) SaveX(ctx context.Context) []*Tag {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TagCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TagCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
