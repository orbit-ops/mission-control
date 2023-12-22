package schema

import (
	"entgo.io/contrib/entoas"
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
		field.String("reason"),
		field.String("requester").Immutable(),
		field.JSON("rocket_config", map[string]string{}).
			Annotations(entoas.Schema(rocketSchema)),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		// edge.From("requests", Approval.Type).Ref("requests").Required().Unique().Immutable(),
		edge.From("mission", Mission.Type).Immutable().Unique().Required().Ref("requests"),
	}
}
