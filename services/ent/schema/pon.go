package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pon holds the schema definition for the Pon entity.
type Pon struct {
	ent.Schema
}

// Fields of the Pon.
func (Pon) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Pon.
func (Pon) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("call", Call.Type).Unique(),
	}
}
