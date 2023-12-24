package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ogen-go/ogen"
	"github.com/orbit-ops/launchpad-core/utils"
	ogauth "github.com/tiagoposse/ogent-auth/authorization"
)

// Rocket holds the schema definition for the Rocket entity.
type Rocket struct {
	ent.Schema
}

// Fields of the Rocket.
func (Rocket) Fields() []ent.Field {
	rocketSchema := ogen.NewSchema().SetType("object")
	rocketSchema.AdditionalProperties = &ogen.AdditionalProperties{Schema: *ogen.String()}

	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
		field.String("code").Optional().Validate(func(s string) error {
			// re := regexp.MustCompile(`(?:.+\/)?([^:]+)(?::.+)?`)
			// if !re.MatchString(s) {
			// 	return fmt.Errorf("%s is not a valid docker image", s)
			// }
			return nil
		}),
		field.JSON("config", map[string]string{}).
			Default(map[string]string{}).
			Annotations(entoas.Schema(rocketSchema)),
	}
}

// Edges of the Rocket.
// func (Rocket) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.To("missions", Mission.Type),
// 	}
// }

// Annotations of the Rocket.
func (Rocket) Annotations() []schema.Annotation {
	return []schema.Annotation{
		ogauth.WithAllScopes(utils.AdminScope),
	}
}
