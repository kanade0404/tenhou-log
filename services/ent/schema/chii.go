package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Chii holds the schema definition for the Chii entity.
type Chii struct {
	ent.Schema
}

// Fields of the Chii.
func (Chii) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Chii.
func (Chii) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("call", Call.Type).Unique(),
	}
}
