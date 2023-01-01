package schema

import "entgo.io/ent"

// GamePlayerPoint holds the schema definition for the GamePlayerPoint entity.
type GamePlayerPoint struct {
	ent.Schema
}

// Fields of the GamePlayerPoint.
func (GamePlayerPoint) Fields() []ent.Field {
	return nil
}

// Edges of the GamePlayerPoint.
func (GamePlayerPoint) Edges() []ent.Edge {
	return nil
}
