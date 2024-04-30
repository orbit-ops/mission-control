package schema

import (
	"time"

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
		field.Time("start_time").Immutable().Annotations(entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude))),
		field.Bool("provisioned").Default(false).Annotations(entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude))),
		field.Bool("rolled_back").Default(false).Annotations(entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude))),
		field.Time("rollback_time").Default(time.Now).Nillable().Annotations(entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude))),
		field.String("rollback_reason").Optional().Nillable().Annotations(entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude))),
		field.Time("expiration").Immutable().Annotations(entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude))),
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
		edge.To("request", Request.Type).Unique().Required().Immutable(),
		edge.From("accessTokens", ActionTokens.Type).Ref("accessTokens"),
	}
}

// // Mixins of the Access.
// func (Access) Mixin() []ent.Mixin {
// 	return []ent.Mixin{
// 		SoftDeleteMixin{},
// 	}
// }
