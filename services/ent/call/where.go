// Code generated by ent, DO NOT EDIT.

package call

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/kanade0404/tenhou-log/services/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Call {
	return predicate.Call(sql.FieldLTE(FieldID, id))
}

// HasEvent applies the HasEdge predicate on the "event" edge.
func HasEvent() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEventWith applies the HasEdge predicate on the "event" edge with a given conditions (other predicates).
func HasEventWith(preds ...predicate.Event) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(EventInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, EventTable, EventColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDiscard applies the HasEdge predicate on the "discard" edge.
func HasDiscard() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DiscardTable, DiscardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDiscardWith applies the HasEdge predicate on the "discard" edge with a given conditions (other predicates).
func HasDiscardWith(preds ...predicate.Discard) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DiscardInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, DiscardTable, DiscardColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChii applies the HasEdge predicate on the "chii" edge.
func HasChii() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ChiiTable, ChiiColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChiiWith applies the HasEdge predicate on the "chii" edge with a given conditions (other predicates).
func HasChiiWith(preds ...predicate.Chii) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChiiInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ChiiTable, ChiiColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChakan applies the HasEdge predicate on the "chakan" edge.
func HasChakan() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ChakanTable, ChakanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChakanWith applies the HasEdge predicate on the "chakan" edge with a given conditions (other predicates).
func HasChakanWith(preds ...predicate.Chakan) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ChakanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ChakanTable, ChakanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasConcealedkan applies the HasEdge predicate on the "concealedkan" edge.
func HasConcealedkan() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ConcealedkanTable, ConcealedkanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConcealedkanWith applies the HasEdge predicate on the "concealedkan" edge with a given conditions (other predicates).
func HasConcealedkanWith(preds ...predicate.ConcealedKan) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConcealedkanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ConcealedkanTable, ConcealedkanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMeldedkan applies the HasEdge predicate on the "meldedkan" edge.
func HasMeldedkan() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MeldedkanTable, MeldedkanColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMeldedkanWith applies the HasEdge predicate on the "meldedkan" edge with a given conditions (other predicates).
func HasMeldedkanWith(preds ...predicate.MeldedKan) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(MeldedkanInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, MeldedkanTable, MeldedkanColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPon applies the HasEdge predicate on the "pon" edge.
func HasPon() predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PonTable, PonColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPonWith applies the HasEdge predicate on the "pon" edge with a given conditions (other predicates).
func HasPonWith(preds ...predicate.Pon) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PonInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PonTable, PonColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Call) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Call) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
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
func Not(p predicate.Call) predicate.Call {
	return predicate.Call(func(s *sql.Selector) {
		p(s.Not())
	})
}