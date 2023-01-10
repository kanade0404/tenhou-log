package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Win holds the schema definition for the Win entity.
type Win struct {
	ent.Schema
}

// Fields of the Win.
func (Win) Fields() []ent.Field {
	return nil
}

// Edges of the Win.
func (Win) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("win").Unique().Required(),
	}
}
