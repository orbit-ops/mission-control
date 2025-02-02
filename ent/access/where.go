// Code generated by ent, DO NOT EDIT.

package access

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldID, id))
}

// StartTime applies equality check predicate on the "start_time" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldStartTime, v))
}

// RolledBack applies equality check predicate on the "rolled_back" field. It's identical to RolledBackEQ.
func RolledBack(v bool) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRolledBack, v))
}

// RollbackTime applies equality check predicate on the "rollback_time" field. It's identical to RollbackTimeEQ.
func RollbackTime(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRollbackTime, v))
}

// RollbackReason applies equality check predicate on the "rollback_reason" field. It's identical to RollbackReasonEQ.
func RollbackReason(v string) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRollbackReason, v))
}

// Expiration applies equality check predicate on the "expiration" field. It's identical to ExpirationEQ.
func Expiration(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldExpiration, v))
}

// StartTimeEQ applies the EQ predicate on the "start_time" field.
func StartTimeEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "start_time" field.
func StartTimeNEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "start_time" field.
func StartTimeIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "start_time" field.
func StartTimeNotIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "start_time" field.
func StartTimeGT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "start_time" field.
func StartTimeGTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "start_time" field.
func StartTimeLT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "start_time" field.
func StartTimeLTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldStartTime, v))
}

// RolledBackEQ applies the EQ predicate on the "rolled_back" field.
func RolledBackEQ(v bool) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRolledBack, v))
}

// RolledBackNEQ applies the NEQ predicate on the "rolled_back" field.
func RolledBackNEQ(v bool) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldRolledBack, v))
}

// RollbackTimeEQ applies the EQ predicate on the "rollback_time" field.
func RollbackTimeEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRollbackTime, v))
}

// RollbackTimeNEQ applies the NEQ predicate on the "rollback_time" field.
func RollbackTimeNEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldRollbackTime, v))
}

// RollbackTimeIn applies the In predicate on the "rollback_time" field.
func RollbackTimeIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldRollbackTime, vs...))
}

// RollbackTimeNotIn applies the NotIn predicate on the "rollback_time" field.
func RollbackTimeNotIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldRollbackTime, vs...))
}

// RollbackTimeGT applies the GT predicate on the "rollback_time" field.
func RollbackTimeGT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldRollbackTime, v))
}

// RollbackTimeGTE applies the GTE predicate on the "rollback_time" field.
func RollbackTimeGTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldRollbackTime, v))
}

// RollbackTimeLT applies the LT predicate on the "rollback_time" field.
func RollbackTimeLT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldRollbackTime, v))
}

// RollbackTimeLTE applies the LTE predicate on the "rollback_time" field.
func RollbackTimeLTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldRollbackTime, v))
}

// RollbackReasonEQ applies the EQ predicate on the "rollback_reason" field.
func RollbackReasonEQ(v string) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldRollbackReason, v))
}

// RollbackReasonNEQ applies the NEQ predicate on the "rollback_reason" field.
func RollbackReasonNEQ(v string) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldRollbackReason, v))
}

// RollbackReasonIn applies the In predicate on the "rollback_reason" field.
func RollbackReasonIn(vs ...string) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldRollbackReason, vs...))
}

// RollbackReasonNotIn applies the NotIn predicate on the "rollback_reason" field.
func RollbackReasonNotIn(vs ...string) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldRollbackReason, vs...))
}

// RollbackReasonGT applies the GT predicate on the "rollback_reason" field.
func RollbackReasonGT(v string) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldRollbackReason, v))
}

// RollbackReasonGTE applies the GTE predicate on the "rollback_reason" field.
func RollbackReasonGTE(v string) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldRollbackReason, v))
}

// RollbackReasonLT applies the LT predicate on the "rollback_reason" field.
func RollbackReasonLT(v string) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldRollbackReason, v))
}

// RollbackReasonLTE applies the LTE predicate on the "rollback_reason" field.
func RollbackReasonLTE(v string) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldRollbackReason, v))
}

// RollbackReasonContains applies the Contains predicate on the "rollback_reason" field.
func RollbackReasonContains(v string) predicate.Access {
	return predicate.Access(sql.FieldContains(FieldRollbackReason, v))
}

// RollbackReasonHasPrefix applies the HasPrefix predicate on the "rollback_reason" field.
func RollbackReasonHasPrefix(v string) predicate.Access {
	return predicate.Access(sql.FieldHasPrefix(FieldRollbackReason, v))
}

// RollbackReasonHasSuffix applies the HasSuffix predicate on the "rollback_reason" field.
func RollbackReasonHasSuffix(v string) predicate.Access {
	return predicate.Access(sql.FieldHasSuffix(FieldRollbackReason, v))
}

// RollbackReasonIsNil applies the IsNil predicate on the "rollback_reason" field.
func RollbackReasonIsNil() predicate.Access {
	return predicate.Access(sql.FieldIsNull(FieldRollbackReason))
}

// RollbackReasonNotNil applies the NotNil predicate on the "rollback_reason" field.
func RollbackReasonNotNil() predicate.Access {
	return predicate.Access(sql.FieldNotNull(FieldRollbackReason))
}

// RollbackReasonEqualFold applies the EqualFold predicate on the "rollback_reason" field.
func RollbackReasonEqualFold(v string) predicate.Access {
	return predicate.Access(sql.FieldEqualFold(FieldRollbackReason, v))
}

// RollbackReasonContainsFold applies the ContainsFold predicate on the "rollback_reason" field.
func RollbackReasonContainsFold(v string) predicate.Access {
	return predicate.Access(sql.FieldContainsFold(FieldRollbackReason, v))
}

// ExpirationEQ applies the EQ predicate on the "expiration" field.
func ExpirationEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldEQ(FieldExpiration, v))
}

// ExpirationNEQ applies the NEQ predicate on the "expiration" field.
func ExpirationNEQ(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldNEQ(FieldExpiration, v))
}

// ExpirationIn applies the In predicate on the "expiration" field.
func ExpirationIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldIn(FieldExpiration, vs...))
}

// ExpirationNotIn applies the NotIn predicate on the "expiration" field.
func ExpirationNotIn(vs ...time.Time) predicate.Access {
	return predicate.Access(sql.FieldNotIn(FieldExpiration, vs...))
}

// ExpirationGT applies the GT predicate on the "expiration" field.
func ExpirationGT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGT(FieldExpiration, v))
}

// ExpirationGTE applies the GTE predicate on the "expiration" field.
func ExpirationGTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldGTE(FieldExpiration, v))
}

// ExpirationLT applies the LT predicate on the "expiration" field.
func ExpirationLT(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLT(FieldExpiration, v))
}

// ExpirationLTE applies the LTE predicate on the "expiration" field.
func ExpirationLTE(v time.Time) predicate.Access {
	return predicate.Access(sql.FieldLTE(FieldExpiration, v))
}

// HasApprovals applies the HasEdge predicate on the "approvals" edge.
func HasApprovals() predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ApprovalsTable, ApprovalsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasApprovalsWith applies the HasEdge predicate on the "approvals" edge with a given conditions (other predicates).
func HasApprovalsWith(preds ...predicate.Approval) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := newApprovalsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRequest applies the HasEdge predicate on the "request" edge.
func HasRequest() predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RequestTable, RequestColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRequestWith applies the HasEdge predicate on the "request" edge with a given conditions (other predicates).
func HasRequestWith(preds ...predicate.Request) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := newRequestStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAccessTokens applies the HasEdge predicate on the "accessTokens" edge.
func HasAccessTokens() predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, AccessTokensTable, AccessTokensColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccessTokensWith applies the HasEdge predicate on the "accessTokens" edge with a given conditions (other predicates).
func HasAccessTokensWith(preds ...predicate.ActionTokens) predicate.Access {
	return predicate.Access(func(s *sql.Selector) {
		step := newAccessTokensStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Access) predicate.Access {
	return predicate.Access(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Access) predicate.Access {
	return predicate.Access(sql.NotPredicates(p))
}
