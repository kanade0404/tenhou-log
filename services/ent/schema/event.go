package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return nil
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("turns", Turn.Type).Unique().Required(),
	}
}
