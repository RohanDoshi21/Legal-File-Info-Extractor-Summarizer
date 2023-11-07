// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// Log is an object representing the database table.
type Log struct {
	ID          int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt   null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	DocumentID  int       `boil:"document_id" json:"document_id" toml:"document_id" yaml:"document_id"`
	PrevContent null.JSON `boil:"prev_content" json:"prev_content,omitempty" toml:"prev_content" yaml:"prev_content,omitempty"`
	NewContent  null.JSON `boil:"new_content" json:"new_content,omitempty" toml:"new_content" yaml:"new_content,omitempty"`
	EditedBy    int       `boil:"edited_by" json:"edited_by" toml:"edited_by" yaml:"edited_by"`

	R *logR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L logL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LogColumns = struct {
	ID          string
	CreatedAt   string
	DocumentID  string
	PrevContent string
	NewContent  string
	EditedBy    string
}{
	ID:          "id",
	CreatedAt:   "created_at",
	DocumentID:  "document_id",
	PrevContent: "prev_content",
	NewContent:  "new_content",
	EditedBy:    "edited_by",
}

var LogTableColumns = struct {
	ID          string
	CreatedAt   string
	DocumentID  string
	PrevContent string
	NewContent  string
	EditedBy    string
}{
	ID:          "logs.id",
	CreatedAt:   "logs.created_at",
	DocumentID:  "logs.document_id",
	PrevContent: "logs.prev_content",
	NewContent:  "logs.new_content",
	EditedBy:    "logs.edited_by",
}

// Generated where

var LogWhere = struct {
	ID          whereHelperint
	CreatedAt   whereHelpernull_Time
	DocumentID  whereHelperint
	PrevContent whereHelpernull_JSON
	NewContent  whereHelpernull_JSON
	EditedBy    whereHelperint
}{
	ID:          whereHelperint{field: "\"logs\".\"id\""},
	CreatedAt:   whereHelpernull_Time{field: "\"logs\".\"created_at\""},
	DocumentID:  whereHelperint{field: "\"logs\".\"document_id\""},
	PrevContent: whereHelpernull_JSON{field: "\"logs\".\"prev_content\""},
	NewContent:  whereHelpernull_JSON{field: "\"logs\".\"new_content\""},
	EditedBy:    whereHelperint{field: "\"logs\".\"edited_by\""},
}

// LogRels is where relationship names are stored.
var LogRels = struct {
	Document string
}{
	Document: "Document",
}

// logR is where relationships are stored.
type logR struct {
	Document *Document `boil:"Document" json:"Document" toml:"Document" yaml:"Document"`
}

// NewStruct creates a new relationship struct
func (*logR) NewStruct() *logR {
	return &logR{}
}

func (r *logR) GetDocument() *Document {
	if r == nil {
		return nil
	}
	return r.Document
}

// logL is where Load methods for each relationship are stored.
type logL struct{}

var (
	logAllColumns            = []string{"id", "created_at", "document_id", "prev_content", "new_content", "edited_by"}
	logColumnsWithoutDefault = []string{"document_id", "edited_by"}
	logColumnsWithDefault    = []string{"id", "created_at", "prev_content", "new_content"}
	logPrimaryKeyColumns     = []string{"id"}
	logGeneratedColumns      = []string{}
)

type (
	// LogSlice is an alias for a slice of pointers to Log.
	// This should almost always be used instead of []Log.
	LogSlice []*Log
	// LogHook is the signature for custom Log hook methods
	LogHook func(context.Context, boil.ContextExecutor, *Log) error

	logQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	logType                 = reflect.TypeOf(&Log{})
	logMapping              = queries.MakeStructMapping(logType)
	logPrimaryKeyMapping, _ = queries.BindMapping(logType, logMapping, logPrimaryKeyColumns)
	logInsertCacheMut       sync.RWMutex
	logInsertCache          = make(map[string]insertCache)
	logUpdateCacheMut       sync.RWMutex
	logUpdateCache          = make(map[string]updateCache)
	logUpsertCacheMut       sync.RWMutex
	logUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var logAfterSelectHooks []LogHook

var logBeforeInsertHooks []LogHook
var logAfterInsertHooks []LogHook

var logBeforeUpdateHooks []LogHook
var logAfterUpdateHooks []LogHook

var logBeforeDeleteHooks []LogHook
var logAfterDeleteHooks []LogHook

var logBeforeUpsertHooks []LogHook
var logAfterUpsertHooks []LogHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Log) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Log) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Log) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Log) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Log) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Log) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Log) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Log) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Log) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range logAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLogHook registers your hook function for all future operations.
func AddLogHook(hookPoint boil.HookPoint, logHook LogHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		logAfterSelectHooks = append(logAfterSelectHooks, logHook)
	case boil.BeforeInsertHook:
		logBeforeInsertHooks = append(logBeforeInsertHooks, logHook)
	case boil.AfterInsertHook:
		logAfterInsertHooks = append(logAfterInsertHooks, logHook)
	case boil.BeforeUpdateHook:
		logBeforeUpdateHooks = append(logBeforeUpdateHooks, logHook)
	case boil.AfterUpdateHook:
		logAfterUpdateHooks = append(logAfterUpdateHooks, logHook)
	case boil.BeforeDeleteHook:
		logBeforeDeleteHooks = append(logBeforeDeleteHooks, logHook)
	case boil.AfterDeleteHook:
		logAfterDeleteHooks = append(logAfterDeleteHooks, logHook)
	case boil.BeforeUpsertHook:
		logBeforeUpsertHooks = append(logBeforeUpsertHooks, logHook)
	case boil.AfterUpsertHook:
		logAfterUpsertHooks = append(logAfterUpsertHooks, logHook)
	}
}

// One returns a single log record from the query.
func (q logQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Log, error) {
	o := &Log{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for logs")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Log records from the query.
func (q logQuery) All(ctx context.Context, exec boil.ContextExecutor) (LogSlice, error) {
	var o []*Log

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Log slice")
	}

	if len(logAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Log records in the query.
func (q logQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count logs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q logQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if logs exists")
	}

	return count > 0, nil
}

// Document pointed to by the foreign key.
func (o *Log) Document(mods ...qm.QueryMod) documentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DocumentID),
	}

	queryMods = append(queryMods, mods...)

	return Documents(queryMods...)
}

// LoadDocument allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (logL) LoadDocument(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLog interface{}, mods queries.Applicator) error {
	var slice []*Log
	var object *Log

	if singular {
		var ok bool
		object, ok = maybeLog.(*Log)
		if !ok {
			object = new(Log)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeLog))
			}
		}
	} else {
		s, ok := maybeLog.(*[]*Log)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeLog)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeLog))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &logR{}
		}
		args = append(args, object.DocumentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &logR{}
			}

			for _, a := range args {
				if a == obj.DocumentID {
					continue Outer
				}
			}

			args = append(args, obj.DocumentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`documents`),
		qm.WhereIn(`documents.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Document")
	}

	var resultSlice []*Document
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Document")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for documents")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for documents")
	}

	if len(documentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Document = foreign
		if foreign.R == nil {
			foreign.R = &documentR{}
		}
		foreign.R.Logs = append(foreign.R.Logs, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DocumentID == foreign.ID {
				local.R.Document = foreign
				if foreign.R == nil {
					foreign.R = &documentR{}
				}
				foreign.R.Logs = append(foreign.R.Logs, local)
				break
			}
		}
	}

	return nil
}

// SetDocument of the log to the related item.
// Sets o.R.Document to related.
// Adds o to related.R.Logs.
func (o *Log) SetDocument(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Document) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"document_id"}),
		strmangle.WhereClause("\"", "\"", 2, logPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DocumentID = related.ID
	if o.R == nil {
		o.R = &logR{
			Document: related,
		}
	} else {
		o.R.Document = related
	}

	if related.R == nil {
		related.R = &documentR{
			Logs: LogSlice{o},
		}
	} else {
		related.R.Logs = append(related.R.Logs, o)
	}

	return nil
}

// Logs retrieves all the records using an executor.
func Logs(mods ...qm.QueryMod) logQuery {
	mods = append(mods, qm.From("\"logs\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"logs\".*"})
	}

	return logQuery{q}
}

// FindLog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLog(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Log, error) {
	logObj := &Log{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"logs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, logObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from logs")
	}

	if err = logObj.doAfterSelectHooks(ctx, exec); err != nil {
		return logObj, err
	}

	return logObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Log) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no logs provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(logColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	logInsertCacheMut.RLock()
	cache, cached := logInsertCache[key]
	logInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			logAllColumns,
			logColumnsWithDefault,
			logColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(logType, logMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(logType, logMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"logs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"logs\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into logs")
	}

	if !cached {
		logInsertCacheMut.Lock()
		logInsertCache[key] = cache
		logInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Log.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Log) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	logUpdateCacheMut.RLock()
	cache, cached := logUpdateCache[key]
	logUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			logAllColumns,
			logPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update logs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"logs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, logPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(logType, logMapping, append(wl, logPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update logs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for logs")
	}

	if !cached {
		logUpdateCacheMut.Lock()
		logUpdateCache[key] = cache
		logUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q logQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for logs")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LogSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"logs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, logPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in log slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all log")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Log) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no logs provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if queries.MustTime(o.CreatedAt).IsZero() {
			queries.SetScanner(&o.CreatedAt, currTime)
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(logColumnsWithDefault, o)

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

	logUpsertCacheMut.RLock()
	cache, cached := logUpsertCache[key]
	logUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			logAllColumns,
			logColumnsWithDefault,
			logColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			logAllColumns,
			logPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert logs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(logPrimaryKeyColumns))
			copy(conflict, logPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"logs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(logType, logMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(logType, logMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert logs")
	}

	if !cached {
		logUpsertCacheMut.Lock()
		logUpsertCache[key] = cache
		logUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Log record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Log) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Log provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), logPrimaryKeyMapping)
	sql := "DELETE FROM \"logs\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for logs")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q logQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no logQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from logs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for logs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LogSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(logBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, logPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from log slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for logs")
	}

	if len(logAfterDeleteHooks) != 0 {
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
func (o *Log) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLog(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LogSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), logPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"logs\".* FROM \"logs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, logPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in LogSlice")
	}

	*o = slice

	return nil
}

// LogExists checks if the Log row exists.
func LogExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"logs\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if logs exists")
	}

	return exists, nil
}

// Exists checks if the Log row exists.
func (o *Log) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LogExists(ctx, exec, o.ID)
}