package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// UserVideo holds the schema definition for the UserVideo entity.
type UserVideo struct {
	ent.Schema
}

// Fields of the UserVideo.
func (UserVideo) Fields() []ent.Field {
	return []ent.Field{
		field.Int("video_id").Positive(),
		field.String("video_url"),
	}
}

// Edges of the UserVideo.
func (UserVideo) Edges() []ent.Edge {
	return nil
}

func (UserVideo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "account_uservideo"},
	}
}
