package schema

import "entgo.io/ent"

// Call holds the schema definition for the Call entity.
type Call struct {
	ent.Schema
}

// Fields of the Call.
func (Call) Fields() []ent.Field {
	return nil
}

// Edges of the Call.
func (Call) Edges() []ent.Edge {
	return nil
}
