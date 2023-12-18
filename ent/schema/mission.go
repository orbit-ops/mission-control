package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("description").Optional(),
		field.String("image"),
		field.Int("min_approvers"),
		field.String("rocket_id").Immutable(),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rocket", Rocket.Type).Ref("missions").Unique().Immutable().Required().Field("rocket_id"),
		edge.To("requests", Request.Type),
	}
}
