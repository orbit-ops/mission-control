package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Access holds the schema definition for the Access entity.
type Access struct {
	ent.Schema
}

// Fields of the Access.
func (Access) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Time("access_time").Immutable(),
		field.Bool("approved").Immutable(),
		field.Bool("rolled_back").Default(false),
		field.Time("rollback_time").Optional(),
		field.Time("end_time").Immutable(),
		field.UUID("request_id", uuid.UUID{}),
	}
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("approvals", Access.Type).Unique(),
	}
}
