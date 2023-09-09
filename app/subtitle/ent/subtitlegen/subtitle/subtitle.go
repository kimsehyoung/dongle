// Code generated by ent, DO NOT EDIT.

package subtitle

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subtitle type in the database.
	Label = "subtitle"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVideoID holds the string denoting the video_id field in the database.
	FieldVideoID = "video_id"
	// FieldLanguage holds the string denoting the language field in the database.
	FieldLanguage = "language"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeRole holds the string denoting the role edge name in mutations.
	EdgeRole = "role"
	// Table holds the table name of the subtitle in the database.
	Table = "subtitle"
	// RoleTable is the table that holds the role relation/edge.
	RoleTable = "subtitle"
	// RoleInverseTable is the table name for the Video entity.
	// It exists in this package in order to avoid circular dependency with the "video" package.
	RoleInverseTable = "video"
	// RoleColumn is the table column denoting the role relation/edge.
	RoleColumn = "video_id"
)

// Columns holds all SQL columns for subtitle fields.
var Columns = []string{
	FieldID,
	FieldVideoID,
	FieldLanguage,
	FieldURL,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
)

// OrderOption defines the ordering options for the Subtitle queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVideoID orders the results by the video_id field.
func ByVideoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoID, opts...).ToFunc()
}

// ByLanguage orders the results by the language field.
func ByLanguage(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguage, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByRoleField orders the results by role field.
func ByRoleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRoleStep(), sql.OrderByField(field, opts...))
	}
}
func newRoleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RoleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RoleTable, RoleColumn),
	)
}