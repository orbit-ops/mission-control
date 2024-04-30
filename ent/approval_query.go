// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/orbit-ops/launchpad-core/ent/access"
	"github.com/orbit-ops/launchpad-core/ent/approval"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
	"github.com/orbit-ops/launchpad-core/ent/request"
)

// ApprovalQuery is the builder for querying Approval entities.
type ApprovalQuery struct {
	config
	ctx         *QueryContext
	order       []approval.OrderOption
	inters      []Interceptor
	predicates  []predicate.Approval
	withRequest *RequestQuery
	withAccess  *AccessQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ApprovalQuery builder.
func (aq *ApprovalQuery) Where(ps ...predicate.Approval) *ApprovalQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *ApprovalQuery) Limit(limit int) *ApprovalQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *ApprovalQuery) Offset(offset int) *ApprovalQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *ApprovalQuery) Unique(unique bool) *ApprovalQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *ApprovalQuery) Order(o ...approval.OrderOption) *ApprovalQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// QueryRequest chains the current query on the "request" edge.
func (aq *ApprovalQuery) QueryRequest() *RequestQuery {
	query := (&RequestClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(approval.Table, approval.FieldID, selector),
			sqlgraph.To(request.Table, request.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, approval.RequestTable, approval.RequestColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAccess chains the current query on the "access" edge.
func (aq *ApprovalQuery) QueryAccess() *AccessQuery {
	query := (&AccessClient{config: aq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := aq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := aq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(approval.Table, approval.FieldID, selector),
			sqlgraph.To(access.Table, access.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, approval.AccessTable, approval.AccessColumn),
		)
		fromU = sqlgraph.SetNeighbors(aq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Approval entity from the query.
// Returns a *NotFoundError when no Approval was found.
func (aq *ApprovalQuery) First(ctx context.Context) (*Approval, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{approval.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *ApprovalQuery) FirstX(ctx context.Context) *Approval {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Approval ID from the query.
// Returns a *NotFoundError when no Approval ID was found.
func (aq *ApprovalQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{approval.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *ApprovalQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Approval entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Approval entity is found.
// Returns a *NotFoundError when no Approval entities are found.
func (aq *ApprovalQuery) Only(ctx context.Context) (*Approval, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{approval.Label}
	default:
		return nil, &NotSingularError{approval.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *ApprovalQuery) OnlyX(ctx context.Context) *Approval {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Approval ID in the query.
// Returns a *NotSingularError when more than one Approval ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *ApprovalQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{approval.Label}
	default:
		err = &NotSingularError{approval.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *ApprovalQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Approvals.
func (aq *ApprovalQuery) All(ctx context.Context) ([]*Approval, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Approval, *ApprovalQuery]()
	return withInterceptors[[]*Approval](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *ApprovalQuery) AllX(ctx context.Context) []*Approval {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Approval IDs.
func (aq *ApprovalQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if aq.ctx.Unique == nil && aq.path != nil {
		aq.Unique(true)
	}
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err = aq.Select(approval.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *ApprovalQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *ApprovalQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*ApprovalQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *ApprovalQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *ApprovalQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *ApprovalQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ApprovalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *ApprovalQuery) Clone() *ApprovalQuery {
	if aq == nil {
		return nil
	}
	return &ApprovalQuery{
		config:      aq.config,
		ctx:         aq.ctx.Clone(),
		order:       append([]approval.OrderOption{}, aq.order...),
		inters:      append([]Interceptor{}, aq.inters...),
		predicates:  append([]predicate.Approval{}, aq.predicates...),
		withRequest: aq.withRequest.Clone(),
		withAccess:  aq.withAccess.Clone(),
		// clone intermediate query.
		sql:  aq.sql.Clone(),
		path: aq.path,
	}
}

// WithRequest tells the query-builder to eager-load the nodes that are connected to
// the "request" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ApprovalQuery) WithRequest(opts ...func(*RequestQuery)) *ApprovalQuery {
	query := (&RequestClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withRequest = query
	return aq
}

// WithAccess tells the query-builder to eager-load the nodes that are connected to
// the "access" edge. The optional arguments are used to configure the query builder of the edge.
func (aq *ApprovalQuery) WithAccess(opts ...func(*AccessQuery)) *ApprovalQuery {
	query := (&AccessClient{config: aq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	aq.withAccess = query
	return aq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Person string `json:"person,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Approval.Query().
//		GroupBy(approval.FieldPerson).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aq *ApprovalQuery) GroupBy(field string, fields ...string) *ApprovalGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ApprovalGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = approval.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Person string `json:"person,omitempty"`
//	}
//
//	client.Approval.Query().
//		Select(approval.FieldPerson).
//		Scan(ctx, &v)
func (aq *ApprovalQuery) Select(fields ...string) *ApprovalSelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &ApprovalSelect{ApprovalQuery: aq}
	sbuild.label = approval.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ApprovalSelect configured with the given aggregations.
func (aq *ApprovalQuery) Aggregate(fns ...AggregateFunc) *ApprovalSelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *ApprovalQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	for _, f := range aq.ctx.Fields {
		if !approval.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.sql = prev
	}
	return nil
}

func (aq *ApprovalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Approval, error) {
	var (
		nodes       = []*Approval{}
		withFKs     = aq.withFKs
		_spec       = aq.querySpec()
		loadedTypes = [2]bool{
			aq.withRequest != nil,
			aq.withAccess != nil,
		}
	)
	if aq.withRequest != nil || aq.withAccess != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, approval.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Approval).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Approval{config: aq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := aq.withRequest; query != nil {
		if err := aq.loadRequest(ctx, query, nodes, nil,
			func(n *Approval, e *Request) { n.Edges.Request = e }); err != nil {
			return nil, err
		}
	}
	if query := aq.withAccess; query != nil {
		if err := aq.loadAccess(ctx, query, nodes, nil,
			func(n *Approval, e *Access) { n.Edges.Access = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (aq *ApprovalQuery) loadRequest(ctx context.Context, query *RequestQuery, nodes []*Approval, init func(*Approval), assign func(*Approval, *Request)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Approval)
	for i := range nodes {
		if nodes[i].approval_request == nil {
			continue
		}
		fk := *nodes[i].approval_request
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(request.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "approval_request" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (aq *ApprovalQuery) loadAccess(ctx context.Context, query *AccessQuery, nodes []*Approval, init func(*Approval), assign func(*Approval, *Access)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Approval)
	for i := range nodes {
		if nodes[i].access_approvals == nil {
			continue
		}
		fk := *nodes[i].access_approvals
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(access.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "access_approvals" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (aq *ApprovalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aq.querySpec()
	_spec.Node.Columns = aq.ctx.Fields
	if len(aq.ctx.Fields) > 0 {
		_spec.Unique = aq.ctx.Unique != nil && *aq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, aq.driver, _spec)
}

func (aq *ApprovalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(approval.Table, approval.Columns, sqlgraph.NewFieldSpec(approval.FieldID, field.TypeUUID))
	_spec.From = aq.sql
	if unique := aq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if aq.path != nil {
		_spec.Unique = true
	}
	if fields := aq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, approval.FieldID)
		for i := range fields {
			if fields[i] != approval.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aq *ApprovalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aq.driver.Dialect())
	t1 := builder.Table(approval.Table)
	columns := aq.ctx.Fields
	if len(columns) == 0 {
		columns = approval.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aq.sql != nil {
		selector = aq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aq.ctx.Unique != nil && *aq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range aq.predicates {
		p(selector)
	}
	for _, p := range aq.order {
		p(selector)
	}
	if offset := aq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ApprovalGroupBy is the group-by builder for Approval entities.
type ApprovalGroupBy struct {
	selector
	build *ApprovalQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *ApprovalGroupBy) Aggregate(fns ...AggregateFunc) *ApprovalGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *ApprovalGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApprovalQuery, *ApprovalGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *ApprovalGroupBy) sqlScan(ctx context.Context, root *ApprovalQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(agb.fns))
	for _, fn := range agb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*agb.flds)+len(agb.fns))
		for _, f := range *agb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*agb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ApprovalSelect is the builder for selecting fields of Approval entities.
type ApprovalSelect struct {
	*ApprovalQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *ApprovalSelect) Aggregate(fns ...AggregateFunc) *ApprovalSelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *ApprovalSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ApprovalQuery, *ApprovalSelect](ctx, as.ApprovalQuery, as, as.inters, v)
}

func (as *ApprovalSelect) sqlScan(ctx context.Context, root *ApprovalQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(as.fns))
	for _, fn := range as.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*as.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := as.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
