package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ProductReview holds the schema definition for the ProductReview entity.
type ProductReview struct {
	ent.Schema
}

// Fields of the ProductReview.
func (ProductReview) Fields() []ent.Field {
	return []ent.Field{
		field.String("playable_video"),
		field.String("review_video"),
	}
}

// Edges of the ProductReview.
func (ProductReview) Edges() []ent.Edge {
	return nil
}

func (ProductReview) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "deepingreview_productreview"},
	}
}
