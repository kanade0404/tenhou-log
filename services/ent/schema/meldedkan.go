package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MeldedKan holds the schema definition for the MeldedKan entity.
type MeldedKan struct {
	ent.Schema
}

// Fields of the MeldedKan.
func (MeldedKan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the MeldedKan.
func (MeldedKan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("call", Call.Type).Unique(),
	}
}
