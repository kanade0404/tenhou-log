package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ConcealedKan holds the schema definition for the ConcealedKan entity.
type ConcealedKan struct {
	ent.Schema
}

// Fields of the ConcealedKan.
func (ConcealedKan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the ConcealedKan.
func (ConcealedKan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("call", Call.Type).Unique(),
	}
}
