package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Chakan holds the schema definition for the Chakan entity.
type Chakan struct {
	ent.Schema
}

// Fields of the Chakan.
func (Chakan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Chakan.
func (Chakan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("call", Call.Type).Unique(),
	}
}
