package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserVideo holds the schema definition for the UserVideo entity.
type ExperVideo struct {
	ent.Schema
}

// Fields of the UserVideo.
func (ExperVideo) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.String("playable_video"),
	}
}

// Edges of the UserVideo.
func (ExperVideo) Edges() []ent.Edge {
	return nil
}

func (ExperVideo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "product_productexpertvideo"},
	}
}
