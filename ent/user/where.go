// Code generated by ent, DO NOT EDIT.

package user

import (
	"nevissGo/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// GameID applies equality check predicate on the "game_id" field. It's identical to GameIDEQ.
func GameID(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGameID, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldDisplayName, v))
}

// GameIDEQ applies the EQ predicate on the "game_id" field.
func GameIDEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldGameID, v))
}

// GameIDNEQ applies the NEQ predicate on the "game_id" field.
func GameIDNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldGameID, v))
}

// GameIDIn applies the In predicate on the "game_id" field.
func GameIDIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldGameID, vs...))
}

// GameIDNotIn applies the NotIn predicate on the "game_id" field.
func GameIDNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldGameID, vs...))
}

// GameIDGT applies the GT predicate on the "game_id" field.
func GameIDGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldGameID, v))
}

// GameIDGTE applies the GTE predicate on the "game_id" field.
func GameIDGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldGameID, v))
}

// GameIDLT applies the LT predicate on the "game_id" field.
func GameIDLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldGameID, v))
}

// GameIDLTE applies the LTE predicate on the "game_id" field.
func GameIDLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldGameID, v))
}

// GameIDContains applies the Contains predicate on the "game_id" field.
func GameIDContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldGameID, v))
}

// GameIDHasPrefix applies the HasPrefix predicate on the "game_id" field.
func GameIDHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldGameID, v))
}

// GameIDHasSuffix applies the HasSuffix predicate on the "game_id" field.
func GameIDHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldGameID, v))
}

// GameIDEqualFold applies the EqualFold predicate on the "game_id" field.
func GameIDEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldGameID, v))
}

// GameIDContainsFold applies the ContainsFold predicate on the "game_id" field.
func GameIDContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldGameID, v))
}

// HasPixels applies the HasEdge predicate on the "pixels" edge.
func HasPixels() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PixelsTable, PixelsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPixelsWith applies the HasEdge predicate on the "pixels" edge with a given conditions (other predicates).
func HasPixelsWith(preds ...predicate.Pixel) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPixelsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHype applies the HasEdge predicate on the "hype" edge.
func HasHype() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, HypeTable, HypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHypeWith applies the HasEdge predicate on the "hype" edge with a given conditions (other predicates).
func HasHypeWith(preds ...predicate.Hype) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newHypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
