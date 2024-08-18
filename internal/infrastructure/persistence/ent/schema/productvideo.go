package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ProductVideo holds the schema definition for the ProductVideo entity.
type ProductVideo struct {
	ent.Schema
}

// Fields of the ProductVideo.
func (ProductVideo) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("playable_video"),
	}
}

// Edges of the ProductVideo.
func (ProductVideo) Edges() []ent.Edge {
	return nil
}

func (ProductVideo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_productvideo"},
	}
}
