package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Wind holds the schema definition for the Wind entity.
type Wind struct {
	ent.Schema
}

// Fields of the Wind.
func (Wind) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
		field.String("name").Unique().Immutable(),
	}
}

// Edges of the Wind.
func (Wind) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.To("rounds", Round.Type),
	}
}
