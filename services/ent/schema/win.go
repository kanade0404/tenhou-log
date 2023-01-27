package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Win holds the schema definition for the Win entity.
type Win struct {
	ent.Schema
}

// Fields of the Win.
func (Win) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Win.
func (Win) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("win").Unique().Required(),
	}
}
