package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Discard holds the schema definition for the Discard entity.
type Discard struct {
	ent.Schema
}

// Fields of the Discard.
func (Discard) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Discard.
func (Discard) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("reach", Reach.Type).Unique(),
		edge.To("call", Call.Type).Unique(),
		edge.To("draw", Drawn.Type).Unique(),
	}
}
