// Code generated by ent, DO NOT EDIT.

package hype

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the hype type in the database.
	Label = "hype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAmountRemaining holds the string denoting the amount_remaining field in the database.
	FieldAmountRemaining = "amount_remaining"
	// FieldMaxHype holds the string denoting the max_hype field in the database.
	FieldMaxHype = "max_hype"
	// FieldLastUpdatedAt holds the string denoting the last_updated_at field in the database.
	FieldLastUpdatedAt = "last_updated_at"
	// FieldHypePerMinute holds the string denoting the hype_per_minute field in the database.
	FieldHypePerMinute = "hype_per_minute"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// Table holds the table name of the hype in the database.
	Table = "hypes"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "hypes"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_hype"
)

// Columns holds all SQL columns for hype fields.
var Columns = []string{
	FieldID,
	FieldAmountRemaining,
	FieldMaxHype,
	FieldLastUpdatedAt,
	FieldHypePerMinute,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "hypes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_hype",
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
	// DefaultLastUpdatedAt holds the default value on creation for the "last_updated_at" field.
	DefaultLastUpdatedAt func() time.Time
	// UpdateDefaultLastUpdatedAt holds the default value on update for the "last_updated_at" field.
	UpdateDefaultLastUpdatedAt func() time.Time
	// DefaultHypePerMinute holds the default value on creation for the "hype_per_minute" field.
	DefaultHypePerMinute int
)

// OrderOption defines the ordering options for the Hype queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAmountRemaining orders the results by the amount_remaining field.
func ByAmountRemaining(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAmountRemaining, opts...).ToFunc()
}

// ByMaxHype orders the results by the max_hype field.
func ByMaxHype(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMaxHype, opts...).ToFunc()
}

// ByLastUpdatedAt orders the results by the last_updated_at field.
func ByLastUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastUpdatedAt, opts...).ToFunc()
}

// ByHypePerMinute orders the results by the hype_per_minute field.
func ByHypePerMinute(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHypePerMinute, opts...).ToFunc()
}

// ByUserField orders the results by user field.
func ByUserField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUserStep(), sql.OrderByField(field, opts...))
	}
}
func newUserStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UserInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
	)
}
