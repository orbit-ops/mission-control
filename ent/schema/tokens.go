package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ActionTokens holds the schema definition for the Otp entity.
type ActionTokens struct {
	ent.Schema
}

// Annotations of the ActionTokens.
func (ActionTokens) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}

// Fields of the ActionTokens.
func (ActionTokens) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Enum("action").Values("create", "remove").Immutable(),
		field.String("token"),
		field.UUID("access_id", uuid.UUID{}),
		field.Time("expiration").Immutable(),
	}
}

// Edges of the ActionTokens.
func (ActionTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("accessTokens", Access.Type).Unique().Required().Field("access_id").
			Annotations(
				entoas.ReadOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
			),
	}
}
