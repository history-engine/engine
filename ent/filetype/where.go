// Code generated by ent, DO NOT EDIT.

package filetype

import (
	"history-engine/engine/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldUserID, v))
}

// Suffix applies equality check predicate on the "suffix" field. It's identical to SuffixEQ.
func Suffix(v string) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldSuffix, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v int) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldType, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int64) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int64) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int64) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldUserID, v))
}

// SuffixEQ applies the EQ predicate on the "suffix" field.
func SuffixEQ(v string) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldSuffix, v))
}

// SuffixNEQ applies the NEQ predicate on the "suffix" field.
func SuffixNEQ(v string) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldSuffix, v))
}

// SuffixIn applies the In predicate on the "suffix" field.
func SuffixIn(vs ...string) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldSuffix, vs...))
}

// SuffixNotIn applies the NotIn predicate on the "suffix" field.
func SuffixNotIn(vs ...string) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldSuffix, vs...))
}

// SuffixGT applies the GT predicate on the "suffix" field.
func SuffixGT(v string) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldSuffix, v))
}

// SuffixGTE applies the GTE predicate on the "suffix" field.
func SuffixGTE(v string) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldSuffix, v))
}

// SuffixLT applies the LT predicate on the "suffix" field.
func SuffixLT(v string) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldSuffix, v))
}

// SuffixLTE applies the LTE predicate on the "suffix" field.
func SuffixLTE(v string) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldSuffix, v))
}

// SuffixContains applies the Contains predicate on the "suffix" field.
func SuffixContains(v string) predicate.FileType {
	return predicate.FileType(sql.FieldContains(FieldSuffix, v))
}

// SuffixHasPrefix applies the HasPrefix predicate on the "suffix" field.
func SuffixHasPrefix(v string) predicate.FileType {
	return predicate.FileType(sql.FieldHasPrefix(FieldSuffix, v))
}

// SuffixHasSuffix applies the HasSuffix predicate on the "suffix" field.
func SuffixHasSuffix(v string) predicate.FileType {
	return predicate.FileType(sql.FieldHasSuffix(FieldSuffix, v))
}

// SuffixEqualFold applies the EqualFold predicate on the "suffix" field.
func SuffixEqualFold(v string) predicate.FileType {
	return predicate.FileType(sql.FieldEqualFold(FieldSuffix, v))
}

// SuffixContainsFold applies the ContainsFold predicate on the "suffix" field.
func SuffixContainsFold(v string) predicate.FileType {
	return predicate.FileType(sql.FieldContainsFold(FieldSuffix, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v int) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v int) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...int) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...int) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v int) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v int) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v int) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v int) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldType, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.FileType {
	return predicate.FileType(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.FileType) predicate.FileType {
	return predicate.FileType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.FileType) predicate.FileType {
	return predicate.FileType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FileType) predicate.FileType {
	return predicate.FileType(sql.NotPredicates(p))
}
