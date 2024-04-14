// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/banner"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/feature"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/predicate"
)

// FeatureUpdate is the builder for updating Feature entities.
type FeatureUpdate struct {
	config
	hooks    []Hook
	mutation *FeatureMutation
}

// Where appends a list predicates to the FeatureUpdate builder.
func (fu *FeatureUpdate) Where(ps ...predicate.Feature) *FeatureUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetFeatureName sets the "feature_name" field.
func (fu *FeatureUpdate) SetFeatureName(s string) *FeatureUpdate {
	fu.mutation.SetFeatureName(s)
	return fu
}

// SetNillableFeatureName sets the "feature_name" field if the given value is not nil.
func (fu *FeatureUpdate) SetNillableFeatureName(s *string) *FeatureUpdate {
	if s != nil {
		fu.SetFeatureName(*s)
	}
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FeatureUpdate) SetUpdatedAt(t time.Time) *FeatureUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fu *FeatureUpdate) SetNillableUpdatedAt(t *time.Time) *FeatureUpdate {
	if t != nil {
		fu.SetUpdatedAt(*t)
	}
	return fu
}

// AddBannerIDs adds the "banners" edge to the Banner entity by IDs.
func (fu *FeatureUpdate) AddBannerIDs(ids ...int) *FeatureUpdate {
	fu.mutation.AddBannerIDs(ids...)
	return fu
}

// AddBanners adds the "banners" edges to the Banner entity.
func (fu *FeatureUpdate) AddBanners(b ...*Banner) *FeatureUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fu.AddBannerIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fu *FeatureUpdate) Mutation() *FeatureMutation {
	return fu.mutation
}

// ClearBanners clears all "banners" edges to the Banner entity.
func (fu *FeatureUpdate) ClearBanners() *FeatureUpdate {
	fu.mutation.ClearBanners()
	return fu
}

// RemoveBannerIDs removes the "banners" edge to Banner entities by IDs.
func (fu *FeatureUpdate) RemoveBannerIDs(ids ...int) *FeatureUpdate {
	fu.mutation.RemoveBannerIDs(ids...)
	return fu
}

// RemoveBanners removes "banners" edges to Banner entities.
func (fu *FeatureUpdate) RemoveBanners(b ...*Banner) *FeatureUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fu.RemoveBannerIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FeatureUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FeatureUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FeatureUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FeatureUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FeatureUpdate) check() error {
	if v, ok := fu.mutation.FeatureName(); ok {
		if err := feature.FeatureNameValidator(v); err != nil {
			return &ValidationError{Name: "feature_name", err: fmt.Errorf(`ent: validator failed for field "Feature.feature_name": %w`, err)}
		}
	}
	return nil
}

func (fu *FeatureUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(feature.Table, feature.Columns, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.FeatureName(); ok {
		_spec.SetField(feature.FieldFeatureName, field.TypeString, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(feature.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.BannersTable,
			Columns: feature.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedBannersIDs(); len(nodes) > 0 && !fu.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.BannersTable,
			Columns: feature.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.BannersTable,
			Columns: feature.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feature.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FeatureUpdateOne is the builder for updating a single Feature entity.
type FeatureUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FeatureMutation
}

// SetFeatureName sets the "feature_name" field.
func (fuo *FeatureUpdateOne) SetFeatureName(s string) *FeatureUpdateOne {
	fuo.mutation.SetFeatureName(s)
	return fuo
}

// SetNillableFeatureName sets the "feature_name" field if the given value is not nil.
func (fuo *FeatureUpdateOne) SetNillableFeatureName(s *string) *FeatureUpdateOne {
	if s != nil {
		fuo.SetFeatureName(*s)
	}
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FeatureUpdateOne) SetUpdatedAt(t time.Time) *FeatureUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fuo *FeatureUpdateOne) SetNillableUpdatedAt(t *time.Time) *FeatureUpdateOne {
	if t != nil {
		fuo.SetUpdatedAt(*t)
	}
	return fuo
}

// AddBannerIDs adds the "banners" edge to the Banner entity by IDs.
func (fuo *FeatureUpdateOne) AddBannerIDs(ids ...int) *FeatureUpdateOne {
	fuo.mutation.AddBannerIDs(ids...)
	return fuo
}

// AddBanners adds the "banners" edges to the Banner entity.
func (fuo *FeatureUpdateOne) AddBanners(b ...*Banner) *FeatureUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fuo.AddBannerIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fuo *FeatureUpdateOne) Mutation() *FeatureMutation {
	return fuo.mutation
}

// ClearBanners clears all "banners" edges to the Banner entity.
func (fuo *FeatureUpdateOne) ClearBanners() *FeatureUpdateOne {
	fuo.mutation.ClearBanners()
	return fuo
}

// RemoveBannerIDs removes the "banners" edge to Banner entities by IDs.
func (fuo *FeatureUpdateOne) RemoveBannerIDs(ids ...int) *FeatureUpdateOne {
	fuo.mutation.RemoveBannerIDs(ids...)
	return fuo
}

// RemoveBanners removes "banners" edges to Banner entities.
func (fuo *FeatureUpdateOne) RemoveBanners(b ...*Banner) *FeatureUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return fuo.RemoveBannerIDs(ids...)
}

// Where appends a list predicates to the FeatureUpdate builder.
func (fuo *FeatureUpdateOne) Where(ps ...predicate.Feature) *FeatureUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FeatureUpdateOne) Select(field string, fields ...string) *FeatureUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Feature entity.
func (fuo *FeatureUpdateOne) Save(ctx context.Context) (*Feature, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FeatureUpdateOne) SaveX(ctx context.Context) *Feature {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FeatureUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FeatureUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FeatureUpdateOne) check() error {
	if v, ok := fuo.mutation.FeatureName(); ok {
		if err := feature.FeatureNameValidator(v); err != nil {
			return &ValidationError{Name: "feature_name", err: fmt.Errorf(`ent: validator failed for field "Feature.feature_name": %w`, err)}
		}
	}
	return nil
}

func (fuo *FeatureUpdateOne) sqlSave(ctx context.Context) (_node *Feature, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(feature.Table, feature.Columns, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Feature.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feature.FieldID)
		for _, f := range fields {
			if !feature.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != feature.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.FeatureName(); ok {
		_spec.SetField(feature.FieldFeatureName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(feature.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.BannersTable,
			Columns: feature.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedBannersIDs(); len(nodes) > 0 && !fuo.mutation.BannersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.BannersTable,
			Columns: feature.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.BannersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   feature.BannersTable,
			Columns: feature.BannersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(banner.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Feature{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{feature.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}