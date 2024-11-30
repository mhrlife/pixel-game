// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"nevissGo/ent/pixel"
	"nevissGo/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PixelCreate is the builder for creating a Pixel entity.
type PixelCreate struct {
	config
	mutation *PixelMutation
	hooks    []Hook
}

// SetColor sets the "color" field.
func (pc *PixelCreate) SetColor(s string) *PixelCreate {
	pc.mutation.SetColor(s)
	return pc
}

// SetUpdatedAt sets the "updated_at" field.
func (pc *PixelCreate) SetUpdatedAt(t time.Time) *PixelCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (pc *PixelCreate) SetNillableUpdatedAt(t *time.Time) *PixelCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PixelCreate) SetID(i int) *PixelCreate {
	pc.mutation.SetID(i)
	return pc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (pc *PixelCreate) AddUserIDs(ids ...int64) *PixelCreate {
	pc.mutation.AddUserIDs(ids...)
	return pc
}

// AddUser adds the "user" edges to the User entity.
func (pc *PixelCreate) AddUser(u ...*User) *PixelCreate {
	ids := make([]int64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddUserIDs(ids...)
}

// Mutation returns the PixelMutation object of the builder.
func (pc *PixelCreate) Mutation() *PixelMutation {
	return pc.mutation
}

// Save creates the Pixel in the database.
func (pc *PixelCreate) Save(ctx context.Context) (*Pixel, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PixelCreate) SaveX(ctx context.Context) *Pixel {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PixelCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PixelCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PixelCreate) defaults() {
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		v := pixel.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PixelCreate) check() error {
	if _, ok := pc.mutation.Color(); !ok {
		return &ValidationError{Name: "color", err: errors.New(`ent: missing required field "Pixel.color"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Pixel.updated_at"`)}
	}
	return nil
}

func (pc *PixelCreate) sqlSave(ctx context.Context) (*Pixel, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PixelCreate) createSpec() (*Pixel, *sqlgraph.CreateSpec) {
	var (
		_node = &Pixel{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(pixel.Table, sqlgraph.NewFieldSpec(pixel.FieldID, field.TypeInt))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Color(); ok {
		_spec.SetField(pixel.FieldColor, field.TypeString, value)
		_node.Color = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(pixel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := pc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   pixel.UserTable,
			Columns: pixel.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PixelCreateBulk is the builder for creating many Pixel entities in bulk.
type PixelCreateBulk struct {
	config
	err      error
	builders []*PixelCreate
}

// Save creates the Pixel entities in the database.
func (pcb *PixelCreateBulk) Save(ctx context.Context) ([]*Pixel, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pixel, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PixelMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PixelCreateBulk) SaveX(ctx context.Context) []*Pixel {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PixelCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PixelCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
