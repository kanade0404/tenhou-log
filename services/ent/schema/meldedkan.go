package schema

import "entgo.io/ent"

// MeldedKan holds the schema definition for the MeldedKan entity.
type MeldedKan struct {
	ent.Schema
}

// Fields of the MeldedKan.
func (MeldedKan) Fields() []ent.Field {
	return nil
}

// Edges of the MeldedKan.
func (MeldedKan) Edges() []ent.Edge {
	return nil
}
