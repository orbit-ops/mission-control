// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/orbit-ops/launchpad-core/ent/mission"
	"github.com/orbit-ops/launchpad-core/ent/predicate"
	"github.com/orbit-ops/launchpad-core/ent/request"
	"github.com/orbit-ops/launchpad-core/ent/rocket"
)

// MissionQuery is the builder for querying Mission entities.
type MissionQuery struct {
	config
	ctx               *QueryContext
	order             []mission.OrderOption
	inters            []Interceptor
	predicates        []predicate.Mission
	withRockets       *RocketQuery
	withRequests      *RequestQuery
	withNamedRockets  map[string]*RocketQuery
	withNamedRequests map[string]*RequestQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MissionQuery builder.
func (mq *MissionQuery) Where(ps ...predicate.Mission) *MissionQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit the number of records to be returned by this query.
func (mq *MissionQuery) Limit(limit int) *MissionQuery {
	mq.ctx.Limit = &limit
	return mq
}

// Offset to start from.
func (mq *MissionQuery) Offset(offset int) *MissionQuery {
	mq.ctx.Offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *MissionQuery) Unique(unique bool) *MissionQuery {
	mq.ctx.Unique = &unique
	return mq
}

// Order specifies how the records should be ordered.
func (mq *MissionQuery) Order(o ...mission.OrderOption) *MissionQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryRockets chains the current query on the "rockets" edge.
func (mq *MissionQuery) QueryRockets() *RocketQuery {
	query := (&RocketClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mission.Table, mission.FieldID, selector),
			sqlgraph.To(rocket.Table, rocket.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, mission.RocketsTable, mission.RocketsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRequests chains the current query on the "requests" edge.
func (mq *MissionQuery) QueryRequests() *RequestQuery {
	query := (&RequestClient{config: mq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(mission.Table, mission.FieldID, selector),
			sqlgraph.To(request.Table, request.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, mission.RequestsTable, mission.RequestsColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Mission entity from the query.
// Returns a *NotFoundError when no Mission was found.
func (mq *MissionQuery) First(ctx context.Context) (*Mission, error) {
	nodes, err := mq.Limit(1).All(setContextOp(ctx, mq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mission.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *MissionQuery) FirstX(ctx context.Context) *Mission {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Mission ID from the query.
// Returns a *NotFoundError when no Mission ID was found.
func (mq *MissionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(1).IDs(setContextOp(ctx, mq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mission.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *MissionQuery) FirstIDX(ctx context.Context) string {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Mission entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Mission entity is found.
// Returns a *NotFoundError when no Mission entities are found.
func (mq *MissionQuery) Only(ctx context.Context) (*Mission, error) {
	nodes, err := mq.Limit(2).All(setContextOp(ctx, mq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mission.Label}
	default:
		return nil, &NotSingularError{mission.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *MissionQuery) OnlyX(ctx context.Context) *Mission {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Mission ID in the query.
// Returns a *NotSingularError when more than one Mission ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *MissionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = mq.Limit(2).IDs(setContextOp(ctx, mq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mission.Label}
	default:
		err = &NotSingularError{mission.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *MissionQuery) OnlyIDX(ctx context.Context) string {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Missions.
func (mq *MissionQuery) All(ctx context.Context) ([]*Mission, error) {
	ctx = setContextOp(ctx, mq.ctx, "All")
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Mission, *MissionQuery]()
	return withInterceptors[[]*Mission](ctx, mq, qr, mq.inters)
}

// AllX is like All, but panics if an error occurs.
func (mq *MissionQuery) AllX(ctx context.Context) []*Mission {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Mission IDs.
func (mq *MissionQuery) IDs(ctx context.Context) (ids []string, err error) {
	if mq.ctx.Unique == nil && mq.path != nil {
		mq.Unique(true)
	}
	ctx = setContextOp(ctx, mq.ctx, "IDs")
	if err = mq.Select(mission.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *MissionQuery) IDsX(ctx context.Context) []string {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *MissionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, mq.ctx, "Count")
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, mq, querierCount[*MissionQuery](), mq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (mq *MissionQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *MissionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, mq.ctx, "Exist")
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *MissionQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MissionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *MissionQuery) Clone() *MissionQuery {
	if mq == nil {
		return nil
	}
	return &MissionQuery{
		config:       mq.config,
		ctx:          mq.ctx.Clone(),
		order:        append([]mission.OrderOption{}, mq.order...),
		inters:       append([]Interceptor{}, mq.inters...),
		predicates:   append([]predicate.Mission{}, mq.predicates...),
		withRockets:  mq.withRockets.Clone(),
		withRequests: mq.withRequests.Clone(),
		// clone intermediate query.
		sql:  mq.sql.Clone(),
		path: mq.path,
	}
}

// WithRockets tells the query-builder to eager-load the nodes that are connected to
// the "rockets" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MissionQuery) WithRockets(opts ...func(*RocketQuery)) *MissionQuery {
	query := (&RocketClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withRockets = query
	return mq
}

// WithRequests tells the query-builder to eager-load the nodes that are connected to
// the "requests" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *MissionQuery) WithRequests(opts ...func(*RequestQuery)) *MissionQuery {
	query := (&RequestClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	mq.withRequests = query
	return mq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Mission.Query().
//		GroupBy(mission.FieldDescription).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mq *MissionQuery) GroupBy(field string, fields ...string) *MissionGroupBy {
	mq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MissionGroupBy{build: mq}
	grbuild.flds = &mq.ctx.Fields
	grbuild.label = mission.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Description string `json:"description,omitempty"`
//	}
//
//	client.Mission.Query().
//		Select(mission.FieldDescription).
//		Scan(ctx, &v)
func (mq *MissionQuery) Select(fields ...string) *MissionSelect {
	mq.ctx.Fields = append(mq.ctx.Fields, fields...)
	sbuild := &MissionSelect{MissionQuery: mq}
	sbuild.label = mission.Label
	sbuild.flds, sbuild.scan = &mq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MissionSelect configured with the given aggregations.
func (mq *MissionQuery) Aggregate(fns ...AggregateFunc) *MissionSelect {
	return mq.Select().Aggregate(fns...)
}

func (mq *MissionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range mq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, mq); err != nil {
				return err
			}
		}
	}
	for _, f := range mq.ctx.Fields {
		if !mission.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *MissionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Mission, error) {
	var (
		nodes       = []*Mission{}
		_spec       = mq.querySpec()
		loadedTypes = [2]bool{
			mq.withRockets != nil,
			mq.withRequests != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Mission).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Mission{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withRockets; query != nil {
		if err := mq.loadRockets(ctx, query, nodes,
			func(n *Mission) { n.Edges.Rockets = []*Rocket{} },
			func(n *Mission, e *Rocket) { n.Edges.Rockets = append(n.Edges.Rockets, e) }); err != nil {
			return nil, err
		}
	}
	if query := mq.withRequests; query != nil {
		if err := mq.loadRequests(ctx, query, nodes,
			func(n *Mission) { n.Edges.Requests = []*Request{} },
			func(n *Mission, e *Request) { n.Edges.Requests = append(n.Edges.Requests, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mq.withNamedRockets {
		if err := mq.loadRockets(ctx, query, nodes,
			func(n *Mission) { n.appendNamedRockets(name) },
			func(n *Mission, e *Rocket) { n.appendNamedRockets(name, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range mq.withNamedRequests {
		if err := mq.loadRequests(ctx, query, nodes,
			func(n *Mission) { n.appendNamedRequests(name) },
			func(n *Mission, e *Request) { n.appendNamedRequests(name, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *MissionQuery) loadRockets(ctx context.Context, query *RocketQuery, nodes []*Mission, init func(*Mission), assign func(*Mission, *Rocket)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[string]*Mission)
	nids := make(map[string]map[*Mission]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(mission.RocketsTable)
		s.Join(joinT).On(s.C(rocket.FieldID), joinT.C(mission.RocketsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(mission.RocketsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(mission.RocketsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullString)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := values[0].(*sql.NullString).String
				inValue := values[1].(*sql.NullString).String
				if nids[inValue] == nil {
					nids[inValue] = map[*Mission]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Rocket](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "rockets" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (mq *MissionQuery) loadRequests(ctx context.Context, query *RequestQuery, nodes []*Mission, init func(*Mission), assign func(*Mission, *Request)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Mission)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(request.FieldMissionID)
	}
	query.Where(predicate.Request(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(mission.RequestsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.MissionID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "mission_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *MissionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.ctx.Fields
	if len(mq.ctx.Fields) > 0 {
		_spec.Unique = mq.ctx.Unique != nil && *mq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *MissionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(mission.Table, mission.Columns, sqlgraph.NewFieldSpec(mission.FieldID, field.TypeString))
	_spec.From = mq.sql
	if unique := mq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if mq.path != nil {
		_spec.Unique = true
	}
	if fields := mq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mission.FieldID)
		for i := range fields {
			if fields[i] != mission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *MissionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(mission.Table)
	columns := mq.ctx.Fields
	if len(columns) == 0 {
		columns = mission.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.ctx.Unique != nil && *mq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedRockets tells the query-builder to eager-load the nodes that are connected to the "rockets"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mq *MissionQuery) WithNamedRockets(name string, opts ...func(*RocketQuery)) *MissionQuery {
	query := (&RocketClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mq.withNamedRockets == nil {
		mq.withNamedRockets = make(map[string]*RocketQuery)
	}
	mq.withNamedRockets[name] = query
	return mq
}

// WithNamedRequests tells the query-builder to eager-load the nodes that are connected to the "requests"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (mq *MissionQuery) WithNamedRequests(name string, opts ...func(*RequestQuery)) *MissionQuery {
	query := (&RequestClient{config: mq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if mq.withNamedRequests == nil {
		mq.withNamedRequests = make(map[string]*RequestQuery)
	}
	mq.withNamedRequests[name] = query
	return mq
}

// MissionGroupBy is the group-by builder for Mission entities.
type MissionGroupBy struct {
	selector
	build *MissionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *MissionGroupBy) Aggregate(fns ...AggregateFunc) *MissionGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the selector query and scans the result into the given value.
func (mgb *MissionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mgb.build.ctx, "GroupBy")
	if err := mgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionQuery, *MissionGroupBy](ctx, mgb.build, mgb, mgb.build.inters, v)
}

func (mgb *MissionGroupBy) sqlScan(ctx context.Context, root *MissionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*mgb.flds)+len(mgb.fns))
		for _, f := range *mgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*mgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MissionSelect is the builder for selecting fields of Mission entities.
type MissionSelect struct {
	*MissionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ms *MissionSelect) Aggregate(fns ...AggregateFunc) *MissionSelect {
	ms.fns = append(ms.fns, fns...)
	return ms
}

// Scan applies the selector query and scans the result into the given value.
func (ms *MissionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ms.ctx, "Select")
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MissionQuery, *MissionSelect](ctx, ms.MissionQuery, ms, ms.inters, v)
}

func (ms *MissionSelect) sqlScan(ctx context.Context, root *MissionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ms.fns))
	for _, fn := range ms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
