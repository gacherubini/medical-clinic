// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Patient is an object representing the database table.
type Patient struct {
	PatientID int `boil:"patient_id" json:"patient_id" toml:"patient_id" yaml:"patient_id"`
	UserID    int `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`

	R *patientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L patientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PatientColumns = struct {
	PatientID string
	UserID    string
}{
	PatientID: "patient_id",
	UserID:    "user_id",
}

var PatientTableColumns = struct {
	PatientID string
	UserID    string
}{
	PatientID: "patients.patient_id",
	UserID:    "patients.user_id",
}

// Generated where

var PatientWhere = struct {
	PatientID whereHelperint
	UserID    whereHelperint
}{
	PatientID: whereHelperint{field: "\"patients\".\"patient_id\""},
	UserID:    whereHelperint{field: "\"patients\".\"user_id\""},
}

// PatientRels is where relationship names are stored.
var PatientRels = struct {
	User string
}{
	User: "User",
}

// patientR is where relationships are stored.
type patientR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*patientR) NewStruct() *patientR {
	return &patientR{}
}

func (r *patientR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// patientL is where Load methods for each relationship are stored.
type patientL struct{}

var (
	patientAllColumns            = []string{"patient_id", "user_id"}
	patientColumnsWithoutDefault = []string{}
	patientColumnsWithDefault    = []string{"patient_id", "user_id"}
	patientPrimaryKeyColumns     = []string{"patient_id"}
	patientGeneratedColumns      = []string{}
)

type (
	// PatientSlice is an alias for a slice of pointers to Patient.
	// This should almost always be used instead of []Patient.
	PatientSlice []*Patient
	// PatientHook is the signature for custom Patient hook methods
	PatientHook func(context.Context, boil.ContextExecutor, *Patient) error

	patientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	patientType                 = reflect.TypeOf(&Patient{})
	patientMapping              = queries.MakeStructMapping(patientType)
	patientPrimaryKeyMapping, _ = queries.BindMapping(patientType, patientMapping, patientPrimaryKeyColumns)
	patientInsertCacheMut       sync.RWMutex
	patientInsertCache          = make(map[string]insertCache)
	patientUpdateCacheMut       sync.RWMutex
	patientUpdateCache          = make(map[string]updateCache)
	patientUpsertCacheMut       sync.RWMutex
	patientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var patientAfterSelectMu sync.Mutex
var patientAfterSelectHooks []PatientHook

var patientBeforeInsertMu sync.Mutex
var patientBeforeInsertHooks []PatientHook
var patientAfterInsertMu sync.Mutex
var patientAfterInsertHooks []PatientHook

var patientBeforeUpdateMu sync.Mutex
var patientBeforeUpdateHooks []PatientHook
var patientAfterUpdateMu sync.Mutex
var patientAfterUpdateHooks []PatientHook

var patientBeforeDeleteMu sync.Mutex
var patientBeforeDeleteHooks []PatientHook
var patientAfterDeleteMu sync.Mutex
var patientAfterDeleteHooks []PatientHook

var patientBeforeUpsertMu sync.Mutex
var patientBeforeUpsertHooks []PatientHook
var patientAfterUpsertMu sync.Mutex
var patientAfterUpsertHooks []PatientHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Patient) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Patient) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Patient) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Patient) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Patient) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Patient) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Patient) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Patient) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Patient) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range patientAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPatientHook registers your hook function for all future operations.
func AddPatientHook(hookPoint boil.HookPoint, patientHook PatientHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		patientAfterSelectMu.Lock()
		patientAfterSelectHooks = append(patientAfterSelectHooks, patientHook)
		patientAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		patientBeforeInsertMu.Lock()
		patientBeforeInsertHooks = append(patientBeforeInsertHooks, patientHook)
		patientBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		patientAfterInsertMu.Lock()
		patientAfterInsertHooks = append(patientAfterInsertHooks, patientHook)
		patientAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		patientBeforeUpdateMu.Lock()
		patientBeforeUpdateHooks = append(patientBeforeUpdateHooks, patientHook)
		patientBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		patientAfterUpdateMu.Lock()
		patientAfterUpdateHooks = append(patientAfterUpdateHooks, patientHook)
		patientAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		patientBeforeDeleteMu.Lock()
		patientBeforeDeleteHooks = append(patientBeforeDeleteHooks, patientHook)
		patientBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		patientAfterDeleteMu.Lock()
		patientAfterDeleteHooks = append(patientAfterDeleteHooks, patientHook)
		patientAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		patientBeforeUpsertMu.Lock()
		patientBeforeUpsertHooks = append(patientBeforeUpsertHooks, patientHook)
		patientBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		patientAfterUpsertMu.Lock()
		patientAfterUpsertHooks = append(patientAfterUpsertHooks, patientHook)
		patientAfterUpsertMu.Unlock()
	}
}

// One returns a single patient record from the query.
func (q patientQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Patient, error) {
	o := &Patient{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for patients")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Patient records from the query.
func (q patientQuery) All(ctx context.Context, exec boil.ContextExecutor) (PatientSlice, error) {
	var o []*Patient

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Patient slice")
	}

	if len(patientAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Patient records in the query.
func (q patientQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count patients rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q patientQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if patients exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Patient) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"user_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (patientL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePatient interface{}, mods queries.Applicator) error {
	var slice []*Patient
	var object *Patient

	if singular {
		var ok bool
		object, ok = maybePatient.(*Patient)
		if !ok {
			object = new(Patient)
			ok = queries.SetFromEmbeddedStruct(&object, &maybePatient)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybePatient))
			}
		}
	} else {
		s, ok := maybePatient.(*[]*Patient)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybePatient)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybePatient))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &patientR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &patientR{}
			}

			args[obj.UserID] = struct{}{}

		}
	}

	if len(args) == 0 {
		return nil
	}

	argsSlice := make([]interface{}, len(args))
	i := 0
	for arg := range args {
		argsSlice[i] = arg
		i++
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.user_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(userAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Patients = append(foreign.R.Patients, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Patients = append(foreign.R.Patients, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the patient to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Patients.
func (o *Patient) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"patients\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, patientPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.PatientID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.UserID
	if o.R == nil {
		o.R = &patientR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Patients: PatientSlice{o},
		}
	} else {
		related.R.Patients = append(related.R.Patients, o)
	}

	return nil
}

// Patients retrieves all the records using an executor.
func Patients(mods ...qm.QueryMod) patientQuery {
	mods = append(mods, qm.From("\"patients\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"patients\".*"})
	}

	return patientQuery{q}
}

// FindPatient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPatient(ctx context.Context, exec boil.ContextExecutor, patientID int, selectCols ...string) (*Patient, error) {
	patientObj := &Patient{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"patients\" where \"patient_id\"=$1", sel,
	)

	q := queries.Raw(query, patientID)

	err := q.Bind(ctx, exec, patientObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from patients")
	}

	if err = patientObj.doAfterSelectHooks(ctx, exec); err != nil {
		return patientObj, err
	}

	return patientObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Patient) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no patients provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(patientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	patientInsertCacheMut.RLock()
	cache, cached := patientInsertCache[key]
	patientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			patientAllColumns,
			patientColumnsWithDefault,
			patientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(patientType, patientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(patientType, patientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"patients\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"patients\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into patients")
	}

	if !cached {
		patientInsertCacheMut.Lock()
		patientInsertCache[key] = cache
		patientInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Patient.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Patient) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	patientUpdateCacheMut.RLock()
	cache, cached := patientUpdateCache[key]
	patientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			patientAllColumns,
			patientPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update patients, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"patients\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, patientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(patientType, patientMapping, append(wl, patientPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update patients row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for patients")
	}

	if !cached {
		patientUpdateCacheMut.Lock()
		patientUpdateCache[key] = cache
		patientUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q patientQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for patients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for patients")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PatientSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), patientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"patients\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, patientPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in patient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all patient")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Patient) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no patients provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(patientColumnsWithDefault, o)

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

	patientUpsertCacheMut.RLock()
	cache, cached := patientUpsertCache[key]
	patientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			patientAllColumns,
			patientColumnsWithDefault,
			patientColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			patientAllColumns,
			patientPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert patients, could not build update column list")
		}

		ret := strmangle.SetComplement(patientAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(patientPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert patients, could not build conflict column list")
			}

			conflict = make([]string, len(patientPrimaryKeyColumns))
			copy(conflict, patientPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"patients\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(patientType, patientMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(patientType, patientMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert patients")
	}

	if !cached {
		patientUpsertCacheMut.Lock()
		patientUpsertCache[key] = cache
		patientUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Patient record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Patient) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Patient provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), patientPrimaryKeyMapping)
	sql := "DELETE FROM \"patients\" WHERE \"patient_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from patients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for patients")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q patientQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no patientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from patients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for patients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PatientSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(patientBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), patientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"patients\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, patientPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from patient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for patients")
	}

	if len(patientAfterDeleteHooks) != 0 {
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
func (o *Patient) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPatient(ctx, exec, o.PatientID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PatientSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PatientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), patientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"patients\".* FROM \"patients\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, patientPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PatientSlice")
	}

	*o = slice

	return nil
}

// PatientExists checks if the Patient row exists.
func PatientExists(ctx context.Context, exec boil.ContextExecutor, patientID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"patients\" where \"patient_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, patientID)
	}
	row := exec.QueryRowContext(ctx, sql, patientID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if patients exists")
	}

	return exists, nil
}

// Exists checks if the Patient row exists.
func (o *Patient) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PatientExists(ctx, exec, o.PatientID)
}
