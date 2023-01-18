package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Call holds the schema definition for the Call entity.
type Call struct {
	ent.Schema
}

// Fields of the Call.
func (Call) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
	}
}

// Edges of the Call.
func (Call) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).Ref("call").Unique().Required(),
		edge.From("discard", Discard.Type).Ref("call").Unique().Required(),
		edge.From("chii", Chii.Type).Ref("call").Unique().Required(),
		edge.From("chakan", Chakan.Type).Ref("call").Unique().Required(),
		edge.From("concealedkan", ConcealedKan.Type).Ref("call").Unique().Required(),
		edge.From("meldedkan", MeldedKan.Type).Ref("call").Unique().Required(),
		edge.From("pon", Pon.Type).Ref("call").Unique().Required(),
	}
}
