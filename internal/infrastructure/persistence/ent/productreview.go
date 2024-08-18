// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/productreview"
)

// ProductReview is the model entity for the ProductReview schema.
type ProductReview struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// PlayableVideo holds the value of the "playable_video" field.
	PlayableVideo string `json:"playable_video,omitempty"`
	// ReviewVideo holds the value of the "review_video" field.
	ReviewVideo  string `json:"review_video,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductReview) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case productreview.FieldID:
			values[i] = new(sql.NullInt64)
		case productreview.FieldPlayableVideo, productreview.FieldReviewVideo:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductReview fields.
func (pr *ProductReview) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case productreview.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case productreview.FieldPlayableVideo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field playable_video", values[i])
			} else if value.Valid {
				pr.PlayableVideo = value.String
			}
		case productreview.FieldReviewVideo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field review_video", values[i])
			} else if value.Valid {
				pr.ReviewVideo = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductReview.
// This includes values selected through modifiers, order, etc.
func (pr *ProductReview) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// Update returns a builder for updating this ProductReview.
// Note that you need to call ProductReview.Unwrap() before calling this method if this ProductReview
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *ProductReview) Update() *ProductReviewUpdateOne {
	return NewProductReviewClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the ProductReview entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *ProductReview) Unwrap() *ProductReview {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductReview is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *ProductReview) String() string {
	var builder strings.Builder
	builder.WriteString("ProductReview(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("playable_video=")
	builder.WriteString(pr.PlayableVideo)
	builder.WriteString(", ")
	builder.WriteString("review_video=")
	builder.WriteString(pr.ReviewVideo)
	builder.WriteByte(')')
	return builder.String()
}

// ProductReviews is a parsable slice of ProductReview.
type ProductReviews []*ProductReview
