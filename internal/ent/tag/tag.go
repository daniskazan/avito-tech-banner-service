// Code generated by ent, DO NOT EDIT.

package tag

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tag type in the database.
	Label = "tag"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTagName holds the string denoting the tag_name field in the database.
	FieldTagName = "tag_name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// EdgeBanners holds the string denoting the banners edge name in mutations.
	EdgeBanners = "banners"
	// Table holds the table name of the tag in the database.
	Table = "tags"
	// BannersTable is the table that holds the banners relation/edge. The primary key declared below.
	BannersTable = "banner_tags"
	// BannersInverseTable is the table name for the Banner entity.
	// It exists in this package in order to avoid circular dependency with the "banner" package.
	BannersInverseTable = "banners"
)

// Columns holds all SQL columns for tag fields.
var Columns = []string{
	FieldID,
	FieldTagName,
	FieldCreatedAt,
	FieldUpdatedAt,
}

var (
	// BannersPrimaryKey and BannersColumn2 are the table columns denoting the
	// primary key for the banners relation (M2M).
	BannersPrimaryKey = []string{"banner_id", "tag_id"}
)

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
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Tag queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTagName orders the results by the tag_name field.
func ByTagName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByBannersCount orders the results by banners count.
func ByBannersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newBannersStep(), opts...)
	}
}

// ByBanners orders the results by banners terms.
func ByBanners(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBannersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newBannersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BannersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, BannersTable, BannersPrimaryKey...),
	)
}