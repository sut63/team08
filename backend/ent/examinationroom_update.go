// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/examinationroom"
	"github.com/sut63/team08/ent/operativerecord"
	"github.com/sut63/team08/ent/predicate"
)

// ExaminationroomUpdate is the builder for updating Examinationroom entities.
type ExaminationroomUpdate struct {
	config
	hooks      []Hook
	mutation   *ExaminationroomMutation
	predicates []predicate.Examinationroom
}

// Where adds a new predicate for the builder.
func (eu *ExaminationroomUpdate) Where(ps ...predicate.Examinationroom) *ExaminationroomUpdate {
	eu.predicates = append(eu.predicates, ps...)
	return eu
}

// SetExaminationroomName sets the examinationroom_name field.
func (eu *ExaminationroomUpdate) SetExaminationroomName(s string) *ExaminationroomUpdate {
	eu.mutation.SetExaminationroomName(s)
	return eu
}

// AddExaminationroomOperativerecordIDs adds the Examinationroom_Operativerecord edge to Operativerecord by ids.
func (eu *ExaminationroomUpdate) AddExaminationroomOperativerecordIDs(ids ...int) *ExaminationroomUpdate {
	eu.mutation.AddExaminationroomOperativerecordIDs(ids...)
	return eu
}

// AddExaminationroomOperativerecord adds the Examinationroom_Operativerecord edges to Operativerecord.
func (eu *ExaminationroomUpdate) AddExaminationroomOperativerecord(o ...*Operativerecord) *ExaminationroomUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eu.AddExaminationroomOperativerecordIDs(ids...)
}

// Mutation returns the ExaminationroomMutation object of the builder.
func (eu *ExaminationroomUpdate) Mutation() *ExaminationroomMutation {
	return eu.mutation
}

// RemoveExaminationroomOperativerecordIDs removes the Examinationroom_Operativerecord edge to Operativerecord by ids.
func (eu *ExaminationroomUpdate) RemoveExaminationroomOperativerecordIDs(ids ...int) *ExaminationroomUpdate {
	eu.mutation.RemoveExaminationroomOperativerecordIDs(ids...)
	return eu
}

// RemoveExaminationroomOperativerecord removes Examinationroom_Operativerecord edges to Operativerecord.
func (eu *ExaminationroomUpdate) RemoveExaminationroomOperativerecord(o ...*Operativerecord) *ExaminationroomUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return eu.RemoveExaminationroomOperativerecordIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (eu *ExaminationroomUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := eu.mutation.ExaminationroomName(); ok {
		if err := examinationroom.ExaminationroomNameValidator(v); err != nil {
			return 0, &ValidationError{Name: "examinationroom_name", err: fmt.Errorf("ent: validator failed for field \"examinationroom_name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExaminationroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *ExaminationroomUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *ExaminationroomUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *ExaminationroomUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *ExaminationroomUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   examinationroom.Table,
			Columns: examinationroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: examinationroom.FieldID,
			},
		},
	}
	if ps := eu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.ExaminationroomName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: examinationroom.FieldExaminationroomName,
		})
	}
	if nodes := eu.mutation.RemovedExaminationroomOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examinationroom.ExaminationroomOperativerecordTable,
			Columns: []string{examinationroom.ExaminationroomOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ExaminationroomOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examinationroom.ExaminationroomOperativerecordTable,
			Columns: []string{examinationroom.ExaminationroomOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examinationroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ExaminationroomUpdateOne is the builder for updating a single Examinationroom entity.
type ExaminationroomUpdateOne struct {
	config
	hooks    []Hook
	mutation *ExaminationroomMutation
}

// SetExaminationroomName sets the examinationroom_name field.
func (euo *ExaminationroomUpdateOne) SetExaminationroomName(s string) *ExaminationroomUpdateOne {
	euo.mutation.SetExaminationroomName(s)
	return euo
}

// AddExaminationroomOperativerecordIDs adds the Examinationroom_Operativerecord edge to Operativerecord by ids.
func (euo *ExaminationroomUpdateOne) AddExaminationroomOperativerecordIDs(ids ...int) *ExaminationroomUpdateOne {
	euo.mutation.AddExaminationroomOperativerecordIDs(ids...)
	return euo
}

// AddExaminationroomOperativerecord adds the Examinationroom_Operativerecord edges to Operativerecord.
func (euo *ExaminationroomUpdateOne) AddExaminationroomOperativerecord(o ...*Operativerecord) *ExaminationroomUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return euo.AddExaminationroomOperativerecordIDs(ids...)
}

// Mutation returns the ExaminationroomMutation object of the builder.
func (euo *ExaminationroomUpdateOne) Mutation() *ExaminationroomMutation {
	return euo.mutation
}

// RemoveExaminationroomOperativerecordIDs removes the Examinationroom_Operativerecord edge to Operativerecord by ids.
func (euo *ExaminationroomUpdateOne) RemoveExaminationroomOperativerecordIDs(ids ...int) *ExaminationroomUpdateOne {
	euo.mutation.RemoveExaminationroomOperativerecordIDs(ids...)
	return euo
}

// RemoveExaminationroomOperativerecord removes Examinationroom_Operativerecord edges to Operativerecord.
func (euo *ExaminationroomUpdateOne) RemoveExaminationroomOperativerecord(o ...*Operativerecord) *ExaminationroomUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return euo.RemoveExaminationroomOperativerecordIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (euo *ExaminationroomUpdateOne) Save(ctx context.Context) (*Examinationroom, error) {
	if v, ok := euo.mutation.ExaminationroomName(); ok {
		if err := examinationroom.ExaminationroomNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "examinationroom_name", err: fmt.Errorf("ent: validator failed for field \"examinationroom_name\": %w", err)}
		}
	}

	var (
		err  error
		node *Examinationroom
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExaminationroomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *ExaminationroomUpdateOne) SaveX(ctx context.Context) *Examinationroom {
	e, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return e
}

// Exec executes the query on the entity.
func (euo *ExaminationroomUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *ExaminationroomUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *ExaminationroomUpdateOne) sqlSave(ctx context.Context) (e *Examinationroom, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   examinationroom.Table,
			Columns: examinationroom.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: examinationroom.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Examinationroom.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := euo.mutation.ExaminationroomName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: examinationroom.FieldExaminationroomName,
		})
	}
	if nodes := euo.mutation.RemovedExaminationroomOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examinationroom.ExaminationroomOperativerecordTable,
			Columns: []string{examinationroom.ExaminationroomOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ExaminationroomOperativerecordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   examinationroom.ExaminationroomOperativerecordTable,
			Columns: []string{examinationroom.ExaminationroomOperativerecordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: operativerecord.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	e = &Examinationroom{config: euo.config}
	_spec.Assign = e.assignValues
	_spec.ScanValues = e.scanValues()
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{examinationroom.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return e, nil
}