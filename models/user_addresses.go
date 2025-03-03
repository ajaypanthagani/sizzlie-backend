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

// UserAddress is an object representing the database table.
type UserAddress struct {
	ID        int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    null.Int `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	AddressID null.Int `boil:"address_id" json:"address_id,omitempty" toml:"address_id" yaml:"address_id,omitempty"`

	R *userAddressR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L userAddressL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var UserAddressColumns = struct {
	ID        string
	UserID    string
	AddressID string
}{
	ID:        "id",
	UserID:    "user_id",
	AddressID: "address_id",
}

var UserAddressTableColumns = struct {
	ID        string
	UserID    string
	AddressID string
}{
	ID:        "user_addresses.id",
	UserID:    "user_addresses.user_id",
	AddressID: "user_addresses.address_id",
}

// Generated where

var UserAddressWhere = struct {
	ID        whereHelperint
	UserID    whereHelpernull_Int
	AddressID whereHelpernull_Int
}{
	ID:        whereHelperint{field: "`user_addresses`.`id`"},
	UserID:    whereHelpernull_Int{field: "`user_addresses`.`user_id`"},
	AddressID: whereHelpernull_Int{field: "`user_addresses`.`address_id`"},
}

// UserAddressRels is where relationship names are stored.
var UserAddressRels = struct {
}{}

// userAddressR is where relationships are stored.
type userAddressR struct {
}

// NewStruct creates a new relationship struct
func (*userAddressR) NewStruct() *userAddressR {
	return &userAddressR{}
}

// userAddressL is where Load methods for each relationship are stored.
type userAddressL struct{}

var (
	userAddressAllColumns            = []string{"id", "user_id", "address_id"}
	userAddressColumnsWithoutDefault = []string{"user_id", "address_id"}
	userAddressColumnsWithDefault    = []string{"id"}
	userAddressPrimaryKeyColumns     = []string{"id"}
	userAddressGeneratedColumns      = []string{}
)

type (
	// UserAddressSlice is an alias for a slice of pointers to UserAddress.
	// This should almost always be used instead of []UserAddress.
	UserAddressSlice []*UserAddress
	// UserAddressHook is the signature for custom UserAddress hook methods
	UserAddressHook func(context.Context, boil.ContextExecutor, *UserAddress) error

	userAddressQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	userAddressType                 = reflect.TypeOf(&UserAddress{})
	userAddressMapping              = queries.MakeStructMapping(userAddressType)
	userAddressPrimaryKeyMapping, _ = queries.BindMapping(userAddressType, userAddressMapping, userAddressPrimaryKeyColumns)
	userAddressInsertCacheMut       sync.RWMutex
	userAddressInsertCache          = make(map[string]insertCache)
	userAddressUpdateCacheMut       sync.RWMutex
	userAddressUpdateCache          = make(map[string]updateCache)
	userAddressUpsertCacheMut       sync.RWMutex
	userAddressUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var userAddressAfterSelectMu sync.Mutex
var userAddressAfterSelectHooks []UserAddressHook

var userAddressBeforeInsertMu sync.Mutex
var userAddressBeforeInsertHooks []UserAddressHook
var userAddressAfterInsertMu sync.Mutex
var userAddressAfterInsertHooks []UserAddressHook

var userAddressBeforeUpdateMu sync.Mutex
var userAddressBeforeUpdateHooks []UserAddressHook
var userAddressAfterUpdateMu sync.Mutex
var userAddressAfterUpdateHooks []UserAddressHook

var userAddressBeforeDeleteMu sync.Mutex
var userAddressBeforeDeleteHooks []UserAddressHook
var userAddressAfterDeleteMu sync.Mutex
var userAddressAfterDeleteHooks []UserAddressHook

var userAddressBeforeUpsertMu sync.Mutex
var userAddressBeforeUpsertHooks []UserAddressHook
var userAddressAfterUpsertMu sync.Mutex
var userAddressAfterUpsertHooks []UserAddressHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *UserAddress) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *UserAddress) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *UserAddress) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *UserAddress) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *UserAddress) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *UserAddress) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *UserAddress) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *UserAddress) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *UserAddress) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range userAddressAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddUserAddressHook registers your hook function for all future operations.
func AddUserAddressHook(hookPoint boil.HookPoint, userAddressHook UserAddressHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		userAddressAfterSelectMu.Lock()
		userAddressAfterSelectHooks = append(userAddressAfterSelectHooks, userAddressHook)
		userAddressAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		userAddressBeforeInsertMu.Lock()
		userAddressBeforeInsertHooks = append(userAddressBeforeInsertHooks, userAddressHook)
		userAddressBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		userAddressAfterInsertMu.Lock()
		userAddressAfterInsertHooks = append(userAddressAfterInsertHooks, userAddressHook)
		userAddressAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		userAddressBeforeUpdateMu.Lock()
		userAddressBeforeUpdateHooks = append(userAddressBeforeUpdateHooks, userAddressHook)
		userAddressBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		userAddressAfterUpdateMu.Lock()
		userAddressAfterUpdateHooks = append(userAddressAfterUpdateHooks, userAddressHook)
		userAddressAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		userAddressBeforeDeleteMu.Lock()
		userAddressBeforeDeleteHooks = append(userAddressBeforeDeleteHooks, userAddressHook)
		userAddressBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		userAddressAfterDeleteMu.Lock()
		userAddressAfterDeleteHooks = append(userAddressAfterDeleteHooks, userAddressHook)
		userAddressAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		userAddressBeforeUpsertMu.Lock()
		userAddressBeforeUpsertHooks = append(userAddressBeforeUpsertHooks, userAddressHook)
		userAddressBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		userAddressAfterUpsertMu.Lock()
		userAddressAfterUpsertHooks = append(userAddressAfterUpsertHooks, userAddressHook)
		userAddressAfterUpsertMu.Unlock()
	}
}

// One returns a single userAddress record from the query.
func (q userAddressQuery) One(ctx context.Context, exec boil.ContextExecutor) (*UserAddress, error) {
	o := &UserAddress{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for user_addresses")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all UserAddress records from the query.
func (q userAddressQuery) All(ctx context.Context, exec boil.ContextExecutor) (UserAddressSlice, error) {
	var o []*UserAddress

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to UserAddress slice")
	}

	if len(userAddressAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all UserAddress records in the query.
func (q userAddressQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count user_addresses rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q userAddressQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if user_addresses exists")
	}

	return count > 0, nil
}

// UserAddresses retrieves all the records using an executor.
func UserAddresses(mods ...qm.QueryMod) userAddressQuery {
	mods = append(mods, qm.From("`user_addresses`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`user_addresses`.*"})
	}

	return userAddressQuery{q}
}

// FindUserAddress retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindUserAddress(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*UserAddress, error) {
	userAddressObj := &UserAddress{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `user_addresses` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, userAddressObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from user_addresses")
	}

	if err = userAddressObj.doAfterSelectHooks(ctx, exec); err != nil {
		return userAddressObj, err
	}

	return userAddressObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *UserAddress) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_addresses provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userAddressColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	userAddressInsertCacheMut.RLock()
	cache, cached := userAddressInsertCache[key]
	userAddressInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			userAddressAllColumns,
			userAddressColumnsWithDefault,
			userAddressColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(userAddressType, userAddressMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(userAddressType, userAddressMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `user_addresses` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `user_addresses` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `user_addresses` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, userAddressPrimaryKeyColumns))
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into user_addresses")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userAddressMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for user_addresses")
	}

CacheNoHooks:
	if !cached {
		userAddressInsertCacheMut.Lock()
		userAddressInsertCache[key] = cache
		userAddressInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the UserAddress.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *UserAddress) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	userAddressUpdateCacheMut.RLock()
	cache, cached := userAddressUpdateCache[key]
	userAddressUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			userAddressAllColumns,
			userAddressPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update user_addresses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `user_addresses` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, userAddressPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(userAddressType, userAddressMapping, append(wl, userAddressPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update user_addresses row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for user_addresses")
	}

	if !cached {
		userAddressUpdateCacheMut.Lock()
		userAddressUpdateCache[key] = cache
		userAddressUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q userAddressQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for user_addresses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for user_addresses")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o UserAddressSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `user_addresses` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userAddressPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in userAddress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all userAddress")
	}
	return rowsAff, nil
}

var mySQLUserAddressUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *UserAddress) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no user_addresses provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(userAddressColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLUserAddressUniqueColumns, o)

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

	userAddressUpsertCacheMut.RLock()
	cache, cached := userAddressUpsertCache[key]
	userAddressUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			userAddressAllColumns,
			userAddressColumnsWithDefault,
			userAddressColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			userAddressAllColumns,
			userAddressPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert user_addresses, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`user_addresses`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `user_addresses` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(userAddressType, userAddressMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(userAddressType, userAddressMapping, ret)
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
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for user_addresses")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == userAddressMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(userAddressType, userAddressMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for user_addresses")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for user_addresses")
	}

CacheNoHooks:
	if !cached {
		userAddressUpsertCacheMut.Lock()
		userAddressUpsertCache[key] = cache
		userAddressUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single UserAddress record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *UserAddress) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no UserAddress provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), userAddressPrimaryKeyMapping)
	sql := "DELETE FROM `user_addresses` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from user_addresses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for user_addresses")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q userAddressQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no userAddressQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from user_addresses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_addresses")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o UserAddressSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(userAddressBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `user_addresses` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userAddressPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from userAddress slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for user_addresses")
	}

	if len(userAddressAfterDeleteHooks) != 0 {
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
func (o *UserAddress) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindUserAddress(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *UserAddressSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := UserAddressSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), userAddressPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `user_addresses`.* FROM `user_addresses` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, userAddressPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in UserAddressSlice")
	}

	*o = slice

	return nil
}

// UserAddressExists checks if the UserAddress row exists.
func UserAddressExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `user_addresses` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if user_addresses exists")
	}

	return exists, nil
}

// Exists checks if the UserAddress row exists.
func (o *UserAddress) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return UserAddressExists(ctx, exec, o.ID)
}
