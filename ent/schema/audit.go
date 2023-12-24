package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Audit holds the schema definition for the Audit entity.
type Audit struct {
	ent.Schema
}

// Annotations of the User.
func (Audit) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.CreateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.UpdateOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}

// Fields of the User.
func (Audit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Enum("action").Values(
			"ApproveRequest",
			"RevokeApprovalRequest",
			"RejectRequest",
			"CreateAccess",
			"RemoveAccess",
		).Immutable(),
		field.String("author").NotEmpty().Immutable(),
		field.Time("timestamp").Immutable(),
	}
}

// Edges of the Audit.
func (Audit) Edges() []ent.Edge {
	return []ent.Edge{}
}
