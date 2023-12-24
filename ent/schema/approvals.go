package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Approval holds the schema definition for the Approval entity.
type Approval struct {
	ent.Schema
}

// Fields of the Approval.
func (Approval) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("person").Immutable(),
		field.Time("approved_time").Immutable(),
		field.Bool("approved").Immutable(),
		field.Bool("revoked").Default(false),
		field.Time("revoked_time").Optional(),
	}
}

// Edges of the Approval.
func (Approval) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("request", Request.Type).Unique().Required().
			Annotations(
				entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
		edge.From("access", Access.Type).Ref("approvals"),
	}
}
