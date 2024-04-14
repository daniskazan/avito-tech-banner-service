package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"time"
)

// Banner holds the schema definition for the Banner entity.
type Banner struct {
	ent.Schema
}

// Fields of the Banner.
func (Banner) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("content", map[string]interface{}{}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now),
		field.Bool("is_active").Default(true),
	}
}

// Edges of the Banner.
func (Banner) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feature", Feature.Type).Ref("banners").Unique(), // Связь "много к одному" с фичей.
		edge.From("tags", Tag.Type).Ref("banners"),                 // Связь "много ко многим" с тегами.
	}
}
