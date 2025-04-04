// Code generated by SQLBoiler 4.18.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// EconomyCreateitem is an object representing the database table.
type EconomyCreateitem struct {
	GuildID     string      `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	UserID      string      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Name        null.String `boil:"name" json:"name,omitempty" toml:"name" yaml:"name,omitempty"`
	Description null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Price       null.Int64  `boil:"price" json:"price,omitempty" toml:"price" yaml:"price,omitempty"`
	Quantity    null.Int64  `boil:"quantity" json:"quantity,omitempty" toml:"quantity" yaml:"quantity,omitempty"`
	Role        null.String `boil:"role" json:"role,omitempty" toml:"role" yaml:"role,omitempty"`
	Reply       null.String `boil:"reply" json:"reply,omitempty" toml:"reply" yaml:"reply,omitempty"`
	MSGID       string      `boil:"msg_id" json:"msg_id" toml:"msg_id" yaml:"msg_id"`

	R *economyCreateitemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L economyCreateitemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EconomyCreateitemColumns = struct {
	GuildID     string
	UserID      string
	Name        string
	Description string
	Price       string
	Quantity    string
	Role        string
	Reply       string
	MSGID       string
}{
	GuildID:     "guild_id",
	UserID:      "user_id",
	Name:        "name",
	Description: "description",
	Price:       "price",
	Quantity:    "quantity",
	Role:        "role",
	Reply:       "reply",
	MSGID:       "msg_id",
}

var EconomyCreateitemTableColumns = struct {
	GuildID     string
	UserID      string
	Name        string
	Description string
	Price       string
	Quantity    string
	Role        string
	Reply       string
	MSGID       string
}{
	GuildID:     "economy_createitem.guild_id",
	UserID:      "economy_createitem.user_id",
	Name:        "economy_createitem.name",
	Description: "economy_createitem.description",
	Price:       "economy_createitem.price",
	Quantity:    "economy_createitem.quantity",
	Role:        "economy_createitem.role",
	Reply:       "economy_createitem.reply",
	MSGID:       "economy_createitem.msg_id",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) ILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" ILIKE ?", x)
}
func (w whereHelpernull_String) NILIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT ILIKE ?", x)
}
func (w whereHelpernull_String) SIMILAR(x null.String) qm.QueryMod {
	return qm.Where(w.field+" SIMILAR TO ?", x)
}
func (w whereHelpernull_String) NSIMILAR(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT SIMILAR TO ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var EconomyCreateitemWhere = struct {
	GuildID     whereHelperstring
	UserID      whereHelperstring
	Name        whereHelpernull_String
	Description whereHelpernull_String
	Price       whereHelpernull_Int64
	Quantity    whereHelpernull_Int64
	Role        whereHelpernull_String
	Reply       whereHelpernull_String
	MSGID       whereHelperstring
}{
	GuildID:     whereHelperstring{field: "\"economy_createitem\".\"guild_id\""},
	UserID:      whereHelperstring{field: "\"economy_createitem\".\"user_id\""},
	Name:        whereHelpernull_String{field: "\"economy_createitem\".\"name\""},
	Description: whereHelpernull_String{field: "\"economy_createitem\".\"description\""},
	Price:       whereHelpernull_Int64{field: "\"economy_createitem\".\"price\""},
	Quantity:    whereHelpernull_Int64{field: "\"economy_createitem\".\"quantity\""},
	Role:        whereHelpernull_String{field: "\"economy_createitem\".\"role\""},
	Reply:       whereHelpernull_String{field: "\"economy_createitem\".\"reply\""},
	MSGID:       whereHelperstring{field: "\"economy_createitem\".\"msg_id\""},
}

// EconomyCreateitemRels is where relationship names are stored.
var EconomyCreateitemRels = struct {
	Guild string
}{
	Guild: "Guild",
}

// economyCreateitemR is where relationships are stored.
type economyCreateitemR struct {
	Guild *EconomyConfig `boil:"Guild" json:"Guild" toml:"Guild" yaml:"Guild"`
}

// NewStruct creates a new relationship struct
func (*economyCreateitemR) NewStruct() *economyCreateitemR {
	return &economyCreateitemR{}
}

func (r *economyCreateitemR) GetGuild() *EconomyConfig {
	if r == nil {
		return nil
	}
	return r.Guild
}

// economyCreateitemL is where Load methods for each relationship are stored.
type economyCreateitemL struct{}

var (
	economyCreateitemAllColumns            = []string{"guild_id", "user_id", "name", "description", "price", "quantity", "role", "reply", "msg_id"}
	economyCreateitemColumnsWithoutDefault = []string{"guild_id", "user_id", "msg_id"}
	economyCreateitemColumnsWithDefault    = []string{"name", "description", "price", "quantity", "role", "reply"}
	economyCreateitemPrimaryKeyColumns     = []string{"guild_id", "user_id"}
	economyCreateitemGeneratedColumns      = []string{}
)

type (
	// EconomyCreateitemSlice is an alias for a slice of pointers to EconomyCreateitem.
	// This should almost always be used instead of []EconomyCreateitem.
	EconomyCreateitemSlice []*EconomyCreateitem

	economyCreateitemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	economyCreateitemType                 = reflect.TypeOf(&EconomyCreateitem{})
	economyCreateitemMapping              = queries.MakeStructMapping(economyCreateitemType)
	economyCreateitemPrimaryKeyMapping, _ = queries.BindMapping(economyCreateitemType, economyCreateitemMapping, economyCreateitemPrimaryKeyColumns)
	economyCreateitemInsertCacheMut       sync.RWMutex
	economyCreateitemInsertCache          = make(map[string]insertCache)
	economyCreateitemUpdateCacheMut       sync.RWMutex
	economyCreateitemUpdateCache          = make(map[string]updateCache)
	economyCreateitemUpsertCacheMut       sync.RWMutex
	economyCreateitemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single economyCreateitem record from the query using the global executor.
func (q economyCreateitemQuery) OneG(ctx context.Context) (*EconomyCreateitem, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single economyCreateitem record from the query.
func (q economyCreateitemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EconomyCreateitem, error) {
	o := &EconomyCreateitem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for economy_createitem")
	}

	return o, nil
}

// AllG returns all EconomyCreateitem records from the query using the global executor.
func (q economyCreateitemQuery) AllG(ctx context.Context) (EconomyCreateitemSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all EconomyCreateitem records from the query.
func (q economyCreateitemQuery) All(ctx context.Context, exec boil.ContextExecutor) (EconomyCreateitemSlice, error) {
	var o []*EconomyCreateitem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EconomyCreateitem slice")
	}

	return o, nil
}

// CountG returns the count of all EconomyCreateitem records in the query using the global executor
func (q economyCreateitemQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all EconomyCreateitem records in the query.
func (q economyCreateitemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count economy_createitem rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q economyCreateitemQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q economyCreateitemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if economy_createitem exists")
	}

	return count > 0, nil
}

// Guild pointed to by the foreign key.
func (o *EconomyCreateitem) Guild(mods ...qm.QueryMod) economyConfigQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"guild_id\" = ?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	return EconomyConfigs(queryMods...)
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (economyCreateitemL) LoadGuild(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEconomyCreateitem interface{}, mods queries.Applicator) error {
	var slice []*EconomyCreateitem
	var object *EconomyCreateitem

	if singular {
		var ok bool
		object, ok = maybeEconomyCreateitem.(*EconomyCreateitem)
		if !ok {
			object = new(EconomyCreateitem)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEconomyCreateitem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEconomyCreateitem))
			}
		}
	} else {
		s, ok := maybeEconomyCreateitem.(*[]*EconomyCreateitem)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEconomyCreateitem)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEconomyCreateitem))
			}
		}
	}

	args := make(map[interface{}]struct{})
	if singular {
		if object.R == nil {
			object.R = &economyCreateitemR{}
		}
		args[object.GuildID] = struct{}{}

	} else {
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &economyCreateitemR{}
			}

			args[obj.GuildID] = struct{}{}

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
		qm.From(`economy_config`),
		qm.WhereIn(`economy_config.guild_id in ?`, argsSlice...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load EconomyConfig")
	}

	var resultSlice []*EconomyConfig
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice EconomyConfig")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for economy_config")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for economy_config")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Guild = foreign
		if foreign.R == nil {
			foreign.R = &economyConfigR{}
		}
		foreign.R.GuildEconomyCreateitems = append(foreign.R.GuildEconomyCreateitems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GuildID == foreign.GuildID {
				local.R.Guild = foreign
				if foreign.R == nil {
					foreign.R = &economyConfigR{}
				}
				foreign.R.GuildEconomyCreateitems = append(foreign.R.GuildEconomyCreateitems, local)
				break
			}
		}
	}

	return nil
}

// SetGuildG of the economyCreateitem to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildEconomyCreateitems.
// Uses the global database handle.
func (o *EconomyCreateitem) SetGuildG(ctx context.Context, insert bool, related *EconomyConfig) error {
	return o.SetGuild(ctx, boil.GetContextDB(), insert, related)
}

// SetGuild of the economyCreateitem to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.GuildEconomyCreateitems.
func (o *EconomyCreateitem) SetGuild(ctx context.Context, exec boil.ContextExecutor, insert bool, related *EconomyConfig) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"economy_createitem\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, economyCreateitemPrimaryKeyColumns),
	)
	values := []interface{}{related.GuildID, o.GuildID, o.UserID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.GuildID
	if o.R == nil {
		o.R = &economyCreateitemR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &economyConfigR{
			GuildEconomyCreateitems: EconomyCreateitemSlice{o},
		}
	} else {
		related.R.GuildEconomyCreateitems = append(related.R.GuildEconomyCreateitems, o)
	}

	return nil
}

// EconomyCreateitems retrieves all the records using an executor.
func EconomyCreateitems(mods ...qm.QueryMod) economyCreateitemQuery {
	mods = append(mods, qm.From("\"economy_createitem\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"economy_createitem\".*"})
	}

	return economyCreateitemQuery{q}
}

// FindEconomyCreateitemG retrieves a single record by ID.
func FindEconomyCreateitemG(ctx context.Context, guildID string, userID string, selectCols ...string) (*EconomyCreateitem, error) {
	return FindEconomyCreateitem(ctx, boil.GetContextDB(), guildID, userID, selectCols...)
}

// FindEconomyCreateitem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEconomyCreateitem(ctx context.Context, exec boil.ContextExecutor, guildID string, userID string, selectCols ...string) (*EconomyCreateitem, error) {
	economyCreateitemObj := &EconomyCreateitem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"economy_createitem\" where \"guild_id\"=$1 AND \"user_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, userID)

	err := q.Bind(ctx, exec, economyCreateitemObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from economy_createitem")
	}

	return economyCreateitemObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EconomyCreateitem) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EconomyCreateitem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_createitem provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(economyCreateitemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	economyCreateitemInsertCacheMut.RLock()
	cache, cached := economyCreateitemInsertCache[key]
	economyCreateitemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			economyCreateitemAllColumns,
			economyCreateitemColumnsWithDefault,
			economyCreateitemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(economyCreateitemType, economyCreateitemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(economyCreateitemType, economyCreateitemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"economy_createitem\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"economy_createitem\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into economy_createitem")
	}

	if !cached {
		economyCreateitemInsertCacheMut.Lock()
		economyCreateitemInsertCache[key] = cache
		economyCreateitemInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single EconomyCreateitem record using the global executor.
// See Update for more documentation.
func (o *EconomyCreateitem) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the EconomyCreateitem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EconomyCreateitem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	economyCreateitemUpdateCacheMut.RLock()
	cache, cached := economyCreateitemUpdateCache[key]
	economyCreateitemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			economyCreateitemAllColumns,
			economyCreateitemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update economy_createitem, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"economy_createitem\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, economyCreateitemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(economyCreateitemType, economyCreateitemMapping, append(wl, economyCreateitemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update economy_createitem row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for economy_createitem")
	}

	if !cached {
		economyCreateitemUpdateCacheMut.Lock()
		economyCreateitemUpdateCache[key] = cache
		economyCreateitemUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q economyCreateitemQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q economyCreateitemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for economy_createitem")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for economy_createitem")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EconomyCreateitemSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EconomyCreateitemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyCreateitemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"economy_createitem\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, economyCreateitemPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in economyCreateitem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all economyCreateitem")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EconomyCreateitem) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EconomyCreateitem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no economy_createitem provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(economyCreateitemColumnsWithDefault, o)

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

	economyCreateitemUpsertCacheMut.RLock()
	cache, cached := economyCreateitemUpsertCache[key]
	economyCreateitemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			economyCreateitemAllColumns,
			economyCreateitemColumnsWithDefault,
			economyCreateitemColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			economyCreateitemAllColumns,
			economyCreateitemPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert economy_createitem, could not build update column list")
		}

		ret := strmangle.SetComplement(economyCreateitemAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(economyCreateitemPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert economy_createitem, could not build conflict column list")
			}

			conflict = make([]string, len(economyCreateitemPrimaryKeyColumns))
			copy(conflict, economyCreateitemPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"economy_createitem\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(economyCreateitemType, economyCreateitemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(economyCreateitemType, economyCreateitemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert economy_createitem")
	}

	if !cached {
		economyCreateitemUpsertCacheMut.Lock()
		economyCreateitemUpsertCache[key] = cache
		economyCreateitemUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single EconomyCreateitem record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EconomyCreateitem) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single EconomyCreateitem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EconomyCreateitem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyCreateitem provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), economyCreateitemPrimaryKeyMapping)
	sql := "DELETE FROM \"economy_createitem\" WHERE \"guild_id\"=$1 AND \"user_id\"=$2"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from economy_createitem")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for economy_createitem")
	}

	return rowsAff, nil
}

func (q economyCreateitemQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q economyCreateitemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no economyCreateitemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economy_createitem")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_createitem")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EconomyCreateitemSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EconomyCreateitemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyCreateitemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"economy_createitem\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyCreateitemPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economyCreateitem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_createitem")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EconomyCreateitem) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no EconomyCreateitem provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EconomyCreateitem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEconomyCreateitem(ctx, exec, o.GuildID, o.UserID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyCreateitemSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty EconomyCreateitemSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyCreateitemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EconomyCreateitemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyCreateitemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"economy_createitem\".* FROM \"economy_createitem\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyCreateitemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EconomyCreateitemSlice")
	}

	*o = slice

	return nil
}

// EconomyCreateitemExistsG checks if the EconomyCreateitem row exists.
func EconomyCreateitemExistsG(ctx context.Context, guildID string, userID string) (bool, error) {
	return EconomyCreateitemExists(ctx, boil.GetContextDB(), guildID, userID)
}

// EconomyCreateitemExists checks if the EconomyCreateitem row exists.
func EconomyCreateitemExists(ctx context.Context, exec boil.ContextExecutor, guildID string, userID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"economy_createitem\" where \"guild_id\"=$1 AND \"user_id\"=$2 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, guildID, userID)
	}
	row := exec.QueryRowContext(ctx, sql, guildID, userID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if economy_createitem exists")
	}

	return exists, nil
}

// Exists checks if the EconomyCreateitem row exists.
func (o *EconomyCreateitem) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EconomyCreateitemExists(ctx, exec, o.GuildID, o.UserID)
}
