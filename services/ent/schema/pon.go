package schema

import "entgo.io/ent"

// Pon holds the schema definition for the Pon entity.
type Pon struct {
	ent.Schema
}

// Fields of the Pon.
func (Pon) Fields() []ent.Field {
	return nil
}

// Edges of the Pon.
func (Pon) Edges() []ent.Edge {
	return nil
}
