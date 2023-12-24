package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	rocketSchema := ogen.NewSchema().SetType("object")
	rocketSchema.AdditionalProperties = &ogen.AdditionalProperties{Schema: *ogen.String()}

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("reason").Immutable(),
		field.String("requester").Immutable(),
		// field.UUID("mission_id", uuid.UUID{}).Immutable(),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("approvals", Approval.Type).Ref("request").Immutable(),
		edge.To("mission", Mission.Type).Immutable().Unique().Required(),
	}
}
