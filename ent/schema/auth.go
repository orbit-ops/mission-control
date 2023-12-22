package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/orbit-ops/launchpad-core/utils"

	ogauthz "github.com/tiagoposse/ogent-auth/authorization"
)

// ApiKey holds the schema definition for the ApiKey entity.
type ApiKey struct {
	ent.Schema
}

// Annotations of the ApiKey.
func (ApiKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		ogauthz.WithAllScopes(utils.AdminScope),
	}
}

// Fields of the ApiKey.
func (ApiKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Immutable(),
		field.String("key").Immutable().Sensitive(),
	}
}

// Edges of the ApiKey.
func (ApiKey) Edges() []ent.Edge {
	return []ent.Edge{}
}
