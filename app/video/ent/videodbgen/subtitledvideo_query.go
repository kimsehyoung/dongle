// Code generated by ent, DO NOT EDIT.

package videodbgen

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/originalvideo"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/predicate"
	"github.com/kimsehyoung/dongle/app/video/ent/videodbgen/subtitledvideo"
)

// SubtitledVideoQuery is the builder for querying SubtitledVideo entities.
type SubtitledVideoQuery struct {
	config
	ctx               *QueryContext
	order             []subtitledvideo.OrderOption
	inters            []Interceptor
	predicates        []predicate.SubtitledVideo
	withOriginalVideo *OriginalVideoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SubtitledVideoQuery builder.
func (svq *SubtitledVideoQuery) Where(ps ...predicate.SubtitledVideo) *SubtitledVideoQuery {
	svq.predicates = append(svq.predicates, ps...)
	return svq
}

// Limit the number of records to be returned by this query.
func (svq *SubtitledVideoQuery) Limit(limit int) *SubtitledVideoQuery {
	svq.ctx.Limit = &limit
	return svq
}

// Offset to start from.
func (svq *SubtitledVideoQuery) Offset(offset int) *SubtitledVideoQuery {
	svq.ctx.Offset = &offset
	return svq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (svq *SubtitledVideoQuery) Unique(unique bool) *SubtitledVideoQuery {
	svq.ctx.Unique = &unique
	return svq
}

// Order specifies how the records should be ordered.
func (svq *SubtitledVideoQuery) Order(o ...subtitledvideo.OrderOption) *SubtitledVideoQuery {
	svq.order = append(svq.order, o...)
	return svq
}

// QueryOriginalVideo chains the current query on the "original_video" edge.
func (svq *SubtitledVideoQuery) QueryOriginalVideo() *OriginalVideoQuery {
	query := (&OriginalVideoClient{config: svq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := svq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := svq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(subtitledvideo.Table, subtitledvideo.FieldID, selector),
			sqlgraph.To(originalvideo.Table, originalvideo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, subtitledvideo.OriginalVideoTable, subtitledvideo.OriginalVideoColumn),
		)
		fromU = sqlgraph.SetNeighbors(svq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SubtitledVideo entity from the query.
// Returns a *NotFoundError when no SubtitledVideo was found.
func (svq *SubtitledVideoQuery) First(ctx context.Context) (*SubtitledVideo, error) {
	nodes, err := svq.Limit(1).All(setContextOp(ctx, svq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{subtitledvideo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (svq *SubtitledVideoQuery) FirstX(ctx context.Context) *SubtitledVideo {
	node, err := svq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SubtitledVideo ID from the query.
// Returns a *NotFoundError when no SubtitledVideo ID was found.
func (svq *SubtitledVideoQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = svq.Limit(1).IDs(setContextOp(ctx, svq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{subtitledvideo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (svq *SubtitledVideoQuery) FirstIDX(ctx context.Context) int {
	id, err := svq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SubtitledVideo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SubtitledVideo entity is found.
// Returns a *NotFoundError when no SubtitledVideo entities are found.
func (svq *SubtitledVideoQuery) Only(ctx context.Context) (*SubtitledVideo, error) {
	nodes, err := svq.Limit(2).All(setContextOp(ctx, svq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{subtitledvideo.Label}
	default:
		return nil, &NotSingularError{subtitledvideo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (svq *SubtitledVideoQuery) OnlyX(ctx context.Context) *SubtitledVideo {
	node, err := svq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SubtitledVideo ID in the query.
// Returns a *NotSingularError when more than one SubtitledVideo ID is found.
// Returns a *NotFoundError when no entities are found.
func (svq *SubtitledVideoQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = svq.Limit(2).IDs(setContextOp(ctx, svq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{subtitledvideo.Label}
	default:
		err = &NotSingularError{subtitledvideo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (svq *SubtitledVideoQuery) OnlyIDX(ctx context.Context) int {
	id, err := svq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SubtitledVideos.
func (svq *SubtitledVideoQuery) All(ctx context.Context) ([]*SubtitledVideo, error) {
	ctx = setContextOp(ctx, svq.ctx, "All")
	if err := svq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SubtitledVideo, *SubtitledVideoQuery]()
	return withInterceptors[[]*SubtitledVideo](ctx, svq, qr, svq.inters)
}

// AllX is like All, but panics if an error occurs.
func (svq *SubtitledVideoQuery) AllX(ctx context.Context) []*SubtitledVideo {
	nodes, err := svq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SubtitledVideo IDs.
func (svq *SubtitledVideoQuery) IDs(ctx context.Context) (ids []int, err error) {
	if svq.ctx.Unique == nil && svq.path != nil {
		svq.Unique(true)
	}
	ctx = setContextOp(ctx, svq.ctx, "IDs")
	if err = svq.Select(subtitledvideo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (svq *SubtitledVideoQuery) IDsX(ctx context.Context) []int {
	ids, err := svq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (svq *SubtitledVideoQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, svq.ctx, "Count")
	if err := svq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, svq, querierCount[*SubtitledVideoQuery](), svq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (svq *SubtitledVideoQuery) CountX(ctx context.Context) int {
	count, err := svq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (svq *SubtitledVideoQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, svq.ctx, "Exist")
	switch _, err := svq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("videodbgen: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (svq *SubtitledVideoQuery) ExistX(ctx context.Context) bool {
	exist, err := svq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SubtitledVideoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (svq *SubtitledVideoQuery) Clone() *SubtitledVideoQuery {
	if svq == nil {
		return nil
	}
	return &SubtitledVideoQuery{
		config:            svq.config,
		ctx:               svq.ctx.Clone(),
		order:             append([]subtitledvideo.OrderOption{}, svq.order...),
		inters:            append([]Interceptor{}, svq.inters...),
		predicates:        append([]predicate.SubtitledVideo{}, svq.predicates...),
		withOriginalVideo: svq.withOriginalVideo.Clone(),
		// clone intermediate query.
		sql:  svq.sql.Clone(),
		path: svq.path,
	}
}

// WithOriginalVideo tells the query-builder to eager-load the nodes that are connected to
// the "original_video" edge. The optional arguments are used to configure the query builder of the edge.
func (svq *SubtitledVideoQuery) WithOriginalVideo(opts ...func(*OriginalVideoQuery)) *SubtitledVideoQuery {
	query := (&OriginalVideoClient{config: svq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	svq.withOriginalVideo = query
	return svq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OriginalVideoID int `json:"original_video_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SubtitledVideo.Query().
//		GroupBy(subtitledvideo.FieldOriginalVideoID).
//		Aggregate(videodbgen.Count()).
//		Scan(ctx, &v)
func (svq *SubtitledVideoQuery) GroupBy(field string, fields ...string) *SubtitledVideoGroupBy {
	svq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SubtitledVideoGroupBy{build: svq}
	grbuild.flds = &svq.ctx.Fields
	grbuild.label = subtitledvideo.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OriginalVideoID int `json:"original_video_id,omitempty"`
//	}
//
//	client.SubtitledVideo.Query().
//		Select(subtitledvideo.FieldOriginalVideoID).
//		Scan(ctx, &v)
func (svq *SubtitledVideoQuery) Select(fields ...string) *SubtitledVideoSelect {
	svq.ctx.Fields = append(svq.ctx.Fields, fields...)
	sbuild := &SubtitledVideoSelect{SubtitledVideoQuery: svq}
	sbuild.label = subtitledvideo.Label
	sbuild.flds, sbuild.scan = &svq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SubtitledVideoSelect configured with the given aggregations.
func (svq *SubtitledVideoQuery) Aggregate(fns ...AggregateFunc) *SubtitledVideoSelect {
	return svq.Select().Aggregate(fns...)
}

func (svq *SubtitledVideoQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range svq.inters {
		if inter == nil {
			return fmt.Errorf("videodbgen: uninitialized interceptor (forgotten import videodbgen/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, svq); err != nil {
				return err
			}
		}
	}
	for _, f := range svq.ctx.Fields {
		if !subtitledvideo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("videodbgen: invalid field %q for query", f)}
		}
	}
	if svq.path != nil {
		prev, err := svq.path(ctx)
		if err != nil {
			return err
		}
		svq.sql = prev
	}
	return nil
}

func (svq *SubtitledVideoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SubtitledVideo, error) {
	var (
		nodes       = []*SubtitledVideo{}
		_spec       = svq.querySpec()
		loadedTypes = [1]bool{
			svq.withOriginalVideo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SubtitledVideo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SubtitledVideo{config: svq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, svq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := svq.withOriginalVideo; query != nil {
		if err := svq.loadOriginalVideo(ctx, query, nodes, nil,
			func(n *SubtitledVideo, e *OriginalVideo) { n.Edges.OriginalVideo = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (svq *SubtitledVideoQuery) loadOriginalVideo(ctx context.Context, query *OriginalVideoQuery, nodes []*SubtitledVideo, init func(*SubtitledVideo), assign func(*SubtitledVideo, *OriginalVideo)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*SubtitledVideo)
	for i := range nodes {
		fk := nodes[i].OriginalVideoID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(originalvideo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "original_video_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (svq *SubtitledVideoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := svq.querySpec()
	_spec.Node.Columns = svq.ctx.Fields
	if len(svq.ctx.Fields) > 0 {
		_spec.Unique = svq.ctx.Unique != nil && *svq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, svq.driver, _spec)
}

func (svq *SubtitledVideoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(subtitledvideo.Table, subtitledvideo.Columns, sqlgraph.NewFieldSpec(subtitledvideo.FieldID, field.TypeInt))
	_spec.From = svq.sql
	if unique := svq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if svq.path != nil {
		_spec.Unique = true
	}
	if fields := svq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, subtitledvideo.FieldID)
		for i := range fields {
			if fields[i] != subtitledvideo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if svq.withOriginalVideo != nil {
			_spec.Node.AddColumnOnce(subtitledvideo.FieldOriginalVideoID)
		}
	}
	if ps := svq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := svq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := svq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := svq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (svq *SubtitledVideoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(svq.driver.Dialect())
	t1 := builder.Table(subtitledvideo.Table)
	columns := svq.ctx.Fields
	if len(columns) == 0 {
		columns = subtitledvideo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if svq.sql != nil {
		selector = svq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if svq.ctx.Unique != nil && *svq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range svq.predicates {
		p(selector)
	}
	for _, p := range svq.order {
		p(selector)
	}
	if offset := svq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := svq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SubtitledVideoGroupBy is the group-by builder for SubtitledVideo entities.
type SubtitledVideoGroupBy struct {
	selector
	build *SubtitledVideoQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (svgb *SubtitledVideoGroupBy) Aggregate(fns ...AggregateFunc) *SubtitledVideoGroupBy {
	svgb.fns = append(svgb.fns, fns...)
	return svgb
}

// Scan applies the selector query and scans the result into the given value.
func (svgb *SubtitledVideoGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, svgb.build.ctx, "GroupBy")
	if err := svgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubtitledVideoQuery, *SubtitledVideoGroupBy](ctx, svgb.build, svgb, svgb.build.inters, v)
}

func (svgb *SubtitledVideoGroupBy) sqlScan(ctx context.Context, root *SubtitledVideoQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(svgb.fns))
	for _, fn := range svgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*svgb.flds)+len(svgb.fns))
		for _, f := range *svgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*svgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := svgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SubtitledVideoSelect is the builder for selecting fields of SubtitledVideo entities.
type SubtitledVideoSelect struct {
	*SubtitledVideoQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (svs *SubtitledVideoSelect) Aggregate(fns ...AggregateFunc) *SubtitledVideoSelect {
	svs.fns = append(svs.fns, fns...)
	return svs
}

// Scan applies the selector query and scans the result into the given value.
func (svs *SubtitledVideoSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, svs.ctx, "Select")
	if err := svs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SubtitledVideoQuery, *SubtitledVideoSelect](ctx, svs.SubtitledVideoQuery, svs, svs.inters, v)
}

func (svs *SubtitledVideoSelect) sqlScan(ctx context.Context, root *SubtitledVideoQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(svs.fns))
	for _, fn := range svs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*svs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := svs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}