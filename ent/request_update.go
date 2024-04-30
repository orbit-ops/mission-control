// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

// RequestUpdate is the builder for updating Request entities.
type RequestUpdate struct {
	config
	hooks    []Hook
	mutation *RequestMutation
}

// Where appends a list predicates to the RequestUpdate builder.
func (ru *RequestUpdate) Where(ps ...predicate.Request) *RequestUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetCancelledTime sets the "cancelled_time" field.
func (ru *RequestUpdate) SetCancelledTime(t time.Time) *RequestUpdate {
	ru.mutation.SetCancelledTime(t)
	return ru
}

// SetNillableCancelledTime sets the "cancelled_time" field if the given value is not nil.
func (ru *RequestUpdate) SetNillableCancelledTime(t *time.Time) *RequestUpdate {
	if t != nil {
		ru.SetCancelledTime(*t)
	}
	return ru
}

// ClearCancelledTime clears the value of the "cancelled_time" field.
func (ru *RequestUpdate) ClearCancelledTime() *RequestUpdate {
	ru.mutation.ClearCancelledTime()
	return ru
}

// SetCancelled sets the "cancelled" field.
func (ru *RequestUpdate) SetCancelled(b bool) *RequestUpdate {
	ru.mutation.SetCancelled(b)
	return ru
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (ru *RequestUpdate) SetNillableCancelled(b *bool) *RequestUpdate {
	if b != nil {
		ru.SetCancelled(*b)
	}
	return ru
}

// Mutation returns the RequestMutation object of the builder.
func (ru *RequestUpdate) Mutation() *RequestMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RequestUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RequestUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RequestUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RequestUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RequestUpdate) check() error {
	if _, ok := ru.mutation.MissionID(); ru.mutation.MissionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Request.mission"`)
	}
	return nil
}

func (ru *RequestUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(request.Table, request.Columns, sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.CancelledTime(); ok {
		_spec.SetField(request.FieldCancelledTime, field.TypeTime, value)
	}
	if ru.mutation.CancelledTimeCleared() {
		_spec.ClearField(request.FieldCancelledTime, field.TypeTime)
	}
	if value, ok := ru.mutation.Cancelled(); ok {
		_spec.SetField(request.FieldCancelled, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{request.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RequestUpdateOne is the builder for updating a single Request entity.
type RequestUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RequestMutation
}

// SetCancelledTime sets the "cancelled_time" field.
func (ruo *RequestUpdateOne) SetCancelledTime(t time.Time) *RequestUpdateOne {
	ruo.mutation.SetCancelledTime(t)
	return ruo
}

// SetNillableCancelledTime sets the "cancelled_time" field if the given value is not nil.
func (ruo *RequestUpdateOne) SetNillableCancelledTime(t *time.Time) *RequestUpdateOne {
	if t != nil {
		ruo.SetCancelledTime(*t)
	}
	return ruo
}

// ClearCancelledTime clears the value of the "cancelled_time" field.
func (ruo *RequestUpdateOne) ClearCancelledTime() *RequestUpdateOne {
	ruo.mutation.ClearCancelledTime()
	return ruo
}

// SetCancelled sets the "cancelled" field.
func (ruo *RequestUpdateOne) SetCancelled(b bool) *RequestUpdateOne {
	ruo.mutation.SetCancelled(b)
	return ruo
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (ruo *RequestUpdateOne) SetNillableCancelled(b *bool) *RequestUpdateOne {
	if b != nil {
		ruo.SetCancelled(*b)
	}
	return ruo
}

// Mutation returns the RequestMutation object of the builder.
func (ruo *RequestUpdateOne) Mutation() *RequestMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RequestUpdate builder.
func (ruo *RequestUpdateOne) Where(ps ...predicate.Request) *RequestUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RequestUpdateOne) Select(field string, fields ...string) *RequestUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Request entity.
func (ruo *RequestUpdateOne) Save(ctx context.Context) (*Request, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RequestUpdateOne) SaveX(ctx context.Context) *Request {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RequestUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RequestUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RequestUpdateOne) check() error {
	if _, ok := ruo.mutation.MissionID(); ruo.mutation.MissionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Request.mission"`)
	}
	return nil
}

func (ruo *RequestUpdateOne) sqlSave(ctx context.Context) (_node *Request, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(request.Table, request.Columns, sqlgraph.NewFieldSpec(request.FieldID, field.TypeUUID))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Request.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, request.FieldID)
		for _, f := range fields {
			if !request.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != request.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.CancelledTime(); ok {
		_spec.SetField(request.FieldCancelledTime, field.TypeTime, value)
	}
	if ruo.mutation.CancelledTimeCleared() {
		_spec.ClearField(request.FieldCancelledTime, field.TypeTime)
	}
	if value, ok := ruo.mutation.Cancelled(); ok {
		_spec.SetField(request.FieldCancelled, field.TypeBool, value)
	}
	_node = &Request{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{request.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
