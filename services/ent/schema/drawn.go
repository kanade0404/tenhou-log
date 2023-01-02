package schema

import "entgo.io/ent"

// Drawn holds the schema definition for the Drawn entity.
type Drawn struct {
	ent.Schema
}

// Fields of the Drawn.
func (Drawn) Fields() []ent.Field {
	return nil
}

// Edges of the Drawn.
func (Drawn) Edges() []ent.Edge {
	return nil
}
