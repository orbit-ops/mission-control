package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("reason"),
		field.String("requester").Immutable(),
		field.String("mission_id").Immutable(),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("requests", Approval.Type).Ref("requests").Required().Unique().Immutable(),
		edge.From("mission", Mission.Type).Immutable().Unique().Required().Ref("requests"),
	}
}
