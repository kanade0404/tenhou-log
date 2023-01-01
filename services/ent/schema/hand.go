package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Hand holds the schema definition for the Hand entity.
type Hand struct {
	ent.Schema
}

// Fields of the Hand.
func (Hand) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).StorageKey("oid"),
	}
}

// Edges of the Hand.
func (Hand) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("rounds", Round.Type).Ref("hands"),
		//edge.To("go_arounds", GoAround.Type),
	}
}
