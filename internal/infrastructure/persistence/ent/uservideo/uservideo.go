// Code generated by ent, DO NOT EDIT.

package uservideo

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the uservideo type in the database.
	Label = "user_video"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVideoID holds the string denoting the video_id field in the database.
	FieldVideoID = "video_id"
	// FieldVideoURL holds the string denoting the video_url field in the database.
	FieldVideoURL = "video_url"
	// FieldPlayableVideo holds the string denoting the playable_video field in the database.
	FieldPlayableVideo = "playable_video"
	// Table holds the table name of the uservideo in the database.
	Table = "account_uservideo"
)

// Columns holds all SQL columns for uservideo fields.
var Columns = []string{
	FieldID,
	FieldVideoID,
	FieldVideoURL,
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

var (
	// VideoIDValidator is a validator for the "video_id" field. It is called by the builders before save.
	VideoIDValidator func(int) error
)

// OrderOption defines the ordering options for the UserVideo queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVideoID orders the results by the video_id field.
func ByVideoID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoID, opts...).ToFunc()
}

// ByVideoURL orders the results by the video_url field.
func ByVideoURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoURL, opts...).ToFunc()
}

// ByPlayableVideo orders the results by the playable_video field.
func ByPlayableVideo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlayableVideo, opts...).ToFunc()
}
