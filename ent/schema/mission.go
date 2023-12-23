package schema

import (
	"errors"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/utils"
	ogauth "github.com/tiagoposse/ogent-auth/authorization"
)

// Mission holds the schema definition for the Mission entity.
type Mission struct {
	ent.Schema
}

// Fields of the Mission.
func (Mission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique(),
		field.String("name").Unique(),
		field.String("description").Optional(),
		field.Int("min_approvers").Validate(func(n int) error {
			if n < 1 {
				return errors.New("minimum approvers must be bigger than 1")
			}
			return nil
		}),
		field.Strings("possible_approvers"),
	}
}

// Edges of the Mission.
func (Mission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rockets", Rocket.Type).Ref("missions").Immutable().Required(),
		edge.To("requests", Request.Type),
	}
}

// Annotations of the Mission.
func (Mission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ogauth.WithAllScopes(utils.AdminScope),
	}
}
