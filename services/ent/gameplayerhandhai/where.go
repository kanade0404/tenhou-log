// Code generated by ent, DO NOT EDIT.

package gameplayerhandhai

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(sql.FieldLTE(FieldID, id))
}

// HasTurn applies the HasEdge predicate on the "turn" edge.
func HasTurn() predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TurnTable, TurnColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTurnWith applies the HasEdge predicate on the "turn" edge with a given conditions (other predicates).
func HasTurnWith(preds ...predicate.Turn) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TurnInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, TurnTable, TurnColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GamePlayerHandHai) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GamePlayerHandHai) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.GamePlayerHandHai) predicate.GamePlayerHandHai {
	return predicate.GamePlayerHandHai(func(s *sql.Selector) {
		p(s.Not())
	})
}
