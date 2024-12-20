// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nevissGo/ent/pixel"
	"nevissGo/ent/predicate"
	"nevissGo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PixelUpdate is the builder for updating Pixel entities.
type PixelUpdate struct {
	config
	hooks    []Hook
	mutation *PixelMutation
}

// Where appends a list predicates to the PixelUpdate builder.
func (pu *PixelUpdate) Where(ps ...predicate.Pixel) *PixelUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetColor sets the "color" field.
func (pu *PixelUpdate) SetColor(s string) *PixelUpdate {
	pu.mutation.SetColor(s)
	return pu
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (pu *PixelUpdate) SetNillableColor(s *string) *PixelUpdate {
	if s != nil {
		pu.SetColor(*s)
	}
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PixelUpdate) SetUpdatedAt(t time.Time) *PixelUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (pu *PixelUpdate) SetUserID(id int64) *PixelUpdate {
	pu.mutation.SetUserID(id)
	return pu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (pu *PixelUpdate) SetNillableUserID(id *int64) *PixelUpdate {
	if id != nil {
		pu = pu.SetUserID(*id)
	}
	return pu
}

// SetUser sets the "user" edge to the User entity.
func (pu *PixelUpdate) SetUser(u *User) *PixelUpdate {
	return pu.SetUserID(u.ID)
}

// Mutation returns the PixelMutation object of the builder.
func (pu *PixelUpdate) Mutation() *PixelMutation {
	return pu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (pu *PixelUpdate) ClearUser() *PixelUpdate {
	pu.mutation.ClearUser()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PixelUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PixelUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PixelUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PixelUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PixelUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		v := pixel.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

func (pu *PixelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pixel.Table, pixel.Columns, sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Color(); ok {
		_spec.SetField(pixel.FieldColor, field.TypeString, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(pixel.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pixel.UserTable,
			Columns: []string{pixel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pixel.UserTable,
			Columns: []string{pixel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pixel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PixelUpdateOne is the builder for updating a single Pixel entity.
type PixelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PixelMutation
}

// SetColor sets the "color" field.
func (puo *PixelUpdateOne) SetColor(s string) *PixelUpdateOne {
	puo.mutation.SetColor(s)
	return puo
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (puo *PixelUpdateOne) SetNillableColor(s *string) *PixelUpdateOne {
	if s != nil {
		puo.SetColor(*s)
	}
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PixelUpdateOne) SetUpdatedAt(t time.Time) *PixelUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (puo *PixelUpdateOne) SetUserID(id int64) *PixelUpdateOne {
	puo.mutation.SetUserID(id)
	return puo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (puo *PixelUpdateOne) SetNillableUserID(id *int64) *PixelUpdateOne {
	if id != nil {
		puo = puo.SetUserID(*id)
	}
	return puo
}

// SetUser sets the "user" edge to the User entity.
func (puo *PixelUpdateOne) SetUser(u *User) *PixelUpdateOne {
	return puo.SetUserID(u.ID)
}

// Mutation returns the PixelMutation object of the builder.
func (puo *PixelUpdateOne) Mutation() *PixelMutation {
	return puo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (puo *PixelUpdateOne) ClearUser() *PixelUpdateOne {
	puo.mutation.ClearUser()
	return puo
}

// Where appends a list predicates to the PixelUpdate builder.
func (puo *PixelUpdateOne) Where(ps ...predicate.Pixel) *PixelUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PixelUpdateOne) Select(field string, fields ...string) *PixelUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pixel entity.
func (puo *PixelUpdateOne) Save(ctx context.Context) (*Pixel, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PixelUpdateOne) SaveX(ctx context.Context) *Pixel {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PixelUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PixelUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PixelUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		v := pixel.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

func (puo *PixelUpdateOne) sqlSave(ctx context.Context) (_node *Pixel, err error) {
	_spec := sqlgraph.NewUpdateSpec(pixel.Table, pixel.Columns, sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pixel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pixel.FieldID)
		for _, f := range fields {
			if !pixel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pixel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Color(); ok {
		_spec.SetField(pixel.FieldColor, field.TypeString, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(pixel.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pixel.UserTable,
			Columns: []string{pixel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pixel.UserTable,
			Columns: []string{pixel.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Pixel{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pixel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
