package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Feature holds the schema definition for the Feature entity.
type Feature struct {
	ent.Schema
}

// Fields of the Feature.
func (Feature) Fields() []ent.Field {
	return []ent.Field{
		field.Text("feature_name").NotEmpty().Unique(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
	}
}

// Edges of the Feature.
func (Feature) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("banners", Banner.Type), // Связь "один ко многим" с баннерами.
	}
}
