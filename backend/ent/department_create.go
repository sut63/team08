// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/department"
	"github.com/sut63/team08/ent/diagnose"
)

// DepartmentCreate is the builder for creating a Department entity.
type DepartmentCreate struct {
	config
	mutation *DepartmentMutation
	hooks    []Hook
}

// SetDepartmentName sets the Department_Name field.
func (dc *DepartmentCreate) SetDepartmentName(s string) *DepartmentCreate {
	dc.mutation.SetDepartmentName(s)
	return dc
}

// AddDepartmentDiagnoseIDs adds the department_diagnose edge to Diagnose by ids.
func (dc *DepartmentCreate) AddDepartmentDiagnoseIDs(ids ...int) *DepartmentCreate {
	dc.mutation.AddDepartmentDiagnoseIDs(ids...)
	return dc
}

// AddDepartmentDiagnose adds the department_diagnose edges to Diagnose.
func (dc *DepartmentCreate) AddDepartmentDiagnose(d ...*Diagnose) *DepartmentCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return dc.AddDepartmentDiagnoseIDs(ids...)
}

// Mutation returns the DepartmentMutation object of the builder.
func (dc *DepartmentCreate) Mutation() *DepartmentMutation {
	return dc.mutation
}

// Save creates the Department in the database.
func (dc *DepartmentCreate) Save(ctx context.Context) (*Department, error) {
	if _, ok := dc.mutation.DepartmentName(); !ok {
		return nil, &ValidationError{Name: "Department_Name", err: errors.New("ent: missing required field \"Department_Name\"")}
	}
	if v, ok := dc.mutation.DepartmentName(); ok {
		if err := department.DepartmentNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Department_Name", err: fmt.Errorf("ent: validator failed for field \"Department_Name\": %w", err)}
		}
	}
	var (
		err  error
		node *Department
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DepartmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DepartmentCreate) SaveX(ctx context.Context) *Department {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DepartmentCreate) sqlSave(ctx context.Context) (*Department, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DepartmentCreate) createSpec() (*Department, *sqlgraph.CreateSpec) {
	var (
		d     = &Department{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: department.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: department.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.DepartmentName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: department.FieldDepartmentName,
		})
		d.DepartmentName = value
	}
	if nodes := dc.mutation.DepartmentDiagnoseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   department.DepartmentDiagnoseTable,
			Columns: []string{department.DepartmentDiagnoseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: diagnose.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}
