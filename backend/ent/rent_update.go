// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/nurse"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/predicate"
	"github.com/sut63/team08/ent/rent"
	"github.com/sut63/team08/ent/room"
)

// RentUpdate is the builder for updating Rent entities.
type RentUpdate struct {
	config
	hooks      []Hook
	mutation   *RentMutation
	predicates []predicate.Rent
}

// Where adds a new predicate for the builder.
func (ru *RentUpdate) Where(ps ...predicate.Rent) *RentUpdate {
	ru.predicates = append(ru.predicates, ps...)
	return ru
}

// SetAddedTime sets the added_time field.
func (ru *RentUpdate) SetAddedTime(t time.Time) *RentUpdate {
	ru.mutation.SetAddedTime(t)
	return ru
}

// SetRoomID sets the room edge to Room by id.
func (ru *RentUpdate) SetRoomID(id int) *RentUpdate {
	ru.mutation.SetRoomID(id)
	return ru
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (ru *RentUpdate) SetNillableRoomID(id *int) *RentUpdate {
	if id != nil {
		ru = ru.SetRoomID(*id)
	}
	return ru
}

// SetRoom sets the room edge to Room.
func (ru *RentUpdate) SetRoom(r *Room) *RentUpdate {
	return ru.SetRoomID(r.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (ru *RentUpdate) SetPatientID(id int) *RentUpdate {
	ru.mutation.SetPatientID(id)
	return ru
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (ru *RentUpdate) SetNillablePatientID(id *int) *RentUpdate {
	if id != nil {
		ru = ru.SetPatientID(*id)
	}
	return ru
}

// SetPatient sets the patient edge to Patient.
func (ru *RentUpdate) SetPatient(p *Patient) *RentUpdate {
	return ru.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (ru *RentUpdate) SetNurseID(id int) *RentUpdate {
	ru.mutation.SetNurseID(id)
	return ru
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (ru *RentUpdate) SetNillableNurseID(id *int) *RentUpdate {
	if id != nil {
		ru = ru.SetNurseID(*id)
	}
	return ru
}

// SetNurse sets the nurse edge to Nurse.
func (ru *RentUpdate) SetNurse(n *Nurse) *RentUpdate {
	return ru.SetNurseID(n.ID)
}

// Mutation returns the RentMutation object of the builder.
func (ru *RentUpdate) Mutation() *RentMutation {
	return ru.mutation
}

// ClearRoom clears the room edge to Room.
func (ru *RentUpdate) ClearRoom() *RentUpdate {
	ru.mutation.ClearRoom()
	return ru
}

// ClearPatient clears the patient edge to Patient.
func (ru *RentUpdate) ClearPatient() *RentUpdate {
	ru.mutation.ClearPatient()
	return ru
}

// ClearNurse clears the nurse edge to Nurse.
func (ru *RentUpdate) ClearNurse() *RentUpdate {
	ru.mutation.ClearNurse()
	return ru
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ru *RentUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RentUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RentUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RentUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *RentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rent.Table,
			Columns: rent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rent.FieldID,
			},
		},
	}
	if ps := ru.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.AddedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rent.FieldAddedTime,
		})
	}
	if ru.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.RoomTable,
			Columns: []string{rent.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.RoomTable,
			Columns: []string{rent.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.PatientTable,
			Columns: []string{rent.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.PatientTable,
			Columns: []string{rent.PatientColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rent.NurseTable,
			Columns: []string{rent.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rent.NurseTable,
			Columns: []string{rent.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// RentUpdateOne is the builder for updating a single Rent entity.
type RentUpdateOne struct {
	config
	hooks    []Hook
	mutation *RentMutation
}

// SetAddedTime sets the added_time field.
func (ruo *RentUpdateOne) SetAddedTime(t time.Time) *RentUpdateOne {
	ruo.mutation.SetAddedTime(t)
	return ruo
}

// SetRoomID sets the room edge to Room by id.
func (ruo *RentUpdateOne) SetRoomID(id int) *RentUpdateOne {
	ruo.mutation.SetRoomID(id)
	return ruo
}

// SetNillableRoomID sets the room edge to Room by id if the given value is not nil.
func (ruo *RentUpdateOne) SetNillableRoomID(id *int) *RentUpdateOne {
	if id != nil {
		ruo = ruo.SetRoomID(*id)
	}
	return ruo
}

// SetRoom sets the room edge to Room.
func (ruo *RentUpdateOne) SetRoom(r *Room) *RentUpdateOne {
	return ruo.SetRoomID(r.ID)
}

// SetPatientID sets the patient edge to Patient by id.
func (ruo *RentUpdateOne) SetPatientID(id int) *RentUpdateOne {
	ruo.mutation.SetPatientID(id)
	return ruo
}

// SetNillablePatientID sets the patient edge to Patient by id if the given value is not nil.
func (ruo *RentUpdateOne) SetNillablePatientID(id *int) *RentUpdateOne {
	if id != nil {
		ruo = ruo.SetPatientID(*id)
	}
	return ruo
}

// SetPatient sets the patient edge to Patient.
func (ruo *RentUpdateOne) SetPatient(p *Patient) *RentUpdateOne {
	return ruo.SetPatientID(p.ID)
}

// SetNurseID sets the nurse edge to Nurse by id.
func (ruo *RentUpdateOne) SetNurseID(id int) *RentUpdateOne {
	ruo.mutation.SetNurseID(id)
	return ruo
}

// SetNillableNurseID sets the nurse edge to Nurse by id if the given value is not nil.
func (ruo *RentUpdateOne) SetNillableNurseID(id *int) *RentUpdateOne {
	if id != nil {
		ruo = ruo.SetNurseID(*id)
	}
	return ruo
}

// SetNurse sets the nurse edge to Nurse.
func (ruo *RentUpdateOne) SetNurse(n *Nurse) *RentUpdateOne {
	return ruo.SetNurseID(n.ID)
}

// Mutation returns the RentMutation object of the builder.
func (ruo *RentUpdateOne) Mutation() *RentMutation {
	return ruo.mutation
}

// ClearRoom clears the room edge to Room.
func (ruo *RentUpdateOne) ClearRoom() *RentUpdateOne {
	ruo.mutation.ClearRoom()
	return ruo
}

// ClearPatient clears the patient edge to Patient.
func (ruo *RentUpdateOne) ClearPatient() *RentUpdateOne {
	ruo.mutation.ClearPatient()
	return ruo
}

// ClearNurse clears the nurse edge to Nurse.
func (ruo *RentUpdateOne) ClearNurse() *RentUpdateOne {
	ruo.mutation.ClearNurse()
	return ruo
}

// Save executes the query and returns the updated entity.
func (ruo *RentUpdateOne) Save(ctx context.Context) (*Rent, error) {

	var (
		err  error
		node *Rent
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RentUpdateOne) SaveX(ctx context.Context) *Rent {
	r, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return r
}

// Exec executes the query on the entity.
func (ruo *RentUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RentUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *RentUpdateOne) sqlSave(ctx context.Context) (r *Rent, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rent.Table,
			Columns: rent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rent.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Rent.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ruo.mutation.AddedTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rent.FieldAddedTime,
		})
	}
	if ruo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.RoomTable,
			Columns: []string{rent.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.RoomTable,
			Columns: []string{rent.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: room.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.PatientTable,
			Columns: []string{rent.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patient.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   rent.PatientTable,
			Columns: []string{rent.PatientColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.NurseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rent.NurseTable,
			Columns: []string{rent.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.NurseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rent.NurseTable,
			Columns: []string{rent.NurseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: nurse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	r = &Rent{config: ruo.config}
	_spec.Assign = r.assignValues
	_spec.ScanValues = r.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rent.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return r, nil
}