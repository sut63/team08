// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/department"
	"github.com/sut63/team08/ent/diagnose"
	"github.com/sut63/team08/ent/disease"
	"github.com/sut63/team08/ent/doctor"
	"github.com/sut63/team08/ent/patient"
	"github.com/sut63/team08/ent/predicate"
)

// DiagnoseQuery is the builder for querying Diagnose entities.
type DiagnoseQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Diagnose
	// eager-loading edges.
	withDisease    *DiseaseQuery
	withDepartment *DepartmentQuery
	withPatient    *PatientQuery
	withDoctor     *DoctorQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (dq *DiagnoseQuery) Where(ps ...predicate.Diagnose) *DiagnoseQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DiagnoseQuery) Limit(limit int) *DiagnoseQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DiagnoseQuery) Offset(offset int) *DiagnoseQuery {
	dq.offset = &offset
	return dq
}

// Order adds an order step to the query.
func (dq *DiagnoseQuery) Order(o ...OrderFunc) *DiagnoseQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDisease chains the current query on the disease edge.
func (dq *DiagnoseQuery) QueryDisease() *DiseaseQuery {
	query := &DiseaseQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(diagnose.Table, diagnose.FieldID, dq.sqlQuery()),
			sqlgraph.To(disease.Table, disease.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, diagnose.DiseaseTable, diagnose.DiseaseColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDepartment chains the current query on the department edge.
func (dq *DiagnoseQuery) QueryDepartment() *DepartmentQuery {
	query := &DepartmentQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(diagnose.Table, diagnose.FieldID, dq.sqlQuery()),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, diagnose.DepartmentTable, diagnose.DepartmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryPatient chains the current query on the patient edge.
func (dq *DiagnoseQuery) QueryPatient() *PatientQuery {
	query := &PatientQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(diagnose.Table, diagnose.FieldID, dq.sqlQuery()),
			sqlgraph.To(patient.Table, patient.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, diagnose.PatientTable, diagnose.PatientColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryDoctor chains the current query on the doctor edge.
func (dq *DiagnoseQuery) QueryDoctor() *DoctorQuery {
	query := &DoctorQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(diagnose.Table, diagnose.FieldID, dq.sqlQuery()),
			sqlgraph.To(doctor.Table, doctor.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, diagnose.DoctorTable, diagnose.DoctorColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Diagnose entity in the query. Returns *NotFoundError when no diagnose was found.
func (dq *DiagnoseQuery) First(ctx context.Context) (*Diagnose, error) {
	ds, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ds) == 0 {
		return nil, &NotFoundError{diagnose.Label}
	}
	return ds[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DiagnoseQuery) FirstX(ctx context.Context) *Diagnose {
	d, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return d
}

// FirstID returns the first Diagnose id in the query. Returns *NotFoundError when no id was found.
func (dq *DiagnoseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{diagnose.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (dq *DiagnoseQuery) FirstXID(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Diagnose entity in the query, returns an error if not exactly one entity was returned.
func (dq *DiagnoseQuery) Only(ctx context.Context) (*Diagnose, error) {
	ds, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ds) {
	case 1:
		return ds[0], nil
	case 0:
		return nil, &NotFoundError{diagnose.Label}
	default:
		return nil, &NotSingularError{diagnose.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DiagnoseQuery) OnlyX(ctx context.Context) *Diagnose {
	d, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return d
}

// OnlyID returns the only Diagnose id in the query, returns an error if not exactly one id was returned.
func (dq *DiagnoseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = &NotSingularError{diagnose.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DiagnoseQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Diagnoses.
func (dq *DiagnoseQuery) All(ctx context.Context) ([]*Diagnose, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DiagnoseQuery) AllX(ctx context.Context) []*Diagnose {
	ds, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ds
}

// IDs executes the query and returns a list of Diagnose ids.
func (dq *DiagnoseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := dq.Select(diagnose.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DiagnoseQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DiagnoseQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DiagnoseQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DiagnoseQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DiagnoseQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DiagnoseQuery) Clone() *DiagnoseQuery {
	return &DiagnoseQuery{
		config:     dq.config,
		limit:      dq.limit,
		offset:     dq.offset,
		order:      append([]OrderFunc{}, dq.order...),
		unique:     append([]string{}, dq.unique...),
		predicates: append([]predicate.Diagnose{}, dq.predicates...),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

//  WithDisease tells the query-builder to eager-loads the nodes that are connected to
// the "disease" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DiagnoseQuery) WithDisease(opts ...func(*DiseaseQuery)) *DiagnoseQuery {
	query := &DiseaseQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDisease = query
	return dq
}

//  WithDepartment tells the query-builder to eager-loads the nodes that are connected to
// the "department" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DiagnoseQuery) WithDepartment(opts ...func(*DepartmentQuery)) *DiagnoseQuery {
	query := &DepartmentQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDepartment = query
	return dq
}

//  WithPatient tells the query-builder to eager-loads the nodes that are connected to
// the "patient" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DiagnoseQuery) WithPatient(opts ...func(*PatientQuery)) *DiagnoseQuery {
	query := &PatientQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withPatient = query
	return dq
}

//  WithDoctor tells the query-builder to eager-loads the nodes that are connected to
// the "doctor" edge. The optional arguments used to configure the query builder of the edge.
func (dq *DiagnoseQuery) WithDoctor(opts ...func(*DoctorQuery)) *DiagnoseQuery {
	query := &DoctorQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withDoctor = query
	return dq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (dq *DiagnoseQuery) GroupBy(field string, fields ...string) *DiagnoseGroupBy {
	group := &DiagnoseGroupBy{config: dq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
func (dq *DiagnoseQuery) Select(field string, fields ...string) *DiagnoseSelect {
	selector := &DiagnoseSelect{config: dq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(), nil
	}
	return selector
}

func (dq *DiagnoseQuery) prepareQuery(ctx context.Context) error {
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DiagnoseQuery) sqlAll(ctx context.Context) ([]*Diagnose, error) {
	var (
		nodes       = []*Diagnose{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [4]bool{
			dq.withDisease != nil,
			dq.withDepartment != nil,
			dq.withPatient != nil,
			dq.withDoctor != nil,
		}
	)
	if dq.withDisease != nil || dq.withDepartment != nil || dq.withPatient != nil || dq.withDoctor != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, diagnose.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Diagnose{config: dq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := dq.withDisease; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Diagnose)
		for i := range nodes {
			if fk := nodes[i].disease_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(disease.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "disease_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Disease = n
			}
		}
	}

	if query := dq.withDepartment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Diagnose)
		for i := range nodes {
			if fk := nodes[i].department_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(department.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "department_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Department = n
			}
		}
	}

	if query := dq.withPatient; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Diagnose)
		for i := range nodes {
			if fk := nodes[i].patient_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(patient.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "patient_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Patient = n
			}
		}
	}

	if query := dq.withDoctor; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Diagnose)
		for i := range nodes {
			if fk := nodes[i].doctor_id; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(doctor.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "doctor_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Doctor = n
			}
		}
	}

	return nodes, nil
}

func (dq *DiagnoseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DiagnoseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (dq *DiagnoseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   diagnose.Table,
			Columns: diagnose.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: diagnose.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DiagnoseQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(diagnose.Table)
	selector := builder.Select(t1.Columns(diagnose.Columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(diagnose.Columns...)...)
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DiagnoseGroupBy is the builder for group-by Diagnose entities.
type DiagnoseGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DiagnoseGroupBy) Aggregate(fns ...AggregateFunc) *DiagnoseGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scan the result into the given value.
func (dgb *DiagnoseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := dgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DiagnoseGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) StringsX(ctx context.Context) []string {
	v, err := dgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = dgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) StringX(ctx context.Context) string {
	v, err := dgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DiagnoseGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) IntsX(ctx context.Context) []int {
	v, err := dgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = dgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) IntX(ctx context.Context) int {
	v, err := dgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DiagnoseGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := dgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = dgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) Float64X(ctx context.Context) float64 {
	v, err := dgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(dgb.fields) > 1 {
		return nil, errors.New("ent: DiagnoseGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := dgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := dgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (dgb *DiagnoseGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = dgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (dgb *DiagnoseGroupBy) BoolX(ctx context.Context) bool {
	v, err := dgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dgb *DiagnoseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dgb.sqlQuery().Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DiagnoseGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql
	columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
	columns = append(columns, dgb.fields...)
	for _, fn := range dgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(dgb.fields...)
}

// DiagnoseSelect is the builder for select fields of Diagnose entities.
type DiagnoseSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ds *DiagnoseSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ds.path(ctx)
	if err != nil {
		return err
	}
	ds.sql = query
	return ds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ds *DiagnoseSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DiagnoseSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ds *DiagnoseSelect) StringsX(ctx context.Context) []string {
	v, err := ds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ds *DiagnoseSelect) StringX(ctx context.Context) string {
	v, err := ds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DiagnoseSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ds *DiagnoseSelect) IntsX(ctx context.Context) []int {
	v, err := ds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ds *DiagnoseSelect) IntX(ctx context.Context) int {
	v, err := ds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DiagnoseSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ds *DiagnoseSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ds *DiagnoseSelect) Float64X(ctx context.Context) float64 {
	v, err := ds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ds.fields) > 1 {
		return nil, errors.New("ent: DiagnoseSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ds *DiagnoseSelect) BoolsX(ctx context.Context) []bool {
	v, err := ds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ds *DiagnoseSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{diagnose.Label}
	default:
		err = fmt.Errorf("ent: DiagnoseSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ds *DiagnoseSelect) BoolX(ctx context.Context) bool {
	v, err := ds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ds *DiagnoseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sqlQuery().Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ds *DiagnoseSelect) sqlQuery() sql.Querier {
	selector := ds.sql
	selector.Select(selector.Columns(ds.fields...)...)
	return selector
}
