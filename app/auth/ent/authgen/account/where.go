// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/kimsehyoung/dongle/app/auth/ent/authgen/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldID, id))
}

// RoleID applies equality check predicate on the "role_id" field. It's identical to RoleIDEQ.
func RoleID(v int32) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldRoleID, v))
}

// LoginID applies equality check predicate on the "login_id" field. It's identical to LoginIDEQ.
func LoginID(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldLoginID, v))
}

// HashedPassword applies equality check predicate on the "hashed_password" field. It's identical to HashedPasswordEQ.
func HashedPassword(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldHashedPassword, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldName, v))
}

// Email applies equality check predicate on the "email" field. It's identical to EmailEQ.
func Email(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// PhoneNumber applies equality check predicate on the "phone_number" field. It's identical to PhoneNumberEQ.
func PhoneNumber(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPhoneNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreatedAt, v))
}

// RoleIDEQ applies the EQ predicate on the "role_id" field.
func RoleIDEQ(v int32) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldRoleID, v))
}

// RoleIDNEQ applies the NEQ predicate on the "role_id" field.
func RoleIDNEQ(v int32) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldRoleID, v))
}

// RoleIDIn applies the In predicate on the "role_id" field.
func RoleIDIn(vs ...int32) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldRoleID, vs...))
}

// RoleIDNotIn applies the NotIn predicate on the "role_id" field.
func RoleIDNotIn(vs ...int32) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldRoleID, vs...))
}

// LoginIDEQ applies the EQ predicate on the "login_id" field.
func LoginIDEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldLoginID, v))
}

// LoginIDNEQ applies the NEQ predicate on the "login_id" field.
func LoginIDNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldLoginID, v))
}

// LoginIDIn applies the In predicate on the "login_id" field.
func LoginIDIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldLoginID, vs...))
}

// LoginIDNotIn applies the NotIn predicate on the "login_id" field.
func LoginIDNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldLoginID, vs...))
}

// LoginIDGT applies the GT predicate on the "login_id" field.
func LoginIDGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldLoginID, v))
}

// LoginIDGTE applies the GTE predicate on the "login_id" field.
func LoginIDGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldLoginID, v))
}

// LoginIDLT applies the LT predicate on the "login_id" field.
func LoginIDLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldLoginID, v))
}

// LoginIDLTE applies the LTE predicate on the "login_id" field.
func LoginIDLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldLoginID, v))
}

// LoginIDContains applies the Contains predicate on the "login_id" field.
func LoginIDContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldLoginID, v))
}

// LoginIDHasPrefix applies the HasPrefix predicate on the "login_id" field.
func LoginIDHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldLoginID, v))
}

// LoginIDHasSuffix applies the HasSuffix predicate on the "login_id" field.
func LoginIDHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldLoginID, v))
}

// LoginIDEqualFold applies the EqualFold predicate on the "login_id" field.
func LoginIDEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldLoginID, v))
}

// LoginIDContainsFold applies the ContainsFold predicate on the "login_id" field.
func LoginIDContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldLoginID, v))
}

// HashedPasswordEQ applies the EQ predicate on the "hashed_password" field.
func HashedPasswordEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldHashedPassword, v))
}

// HashedPasswordNEQ applies the NEQ predicate on the "hashed_password" field.
func HashedPasswordNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldHashedPassword, v))
}

// HashedPasswordIn applies the In predicate on the "hashed_password" field.
func HashedPasswordIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldHashedPassword, vs...))
}

// HashedPasswordNotIn applies the NotIn predicate on the "hashed_password" field.
func HashedPasswordNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldHashedPassword, vs...))
}

// HashedPasswordGT applies the GT predicate on the "hashed_password" field.
func HashedPasswordGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldHashedPassword, v))
}

// HashedPasswordGTE applies the GTE predicate on the "hashed_password" field.
func HashedPasswordGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldHashedPassword, v))
}

// HashedPasswordLT applies the LT predicate on the "hashed_password" field.
func HashedPasswordLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldHashedPassword, v))
}

// HashedPasswordLTE applies the LTE predicate on the "hashed_password" field.
func HashedPasswordLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldHashedPassword, v))
}

// HashedPasswordContains applies the Contains predicate on the "hashed_password" field.
func HashedPasswordContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldHashedPassword, v))
}

// HashedPasswordHasPrefix applies the HasPrefix predicate on the "hashed_password" field.
func HashedPasswordHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldHashedPassword, v))
}

// HashedPasswordHasSuffix applies the HasSuffix predicate on the "hashed_password" field.
func HashedPasswordHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldHashedPassword, v))
}

// HashedPasswordEqualFold applies the EqualFold predicate on the "hashed_password" field.
func HashedPasswordEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldHashedPassword, v))
}

// HashedPasswordContainsFold applies the ContainsFold predicate on the "hashed_password" field.
func HashedPasswordContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldHashedPassword, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldName, v))
}

// EmailEQ applies the EQ predicate on the "email" field.
func EmailEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldEmail, v))
}

// EmailNEQ applies the NEQ predicate on the "email" field.
func EmailNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldEmail, v))
}

// EmailIn applies the In predicate on the "email" field.
func EmailIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldEmail, vs...))
}

// EmailNotIn applies the NotIn predicate on the "email" field.
func EmailNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldEmail, vs...))
}

// EmailGT applies the GT predicate on the "email" field.
func EmailGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldEmail, v))
}

// EmailGTE applies the GTE predicate on the "email" field.
func EmailGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldEmail, v))
}

// EmailLT applies the LT predicate on the "email" field.
func EmailLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldEmail, v))
}

// EmailLTE applies the LTE predicate on the "email" field.
func EmailLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldEmail, v))
}

// EmailContains applies the Contains predicate on the "email" field.
func EmailContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldEmail, v))
}

// EmailHasPrefix applies the HasPrefix predicate on the "email" field.
func EmailHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldEmail, v))
}

// EmailHasSuffix applies the HasSuffix predicate on the "email" field.
func EmailHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldEmail, v))
}

// EmailEqualFold applies the EqualFold predicate on the "email" field.
func EmailEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldEmail, v))
}

// EmailContainsFold applies the ContainsFold predicate on the "email" field.
func EmailContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldEmail, v))
}

// PhoneNumberEQ applies the EQ predicate on the "phone_number" field.
func PhoneNumberEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldPhoneNumber, v))
}

// PhoneNumberNEQ applies the NEQ predicate on the "phone_number" field.
func PhoneNumberNEQ(v string) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldPhoneNumber, v))
}

// PhoneNumberIn applies the In predicate on the "phone_number" field.
func PhoneNumberIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldPhoneNumber, vs...))
}

// PhoneNumberNotIn applies the NotIn predicate on the "phone_number" field.
func PhoneNumberNotIn(vs ...string) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldPhoneNumber, vs...))
}

// PhoneNumberGT applies the GT predicate on the "phone_number" field.
func PhoneNumberGT(v string) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldPhoneNumber, v))
}

// PhoneNumberGTE applies the GTE predicate on the "phone_number" field.
func PhoneNumberGTE(v string) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldPhoneNumber, v))
}

// PhoneNumberLT applies the LT predicate on the "phone_number" field.
func PhoneNumberLT(v string) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldPhoneNumber, v))
}

// PhoneNumberLTE applies the LTE predicate on the "phone_number" field.
func PhoneNumberLTE(v string) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldPhoneNumber, v))
}

// PhoneNumberContains applies the Contains predicate on the "phone_number" field.
func PhoneNumberContains(v string) predicate.Account {
	return predicate.Account(sql.FieldContains(FieldPhoneNumber, v))
}

// PhoneNumberHasPrefix applies the HasPrefix predicate on the "phone_number" field.
func PhoneNumberHasPrefix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasPrefix(FieldPhoneNumber, v))
}

// PhoneNumberHasSuffix applies the HasSuffix predicate on the "phone_number" field.
func PhoneNumberHasSuffix(v string) predicate.Account {
	return predicate.Account(sql.FieldHasSuffix(FieldPhoneNumber, v))
}

// PhoneNumberEqualFold applies the EqualFold predicate on the "phone_number" field.
func PhoneNumberEqualFold(v string) predicate.Account {
	return predicate.Account(sql.FieldEqualFold(FieldPhoneNumber, v))
}

// PhoneNumberContainsFold applies the ContainsFold predicate on the "phone_number" field.
func PhoneNumberContainsFold(v string) predicate.Account {
	return predicate.Account(sql.FieldContainsFold(FieldPhoneNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Account {
	return predicate.Account(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Account {
	return predicate.Account(sql.FieldLTE(FieldCreatedAt, v))
}

// HasRole applies the HasEdge predicate on the "role" edge.
func HasRole() predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoleWith applies the HasEdge predicate on the "role" edge with a given conditions (other predicates).
func HasRoleWith(preds ...predicate.Role) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		step := newRoleStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
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
func Not(p predicate.Account) predicate.Account {
	return predicate.Account(func(s *sql.Selector) {
		p(s.Not())
	})
}
