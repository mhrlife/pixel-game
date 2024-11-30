// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nevissGo/ent/hype"
	"nevissGo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// HypeCreate is the builder for creating a Hype entity.
type HypeCreate struct {
	config
	mutation *HypeMutation
	hooks    []Hook
}

// SetAmountRemaining sets the "amount_remaining" field.
func (hc *HypeCreate) SetAmountRemaining(i int) *HypeCreate {
	hc.mutation.SetAmountRemaining(i)
	return hc
}

// SetMaxHype sets the "max_hype" field.
func (hc *HypeCreate) SetMaxHype(i int) *HypeCreate {
	hc.mutation.SetMaxHype(i)
	return hc
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (hc *HypeCreate) SetLastUpdatedAt(t time.Time) *HypeCreate {
	hc.mutation.SetLastUpdatedAt(t)
	return hc
}

// SetNillableLastUpdatedAt sets the "last_updated_at" field if the given value is not nil.
func (hc *HypeCreate) SetNillableLastUpdatedAt(t *time.Time) *HypeCreate {
	if t != nil {
		hc.SetLastUpdatedAt(*t)
	}
	return hc
}

// SetHypePerMinute sets the "hype_per_minute" field.
func (hc *HypeCreate) SetHypePerMinute(i int) *HypeCreate {
	hc.mutation.SetHypePerMinute(i)
	return hc
}

// SetNillableHypePerMinute sets the "hype_per_minute" field if the given value is not nil.
func (hc *HypeCreate) SetNillableHypePerMinute(i *int) *HypeCreate {
	if i != nil {
		hc.SetHypePerMinute(*i)
	}
	return hc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (hc *HypeCreate) SetUserID(id int64) *HypeCreate {
	hc.mutation.SetUserID(id)
	return hc
}

// SetUser sets the "user" edge to the User entity.
func (hc *HypeCreate) SetUser(u *User) *HypeCreate {
	return hc.SetUserID(u.ID)
}

// Mutation returns the HypeMutation object of the builder.
func (hc *HypeCreate) Mutation() *HypeMutation {
	return hc.mutation
}

// Save creates the Hype in the database.
func (hc *HypeCreate) Save(ctx context.Context) (*Hype, error) {
	hc.defaults()
	return withHooks(ctx, hc.sqlSave, hc.mutation, hc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (hc *HypeCreate) SaveX(ctx context.Context) *Hype {
	v, err := hc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hc *HypeCreate) Exec(ctx context.Context) error {
	_, err := hc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hc *HypeCreate) ExecX(ctx context.Context) {
	if err := hc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (hc *HypeCreate) defaults() {
	if _, ok := hc.mutation.LastUpdatedAt(); !ok {
		v := hype.DefaultLastUpdatedAt()
		hc.mutation.SetLastUpdatedAt(v)
	}
	if _, ok := hc.mutation.HypePerMinute(); !ok {
		v := hype.DefaultHypePerMinute
		hc.mutation.SetHypePerMinute(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (hc *HypeCreate) check() error {
	if _, ok := hc.mutation.AmountRemaining(); !ok {
		return &ValidationError{Name: "amount_remaining", err: errors.New(`ent: missing required field "Hype.amount_remaining"`)}
	}
	if _, ok := hc.mutation.MaxHype(); !ok {
		return &ValidationError{Name: "max_hype", err: errors.New(`ent: missing required field "Hype.max_hype"`)}
	}
	if _, ok := hc.mutation.LastUpdatedAt(); !ok {
		return &ValidationError{Name: "last_updated_at", err: errors.New(`ent: missing required field "Hype.last_updated_at"`)}
	}
	if _, ok := hc.mutation.HypePerMinute(); !ok {
		return &ValidationError{Name: "hype_per_minute", err: errors.New(`ent: missing required field "Hype.hype_per_minute"`)}
	}
	if len(hc.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Hype.user"`)}
	}
	return nil
}

func (hc *HypeCreate) sqlSave(ctx context.Context) (*Hype, error) {
	if err := hc.check(); err != nil {
		return nil, err
	}
	_node, _spec := hc.createSpec()
	if err := sqlgraph.CreateNode(ctx, hc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	hc.mutation.id = &_node.ID
	hc.mutation.done = true
	return _node, nil
}

func (hc *HypeCreate) createSpec() (*Hype, *sqlgraph.CreateSpec) {
	var (
		_node = &Hype{config: hc.config}
		_spec = sqlgraph.NewCreateSpec(hype.Table, sqlgraph.NewFieldSpec(hype.FieldID, field.TypeInt))
	)
	if value, ok := hc.mutation.AmountRemaining(); ok {
		_spec.SetField(hype.FieldAmountRemaining, field.TypeInt, value)
		_node.AmountRemaining = value
	}
	if value, ok := hc.mutation.MaxHype(); ok {
		_spec.SetField(hype.FieldMaxHype, field.TypeInt, value)
		_node.MaxHype = value
	}
	if value, ok := hc.mutation.LastUpdatedAt(); ok {
		_spec.SetField(hype.FieldLastUpdatedAt, field.TypeTime, value)
		_node.LastUpdatedAt = value
	}
	if value, ok := hc.mutation.HypePerMinute(); ok {
		_spec.SetField(hype.FieldHypePerMinute, field.TypeInt, value)
		_node.HypePerMinute = value
	}
	if nodes := hc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   hype.UserTable,
			Columns: []string{hype.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_hype = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// HypeCreateBulk is the builder for creating many Hype entities in bulk.
type HypeCreateBulk struct {
	config
	err      error
	builders []*HypeCreate
}

// Save creates the Hype entities in the database.
func (hcb *HypeCreateBulk) Save(ctx context.Context) ([]*Hype, error) {
	if hcb.err != nil {
		return nil, hcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(hcb.builders))
	nodes := make([]*Hype, len(hcb.builders))
	mutators := make([]Mutator, len(hcb.builders))
	for i := range hcb.builders {
		func(i int, root context.Context) {
			builder := hcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*HypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, hcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, hcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, hcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (hcb *HypeCreateBulk) SaveX(ctx context.Context) []*Hype {
	v, err := hcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (hcb *HypeCreateBulk) Exec(ctx context.Context) error {
	_, err := hcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hcb *HypeCreateBulk) ExecX(ctx context.Context) {
	if err := hcb.Exec(ctx); err != nil {
		panic(err)
	}
}