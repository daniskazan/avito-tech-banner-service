// Code generated by ent, DO NOT EDIT.

package feature

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/daniskazan/avito-tech-banner-service/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Feature {
	return predicate.Feature(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Feature {
	return predicate.Feature(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Feature {
	return predicate.Feature(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Feature {
	return predicate.Feature(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Feature {
	return predicate.Feature(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Feature {
	return predicate.Feature(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Feature {
	return predicate.Feature(sql.FieldLTE(FieldID, id))
}

// FeatureName applies equality check predicate on the "feature_name" field. It's identical to FeatureNameEQ.
func FeatureName(v string) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldFeatureName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldUpdatedAt, v))
}

// FeatureNameEQ applies the EQ predicate on the "feature_name" field.
func FeatureNameEQ(v string) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldFeatureName, v))
}

// FeatureNameNEQ applies the NEQ predicate on the "feature_name" field.
func FeatureNameNEQ(v string) predicate.Feature {
	return predicate.Feature(sql.FieldNEQ(FieldFeatureName, v))
}

// FeatureNameIn applies the In predicate on the "feature_name" field.
func FeatureNameIn(vs ...string) predicate.Feature {
	return predicate.Feature(sql.FieldIn(FieldFeatureName, vs...))
}

// FeatureNameNotIn applies the NotIn predicate on the "feature_name" field.
func FeatureNameNotIn(vs ...string) predicate.Feature {
	return predicate.Feature(sql.FieldNotIn(FieldFeatureName, vs...))
}

// FeatureNameGT applies the GT predicate on the "feature_name" field.
func FeatureNameGT(v string) predicate.Feature {
	return predicate.Feature(sql.FieldGT(FieldFeatureName, v))
}

// FeatureNameGTE applies the GTE predicate on the "feature_name" field.
func FeatureNameGTE(v string) predicate.Feature {
	return predicate.Feature(sql.FieldGTE(FieldFeatureName, v))
}

// FeatureNameLT applies the LT predicate on the "feature_name" field.
func FeatureNameLT(v string) predicate.Feature {
	return predicate.Feature(sql.FieldLT(FieldFeatureName, v))
}

// FeatureNameLTE applies the LTE predicate on the "feature_name" field.
func FeatureNameLTE(v string) predicate.Feature {
	return predicate.Feature(sql.FieldLTE(FieldFeatureName, v))
}

// FeatureNameContains applies the Contains predicate on the "feature_name" field.
func FeatureNameContains(v string) predicate.Feature {
	return predicate.Feature(sql.FieldContains(FieldFeatureName, v))
}

// FeatureNameHasPrefix applies the HasPrefix predicate on the "feature_name" field.
func FeatureNameHasPrefix(v string) predicate.Feature {
	return predicate.Feature(sql.FieldHasPrefix(FieldFeatureName, v))
}

// FeatureNameHasSuffix applies the HasSuffix predicate on the "feature_name" field.
func FeatureNameHasSuffix(v string) predicate.Feature {
	return predicate.Feature(sql.FieldHasSuffix(FieldFeatureName, v))
}

// FeatureNameEqualFold applies the EqualFold predicate on the "feature_name" field.
func FeatureNameEqualFold(v string) predicate.Feature {
	return predicate.Feature(sql.FieldEqualFold(FieldFeatureName, v))
}

// FeatureNameContainsFold applies the ContainsFold predicate on the "feature_name" field.
func FeatureNameContainsFold(v string) predicate.Feature {
	return predicate.Feature(sql.FieldContainsFold(FieldFeatureName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Feature {
	return predicate.Feature(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasBanners applies the HasEdge predicate on the "banners" edge.
func HasBanners() predicate.Feature {
	return predicate.Feature(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, BannersTable, BannersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBannersWith applies the HasEdge predicate on the "banners" edge with a given conditions (other predicates).
func HasBannersWith(preds ...predicate.Banner) predicate.Feature {
	return predicate.Feature(func(s *sql.Selector) {
		step := newBannersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Feature) predicate.Feature {
	return predicate.Feature(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Feature) predicate.Feature {
	return predicate.Feature(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Feature) predicate.Feature {
	return predicate.Feature(sql.NotPredicates(p))
}
