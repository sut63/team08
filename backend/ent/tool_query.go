// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team08/ent/operativerecord"
	"github.com/sut63/team08/ent/predicate"
	"github.com/sut63/team08/ent/tool"
)

// ToolQuery is the builder for querying Tool entities.
type ToolQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.Tool
	// eager-loading edges.
	withToolOperativerecord *OperativerecordQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (tq *ToolQuery) Where(ps ...predicate.Tool) *ToolQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *ToolQuery) Limit(limit int) *ToolQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *ToolQuery) Offset(offset int) *ToolQuery {
	tq.offset = &offset
	return tq
}

// Order adds an order step to the query.
func (tq *ToolQuery) Order(o ...OrderFunc) *ToolQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryToolOperativerecord chains the current query on the Tool_Operativerecord edge.
func (tq *ToolQuery) QueryToolOperativerecord() *OperativerecordQuery {
	query := &OperativerecordQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tool.Table, tool.FieldID, tq.sqlQuery()),
			sqlgraph.To(operativerecord.Table, operativerecord.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tool.ToolOperativerecordTable, tool.ToolOperativerecordColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tool entity in the query. Returns *NotFoundError when no tool was found.
func (tq *ToolQuery) First(ctx context.Context) (*Tool, error) {
	ts, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(ts) == 0 {
		return nil, &NotFoundError{tool.Label}
	}
	return ts[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *ToolQuery) FirstX(ctx context.Context) *Tool {
	t, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return t
}

// FirstID returns the first Tool id in the query. Returns *NotFoundError when no id was found.
func (tq *ToolQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tool.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (tq *ToolQuery) FirstXID(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Tool entity in the query, returns an error if not exactly one entity was returned.
func (tq *ToolQuery) Only(ctx context.Context) (*Tool, error) {
	ts, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(ts) {
	case 1:
		return ts[0], nil
	case 0:
		return nil, &NotFoundError{tool.Label}
	default:
		return nil, &NotSingularError{tool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *ToolQuery) OnlyX(ctx context.Context) *Tool {
	t, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// OnlyID returns the only Tool id in the query, returns an error if not exactly one id was returned.
func (tq *ToolQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = &NotSingularError{tool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *ToolQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tools.
func (tq *ToolQuery) All(ctx context.Context) ([]*Tool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *ToolQuery) AllX(ctx context.Context) []*Tool {
	ts, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return ts
}

// IDs executes the query and returns a list of Tool ids.
func (tq *ToolQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(tool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *ToolQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *ToolQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *ToolQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *ToolQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *ToolQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *ToolQuery) Clone() *ToolQuery {
	return &ToolQuery{
		config:     tq.config,
		limit:      tq.limit,
		offset:     tq.offset,
		order:      append([]OrderFunc{}, tq.order...),
		unique:     append([]string{}, tq.unique...),
		predicates: append([]predicate.Tool{}, tq.predicates...),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

//  WithToolOperativerecord tells the query-builder to eager-loads the nodes that are connected to
// the "Tool_Operativerecord" edge. The optional arguments used to configure the query builder of the edge.
func (tq *ToolQuery) WithToolOperativerecord(opts ...func(*OperativerecordQuery)) *ToolQuery {
	query := &OperativerecordQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withToolOperativerecord = query
	return tq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ToolName string `json:"Tool_Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Tool.Query().
//		GroupBy(tool.FieldToolName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *ToolQuery) GroupBy(field string, fields ...string) *ToolGroupBy {
	group := &ToolGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		ToolName string `json:"Tool_Name,omitempty"`
//	}
//
//	client.Tool.Query().
//		Select(tool.FieldToolName).
//		Scan(ctx, &v)
//
func (tq *ToolQuery) Select(field string, fields ...string) *ToolSelect {
	selector := &ToolSelect{config: tq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(), nil
	}
	return selector
}

func (tq *ToolQuery) prepareQuery(ctx context.Context) error {
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *ToolQuery) sqlAll(ctx context.Context) ([]*Tool, error) {
	var (
		nodes       = []*Tool{}
		_spec       = tq.querySpec()
		loadedTypes = [1]bool{
			tq.withToolOperativerecord != nil,
		}
	)
	_spec.ScanValues = func() []interface{} {
		node := &Tool{config: tq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
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
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withToolOperativerecord; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Tool)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Operativerecord(func(s *sql.Selector) {
			s.Where(sql.InValues(tool.ToolOperativerecordColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.Tool_id
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "Tool_id" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "Tool_id" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.ToolOperativerecord = append(node.Edges.ToolOperativerecord, n)
		}
	}

	return nodes, nil
}

func (tq *ToolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *ToolQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (tq *ToolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tool.Table,
			Columns: tool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tool.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *ToolQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tool.Table)
	selector := builder.Select(t1.Columns(tool.Columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(tool.Columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ToolGroupBy is the builder for group-by Tool entities.
type ToolGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *ToolGroupBy) Aggregate(fns ...AggregateFunc) *ToolGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scan the result into the given value.
func (tgb *ToolGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *ToolGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: ToolGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *ToolGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tgb *ToolGroupBy) StringX(ctx context.Context) string {
	v, err := tgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: ToolGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *ToolGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tgb *ToolGroupBy) IntX(ctx context.Context) int {
	v, err := tgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: ToolGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *ToolGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tgb *ToolGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: ToolGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *ToolGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (tgb *ToolGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tgb *ToolGroupBy) BoolX(ctx context.Context) bool {
	v, err := tgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *ToolGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tgb.sqlQuery().Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *ToolGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql
	columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
	columns = append(columns, tgb.fields...)
	for _, fn := range tgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(tgb.fields...)
}

// ToolSelect is the builder for select fields of Tool entities.
type ToolSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (ts *ToolSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := ts.path(ctx)
	if err != nil {
		return err
	}
	ts.sql = query
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *ToolSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: ToolSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *ToolSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ts *ToolSelect) StringX(ctx context.Context) string {
	v, err := ts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: ToolSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *ToolSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ts *ToolSelect) IntX(ctx context.Context) int {
	v, err := ts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: ToolSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *ToolSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ts *ToolSelect) Float64X(ctx context.Context) float64 {
	v, err := ts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: ToolSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *ToolSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (ts *ToolSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tool.Label}
	default:
		err = fmt.Errorf("ent: ToolSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ts *ToolSelect) BoolX(ctx context.Context) bool {
	v, err := ts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *ToolSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sqlQuery().Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ts *ToolSelect) sqlQuery() sql.Querier {
	selector := ts.sql
	selector.Select(selector.Columns(ts.fields...)...)
	return selector
}
