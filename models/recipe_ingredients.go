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

// RecipeIngredient is an object representing the database table.
type RecipeIngredient struct {
	ID           int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	RecipeID     null.Int `boil:"recipe_id" json:"recipe_id,omitempty" toml:"recipe_id" yaml:"recipe_id,omitempty"`
	IngredientID null.Int `boil:"ingredient_id" json:"ingredient_id,omitempty" toml:"ingredient_id" yaml:"ingredient_id,omitempty"`

	R *recipeIngredientR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L recipeIngredientL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RecipeIngredientColumns = struct {
	ID           string
	RecipeID     string
	IngredientID string
}{
	ID:           "id",
	RecipeID:     "recipe_id",
	IngredientID: "ingredient_id",
}

var RecipeIngredientTableColumns = struct {
	ID           string
	RecipeID     string
	IngredientID string
}{
	ID:           "recipe_ingredients.id",
	RecipeID:     "recipe_ingredients.recipe_id",
	IngredientID: "recipe_ingredients.ingredient_id",
}

// Generated where

var RecipeIngredientWhere = struct {
	ID           whereHelperint
	RecipeID     whereHelpernull_Int
	IngredientID whereHelpernull_Int
}{
	ID:           whereHelperint{field: "`recipe_ingredients`.`id`"},
	RecipeID:     whereHelpernull_Int{field: "`recipe_ingredients`.`recipe_id`"},
	IngredientID: whereHelpernull_Int{field: "`recipe_ingredients`.`ingredient_id`"},
}

// RecipeIngredientRels is where relationship names are stored.
var RecipeIngredientRels = struct {
}{}

// recipeIngredientR is where relationships are stored.
type recipeIngredientR struct {
}

// NewStruct creates a new relationship struct
func (*recipeIngredientR) NewStruct() *recipeIngredientR {
	return &recipeIngredientR{}
}

// recipeIngredientL is where Load methods for each relationship are stored.
type recipeIngredientL struct{}

var (
	recipeIngredientAllColumns            = []string{"id", "recipe_id", "ingredient_id"}
	recipeIngredientColumnsWithoutDefault = []string{"recipe_id", "ingredient_id"}
	recipeIngredientColumnsWithDefault    = []string{"id"}
	recipeIngredientPrimaryKeyColumns     = []string{"id"}
	recipeIngredientGeneratedColumns      = []string{}
)

type (
	// RecipeIngredientSlice is an alias for a slice of pointers to RecipeIngredient.
	// This should almost always be used instead of []RecipeIngredient.
	RecipeIngredientSlice []*RecipeIngredient
	// RecipeIngredientHook is the signature for custom RecipeIngredient hook methods
	RecipeIngredientHook func(context.Context, boil.ContextExecutor, *RecipeIngredient) error

	recipeIngredientQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	recipeIngredientType                 = reflect.TypeOf(&RecipeIngredient{})
	recipeIngredientMapping              = queries.MakeStructMapping(recipeIngredientType)
	recipeIngredientPrimaryKeyMapping, _ = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, recipeIngredientPrimaryKeyColumns)
	recipeIngredientInsertCacheMut       sync.RWMutex
	recipeIngredientInsertCache          = make(map[string]insertCache)
	recipeIngredientUpdateCacheMut       sync.RWMutex
	recipeIngredientUpdateCache          = make(map[string]updateCache)
	recipeIngredientUpsertCacheMut       sync.RWMutex
	recipeIngredientUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var recipeIngredientAfterSelectMu sync.Mutex
var recipeIngredientAfterSelectHooks []RecipeIngredientHook

var recipeIngredientBeforeInsertMu sync.Mutex
var recipeIngredientBeforeInsertHooks []RecipeIngredientHook
var recipeIngredientAfterInsertMu sync.Mutex
var recipeIngredientAfterInsertHooks []RecipeIngredientHook

var recipeIngredientBeforeUpdateMu sync.Mutex
var recipeIngredientBeforeUpdateHooks []RecipeIngredientHook
var recipeIngredientAfterUpdateMu sync.Mutex
var recipeIngredientAfterUpdateHooks []RecipeIngredientHook

var recipeIngredientBeforeDeleteMu sync.Mutex
var recipeIngredientBeforeDeleteHooks []RecipeIngredientHook
var recipeIngredientAfterDeleteMu sync.Mutex
var recipeIngredientAfterDeleteHooks []RecipeIngredientHook

var recipeIngredientBeforeUpsertMu sync.Mutex
var recipeIngredientBeforeUpsertHooks []RecipeIngredientHook
var recipeIngredientAfterUpsertMu sync.Mutex
var recipeIngredientAfterUpsertHooks []RecipeIngredientHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *RecipeIngredient) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *RecipeIngredient) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *RecipeIngredient) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *RecipeIngredient) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *RecipeIngredient) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *RecipeIngredient) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *RecipeIngredient) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *RecipeIngredient) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *RecipeIngredient) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range recipeIngredientAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRecipeIngredientHook registers your hook function for all future operations.
func AddRecipeIngredientHook(hookPoint boil.HookPoint, recipeIngredientHook RecipeIngredientHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		recipeIngredientAfterSelectMu.Lock()
		recipeIngredientAfterSelectHooks = append(recipeIngredientAfterSelectHooks, recipeIngredientHook)
		recipeIngredientAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		recipeIngredientBeforeInsertMu.Lock()
		recipeIngredientBeforeInsertHooks = append(recipeIngredientBeforeInsertHooks, recipeIngredientHook)
		recipeIngredientBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		recipeIngredientAfterInsertMu.Lock()
		recipeIngredientAfterInsertHooks = append(recipeIngredientAfterInsertHooks, recipeIngredientHook)
		recipeIngredientAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		recipeIngredientBeforeUpdateMu.Lock()
		recipeIngredientBeforeUpdateHooks = append(recipeIngredientBeforeUpdateHooks, recipeIngredientHook)
		recipeIngredientBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		recipeIngredientAfterUpdateMu.Lock()
		recipeIngredientAfterUpdateHooks = append(recipeIngredientAfterUpdateHooks, recipeIngredientHook)
		recipeIngredientAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		recipeIngredientBeforeDeleteMu.Lock()
		recipeIngredientBeforeDeleteHooks = append(recipeIngredientBeforeDeleteHooks, recipeIngredientHook)
		recipeIngredientBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		recipeIngredientAfterDeleteMu.Lock()
		recipeIngredientAfterDeleteHooks = append(recipeIngredientAfterDeleteHooks, recipeIngredientHook)
		recipeIngredientAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		recipeIngredientBeforeUpsertMu.Lock()
		recipeIngredientBeforeUpsertHooks = append(recipeIngredientBeforeUpsertHooks, recipeIngredientHook)
		recipeIngredientBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		recipeIngredientAfterUpsertMu.Lock()
		recipeIngredientAfterUpsertHooks = append(recipeIngredientAfterUpsertHooks, recipeIngredientHook)
		recipeIngredientAfterUpsertMu.Unlock()
	}
}

// One returns a single recipeIngredient record from the query.
func (q recipeIngredientQuery) One(ctx context.Context, exec boil.ContextExecutor) (*RecipeIngredient, error) {
	o := &RecipeIngredient{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for recipe_ingredients")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all RecipeIngredient records from the query.
func (q recipeIngredientQuery) All(ctx context.Context, exec boil.ContextExecutor) (RecipeIngredientSlice, error) {
	var o []*RecipeIngredient

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to RecipeIngredient slice")
	}

	if len(recipeIngredientAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all RecipeIngredient records in the query.
func (q recipeIngredientQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count recipe_ingredients rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q recipeIngredientQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if recipe_ingredients exists")
	}

	return count > 0, nil
}

// RecipeIngredients retrieves all the records using an executor.
func RecipeIngredients(mods ...qm.QueryMod) recipeIngredientQuery {
	mods = append(mods, qm.From("`recipe_ingredients`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`recipe_ingredients`.*"})
	}

	return recipeIngredientQuery{q}
}

// FindRecipeIngredient retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRecipeIngredient(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*RecipeIngredient, error) {
	recipeIngredientObj := &RecipeIngredient{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `recipe_ingredients` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, recipeIngredientObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from recipe_ingredients")
	}

	if err = recipeIngredientObj.doAfterSelectHooks(ctx, exec); err != nil {
		return recipeIngredientObj, err
	}

	return recipeIngredientObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *RecipeIngredient) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_ingredients provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recipeIngredientColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	recipeIngredientInsertCacheMut.RLock()
	cache, cached := recipeIngredientInsertCache[key]
	recipeIngredientInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			recipeIngredientAllColumns,
			recipeIngredientColumnsWithDefault,
			recipeIngredientColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `recipe_ingredients` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `recipe_ingredients` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `recipe_ingredients` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, recipeIngredientPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into recipe_ingredients")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == recipeIngredientMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for recipe_ingredients")
	}

CacheNoHooks:
	if !cached {
		recipeIngredientInsertCacheMut.Lock()
		recipeIngredientInsertCache[key] = cache
		recipeIngredientInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the RecipeIngredient.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *RecipeIngredient) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	recipeIngredientUpdateCacheMut.RLock()
	cache, cached := recipeIngredientUpdateCache[key]
	recipeIngredientUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			recipeIngredientAllColumns,
			recipeIngredientPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update recipe_ingredients, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `recipe_ingredients` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, recipeIngredientPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, append(wl, recipeIngredientPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update recipe_ingredients row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for recipe_ingredients")
	}

	if !cached {
		recipeIngredientUpdateCacheMut.Lock()
		recipeIngredientUpdateCache[key] = cache
		recipeIngredientUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q recipeIngredientQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for recipe_ingredients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for recipe_ingredients")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RecipeIngredientSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeIngredientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `recipe_ingredients` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, recipeIngredientPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in recipeIngredient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all recipeIngredient")
	}
	return rowsAff, nil
}

var mySQLRecipeIngredientUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *RecipeIngredient) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no recipe_ingredients provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(recipeIngredientColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLRecipeIngredientUniqueColumns, o)

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

	recipeIngredientUpsertCacheMut.RLock()
	cache, cached := recipeIngredientUpsertCache[key]
	recipeIngredientUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			recipeIngredientAllColumns,
			recipeIngredientColumnsWithDefault,
			recipeIngredientColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			recipeIngredientAllColumns,
			recipeIngredientPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert recipe_ingredients, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`recipe_ingredients`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `recipe_ingredients` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for recipe_ingredients")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == recipeIngredientMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(recipeIngredientType, recipeIngredientMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for recipe_ingredients")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for recipe_ingredients")
	}

CacheNoHooks:
	if !cached {
		recipeIngredientUpsertCacheMut.Lock()
		recipeIngredientUpsertCache[key] = cache
		recipeIngredientUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single RecipeIngredient record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *RecipeIngredient) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no RecipeIngredient provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), recipeIngredientPrimaryKeyMapping)
	sql := "DELETE FROM `recipe_ingredients` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from recipe_ingredients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for recipe_ingredients")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q recipeIngredientQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no recipeIngredientQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipe_ingredients")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_ingredients")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RecipeIngredientSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(recipeIngredientBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeIngredientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `recipe_ingredients` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, recipeIngredientPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from recipeIngredient slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for recipe_ingredients")
	}

	if len(recipeIngredientAfterDeleteHooks) != 0 {
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
func (o *RecipeIngredient) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRecipeIngredient(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RecipeIngredientSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RecipeIngredientSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), recipeIngredientPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `recipe_ingredients`.* FROM `recipe_ingredients` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, recipeIngredientPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RecipeIngredientSlice")
	}

	*o = slice

	return nil
}

// RecipeIngredientExists checks if the RecipeIngredient row exists.
func RecipeIngredientExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `recipe_ingredients` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if recipe_ingredients exists")
	}

	return exists, nil
}

// Exists checks if the RecipeIngredient row exists.
func (o *RecipeIngredient) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return RecipeIngredientExists(ctx, exec, o.ID)
}
