// Code generated by ent, DO NOT EDIT.

package productvideo

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the productvideo type in the database.
	Label = "product_video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldURL holds the string denoting the url field in the database.
	FieldURL = "url"
	// FieldPlayableVideo holds the string denoting the playable_video field in the database.
	FieldPlayableVideo = "playable_video"
	// Table holds the table name of the productvideo in the database.
	Table = "product_productvideo"
)

// Columns holds all SQL columns for productvideo fields.
var Columns = []string{
	FieldID,
	FieldURL,
	FieldPlayableVideo,
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

// OrderOption defines the ordering options for the ProductVideo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByURL orders the results by the url field.
func ByURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldURL, opts...).ToFunc()
}

// ByPlayableVideo orders the results by the playable_video field.
func ByPlayableVideo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlayableVideo, opts...).ToFunc()
}
