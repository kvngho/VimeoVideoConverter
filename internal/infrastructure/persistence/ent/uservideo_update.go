// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/predicate"
	"github.com/kvngho/vimeovideoconverter/internal/infrastructure/persistence/ent/uservideo"
)

// UserVideoUpdate is the builder for updating UserVideo entities.
type UserVideoUpdate struct {
	config
	hooks    []Hook
	mutation *UserVideoMutation
}

// Where appends a list predicates to the UserVideoUpdate builder.
func (uvu *UserVideoUpdate) Where(ps ...predicate.UserVideo) *UserVideoUpdate {
	uvu.mutation.Where(ps...)
	return uvu
}

// SetVideoID sets the "video_id" field.
func (uvu *UserVideoUpdate) SetVideoID(i int) *UserVideoUpdate {
	uvu.mutation.ResetVideoID()
	uvu.mutation.SetVideoID(i)
	return uvu
}

// SetNillableVideoID sets the "video_id" field if the given value is not nil.
func (uvu *UserVideoUpdate) SetNillableVideoID(i *int) *UserVideoUpdate {
	if i != nil {
		uvu.SetVideoID(*i)
	}
	return uvu
}

// AddVideoID adds i to the "video_id" field.
func (uvu *UserVideoUpdate) AddVideoID(i int) *UserVideoUpdate {
	uvu.mutation.AddVideoID(i)
	return uvu
}

// SetVideoURL sets the "video_url" field.
func (uvu *UserVideoUpdate) SetVideoURL(s string) *UserVideoUpdate {
	uvu.mutation.SetVideoURL(s)
	return uvu
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (uvu *UserVideoUpdate) SetNillableVideoURL(s *string) *UserVideoUpdate {
	if s != nil {
		uvu.SetVideoURL(*s)
	}
	return uvu
}

// SetPlayableVideo sets the "playable_video" field.
func (uvu *UserVideoUpdate) SetPlayableVideo(s string) *UserVideoUpdate {
	uvu.mutation.SetPlayableVideo(s)
	return uvu
}

// SetNillablePlayableVideo sets the "playable_video" field if the given value is not nil.
func (uvu *UserVideoUpdate) SetNillablePlayableVideo(s *string) *UserVideoUpdate {
	if s != nil {
		uvu.SetPlayableVideo(*s)
	}
	return uvu
}

// Mutation returns the UserVideoMutation object of the builder.
func (uvu *UserVideoUpdate) Mutation() *UserVideoMutation {
	return uvu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uvu *UserVideoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uvu.sqlSave, uvu.mutation, uvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uvu *UserVideoUpdate) SaveX(ctx context.Context) int {
	affected, err := uvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uvu *UserVideoUpdate) Exec(ctx context.Context) error {
	_, err := uvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvu *UserVideoUpdate) ExecX(ctx context.Context) {
	if err := uvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvu *UserVideoUpdate) check() error {
	if v, ok := uvu.mutation.VideoID(); ok {
		if err := uservideo.VideoIDValidator(v); err != nil {
			return &ValidationError{Name: "video_id", err: fmt.Errorf(`ent: validator failed for field "UserVideo.video_id": %w`, err)}
		}
	}
	return nil
}

func (uvu *UserVideoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(uservideo.Table, uservideo.Columns, sqlgraph.NewFieldSpec(uservideo.FieldID, field.TypeInt))
	if ps := uvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvu.mutation.VideoID(); ok {
		_spec.SetField(uservideo.FieldVideoID, field.TypeInt, value)
	}
	if value, ok := uvu.mutation.AddedVideoID(); ok {
		_spec.AddField(uservideo.FieldVideoID, field.TypeInt, value)
	}
	if value, ok := uvu.mutation.VideoURL(); ok {
		_spec.SetField(uservideo.FieldVideoURL, field.TypeString, value)
	}
	if value, ok := uvu.mutation.PlayableVideo(); ok {
		_spec.SetField(uservideo.FieldPlayableVideo, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uservideo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uvu.mutation.done = true
	return n, nil
}

// UserVideoUpdateOne is the builder for updating a single UserVideo entity.
type UserVideoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserVideoMutation
}

// SetVideoID sets the "video_id" field.
func (uvuo *UserVideoUpdateOne) SetVideoID(i int) *UserVideoUpdateOne {
	uvuo.mutation.ResetVideoID()
	uvuo.mutation.SetVideoID(i)
	return uvuo
}

// SetNillableVideoID sets the "video_id" field if the given value is not nil.
func (uvuo *UserVideoUpdateOne) SetNillableVideoID(i *int) *UserVideoUpdateOne {
	if i != nil {
		uvuo.SetVideoID(*i)
	}
	return uvuo
}

// AddVideoID adds i to the "video_id" field.
func (uvuo *UserVideoUpdateOne) AddVideoID(i int) *UserVideoUpdateOne {
	uvuo.mutation.AddVideoID(i)
	return uvuo
}

// SetVideoURL sets the "video_url" field.
func (uvuo *UserVideoUpdateOne) SetVideoURL(s string) *UserVideoUpdateOne {
	uvuo.mutation.SetVideoURL(s)
	return uvuo
}

// SetNillableVideoURL sets the "video_url" field if the given value is not nil.
func (uvuo *UserVideoUpdateOne) SetNillableVideoURL(s *string) *UserVideoUpdateOne {
	if s != nil {
		uvuo.SetVideoURL(*s)
	}
	return uvuo
}

// SetPlayableVideo sets the "playable_video" field.
func (uvuo *UserVideoUpdateOne) SetPlayableVideo(s string) *UserVideoUpdateOne {
	uvuo.mutation.SetPlayableVideo(s)
	return uvuo
}

// SetNillablePlayableVideo sets the "playable_video" field if the given value is not nil.
func (uvuo *UserVideoUpdateOne) SetNillablePlayableVideo(s *string) *UserVideoUpdateOne {
	if s != nil {
		uvuo.SetPlayableVideo(*s)
	}
	return uvuo
}

// Mutation returns the UserVideoMutation object of the builder.
func (uvuo *UserVideoUpdateOne) Mutation() *UserVideoMutation {
	return uvuo.mutation
}

// Where appends a list predicates to the UserVideoUpdate builder.
func (uvuo *UserVideoUpdateOne) Where(ps ...predicate.UserVideo) *UserVideoUpdateOne {
	uvuo.mutation.Where(ps...)
	return uvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uvuo *UserVideoUpdateOne) Select(field string, fields ...string) *UserVideoUpdateOne {
	uvuo.fields = append([]string{field}, fields...)
	return uvuo
}

// Save executes the query and returns the updated UserVideo entity.
func (uvuo *UserVideoUpdateOne) Save(ctx context.Context) (*UserVideo, error) {
	return withHooks(ctx, uvuo.sqlSave, uvuo.mutation, uvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uvuo *UserVideoUpdateOne) SaveX(ctx context.Context) *UserVideo {
	node, err := uvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uvuo *UserVideoUpdateOne) Exec(ctx context.Context) error {
	_, err := uvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uvuo *UserVideoUpdateOne) ExecX(ctx context.Context) {
	if err := uvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uvuo *UserVideoUpdateOne) check() error {
	if v, ok := uvuo.mutation.VideoID(); ok {
		if err := uservideo.VideoIDValidator(v); err != nil {
			return &ValidationError{Name: "video_id", err: fmt.Errorf(`ent: validator failed for field "UserVideo.video_id": %w`, err)}
		}
	}
	return nil
}

func (uvuo *UserVideoUpdateOne) sqlSave(ctx context.Context) (_node *UserVideo, err error) {
	if err := uvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(uservideo.Table, uservideo.Columns, sqlgraph.NewFieldSpec(uservideo.FieldID, field.TypeInt))
	id, ok := uvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserVideo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, uservideo.FieldID)
		for _, f := range fields {
			if !uservideo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != uservideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uvuo.mutation.VideoID(); ok {
		_spec.SetField(uservideo.FieldVideoID, field.TypeInt, value)
	}
	if value, ok := uvuo.mutation.AddedVideoID(); ok {
		_spec.AddField(uservideo.FieldVideoID, field.TypeInt, value)
	}
	if value, ok := uvuo.mutation.VideoURL(); ok {
		_spec.SetField(uservideo.FieldVideoURL, field.TypeString, value)
	}
	if value, ok := uvuo.mutation.PlayableVideo(); ok {
		_spec.SetField(uservideo.FieldPlayableVideo, field.TypeString, value)
	}
	_node = &UserVideo{config: uvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{uservideo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uvuo.mutation.done = true
	return _node, nil
}
