// Code generated by ent, DO NOT EDIT.

package permission

import (
	"go-ent-project/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// CanRead applies equality check predicate on the "can_read" field. It's identical to CanReadEQ.
func CanRead(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanRead, v))
}

// CanCreate applies equality check predicate on the "can_create" field. It's identical to CanCreateEQ.
func CanCreate(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanCreate, v))
}

// CanUpdate applies equality check predicate on the "can_update" field. It's identical to CanUpdateEQ.
func CanUpdate(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanUpdate, v))
}

// CanDelete applies equality check predicate on the "can_delete" field. It's identical to CanDeleteEQ.
func CanDelete(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanDelete, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Permission {
	return predicate.Permission(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Permission {
	return predicate.Permission(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Permission {
	return predicate.Permission(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Permission {
	return predicate.Permission(sql.FieldContainsFold(FieldName, v))
}

// CanReadEQ applies the EQ predicate on the "can_read" field.
func CanReadEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanRead, v))
}

// CanReadNEQ applies the NEQ predicate on the "can_read" field.
func CanReadNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCanRead, v))
}

// CanCreateEQ applies the EQ predicate on the "can_create" field.
func CanCreateEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanCreate, v))
}

// CanCreateNEQ applies the NEQ predicate on the "can_create" field.
func CanCreateNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCanCreate, v))
}

// CanUpdateEQ applies the EQ predicate on the "can_update" field.
func CanUpdateEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanUpdate, v))
}

// CanUpdateNEQ applies the NEQ predicate on the "can_update" field.
func CanUpdateNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCanUpdate, v))
}

// CanDeleteEQ applies the EQ predicate on the "can_delete" field.
func CanDeleteEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldEQ(FieldCanDelete, v))
}

// CanDeleteNEQ applies the NEQ predicate on the "can_delete" field.
func CanDeleteNEQ(v bool) predicate.Permission {
	return predicate.Permission(sql.FieldNEQ(FieldCanDelete, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Permission) predicate.Permission {
	return predicate.Permission(sql.NotPredicates(p))
}
