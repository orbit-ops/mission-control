// Code generated by ent, DO NOT EDIT.

package approval

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the approval type in the database.
	Label = "approval"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPerson holds the string denoting the person field in the database.
	FieldPerson = "person"
	// FieldApprovedTime holds the string denoting the approved_time field in the database.
	FieldApprovedTime = "approved_time"
	// FieldApproved holds the string denoting the approved field in the database.
	FieldApproved = "approved"
	// FieldRevoked holds the string denoting the revoked field in the database.
	FieldRevoked = "revoked"
	// FieldRevokedTime holds the string denoting the revoked_time field in the database.
	FieldRevokedTime = "revoked_time"
	// EdgeRequest holds the string denoting the request edge name in mutations.
	EdgeRequest = "request"
	// EdgeAccess holds the string denoting the access edge name in mutations.
	EdgeAccess = "access"
	// Table holds the table name of the approval in the database.
	Table = "approvals"
	// RequestTable is the table that holds the request relation/edge.
	RequestTable = "approvals"
	// RequestInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestInverseTable = "requests"
	// RequestColumn is the table column denoting the request relation/edge.
	RequestColumn = "approval_request"
	// AccessTable is the table that holds the access relation/edge.
	AccessTable = "approvals"
	// AccessInverseTable is the table name for the Access entity.
	// It exists in this package in order to avoid circular dependency with the "access" package.
	AccessInverseTable = "accesses"
	// AccessColumn is the table column denoting the access relation/edge.
	AccessColumn = "access_approvals"
)

// Columns holds all SQL columns for approval fields.
var Columns = []string{
	FieldID,
	FieldPerson,
	FieldApprovedTime,
	FieldApproved,
	FieldRevoked,
	FieldRevokedTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "approvals"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"access_approvals",
	"approval_request",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// PersonValidator is a validator for the "person" field. It is called by the builders before save.
	PersonValidator func(string) error
	// DefaultApprovedTime holds the default value on creation for the "approved_time" field.
	DefaultApprovedTime func() time.Time
	// DefaultRevoked holds the default value on creation for the "revoked" field.
	DefaultRevoked bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Approval queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPerson orders the results by the person field.
func ByPerson(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPerson, opts...).ToFunc()
}

// ByApprovedTime orders the results by the approved_time field.
func ByApprovedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApprovedTime, opts...).ToFunc()
}

// ByApproved orders the results by the approved field.
func ByApproved(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldApproved, opts...).ToFunc()
}

// ByRevoked orders the results by the revoked field.
func ByRevoked(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevoked, opts...).ToFunc()
}

// ByRevokedTime orders the results by the revoked_time field.
func ByRevokedTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevokedTime, opts...).ToFunc()
}

// ByRequestField orders the results by request field.
func ByRequestField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequestStep(), sql.OrderByField(field, opts...))
	}
}

// ByAccessField orders the results by access field.
func ByAccessField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAccessStep(), sql.OrderByField(field, opts...))
	}
}
func newRequestStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequestInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RequestTable, RequestColumn),
	)
}
func newAccessStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AccessInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AccessTable, AccessColumn),
	)
}
