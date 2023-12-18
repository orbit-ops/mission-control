package schema

import (
	"fmt"
	"regexp"

	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ogen-go/ogen"
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
		field.String("id").Unique(),
		field.String("description").Optional(),
		field.String("image").Validate(func(s string) error {
			re := regexp.MustCompile(`(?:.+\/)?([^:]+)(?::.+)?`)
			if !re.MatchString(s) {
				return fmt.Errorf("%s is not a valid docker image", s)
			}
			return nil
		}),
		field.JSON("config", map[string]string{}).
			Annotations(entoas.Schema(rocketSchema)),
	}
}

// Edges of the Rocket.
func (Rocket) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("missions", Mission.Type),
	}
}
