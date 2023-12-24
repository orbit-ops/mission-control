package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Time("start_time").Immutable(),
		field.Bool("approved").Immutable(),
		field.Bool("rolled_back").Default(false),
		field.Time("rollback_time").Optional(),
		field.String("rollback_reason").Optional(),
		field.Time("end_time").Immutable(),
		field.UUID("request_id", uuid.UUID{}),
	}
}

// Annotations of the Access.
func (Access) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}

// Edges of the Access.
func (Access) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("approvals", Approval.Type).
			Annotations(
				entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
		edge.From("accessTokens", ActionTokens.Type).Ref("accessTokens"),
	}
}
