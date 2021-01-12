// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/prefix"
)

// PrefixCreate is the builder for creating a Prefix entity.
type PrefixCreate struct {
	config
	mutation *PrefixMutation
	hooks    []Hook
}

// SetPname sets the Pname field.
func (pc *PrefixCreate) SetPname(s string) *PrefixCreate {
	pc.mutation.SetPname(s)
	return pc
}

// AddFromprefixIDs adds the fromprefix edge to Patient by ids.
func (pc *PrefixCreate) AddFromprefixIDs(ids ...int) *PrefixCreate {
	pc.mutation.AddFromprefixIDs(ids...)
	return pc
}

// AddFromprefix adds the fromprefix edges to Patient.
func (pc *PrefixCreate) AddFromprefix(p ...*Patient) *PrefixCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddFromprefixIDs(ids...)
}

// Mutation returns the PrefixMutation object of the builder.
func (pc *PrefixCreate) Mutation() *PrefixMutation {
	return pc.mutation
}

// Save creates the Prefix in the database.
func (pc *PrefixCreate) Save(ctx context.Context) (*Prefix, error) {
	if _, ok := pc.mutation.Pname(); !ok {
		return nil, &ValidationError{Name: "Pname", err: errors.New("ent: missing required field \"Pname\"")}
	}
	if v, ok := pc.mutation.Pname(); ok {
		if err := prefix.PnameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Pname", err: fmt.Errorf("ent: validator failed for field \"Pname\": %w", err)}
		}
	}
	var (
		err  error
		node *Prefix
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrefixMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PrefixCreate) SaveX(ctx context.Context) *Prefix {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PrefixCreate) sqlSave(ctx context.Context) (*Prefix, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *PrefixCreate) createSpec() (*Prefix, *sqlgraph.CreateSpec) {
	var (
		pr    = &Prefix{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: prefix.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prefix.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Pname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: prefix.FieldPname,
		})
		pr.Pname = value
	}
	if nodes := pc.mutation.FromprefixIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   prefix.FromprefixTable,
			Columns: []string{prefix.FromprefixColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pr, _spec
}