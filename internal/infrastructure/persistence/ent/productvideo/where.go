// Code generated by ent, DO NOT EDIT.

package productvideo

import (
	"entgo.io/ent/dialect/sql"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldLTE(FieldID, id))
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEQ(FieldURL, v))
}

// PlayableVideo applies equality check predicate on the "playable_video" field. It's identical to PlayableVideoEQ.
func PlayableVideo(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEQ(FieldPlayableVideo, v))
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEQ(FieldURL, v))
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldNEQ(FieldURL, v))
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldIn(FieldURL, vs...))
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldNotIn(FieldURL, vs...))
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldGT(FieldURL, v))
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldGTE(FieldURL, v))
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldLT(FieldURL, v))
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldLTE(FieldURL, v))
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldContains(FieldURL, v))
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldHasPrefix(FieldURL, v))
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldHasSuffix(FieldURL, v))
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEqualFold(FieldURL, v))
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldContainsFold(FieldURL, v))
}

// PlayableVideoEQ applies the EQ predicate on the "playable_video" field.
func PlayableVideoEQ(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEQ(FieldPlayableVideo, v))
}

// PlayableVideoNEQ applies the NEQ predicate on the "playable_video" field.
func PlayableVideoNEQ(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldNEQ(FieldPlayableVideo, v))
}

// PlayableVideoIn applies the In predicate on the "playable_video" field.
func PlayableVideoIn(vs ...string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldIn(FieldPlayableVideo, vs...))
}

// PlayableVideoNotIn applies the NotIn predicate on the "playable_video" field.
func PlayableVideoNotIn(vs ...string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldNotIn(FieldPlayableVideo, vs...))
}

// PlayableVideoGT applies the GT predicate on the "playable_video" field.
func PlayableVideoGT(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldGT(FieldPlayableVideo, v))
}

// PlayableVideoGTE applies the GTE predicate on the "playable_video" field.
func PlayableVideoGTE(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldGTE(FieldPlayableVideo, v))
}

// PlayableVideoLT applies the LT predicate on the "playable_video" field.
func PlayableVideoLT(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldLT(FieldPlayableVideo, v))
}

// PlayableVideoLTE applies the LTE predicate on the "playable_video" field.
func PlayableVideoLTE(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldLTE(FieldPlayableVideo, v))
}

// PlayableVideoContains applies the Contains predicate on the "playable_video" field.
func PlayableVideoContains(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldContains(FieldPlayableVideo, v))
}

// PlayableVideoHasPrefix applies the HasPrefix predicate on the "playable_video" field.
func PlayableVideoHasPrefix(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldHasPrefix(FieldPlayableVideo, v))
}

// PlayableVideoHasSuffix applies the HasSuffix predicate on the "playable_video" field.
func PlayableVideoHasSuffix(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldHasSuffix(FieldPlayableVideo, v))
}

// PlayableVideoEqualFold applies the EqualFold predicate on the "playable_video" field.
func PlayableVideoEqualFold(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldEqualFold(FieldPlayableVideo, v))
}

// PlayableVideoContainsFold applies the ContainsFold predicate on the "playable_video" field.
func PlayableVideoContainsFold(v string) predicate.ProductVideo {
	return predicate.ProductVideo(sql.FieldContainsFold(FieldPlayableVideo, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ProductVideo) predicate.ProductVideo {
	return predicate.ProductVideo(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ProductVideo) predicate.ProductVideo {
	return predicate.ProductVideo(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ProductVideo) predicate.ProductVideo {
	return predicate.ProductVideo(sql.NotPredicates(p))
}
