// Code generated by ent, DO NOT EDIT.

package mission

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Mission {
	return predicate.Mission(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldDescription, v))
}

// Duration applies equality check predicate on the "duration" field. It's identical to DurationEQ.
func Duration(v int) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldDuration, v))
}

// MinApprovers applies equality check predicate on the "min_approvers" field. It's identical to MinApproversEQ.
func MinApprovers(v int) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldMinApprovers, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Mission {
	return predicate.Mission(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Mission {
	return predicate.Mission(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Mission {
	return predicate.Mission(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Mission {
	return predicate.Mission(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Mission {
	return predicate.Mission(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Mission {
	return predicate.Mission(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Mission {
	return predicate.Mission(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Mission {
	return predicate.Mission(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Mission {
	return predicate.Mission(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Mission {
	return predicate.Mission(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Mission {
	return predicate.Mission(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Mission {
	return predicate.Mission(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Mission {
	return predicate.Mission(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Mission {
	return predicate.Mission(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Mission {
	return predicate.Mission(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Mission {
	return predicate.Mission(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Mission {
	return predicate.Mission(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Mission {
	return predicate.Mission(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Mission {
	return predicate.Mission(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Mission {
	return predicate.Mission(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Mission {
	return predicate.Mission(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Mission {
	return predicate.Mission(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Mission {
	return predicate.Mission(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Mission {
	return predicate.Mission(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Mission {
	return predicate.Mission(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Mission {
	return predicate.Mission(sql.FieldContainsFold(FieldDescription, v))
}

// DurationEQ applies the EQ predicate on the "duration" field.
func DurationEQ(v int) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldDuration, v))
}

// DurationNEQ applies the NEQ predicate on the "duration" field.
func DurationNEQ(v int) predicate.Mission {
	return predicate.Mission(sql.FieldNEQ(FieldDuration, v))
}

// DurationIn applies the In predicate on the "duration" field.
func DurationIn(vs ...int) predicate.Mission {
	return predicate.Mission(sql.FieldIn(FieldDuration, vs...))
}

// DurationNotIn applies the NotIn predicate on the "duration" field.
func DurationNotIn(vs ...int) predicate.Mission {
	return predicate.Mission(sql.FieldNotIn(FieldDuration, vs...))
}

// DurationGT applies the GT predicate on the "duration" field.
func DurationGT(v int) predicate.Mission {
	return predicate.Mission(sql.FieldGT(FieldDuration, v))
}

// DurationGTE applies the GTE predicate on the "duration" field.
func DurationGTE(v int) predicate.Mission {
	return predicate.Mission(sql.FieldGTE(FieldDuration, v))
}

// DurationLT applies the LT predicate on the "duration" field.
func DurationLT(v int) predicate.Mission {
	return predicate.Mission(sql.FieldLT(FieldDuration, v))
}

// DurationLTE applies the LTE predicate on the "duration" field.
func DurationLTE(v int) predicate.Mission {
	return predicate.Mission(sql.FieldLTE(FieldDuration, v))
}

// MinApproversEQ applies the EQ predicate on the "min_approvers" field.
func MinApproversEQ(v int) predicate.Mission {
	return predicate.Mission(sql.FieldEQ(FieldMinApprovers, v))
}

// MinApproversNEQ applies the NEQ predicate on the "min_approvers" field.
func MinApproversNEQ(v int) predicate.Mission {
	return predicate.Mission(sql.FieldNEQ(FieldMinApprovers, v))
}

// MinApproversIn applies the In predicate on the "min_approvers" field.
func MinApproversIn(vs ...int) predicate.Mission {
	return predicate.Mission(sql.FieldIn(FieldMinApprovers, vs...))
}

// MinApproversNotIn applies the NotIn predicate on the "min_approvers" field.
func MinApproversNotIn(vs ...int) predicate.Mission {
	return predicate.Mission(sql.FieldNotIn(FieldMinApprovers, vs...))
}

// MinApproversGT applies the GT predicate on the "min_approvers" field.
func MinApproversGT(v int) predicate.Mission {
	return predicate.Mission(sql.FieldGT(FieldMinApprovers, v))
}

// MinApproversGTE applies the GTE predicate on the "min_approvers" field.
func MinApproversGTE(v int) predicate.Mission {
	return predicate.Mission(sql.FieldGTE(FieldMinApprovers, v))
}

// MinApproversLT applies the LT predicate on the "min_approvers" field.
func MinApproversLT(v int) predicate.Mission {
	return predicate.Mission(sql.FieldLT(FieldMinApprovers, v))
}

// MinApproversLTE applies the LTE predicate on the "min_approvers" field.
func MinApproversLTE(v int) predicate.Mission {
	return predicate.Mission(sql.FieldLTE(FieldMinApprovers, v))
}

// HasRockets applies the HasEdge predicate on the "rockets" edge.
func HasRockets() predicate.Mission {
	return predicate.Mission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RocketsTable, RocketsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRocketsWith applies the HasEdge predicate on the "rockets" edge with a given conditions (other predicates).
func HasRocketsWith(preds ...predicate.Rocket) predicate.Mission {
	return predicate.Mission(func(s *sql.Selector) {
		step := newRocketsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRequests applies the HasEdge predicate on the "requests" edge.
func HasRequests() predicate.Mission {
	return predicate.Mission(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, RequestsTable, RequestsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRequestsWith applies the HasEdge predicate on the "requests" edge with a given conditions (other predicates).
func HasRequestsWith(preds ...predicate.Request) predicate.Mission {
	return predicate.Mission(func(s *sql.Selector) {
		step := newRequestsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Mission) predicate.Mission {
	return predicate.Mission(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Mission) predicate.Mission {
	return predicate.Mission(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Mission) predicate.Mission {
	return predicate.Mission(sql.NotPredicates(p))
}
