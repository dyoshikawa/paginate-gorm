// a

package paginator

import (
	"context"
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

// gorm.DB interface
type DBIface interface {
	// New clone a new db connection without search conditions
	New() *gorm.DB
	// Close close current db connection.  If database connection is not an io.Closer, returns an error.
	Close() error
	// DB get `*sql.DB` from current connection
	// If the underlying database connection is not a *sql.DB, returns nil
	DB() *sql.DB
	// CommonDB return the underlying `*sql.DB` or `*sql.Tx` instance, mainly intended to allow coexistence with legacy non-GORM code.
	CommonDB() gorm.SQLCommon
	// Dialect get dialect
	Dialect() gorm.Dialect
	// Callback return `Callbacks` container, you could add/change/delete callbacks with it
	//     db.Callback().Create().Register("update_created_at", updateCreated)
	// Refer https://jinzhu.github.io/gorm/development.html#callbacks
	Callback() *gorm.Callback
	// SetLogger replace default logger
	//SetLogger(log logger)
	//// LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs
	LogMode(enable bool) *gorm.DB
	// SetNowFuncOverride set the function to be used when creating a new timestamp
	SetNowFuncOverride(nowFuncOverride func() time.Time) *gorm.DB
	// BlockGlobalUpdate if true, generates an error on update/delete without where clause.
	// This is to prevent eventual error with empty objects updates/deletions
	BlockGlobalUpdate(enable bool) *gorm.DB
	// HasBlockGlobalUpdate return state of block
	HasBlockGlobalUpdate() bool
	// SingularTable use singular table by default
	SingularTable(enable bool)
	// NewScope create a scope for current operation
	NewScope(value interface{}) *gorm.Scope
	// QueryExpr returns the query as expr object
	//QueryExpr() *expr
	//// SubQuery returns the query as sub query
	//SubQuery() *expr
	//// Where return a new relation, filter records with given conditions, accepts `map`, `struct` or `string` as conditions, refer http://jinzhu.github.io/gorm/crud.html#query
	Where(query interface{}, args ...interface{}) *gorm.DB
	// Or filter records that match before conditions or this one, similar to `Where`
	Or(query interface{}, args ...interface{}) *gorm.DB
	// Not filter records that don't match current conditions, similar to `Where`
	Not(query interface{}, args ...interface{}) *gorm.DB
	// Limit specify the number of records to be retrieved
	Limit(limit interface{}) *gorm.DB
	// Offset specify the number of records to skip before starting to return the records
	Offset(offset interface{}) *gorm.DB
	// Order specify order when retrieve records from database, set reorder to `true` to overwrite defined conditions
	//     db.Order("name DESC")
	//     db.Order("name DESC", true) // reorder
	//     db.Order(gorm.Expr("name = ? DESC", "first")) // sql expression
	Order(value interface{}, reorder ...bool) *gorm.DB
	// Select specify fields that you want to retrieve from database when querying, by default, will select all fields;
	// When creating/updating, specify fields that you want to save to database
	Select(query interface{}, args ...interface{}) *gorm.DB
	// Omit specify fields that you want to ignore when saving to database for creating, updating
	Omit(columns ...string) *gorm.DB
	// Group specify the group method on the find
	Group(query string) *gorm.DB
	// Having specify HAVING conditions for GROUP BY
	Having(query interface{}, values ...interface{}) *gorm.DB
	// Joins specify Joins conditions
	//     db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
	Joins(query string, args ...interface{}) *gorm.DB
	// Scopes pass current database connection to arguments `func(*gorm.DB) *gorm.DB`, which could be used to add conditions dynamically
	//     func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
	//         return db.Where("amount > ?", 1000)
	//     }
	//
	//     func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
	//         return func (db *gorm.DB) *gorm.DB {
	//             return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
	//         }
	//     }
	//
	//     db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
	// Refer https://jinzhu.github.io/gorm/crud.html#scopes
	Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB
	// Unscoped return all record including deleted record, refer Soft Delete https://jinzhu.github.io/gorm/crud.html#soft-delete
	Unscoped() *gorm.DB
	// Attrs initialize struct with argument if record not found with `FirstOrInit` https://jinzhu.github.io/gorm/crud.html#firstorinit or `FirstOrCreate` https://jinzhu.github.io/gorm/crud.html#firstorcreate
	Attrs(attrs ...interface{}) *gorm.DB
	// Assign assign result with argument regardless it is found or not with `FirstOrInit` https://jinzhu.github.io/gorm/crud.html#firstorinit or `FirstOrCreate` https://jinzhu.github.io/gorm/crud.html#firstorcreate
	Assign(attrs ...interface{}) *gorm.DB
	// First find first record that match given conditions, order by primary key
	First(out interface{}, where ...interface{}) *gorm.DB
	// Take return a record that match given conditions, the order will depend on the database implementation
	Take(out interface{}, where ...interface{}) *gorm.DB
	// Last find last record that match given conditions, order by primary key
	Last(out interface{}, where ...interface{}) *gorm.DB
	// Find find records that match given conditions
	Find(out interface{}, where ...interface{}) *gorm.DB
	//Preloads preloads relations, don`t touch out
	Preloads(out interface{}) *gorm.DB
	// Scan scan value to a struct
	Scan(dest interface{}) *gorm.DB
	// Row return `*sql.Row` with given conditions
	Row() *sql.Row
	// Rows return `*sql.Rows` with given conditions
	Rows() (*sql.Rows, error)
	// ScanRows scan `*sql.Rows` to give struct
	ScanRows(rows *sql.Rows, result interface{}) error
	// Pluck used to query single column from a model as a map
	//     var ages []int64
	//     db.Find(&users).Pluck("age", &ages)
	Pluck(column string, value interface{}) *gorm.DB
	// Count get how many records for a model
	Count(value interface{}) *gorm.DB
	// Related get related associations
	Related(value interface{}, foreignKeys ...string) *gorm.DB
	// FirstOrInit find first matched record or initialize a new one with given conditions (only works with struct, map conditions)
	// https://jinzhu.github.io/gorm/crud.html#firstorinit
	FirstOrInit(out interface{}, where ...interface{}) *gorm.DB
	// FirstOrCreate find first matched record or create a new one with given conditions (only works with struct, map conditions)
	// https://jinzhu.github.io/gorm/crud.html#firstorcreate
	FirstOrCreate(out interface{}, where ...interface{}) *gorm.DB
	// Update update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
	Update(attrs ...interface{}) *gorm.DB
	// Updates update attributes with callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
	Updates(values interface{}, ignoreProtectedAttrs ...bool) *gorm.DB
	// UpdateColumn update attributes without callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
	UpdateColumn(attrs ...interface{}) *gorm.DB
	// UpdateColumns update attributes without callbacks, refer: https://jinzhu.github.io/gorm/crud.html#update
	UpdateColumns(values interface{}) *gorm.DB
	// Save update value in database, if the value doesn't have primary key, will insert it
	Save(value interface{}) *gorm.DB
	// Create insert the value into database
	Create(value interface{}) *gorm.DB
	// Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition
	Delete(value interface{}, where ...interface{}) *gorm.DB
	// Raw use raw sql as conditions, won't run it unless invoked by other methods
	//    db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
	Raw(sql string, values ...interface{}) *gorm.DB
	// Exec execute raw sql
	Exec(sql string, values ...interface{}) *gorm.DB
	// Model specify the model you would like to run db operations
	//    // update all users's name to `hello`
	//    db.Model(&User{}).Update("name", "hello")
	//    // if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
	//    db.Model(&user).Update("name", "hello")
	Model(value interface{}) *gorm.DB
	// Table specify the table you would like to run db operations
	Table(name string) *gorm.DB
	// Debug start debug mode
	Debug() *gorm.DB
	// Begin begins a transaction
	Begin() *gorm.DB
	// BeginTx begins a transaction with options
	BeginTx(ctx context.Context, opts *sql.TxOptions) *gorm.DB
	// Commit commit a transaction
	Commit() *gorm.DB
	// Rollback rollback a transaction
	Rollback() *gorm.DB
	// RollbackUnlessCommitted rollback a transaction if it has not yet been
	// committed.
	RollbackUnlessCommitted() *gorm.DB
	// NewRecord check if value's primary key is blank
	NewRecord(value interface{}) bool
	// RecordNotFound check if returning ErrRecordNotFound error
	RecordNotFound() bool
	// CreateTable create table for models
	CreateTable(models ...interface{}) *gorm.DB
	// DropTable drop table for models
	DropTable(values ...interface{}) *gorm.DB
	// DropTableIfExists drop table if it is exist
	DropTableIfExists(values ...interface{}) *gorm.DB
	// HasTable check has table or not
	HasTable(value interface{}) bool
	// AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data
	AutoMigrate(values ...interface{}) *gorm.DB
	// ModifyColumn modify column to type
	ModifyColumn(column string, typ string) *gorm.DB
	// DropColumn drop a column
	DropColumn(column string) *gorm.DB
	// AddIndex add index for columns with given name
	AddIndex(indexName string, columns ...string) *gorm.DB
	// AddUniqueIndex add unique index for columns with given name
	AddUniqueIndex(indexName string, columns ...string) *gorm.DB
	// RemoveIndex remove index with name
	RemoveIndex(indexName string) *gorm.DB
	// AddForeignKey Add foreign key to the given scope, e.g:
	//     db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
	AddForeignKey(field string, dest string, onDelete string, onUpdate string) *gorm.DB
	// RemoveForeignKey Remove foreign key from the given scope, e.g:
	//     db.Model(&User{}).RemoveForeignKey("city_id", "cities(id)")
	RemoveForeignKey(field string, dest string) *gorm.DB
	// Association start `Association Mode` to handler relations things easir in that mode, refer: https://jinzhu.github.io/gorm/associations.html#association-mode
	Association(column string) *gorm.Association
	// Preload preload associations with given conditions
	//    db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
	Preload(column string, conditions ...interface{}) *gorm.DB
	// Set set setting by name, which could be used in callbacks, will clone a new db, and update its setting
	Set(name string, value interface{}) *gorm.DB
	// InstantSet instant set setting, will affect current db
	InstantSet(name string, value interface{}) *gorm.DB
	// Get get setting by name
	Get(name string) (value interface{}, ok bool)
	// SetJoinTableHandler set a model's join table handler for a relation
	SetJoinTableHandler(source interface{}, column string, handler gorm.JoinTableHandlerInterface)
	// AddError add error to the db
	AddError(err error) error
	// GetErrors get happened errors from the db
	GetErrors() []error
}
