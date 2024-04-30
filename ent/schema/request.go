package schema

import (
	"time"

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
		field.String("reason").Immutable().NotEmpty(),
		field.String("requester").Immutable().NotEmpty(),
		field.Time("timestamp").Immutable().Default(time.Now),
		field.Time("cancelled_time").Optional().Annotations(
			entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		),
		field.Bool("cancelled").Default(false).Annotations(
			entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("approvals", Approval.Type).Ref("request").Immutable().
			Annotations(
				entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
		edge.To("mission", Mission.Type).Unique().Immutable().Required(),
	}
}
