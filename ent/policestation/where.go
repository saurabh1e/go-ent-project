// Code generated by ent, DO NOT EDIT.

package policestation

import (
	"go-ent-project/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldName, v))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldCode, v))
}

// Identifier applies equality check predicate on the "identifier" field. It's identical to IdentifierEQ.
func Identifier(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldIdentifier, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldContainsFold(FieldName, v))
}

// LocationIsNil applies the IsNil predicate on the "location" field.
func LocationIsNil() predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldIsNull(FieldLocation))
}

// LocationNotNil applies the NotNil predicate on the "location" field.
func LocationNotNil() predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNotNull(FieldLocation))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldContainsFold(FieldCode, v))
}

// IdentifierEQ applies the EQ predicate on the "identifier" field.
func IdentifierEQ(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEQ(FieldIdentifier, v))
}

// IdentifierNEQ applies the NEQ predicate on the "identifier" field.
func IdentifierNEQ(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNEQ(FieldIdentifier, v))
}

// IdentifierIn applies the In predicate on the "identifier" field.
func IdentifierIn(vs ...string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldIn(FieldIdentifier, vs...))
}

// IdentifierNotIn applies the NotIn predicate on the "identifier" field.
func IdentifierNotIn(vs ...string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldNotIn(FieldIdentifier, vs...))
}

// IdentifierGT applies the GT predicate on the "identifier" field.
func IdentifierGT(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGT(FieldIdentifier, v))
}

// IdentifierGTE applies the GTE predicate on the "identifier" field.
func IdentifierGTE(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldGTE(FieldIdentifier, v))
}

// IdentifierLT applies the LT predicate on the "identifier" field.
func IdentifierLT(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLT(FieldIdentifier, v))
}

// IdentifierLTE applies the LTE predicate on the "identifier" field.
func IdentifierLTE(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldLTE(FieldIdentifier, v))
}

// IdentifierContains applies the Contains predicate on the "identifier" field.
func IdentifierContains(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldContains(FieldIdentifier, v))
}

// IdentifierHasPrefix applies the HasPrefix predicate on the "identifier" field.
func IdentifierHasPrefix(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldHasPrefix(FieldIdentifier, v))
}

// IdentifierHasSuffix applies the HasSuffix predicate on the "identifier" field.
func IdentifierHasSuffix(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldHasSuffix(FieldIdentifier, v))
}

// IdentifierEqualFold applies the EqualFold predicate on the "identifier" field.
func IdentifierEqualFold(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldEqualFold(FieldIdentifier, v))
}

// IdentifierContainsFold applies the ContainsFold predicate on the "identifier" field.
func IdentifierContainsFold(v string) predicate.PoliceStation {
	return predicate.PoliceStation(sql.FieldContainsFold(FieldIdentifier, v))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.PoliceStation {
	return predicate.PoliceStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UsersTable, UsersColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.PoliceStation {
	return predicate.PoliceStation(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParentStation applies the HasEdge predicate on the "parent_station" edge.
func HasParentStation() predicate.PoliceStation {
	return predicate.PoliceStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentStationTable, ParentStationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentStationWith applies the HasEdge predicate on the "parent_station" edge with a given conditions (other predicates).
func HasParentStationWith(preds ...predicate.PoliceStation) predicate.PoliceStation {
	return predicate.PoliceStation(func(s *sql.Selector) {
		step := newParentStationStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildStations applies the HasEdge predicate on the "child_stations" edge.
func HasChildStations() predicate.PoliceStation {
	return predicate.PoliceStation(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildStationsTable, ChildStationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildStationsWith applies the HasEdge predicate on the "child_stations" edge with a given conditions (other predicates).
func HasChildStationsWith(preds ...predicate.PoliceStation) predicate.PoliceStation {
	return predicate.PoliceStation(func(s *sql.Selector) {
		step := newChildStationsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PoliceStation) predicate.PoliceStation {
	return predicate.PoliceStation(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PoliceStation) predicate.PoliceStation {
	return predicate.PoliceStation(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PoliceStation) predicate.PoliceStation {
	return predicate.PoliceStation(sql.NotPredicates(p))
}
