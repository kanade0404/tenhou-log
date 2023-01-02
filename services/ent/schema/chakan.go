package schema

import "entgo.io/ent"

// Chakan holds the schema definition for the Chakan entity.
type Chakan struct {
	ent.Schema
}

// Fields of the Chakan.
func (Chakan) Fields() []ent.Field {
	return nil
}

// Edges of the Chakan.
func (Chakan) Edges() []ent.Edge {
	return nil
}
