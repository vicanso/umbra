// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/vicanso/cybertect/ent/predicate"
	"github.com/vicanso/cybertect/ent/tcpdetector"
)

// TCPDetectorQuery is the builder for querying TCPDetector entities.
type TCPDetectorQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.TCPDetector
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TCPDetectorQuery builder.
func (tdq *TCPDetectorQuery) Where(ps ...predicate.TCPDetector) *TCPDetectorQuery {
	tdq.predicates = append(tdq.predicates, ps...)
	return tdq
}

// Limit adds a limit step to the query.
func (tdq *TCPDetectorQuery) Limit(limit int) *TCPDetectorQuery {
	tdq.limit = &limit
	return tdq
}

// Offset adds an offset step to the query.
func (tdq *TCPDetectorQuery) Offset(offset int) *TCPDetectorQuery {
	tdq.offset = &offset
	return tdq
}

// Order adds an order step to the query.
func (tdq *TCPDetectorQuery) Order(o ...OrderFunc) *TCPDetectorQuery {
	tdq.order = append(tdq.order, o...)
	return tdq
}

// First returns the first TCPDetector entity from the query.
// Returns a *NotFoundError when no TCPDetector was found.
func (tdq *TCPDetectorQuery) First(ctx context.Context) (*TCPDetector, error) {
	nodes, err := tdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tcpdetector.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tdq *TCPDetectorQuery) FirstX(ctx context.Context) *TCPDetector {
	node, err := tdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TCPDetector ID from the query.
// Returns a *NotFoundError when no TCPDetector ID was found.
func (tdq *TCPDetectorQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tcpdetector.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tdq *TCPDetectorQuery) FirstIDX(ctx context.Context) int {
	id, err := tdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TCPDetector entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one TCPDetector entity is not found.
// Returns a *NotFoundError when no TCPDetector entities are found.
func (tdq *TCPDetectorQuery) Only(ctx context.Context) (*TCPDetector, error) {
	nodes, err := tdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tcpdetector.Label}
	default:
		return nil, &NotSingularError{tcpdetector.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tdq *TCPDetectorQuery) OnlyX(ctx context.Context) *TCPDetector {
	node, err := tdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TCPDetector ID in the query.
// Returns a *NotSingularError when exactly one TCPDetector ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tdq *TCPDetectorQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = &NotSingularError{tcpdetector.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tdq *TCPDetectorQuery) OnlyIDX(ctx context.Context) int {
	id, err := tdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TCPDetectors.
func (tdq *TCPDetectorQuery) All(ctx context.Context) ([]*TCPDetector, error) {
	if err := tdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tdq *TCPDetectorQuery) AllX(ctx context.Context) []*TCPDetector {
	nodes, err := tdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TCPDetector IDs.
func (tdq *TCPDetectorQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tdq.Select(tcpdetector.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tdq *TCPDetectorQuery) IDsX(ctx context.Context) []int {
	ids, err := tdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tdq *TCPDetectorQuery) Count(ctx context.Context) (int, error) {
	if err := tdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tdq *TCPDetectorQuery) CountX(ctx context.Context) int {
	count, err := tdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tdq *TCPDetectorQuery) Exist(ctx context.Context) (bool, error) {
	if err := tdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tdq *TCPDetectorQuery) ExistX(ctx context.Context) bool {
	exist, err := tdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TCPDetectorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tdq *TCPDetectorQuery) Clone() *TCPDetectorQuery {
	if tdq == nil {
		return nil
	}
	return &TCPDetectorQuery{
		config:     tdq.config,
		limit:      tdq.limit,
		offset:     tdq.offset,
		order:      append([]OrderFunc{}, tdq.order...),
		predicates: append([]predicate.TCPDetector{}, tdq.predicates...),
		// clone intermediate query.
		sql:  tdq.sql.Clone(),
		path: tdq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TCPDetector.Query().
//		GroupBy(tcpdetector.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tdq *TCPDetectorQuery) GroupBy(field string, fields ...string) *TCPDetectorGroupBy {
	group := &TCPDetectorGroupBy{config: tdq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tdq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty" sql:"created_at"`
//	}
//
//	client.TCPDetector.Query().
//		Select(tcpdetector.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (tdq *TCPDetectorQuery) Select(field string, fields ...string) *TCPDetectorSelect {
	tdq.fields = append([]string{field}, fields...)
	return &TCPDetectorSelect{TCPDetectorQuery: tdq}
}

func (tdq *TCPDetectorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tdq.fields {
		if !tcpdetector.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tdq.path != nil {
		prev, err := tdq.path(ctx)
		if err != nil {
			return err
		}
		tdq.sql = prev
	}
	return nil
}

func (tdq *TCPDetectorQuery) sqlAll(ctx context.Context) ([]*TCPDetector, error) {
	var (
		nodes = []*TCPDetector{}
		_spec = tdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &TCPDetector{config: tdq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tdq *TCPDetectorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tdq.querySpec()
	return sqlgraph.CountNodes(ctx, tdq.driver, _spec)
}

func (tdq *TCPDetectorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tdq *TCPDetectorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tcpdetector.Table,
			Columns: tcpdetector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tcpdetector.FieldID,
			},
		},
		From:   tdq.sql,
		Unique: true,
	}
	if fields := tdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tcpdetector.FieldID)
		for i := range fields {
			if fields[i] != tcpdetector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, tcpdetector.ValidColumn)
			}
		}
	}
	return _spec
}

func (tdq *TCPDetectorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tdq.driver.Dialect())
	t1 := builder.Table(tcpdetector.Table)
	selector := builder.Select(t1.Columns(tcpdetector.Columns...)...).From(t1)
	if tdq.sql != nil {
		selector = tdq.sql
		selector.Select(selector.Columns(tcpdetector.Columns...)...)
	}
	for _, p := range tdq.predicates {
		p(selector)
	}
	for _, p := range tdq.order {
		p(selector, tcpdetector.ValidColumn)
	}
	if offset := tdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TCPDetectorGroupBy is the group-by builder for TCPDetector entities.
type TCPDetectorGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tdgb *TCPDetectorGroupBy) Aggregate(fns ...AggregateFunc) *TCPDetectorGroupBy {
	tdgb.fns = append(tdgb.fns, fns...)
	return tdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tdgb *TCPDetectorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tdgb.path(ctx)
	if err != nil {
		return err
	}
	tdgb.sql = query
	return tdgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tdgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tdgb.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) StringsX(ctx context.Context) []string {
	v, err := tdgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tdgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) StringX(ctx context.Context) string {
	v, err := tdgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tdgb.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) IntsX(ctx context.Context) []int {
	v, err := tdgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tdgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) IntX(ctx context.Context) int {
	v, err := tdgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tdgb.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tdgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tdgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tdgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tdgb.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tdgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tdgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tdgb *TCPDetectorGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tdgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tdgb *TCPDetectorGroupBy) BoolX(ctx context.Context) bool {
	v, err := tdgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tdgb *TCPDetectorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tdgb.fields {
		if !tcpdetector.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tdgb *TCPDetectorGroupBy) sqlQuery() *sql.Selector {
	selector := tdgb.sql
	columns := make([]string, 0, len(tdgb.fields)+len(tdgb.fns))
	columns = append(columns, tdgb.fields...)
	for _, fn := range tdgb.fns {
		columns = append(columns, fn(selector, tcpdetector.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(tdgb.fields...)
}

// TCPDetectorSelect is the builder for selecting fields of TCPDetector entities.
type TCPDetectorSelect struct {
	*TCPDetectorQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tds *TCPDetectorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tds.prepareQuery(ctx); err != nil {
		return err
	}
	tds.sql = tds.TCPDetectorQuery.sqlQuery(ctx)
	return tds.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tds *TCPDetectorSelect) ScanX(ctx context.Context, v interface{}) {
	if err := tds.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Strings(ctx context.Context) ([]string, error) {
	if len(tds.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := tds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tds *TCPDetectorSelect) StringsX(ctx context.Context) []string {
	v, err := tds.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tds.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tds *TCPDetectorSelect) StringX(ctx context.Context) string {
	v, err := tds.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Ints(ctx context.Context) ([]int, error) {
	if len(tds.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := tds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tds *TCPDetectorSelect) IntsX(ctx context.Context) []int {
	v, err := tds.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tds.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tds *TCPDetectorSelect) IntX(ctx context.Context) int {
	v, err := tds.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(tds.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := tds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tds *TCPDetectorSelect) Float64sX(ctx context.Context) []float64 {
	v, err := tds.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tds.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tds *TCPDetectorSelect) Float64X(ctx context.Context) float64 {
	v, err := tds.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(tds.fields) > 1 {
		return nil, errors.New("ent: TCPDetectorSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := tds.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tds *TCPDetectorSelect) BoolsX(ctx context.Context) []bool {
	v, err := tds.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (tds *TCPDetectorSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tds.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{tcpdetector.Label}
	default:
		err = fmt.Errorf("ent: TCPDetectorSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tds *TCPDetectorSelect) BoolX(ctx context.Context) bool {
	v, err := tds.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tds *TCPDetectorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tds.sqlQuery().Query()
	if err := tds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tds *TCPDetectorSelect) sqlQuery() sql.Querier {
	selector := tds.sql
	selector.Select(selector.Columns(tds.fields...)...)
	return selector
}
