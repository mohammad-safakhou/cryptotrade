// Code generated by SQLBoiler 4.9.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Position is an object representing the database table.
type Position struct {
	ID           int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	MarketPrice  null.Float64 `boil:"market_price" json:"market_price,omitempty" toml:"market_price" yaml:"market_price,omitempty"`
	Side         null.String  `boil:"side" json:"side,omitempty" toml:"side" yaml:"side,omitempty"`
	Leverage     null.String  `boil:"leverage" json:"leverage,omitempty" toml:"leverage" yaml:"leverage,omitempty"`
	PositionSize null.String  `boil:"position_size" json:"position_size,omitempty" toml:"position_size" yaml:"position_size,omitempty"`
	PositionType null.String  `boil:"position_type" json:"position_type,omitempty" toml:"position_type" yaml:"position_type,omitempty"`
	Symbol       null.String  `boil:"symbol" json:"symbol,omitempty" toml:"symbol" yaml:"symbol,omitempty"`
	IsClose      null.Bool    `boil:"is_close" json:"is_close,omitempty" toml:"is_close" yaml:"is_close,omitempty"`
	CreatedAt    null.Time    `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt    null.Time    `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt    null.Time    `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`

	R *positionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L positionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PositionColumns = struct {
	ID           string
	MarketPrice  string
	Side         string
	Leverage     string
	PositionSize string
	PositionType string
	Symbol       string
	IsClose      string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "id",
	MarketPrice:  "market_price",
	Side:         "side",
	Leverage:     "leverage",
	PositionSize: "position_size",
	PositionType: "position_type",
	Symbol:       "symbol",
	IsClose:      "is_close",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

var PositionTableColumns = struct {
	ID           string
	MarketPrice  string
	Side         string
	Leverage     string
	PositionSize string
	PositionType string
	Symbol       string
	IsClose      string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}{
	ID:           "positions.id",
	MarketPrice:  "positions.market_price",
	Side:         "positions.side",
	Leverage:     "positions.leverage",
	PositionSize: "positions.position_size",
	PositionType: "positions.position_type",
	Symbol:       "positions.symbol",
	IsClose:      "positions.is_close",
	CreatedAt:    "positions.created_at",
	UpdatedAt:    "positions.updated_at",
	DeletedAt:    "positions.deleted_at",
}

// Generated where

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var PositionWhere = struct {
	ID           whereHelperint
	MarketPrice  whereHelpernull_Float64
	Side         whereHelpernull_String
	Leverage     whereHelpernull_String
	PositionSize whereHelpernull_String
	PositionType whereHelpernull_String
	Symbol       whereHelpernull_String
	IsClose      whereHelpernull_Bool
	CreatedAt    whereHelpernull_Time
	UpdatedAt    whereHelpernull_Time
	DeletedAt    whereHelpernull_Time
}{
	ID:           whereHelperint{field: "\"positions\".\"id\""},
	MarketPrice:  whereHelpernull_Float64{field: "\"positions\".\"market_price\""},
	Side:         whereHelpernull_String{field: "\"positions\".\"side\""},
	Leverage:     whereHelpernull_String{field: "\"positions\".\"leverage\""},
	PositionSize: whereHelpernull_String{field: "\"positions\".\"position_size\""},
	PositionType: whereHelpernull_String{field: "\"positions\".\"position_type\""},
	Symbol:       whereHelpernull_String{field: "\"positions\".\"symbol\""},
	IsClose:      whereHelpernull_Bool{field: "\"positions\".\"is_close\""},
	CreatedAt:    whereHelpernull_Time{field: "\"positions\".\"created_at\""},
	UpdatedAt:    whereHelpernull_Time{field: "\"positions\".\"updated_at\""},
	DeletedAt:    whereHelpernull_Time{field: "\"positions\".\"deleted_at\""},
}

// PositionRels is where relationship names are stored.
var PositionRels = struct {
}{}

// positionR is where relationships are stored.
type positionR struct {
}

// NewStruct creates a new relationship struct
func (*positionR) NewStruct() *positionR {
	return &positionR{}
}

// positionL is where Load methods for each relationship are stored.
type positionL struct{}

var (
	positionAllColumns            = []string{"id", "market_price", "side", "leverage", "position_size", "position_type", "symbol", "is_close", "created_at", "updated_at", "deleted_at"}
	positionColumnsWithoutDefault = []string{}
	positionColumnsWithDefault    = []string{"id", "market_price", "side", "leverage", "position_size", "position_type", "symbol", "is_close", "created_at", "updated_at", "deleted_at"}
	positionPrimaryKeyColumns     = []string{"id"}
	positionGeneratedColumns      = []string{}
)

type (
	// PositionSlice is an alias for a slice of pointers to Position.
	// This should almost always be used instead of []Position.
	PositionSlice []*Position
	// PositionHook is the signature for custom Position hook methods
	PositionHook func(context.Context, boil.ContextExecutor, *Position) error

	positionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	positionType                 = reflect.TypeOf(&Position{})
	positionMapping              = queries.MakeStructMapping(positionType)
	positionPrimaryKeyMapping, _ = queries.BindMapping(positionType, positionMapping, positionPrimaryKeyColumns)
	positionInsertCacheMut       sync.RWMutex
	positionInsertCache          = make(map[string]insertCache)
	positionUpdateCacheMut       sync.RWMutex
	positionUpdateCache          = make(map[string]updateCache)
	positionUpsertCacheMut       sync.RWMutex
	positionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var positionAfterSelectHooks []PositionHook

var positionBeforeInsertHooks []PositionHook
var positionAfterInsertHooks []PositionHook

var positionBeforeUpdateHooks []PositionHook
var positionAfterUpdateHooks []PositionHook

var positionBeforeDeleteHooks []PositionHook
var positionAfterDeleteHooks []PositionHook

var positionBeforeUpsertHooks []PositionHook
var positionAfterUpsertHooks []PositionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Position) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Position) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Position) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Position) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Position) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Position) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Position) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Position) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Position) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range positionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPositionHook registers your hook function for all future operations.
func AddPositionHook(hookPoint boil.HookPoint, positionHook PositionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		positionAfterSelectHooks = append(positionAfterSelectHooks, positionHook)
	case boil.BeforeInsertHook:
		positionBeforeInsertHooks = append(positionBeforeInsertHooks, positionHook)
	case boil.AfterInsertHook:
		positionAfterInsertHooks = append(positionAfterInsertHooks, positionHook)
	case boil.BeforeUpdateHook:
		positionBeforeUpdateHooks = append(positionBeforeUpdateHooks, positionHook)
	case boil.AfterUpdateHook:
		positionAfterUpdateHooks = append(positionAfterUpdateHooks, positionHook)
	case boil.BeforeDeleteHook:
		positionBeforeDeleteHooks = append(positionBeforeDeleteHooks, positionHook)
	case boil.AfterDeleteHook:
		positionAfterDeleteHooks = append(positionAfterDeleteHooks, positionHook)
	case boil.BeforeUpsertHook:
		positionBeforeUpsertHooks = append(positionBeforeUpsertHooks, positionHook)
	case boil.AfterUpsertHook:
		positionAfterUpsertHooks = append(positionAfterUpsertHooks, positionHook)
	}
}

// One returns a single position record from the query.
func (q positionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Position, error) {
	o := &Position{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for positions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Position records from the query.
func (q positionQuery) All(ctx context.Context, exec boil.ContextExecutor) (PositionSlice, error) {
	var o []*Position

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Position slice")
	}

	if len(positionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Position records in the query.
func (q positionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count positions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q positionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if positions exists")
	}

	return count > 0, nil
}

// Positions retrieves all the records using an executor.
func Positions(mods ...qm.QueryMod) positionQuery {
	mods = append(mods, qm.From("\"positions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"positions\".*"})
	}

	return positionQuery{NewQuery(mods...)}
}

// FindPosition retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPosition(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Position, error) {
	positionObj := &Position{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"positions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, positionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from positions")
	}

	if err = positionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return positionObj, err
	}

	return positionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Position) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no positions provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(positionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	positionInsertCacheMut.RLock()
	cache, cached := positionInsertCache[key]
	positionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			positionAllColumns,
			positionColumnsWithDefault,
			positionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(positionType, positionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(positionType, positionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"positions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"positions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into positions")
	}

	if !cached {
		positionInsertCacheMut.Lock()
		positionInsertCache[key] = cache
		positionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Position.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Position) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	positionUpdateCacheMut.RLock()
	cache, cached := positionUpdateCache[key]
	positionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			positionAllColumns,
			positionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update positions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"positions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, positionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(positionType, positionMapping, append(wl, positionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update positions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for positions")
	}

	if !cached {
		positionUpdateCacheMut.Lock()
		positionUpdateCache[key] = cache
		positionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q positionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for positions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for positions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PositionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), positionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"positions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, positionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in position slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all position")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Position) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no positions provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(positionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	positionUpsertCacheMut.RLock()
	cache, cached := positionUpsertCache[key]
	positionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			positionAllColumns,
			positionColumnsWithDefault,
			positionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			positionAllColumns,
			positionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert positions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(positionPrimaryKeyColumns))
			copy(conflict, positionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"positions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(positionType, positionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(positionType, positionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert positions")
	}

	if !cached {
		positionUpsertCacheMut.Lock()
		positionUpsertCache[key] = cache
		positionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Position record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Position) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Position provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), positionPrimaryKeyMapping)
	sql := "DELETE FROM \"positions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from positions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for positions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q positionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no positionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from positions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for positions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PositionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(positionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), positionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"positions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, positionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from position slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for positions")
	}

	if len(positionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Position) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPosition(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PositionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PositionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), positionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"positions\".* FROM \"positions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, positionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PositionSlice")
	}

	*o = slice

	return nil
}

// PositionExists checks if the Position row exists.
func PositionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"positions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if positions exists")
	}

	return exists, nil
}