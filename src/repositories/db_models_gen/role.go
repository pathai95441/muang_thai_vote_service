// Code generated by SQLBoiler 4.11.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db_models_gen

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

// Role is an object representing the database table.
type Role struct {
	ID          int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	RoleName    string      `boil:"role_name" json:"role_name" toml:"role_name" yaml:"role_name"`
	Description string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	CreatedAt   time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	CreatedBy   string      `boil:"created_by" json:"created_by" toml:"created_by" yaml:"created_by"`
	UpdatedAt   null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	UpdatedBy   null.String `boil:"updated_by" json:"updated_by,omitempty" toml:"updated_by" yaml:"updated_by,omitempty"`
	DeletedAt   null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	DeletedBy   null.String `boil:"deleted_by" json:"deleted_by,omitempty" toml:"deleted_by" yaml:"deleted_by,omitempty"`

	R *roleR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L roleL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RoleColumns = struct {
	ID          string
	RoleName    string
	Description string
	CreatedAt   string
	CreatedBy   string
	UpdatedAt   string
	UpdatedBy   string
	DeletedAt   string
	DeletedBy   string
}{
	ID:          "id",
	RoleName:    "role_name",
	Description: "description",
	CreatedAt:   "created_at",
	CreatedBy:   "created_by",
	UpdatedAt:   "updated_at",
	UpdatedBy:   "updated_by",
	DeletedAt:   "deleted_at",
	DeletedBy:   "deleted_by",
}

var RoleTableColumns = struct {
	ID          string
	RoleName    string
	Description string
	CreatedAt   string
	CreatedBy   string
	UpdatedAt   string
	UpdatedBy   string
	DeletedAt   string
	DeletedBy   string
}{
	ID:          "role.id",
	RoleName:    "role.role_name",
	Description: "role.description",
	CreatedAt:   "role.created_at",
	CreatedBy:   "role.created_by",
	UpdatedAt:   "role.updated_at",
	UpdatedBy:   "role.updated_by",
	DeletedAt:   "role.deleted_at",
	DeletedBy:   "role.deleted_by",
}

// Generated where

var RoleWhere = struct {
	ID          whereHelperint
	RoleName    whereHelperstring
	Description whereHelperstring
	CreatedAt   whereHelpertime_Time
	CreatedBy   whereHelperstring
	UpdatedAt   whereHelpernull_Time
	UpdatedBy   whereHelpernull_String
	DeletedAt   whereHelpernull_Time
	DeletedBy   whereHelpernull_String
}{
	ID:          whereHelperint{field: "`role`.`id`"},
	RoleName:    whereHelperstring{field: "`role`.`role_name`"},
	Description: whereHelperstring{field: "`role`.`description`"},
	CreatedAt:   whereHelpertime_Time{field: "`role`.`created_at`"},
	CreatedBy:   whereHelperstring{field: "`role`.`created_by`"},
	UpdatedAt:   whereHelpernull_Time{field: "`role`.`updated_at`"},
	UpdatedBy:   whereHelpernull_String{field: "`role`.`updated_by`"},
	DeletedAt:   whereHelpernull_Time{field: "`role`.`deleted_at`"},
	DeletedBy:   whereHelpernull_String{field: "`role`.`deleted_by`"},
}

// RoleRels is where relationship names are stored.
var RoleRels = struct {
	Users string
}{
	Users: "Users",
}

// roleR is where relationships are stored.
type roleR struct {
	Users UserSlice `boil:"Users" json:"Users" toml:"Users" yaml:"Users"`
}

// NewStruct creates a new relationship struct
func (*roleR) NewStruct() *roleR {
	return &roleR{}
}

func (r *roleR) GetUsers() UserSlice {
	if r == nil {
		return nil
	}
	return r.Users
}

// roleL is where Load methods for each relationship are stored.
type roleL struct{}

var (
	roleAllColumns            = []string{"id", "role_name", "description", "created_at", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	roleColumnsWithoutDefault = []string{"id", "role_name", "description", "created_by", "updated_at", "updated_by", "deleted_at", "deleted_by"}
	roleColumnsWithDefault    = []string{"created_at"}
	rolePrimaryKeyColumns     = []string{"id"}
	roleGeneratedColumns      = []string{}
)

type (
	// RoleSlice is an alias for a slice of pointers to Role.
	// This should almost always be used instead of []Role.
	RoleSlice []*Role
	// RoleHook is the signature for custom Role hook methods
	RoleHook func(context.Context, boil.ContextExecutor, *Role) error

	roleQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	roleType                 = reflect.TypeOf(&Role{})
	roleMapping              = queries.MakeStructMapping(roleType)
	rolePrimaryKeyMapping, _ = queries.BindMapping(roleType, roleMapping, rolePrimaryKeyColumns)
	roleInsertCacheMut       sync.RWMutex
	roleInsertCache          = make(map[string]insertCache)
	roleUpdateCacheMut       sync.RWMutex
	roleUpdateCache          = make(map[string]updateCache)
	roleUpsertCacheMut       sync.RWMutex
	roleUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var roleAfterSelectHooks []RoleHook

var roleBeforeInsertHooks []RoleHook
var roleAfterInsertHooks []RoleHook

var roleBeforeUpdateHooks []RoleHook
var roleAfterUpdateHooks []RoleHook

var roleBeforeDeleteHooks []RoleHook
var roleAfterDeleteHooks []RoleHook

var roleBeforeUpsertHooks []RoleHook
var roleAfterUpsertHooks []RoleHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Role) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Role) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Role) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Role) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Role) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Role) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Role) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Role) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Role) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range roleAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRoleHook registers your hook function for all future operations.
func AddRoleHook(hookPoint boil.HookPoint, roleHook RoleHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		roleAfterSelectHooks = append(roleAfterSelectHooks, roleHook)
	case boil.BeforeInsertHook:
		roleBeforeInsertHooks = append(roleBeforeInsertHooks, roleHook)
	case boil.AfterInsertHook:
		roleAfterInsertHooks = append(roleAfterInsertHooks, roleHook)
	case boil.BeforeUpdateHook:
		roleBeforeUpdateHooks = append(roleBeforeUpdateHooks, roleHook)
	case boil.AfterUpdateHook:
		roleAfterUpdateHooks = append(roleAfterUpdateHooks, roleHook)
	case boil.BeforeDeleteHook:
		roleBeforeDeleteHooks = append(roleBeforeDeleteHooks, roleHook)
	case boil.AfterDeleteHook:
		roleAfterDeleteHooks = append(roleAfterDeleteHooks, roleHook)
	case boil.BeforeUpsertHook:
		roleBeforeUpsertHooks = append(roleBeforeUpsertHooks, roleHook)
	case boil.AfterUpsertHook:
		roleAfterUpsertHooks = append(roleAfterUpsertHooks, roleHook)
	}
}

// One returns a single role record from the query.
func (q roleQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Role, error) {
	o := &Role{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db_models_gen: failed to execute a one query for role")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Role records from the query.
func (q roleQuery) All(ctx context.Context, exec boil.ContextExecutor) (RoleSlice, error) {
	var o []*Role

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db_models_gen: failed to assign all query results to Role slice")
	}

	if len(roleAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Role records in the query.
func (q roleQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: failed to count role rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q roleQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db_models_gen: failed to check if role exists")
	}

	return count > 0, nil
}

// Users retrieves all the user's Users with an executor.
func (o *Role) Users(mods ...qm.QueryMod) userQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`users`.`role_id`=?", o.ID),
	)

	return Users(queryMods...)
}

// LoadUsers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (roleL) LoadUsers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRole interface{}, mods queries.Applicator) error {
	var slice []*Role
	var object *Role

	if singular {
		object = maybeRole.(*Role)
	} else {
		slice = *maybeRole.(*[]*Role)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &roleR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &roleR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.role_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load users")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice users")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on users")
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
	if singular {
		object.R.Users = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &userR{}
			}
			foreign.R.Role = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RoleID {
				local.R.Users = append(local.R.Users, foreign)
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Role = local
				break
			}
		}
	}

	return nil
}

// AddUsers adds the given related objects to the existing relationships
// of the role, optionally inserting them as new records.
// Appends related to o.R.Users.
// Sets related.R.Role appropriately.
func (o *Role) AddUsers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*User) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RoleID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `users` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"role_id"}),
				strmangle.WhereClause("`", "`", 0, userPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RoleID = o.ID
		}
	}

	if o.R == nil {
		o.R = &roleR{
			Users: related,
		}
	} else {
		o.R.Users = append(o.R.Users, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &userR{
				Role: o,
			}
		} else {
			rel.R.Role = o
		}
	}
	return nil
}

// Roles retrieves all the records using an executor.
func Roles(mods ...qm.QueryMod) roleQuery {
	mods = append(mods, qm.From("`role`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`role`.*"})
	}

	return roleQuery{q}
}

// FindRole retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRole(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*Role, error) {
	roleObj := &Role{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `role` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, roleObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db_models_gen: unable to select from role")
	}

	if err = roleObj.doAfterSelectHooks(ctx, exec); err != nil {
		return roleObj, err
	}

	return roleObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Role) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db_models_gen: no role provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if queries.MustTime(o.UpdatedAt).IsZero() {
			queries.SetScanner(&o.UpdatedAt, currTime)
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	roleInsertCacheMut.RLock()
	cache, cached := roleInsertCache[key]
	roleInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			roleAllColumns,
			roleColumnsWithDefault,
			roleColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(roleType, roleMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `role` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `role` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `role` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, rolePrimaryKeyColumns))
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db_models_gen: unable to insert into role")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "db_models_gen: unable to populate default values for role")
	}

CacheNoHooks:
	if !cached {
		roleInsertCacheMut.Lock()
		roleInsertCache[key] = cache
		roleInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Role.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Role) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	roleUpdateCacheMut.RLock()
	cache, cached := roleUpdateCache[key]
	roleUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			roleAllColumns,
			rolePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db_models_gen: unable to update role, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `role` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, rolePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, append(wl, rolePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "db_models_gen: unable to update role row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: failed to get rows affected by update for role")
	}

	if !cached {
		roleUpdateCacheMut.Lock()
		roleUpdateCache[key] = cache
		roleUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q roleQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to update all for role")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to retrieve rows affected for role")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RoleSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db_models_gen: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `role` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, rolePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to update all in role slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to retrieve rows affected all in update all role")
	}
	return rowsAff, nil
}

var mySQLRoleUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Role) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db_models_gen: no role provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(roleColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRoleUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	roleUpsertCacheMut.RLock()
	cache, cached := roleUpsertCache[key]
	roleUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			roleAllColumns,
			roleColumnsWithDefault,
			roleColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			roleAllColumns,
			rolePrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("db_models_gen: unable to upsert role, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`role`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `role` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(roleType, roleMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(roleType, roleMapping, ret)
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
	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db_models_gen: unable to upsert for role")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(roleType, roleMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "db_models_gen: unable to retrieve unique values for role")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "db_models_gen: unable to populate default values for role")
	}

CacheNoHooks:
	if !cached {
		roleUpsertCacheMut.Lock()
		roleUpsertCache[key] = cache
		roleUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Role record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Role) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db_models_gen: no Role provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), rolePrimaryKeyMapping)
	sql := "DELETE FROM `role` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to delete from role")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: failed to get rows affected by delete for role")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q roleQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db_models_gen: no roleQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to delete all from role")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: failed to get rows affected by deleteall for role")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RoleSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(roleBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `role` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, rolePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: unable to delete all from role slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db_models_gen: failed to get rows affected by deleteall for role")
	}

	if len(roleAfterDeleteHooks) != 0 {
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
func (o *Role) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRole(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RoleSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RoleSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), rolePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `role`.* FROM `role` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, rolePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db_models_gen: unable to reload all in RoleSlice")
	}

	*o = slice

	return nil
}

// RoleExists checks if the Role row exists.
func RoleExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `role` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db_models_gen: unable to check if role exists")
	}

	return exists, nil
}
