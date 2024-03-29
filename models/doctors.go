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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Doctor is an object representing the database table.
type Doctor struct {
	DoctorID          int      `boil:"doctor_id" json:"doctor_id" toml:"doctor_id" yaml:"doctor_id"`
	Specialties       string   `boil:"specialties" json:"specialties" toml:"specialties" yaml:"specialties"`
	UserID            int      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	HealthinsuranceID null.Int `boil:"healthinsurance_id" json:"healthinsurance_id,omitempty" toml:"healthinsurance_id" yaml:"healthinsurance_id,omitempty"`

	R *doctorR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L doctorL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DoctorColumns = struct {
	DoctorID          string
	Specialties       string
	UserID            string
	HealthinsuranceID string
}{
	DoctorID:          "doctor_id",
	Specialties:       "specialties",
	UserID:            "user_id",
	HealthinsuranceID: "healthinsurance_id",
}

var DoctorTableColumns = struct {
	DoctorID          string
	Specialties       string
	UserID            string
	HealthinsuranceID string
}{
	DoctorID:          "doctors.doctor_id",
	Specialties:       "doctors.specialties",
	UserID:            "doctors.user_id",
	HealthinsuranceID: "doctors.healthinsurance_id",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod     { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod    { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) LIKE(x string) qm.QueryMod   { return qm.Where(w.field+" LIKE ?", x) }
func (w whereHelperstring) NLIKE(x string) qm.QueryMod  { return qm.Where(w.field+" NOT LIKE ?", x) }
func (w whereHelperstring) ILIKE(x string) qm.QueryMod  { return qm.Where(w.field+" ILIKE ?", x) }
func (w whereHelperstring) NILIKE(x string) qm.QueryMod { return qm.Where(w.field+" NOT ILIKE ?", x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var DoctorWhere = struct {
	DoctorID          whereHelperint
	Specialties       whereHelperstring
	UserID            whereHelperint
	HealthinsuranceID whereHelpernull_Int
}{
	DoctorID:          whereHelperint{field: "\"doctors\".\"doctor_id\""},
	Specialties:       whereHelperstring{field: "\"doctors\".\"specialties\""},
	UserID:            whereHelperint{field: "\"doctors\".\"user_id\""},
	HealthinsuranceID: whereHelpernull_Int{field: "\"doctors\".\"healthinsurance_id\""},
}

// DoctorRels is where relationship names are stored.
var DoctorRels = struct {
	User            string
	Healthinsurance string
}{
	User:            "User",
	Healthinsurance: "Healthinsurance",
}

// doctorR is where relationships are stored.
type doctorR struct {
	User            *User            `boil:"User" json:"User" toml:"User" yaml:"User"`
	Healthinsurance *Healthinsurance `boil:"Healthinsurance" json:"Healthinsurance" toml:"Healthinsurance" yaml:"Healthinsurance"`
}

// NewStruct creates a new relationship struct
func (*doctorR) NewStruct() *doctorR {
	return &doctorR{}
}

func (r *doctorR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *doctorR) GetHealthinsurance() *Healthinsurance {
	if r == nil {
		return nil
	}
	return r.Healthinsurance
}

// doctorL is where Load methods for each relationship are stored.
type doctorL struct{}

var (
	doctorAllColumns            = []string{"doctor_id", "specialties", "user_id", "healthinsurance_id"}
	doctorColumnsWithoutDefault = []string{"specialties"}
	doctorColumnsWithDefault    = []string{"doctor_id", "user_id", "healthinsurance_id"}
	doctorPrimaryKeyColumns     = []string{"doctor_id"}
	doctorGeneratedColumns      = []string{}
)

type (
	// DoctorSlice is an alias for a slice of pointers to Doctor.
	// This should almost always be used instead of []Doctor.
	DoctorSlice []*Doctor
	// DoctorHook is the signature for custom Doctor hook methods
	DoctorHook func(context.Context, boil.ContextExecutor, *Doctor) error

	doctorQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	doctorType                 = reflect.TypeOf(&Doctor{})
	doctorMapping              = queries.MakeStructMapping(doctorType)
	doctorPrimaryKeyMapping, _ = queries.BindMapping(doctorType, doctorMapping, doctorPrimaryKeyColumns)
	doctorInsertCacheMut       sync.RWMutex
	doctorInsertCache          = make(map[string]insertCache)
	doctorUpdateCacheMut       sync.RWMutex
	doctorUpdateCache          = make(map[string]updateCache)
	doctorUpsertCacheMut       sync.RWMutex
	doctorUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var doctorAfterSelectMu sync.Mutex
var doctorAfterSelectHooks []DoctorHook

var doctorBeforeInsertMu sync.Mutex
var doctorBeforeInsertHooks []DoctorHook
var doctorAfterInsertMu sync.Mutex
var doctorAfterInsertHooks []DoctorHook

var doctorBeforeUpdateMu sync.Mutex
var doctorBeforeUpdateHooks []DoctorHook
var doctorAfterUpdateMu sync.Mutex
var doctorAfterUpdateHooks []DoctorHook

var doctorBeforeDeleteMu sync.Mutex
var doctorBeforeDeleteHooks []DoctorHook
var doctorAfterDeleteMu sync.Mutex
var doctorAfterDeleteHooks []DoctorHook

var doctorBeforeUpsertMu sync.Mutex
var doctorBeforeUpsertHooks []DoctorHook
var doctorAfterUpsertMu sync.Mutex
var doctorAfterUpsertHooks []DoctorHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Doctor) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Doctor) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Doctor) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Doctor) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Doctor) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Doctor) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Doctor) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Doctor) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Doctor) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range doctorAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDoctorHook registers your hook function for all future operations.
func AddDoctorHook(hookPoint boil.HookPoint, doctorHook DoctorHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		doctorAfterSelectMu.Lock()
		doctorAfterSelectHooks = append(doctorAfterSelectHooks, doctorHook)
		doctorAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		doctorBeforeInsertMu.Lock()
		doctorBeforeInsertHooks = append(doctorBeforeInsertHooks, doctorHook)
		doctorBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		doctorAfterInsertMu.Lock()
		doctorAfterInsertHooks = append(doctorAfterInsertHooks, doctorHook)
		doctorAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		doctorBeforeUpdateMu.Lock()
		doctorBeforeUpdateHooks = append(doctorBeforeUpdateHooks, doctorHook)
		doctorBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		doctorAfterUpdateMu.Lock()
		doctorAfterUpdateHooks = append(doctorAfterUpdateHooks, doctorHook)
		doctorAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		doctorBeforeDeleteMu.Lock()
		doctorBeforeDeleteHooks = append(doctorBeforeDeleteHooks, doctorHook)
		doctorBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		doctorAfterDeleteMu.Lock()
		doctorAfterDeleteHooks = append(doctorAfterDeleteHooks, doctorHook)
		doctorAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		doctorBeforeUpsertMu.Lock()
		doctorBeforeUpsertHooks = append(doctorBeforeUpsertHooks, doctorHook)
		doctorBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		doctorAfterUpsertMu.Lock()
		doctorAfterUpsertHooks = append(doctorAfterUpsertHooks, doctorHook)
		doctorAfterUpsertMu.Unlock()
	}
}

// One returns a single doctor record from the query.
func (q doctorQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Doctor, error) {
	o := &Doctor{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for doctors")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Doctor records from the query.
func (q doctorQuery) All(ctx context.Context, exec boil.ContextExecutor) (DoctorSlice, error) {
	var o []*Doctor

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Doctor slice")
	}

	if len(doctorAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Doctor records in the query.
func (q doctorQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count doctors rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q doctorQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if doctors exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Doctor) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"user_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// Healthinsurance pointed to by the foreign key.
func (o *Doctor) Healthinsurance(mods ...qm.QueryMod) healthinsuranceQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"healthinsurance_id\" = ?", o.HealthinsuranceID),
	}

	queryMods = append(queryMods, mods...)

	return Healthinsurances(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (doctorL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDoctor interface{}, mods queries.Applicator) error {
	var slice []*Doctor
	var object *Doctor

	if singular {
		var ok bool
		object, ok = maybeDoctor.(*Doctor)
		if !ok {
			object = new(Doctor)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDoctor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDoctor))
			}
		}
	} else {
		s, ok := maybeDoctor.(*[]*Doctor)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDoctor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDoctor))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &doctorR{}
		}
		args[object.UserID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &doctorR{}
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
		foreign.R.Doctors = append(foreign.R.Doctors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.UserID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Doctors = append(foreign.R.Doctors, local)
				break
			}
		}
	}

	return nil
}

// LoadHealthinsurance allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (doctorL) LoadHealthinsurance(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDoctor interface{}, mods queries.Applicator) error {
	var slice []*Doctor
	var object *Doctor

	if singular {
		var ok bool
		object, ok = maybeDoctor.(*Doctor)
		if !ok {
			object = new(Doctor)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDoctor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDoctor))
			}
		}
	} else {
		s, ok := maybeDoctor.(*[]*Doctor)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDoctor)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDoctor))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &doctorR{}
		}
		if !queries.IsNil(object.HealthinsuranceID) {
			args[object.HealthinsuranceID] = struct{}{}
		}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &doctorR{}
			}

			if !queries.IsNil(obj.HealthinsuranceID) {
				args[obj.HealthinsuranceID] = struct{}{}
			}

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
		qm.From(`healthinsurance`),
		qm.WhereIn(`healthinsurance.healthinsurance_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Healthinsurance")
	}

	var resultSlice []*Healthinsurance
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Healthinsurance")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for healthinsurance")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for healthinsurance")
	}

	if len(healthinsuranceAfterSelectHooks) != 0 {
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
		object.R.Healthinsurance = foreign
		if foreign.R == nil {
			foreign.R = &healthinsuranceR{}
		}
		foreign.R.Doctors = append(foreign.R.Doctors, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.HealthinsuranceID, foreign.HealthinsuranceID) {
				local.R.Healthinsurance = foreign
				if foreign.R == nil {
					foreign.R = &healthinsuranceR{}
				}
				foreign.R.Doctors = append(foreign.R.Doctors, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the doctor to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Doctors.
func (o *Doctor) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"doctors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, doctorPrimaryKeyColumns),
	)
	values := []interface{}{related.UserID, o.DoctorID}

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
		o.R = &doctorR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Doctors: DoctorSlice{o},
		}
	} else {
		related.R.Doctors = append(related.R.Doctors, o)
	}

	return nil
}

// SetHealthinsurance of the doctor to the related item.
// Sets o.R.Healthinsurance to related.
// Adds o to related.R.Doctors.
func (o *Doctor) SetHealthinsurance(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Healthinsurance) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"doctors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"healthinsurance_id"}),
		strmangle.WhereClause("\"", "\"", 2, doctorPrimaryKeyColumns),
	)
	values := []interface{}{related.HealthinsuranceID, o.DoctorID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.HealthinsuranceID, related.HealthinsuranceID)
	if o.R == nil {
		o.R = &doctorR{
			Healthinsurance: related,
		}
	} else {
		o.R.Healthinsurance = related
	}

	if related.R == nil {
		related.R = &healthinsuranceR{
			Doctors: DoctorSlice{o},
		}
	} else {
		related.R.Doctors = append(related.R.Doctors, o)
	}

	return nil
}

// RemoveHealthinsurance relationship.
// Sets o.R.Healthinsurance to nil.
// Removes o from all passed in related items' relationships struct.
func (o *Doctor) RemoveHealthinsurance(ctx context.Context, exec boil.ContextExecutor, related *Healthinsurance) error {
	var err error

	queries.SetScanner(&o.HealthinsuranceID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("healthinsurance_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Healthinsurance = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Doctors {
		if queries.Equal(o.HealthinsuranceID, ri.HealthinsuranceID) {
			continue
		}

		ln := len(related.R.Doctors)
		if ln > 1 && i < ln-1 {
			related.R.Doctors[i] = related.R.Doctors[ln-1]
		}
		related.R.Doctors = related.R.Doctors[:ln-1]
		break
	}
	return nil
}

// Doctors retrieves all the records using an executor.
func Doctors(mods ...qm.QueryMod) doctorQuery {
	mods = append(mods, qm.From("\"doctors\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"doctors\".*"})
	}

	return doctorQuery{q}
}

// FindDoctor retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDoctor(ctx context.Context, exec boil.ContextExecutor, doctorID int, selectCols ...string) (*Doctor, error) {
	doctorObj := &Doctor{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"doctors\" where \"doctor_id\"=$1", sel,
	)

	q := queries.Raw(query, doctorID)

	err := q.Bind(ctx, exec, doctorObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from doctors")
	}

	if err = doctorObj.doAfterSelectHooks(ctx, exec); err != nil {
		return doctorObj, err
	}

	return doctorObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Doctor) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no doctors provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(doctorColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	doctorInsertCacheMut.RLock()
	cache, cached := doctorInsertCache[key]
	doctorInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			doctorAllColumns,
			doctorColumnsWithDefault,
			doctorColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(doctorType, doctorMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(doctorType, doctorMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"doctors\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"doctors\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into doctors")
	}

	if !cached {
		doctorInsertCacheMut.Lock()
		doctorInsertCache[key] = cache
		doctorInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Doctor.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Doctor) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	doctorUpdateCacheMut.RLock()
	cache, cached := doctorUpdateCache[key]
	doctorUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			doctorAllColumns,
			doctorPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update doctors, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"doctors\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, doctorPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(doctorType, doctorMapping, append(wl, doctorPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update doctors row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for doctors")
	}

	if !cached {
		doctorUpdateCacheMut.Lock()
		doctorUpdateCache[key] = cache
		doctorUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q doctorQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for doctors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for doctors")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DoctorSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"doctors\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, doctorPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in doctor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all doctor")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Doctor) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no doctors provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(doctorColumnsWithDefault, o)

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

	doctorUpsertCacheMut.RLock()
	cache, cached := doctorUpsertCache[key]
	doctorUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			doctorAllColumns,
			doctorColumnsWithDefault,
			doctorColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			doctorAllColumns,
			doctorPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert doctors, could not build update column list")
		}

		ret := strmangle.SetComplement(doctorAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(doctorPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert doctors, could not build conflict column list")
			}

			conflict = make([]string, len(doctorPrimaryKeyColumns))
			copy(conflict, doctorPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"doctors\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(doctorType, doctorMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(doctorType, doctorMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert doctors")
	}

	if !cached {
		doctorUpsertCacheMut.Lock()
		doctorUpsertCache[key] = cache
		doctorUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Doctor record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Doctor) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Doctor provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), doctorPrimaryKeyMapping)
	sql := "DELETE FROM \"doctors\" WHERE \"doctor_id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from doctors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for doctors")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q doctorQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no doctorQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from doctors")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for doctors")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DoctorSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(doctorBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"doctors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, doctorPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from doctor slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for doctors")
	}

	if len(doctorAfterDeleteHooks) != 0 {
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
func (o *Doctor) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDoctor(ctx, exec, o.DoctorID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DoctorSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DoctorSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), doctorPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"doctors\".* FROM \"doctors\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, doctorPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DoctorSlice")
	}

	*o = slice

	return nil
}

// DoctorExists checks if the Doctor row exists.
func DoctorExists(ctx context.Context, exec boil.ContextExecutor, doctorID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"doctors\" where \"doctor_id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, doctorID)
	}
	row := exec.QueryRowContext(ctx, sql, doctorID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if doctors exists")
	}

	return exists, nil
}

// Exists checks if the Doctor row exists.
func (o *Doctor) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DoctorExists(ctx, exec, o.DoctorID)
}
