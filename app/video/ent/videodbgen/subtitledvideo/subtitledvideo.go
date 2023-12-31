// Code generated by ent, DO NOT EDIT.

package subtitledvideo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the subtitledvideo type in the database.
	Label = "subtitled_video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldOriginalVideoID holds the string denoting the original_video_id field in the database.
	FieldOriginalVideoID = "original_video_id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeOriginalVideo holds the string denoting the original_video edge name in mutations.
	EdgeOriginalVideo = "original_video"
	// Table holds the table name of the subtitledvideo in the database.
	Table = "subtitled_video"
	// OriginalVideoTable is the table that holds the original_video relation/edge.
	OriginalVideoTable = "subtitled_video"
	// OriginalVideoInverseTable is the table name for the OriginalVideo entity.
	// It exists in this package in order to avoid circular dependency with the "originalvideo" package.
	OriginalVideoInverseTable = "original_video"
	// OriginalVideoColumn is the table column denoting the original_video relation/edge.
	OriginalVideoColumn = "original_video_id"
)

// Columns holds all SQL columns for subtitledvideo fields.
var Columns = []string{
	FieldID,
	FieldOriginalVideoID,
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

// OrderOption defines the ordering options for the SubtitledVideo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByOriginalVideoID orders the results by the original_video_id field.
func ByOriginalVideoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOriginalVideoID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByOriginalVideoField orders the results by original_video field.
func ByOriginalVideoField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOriginalVideoStep(), sql.OrderByField(field, opts...))
	}
}
func newOriginalVideoStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OriginalVideoInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OriginalVideoTable, OriginalVideoColumn),
	)
}
