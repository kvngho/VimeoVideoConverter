package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// DeepingTalk holds the schema definition for the DeepingTalk entity.
type DeepingTalk struct {
	ent.Schema
}

// Fields of the ProductReview.
func (DeepingTalk) Fields() []ent.Field {
	return []ent.Field{
		field.String("talk_video"),
		field.String("playable_video"),
	}
}

// Edges of the ProductReview.
func (DeepingTalk) Edges() []ent.Edge {
	return nil
}

func (DeepingTalk) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "deepingtalk_deepingtalk"},
	}
}
