package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// GamePlayerHandHai holds the schema definition for the GamePlayerHandHai entity.
type GamePlayerHandHai struct {
	ent.Schema
}

// Fields of the GamePlayerHandHai.
func (GamePlayerHandHai) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Ints("hais").Immutable(),
	}
}

// Edges of the GamePlayerHandHai.
func (GamePlayerHandHai) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("turn", Turn.Type).Ref("gameplayerhandhai").Unique().Required(),
	}
}
