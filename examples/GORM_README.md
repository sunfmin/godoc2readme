




* [Variables](#pkg-variables)

* [Expr](#expr)
* [Parse Field Struct For Dialect](#parse-field-struct-for-dialect)
* [Register Dialect](#register-dialect)
* [To DB Name](#to-db-name)
* [Type Association](#type-association)
  * [Append](#association-append)
  * [Clear](#association-clear)
  * [Count](#association-count)
  * [Delete](#association-delete)
  * [Find](#association-find)
  * [Replace](#association-replace)
* [Type Callback](#type-callback)
  * [Create](#callback-create)
  * [Delete](#callback-delete)
  * [Query](#callback-query)
  * [Row Query](#callback-row-query)
  * [Update](#callback-update)
* [Type Callback Processor](#type-callback-processor)
  * [After](#callback-processor-after)
  * [Before](#callback-processor-before)
  * [Get](#callback-processor-get)
  * [Register](#callback-processor-register)
  * [Remove](#callback-processor-remove)
  * [Replace](#callback-processor-replace)
* [Type DB](#type-db)
  * [Open](#db-open)
  * [Add Error](#db-add-error)
  * [Add Foreign Key](#db-add-foreign-key)
  * [Add Index](#db-add-index)
  * [Add Unique Index](#db-add-unique-index)
  * [Assign](#db-assign)
  * [Association](#db-association)
  * [Attrs](#db-attrs)
  * [Auto Migrate](#db-auto-migrate)
  * [Begin](#db-begin)
  * [Callback](#db-callback)
  * [Close](#db-close)
  * [Commit](#db-commit)
  * [Common DB](#db-common-db)
  * [Count](#db-count)
  * [Create](#db-create)
  * [Create Table](#db-create-table)
  * [DB](#db-db)
  * [Debug](#db-debug)
  * [Delete](#db-delete)
  * [Dialect](#db-dialect)
  * [Drop Column](#db-drop-column)
  * [Drop Table](#db-drop-table)
  * [Drop Table If Exists](#db-drop-table-if-exists)
  * [Exec](#db-exec)
  * [Find](#db-find)
  * [First](#db-first)
  * [First Or Create](#db-first-or-create)
  * [First Or Init](#db-first-or-init)
  * [Get](#db-get)
  * [Get Errors](#db-get-errors)
  * [Group](#db-group)
  * [Has Table](#db-has-table)
  * [Having](#db-having)
  * [Instant Set](#db-instant-set)
  * [Joins](#db-joins)
  * [Last](#db-last)
  * [Limit](#db-limit)
  * [Log Mode](#db-log-mode)
  * [Model](#db-model)
  * [Modify Column](#db-modify-column)
  * [New](#db-new)
  * [New Record](#db-new-record)
  * [New Scope](#db-new-scope)
  * [Not](#db-not)
  * [Offset](#db-offset)
  * [Omit](#db-omit)
  * [Or](#db-or)
  * [Order](#db-order)
  * [Pluck](#db-pluck)
  * [Preload](#db-preload)
  * [Raw](#db-raw)
  * [Record Not Found](#db-record-not-found)
  * [Related](#db-related)
  * [Remove Index](#db-remove-index)
  * [Rollback](#db-rollback)
  * [Row](#db-row)
  * [Rows](#db-rows)
  * [Save](#db-save)
  * [Scan](#db-scan)
  * [Scan Rows](#db-scan-rows)
  * [Scopes](#db-scopes)
  * [Select](#db-select)
  * [Set](#db-set)
  * [Set Join Table Handler](#db-set-join-table-handler)
  * [Set Logger](#db-set-logger)
  * [Singular Table](#db-singular-table)
  * [Table](#db-table)
  * [Unscoped](#db-unscoped)
  * [Update](#db-update)
  * [Update Column](#db-update-column)
  * [Update Columns](#db-update-columns)
  * [Updates](#db-updates)
  * [Where](#db-where)
* [Type Default Foreign Key Namer](#type-default-foreign-key-namer)
  * [Build Foreign Key Name](#default-foreign-key-namer-build-foreign-key-name)
* [Type Dialect](#type-dialect)
* [Type Errors](#type-errors)
  * [Add](#errors-add)
  * [Error](#errors-error)
  * [Get Errors](#errors-get-errors)
* [Type Field](#type-field)
  * [Set](#field-set)
* [Type Join Table Foreign Key](#type-join-table-foreign-key)
* [Type Join Table Handler](#type-join-table-handler)
  * [Add](#join-table-handler-add)
  * [Delete](#join-table-handler-delete)
  * [Destination Foreign Keys](#join-table-handler-destination-foreign-keys)
  * [Join With](#join-table-handler-join-with)
  * [Setup](#join-table-handler-setup)
  * [Source Foreign Keys](#join-table-handler-source-foreign-keys)
  * [Table](#join-table-handler-table)
* [Type Join Table Handler Interface](#type-join-table-handler-interface)
* [Type Join Table Source](#type-join-table-source)
* [Type Log Writer](#type-log-writer)
* [Type Logger](#type-logger)
  * [Print](#logger-print)
* [Type Model](#type-model)
* [Type Model Struct](#type-model-struct)
  * [Table Name](#model-struct-table-name)
* [Type Relationship](#type-relationship)
* [Type Scope](#type-scope)
  * [Add To Vars](#scope-add-to-vars)
  * [Begin](#scope-begin)
  * [Call Method](#scope-call-method)
  * [Combined Condition Sql](#scope-combined-condition-sql)
  * [Commit Or Rollback](#scope-commit-or-rollback)
  * [DB](#scope-db)
  * [Dialect](#scope-dialect)
  * [Err](#scope-err)
  * [Exec](#scope-exec)
  * [Field By Name](#scope-field-by-name)
  * [Fields](#scope-fields)
  * [Get](#scope-get)
  * [Get Model Struct](#scope-get-model-struct)
  * [Get Struct Fields](#scope-get-struct-fields)
  * [Has Column](#scope-has-column)
  * [Has Error](#scope-has-error)
  * [Indirect Value](#scope-indirect-value)
  * [Instance Get](#scope-instance-get)
  * [Instance ID](#scope-instance-id)
  * [Instance Set](#scope-instance-set)
  * [Log](#scope-log)
  * [New](#scope-new)
  * [New DB](#scope-new-db)
  * [Omit Attrs](#scope-omit-attrs)
  * [Primary Field](#scope-primary-field)
  * [Primary Fields](#scope-primary-fields)
  * [Primary Key](#scope-primary-key)
  * [Primary Key Value](#scope-primary-key-value)
  * [Primary Key Zero](#scope-primary-key-zero)
  * [Quote](#scope-quote)
  * [Quoted Table Name](#scope-quoted-table-name)
  * [Raw](#scope-raw)
  * [SQ LD B](#scope-sq-ld-b)
  * [Select Attrs](#scope-select-attrs)
  * [Set](#scope-set)
  * [Set Column](#scope-set-column)
  * [Skip Left](#scope-skip-left)
  * [Table Name](#scope-table-name)
* [Type Struct Field](#type-struct-field)

## Expr
``` go
func Expr(expression string, args ...interface{}) *expr
```
Expr generate raw SQL expression, for example:


```go
DB.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
```



## Parse Field Struct For Dialect
``` go
func ParseFieldStructForDialect(field *StructField) (fieldValue reflect.Value, sqlType string, size int, additionalType string)
```
ParseFieldStructForDialect parse field struct for dialect



## Register Dialect
``` go
func RegisterDialect(name string, dialect Dialect)
```
RegisterDialect register new dialect



## To DB Name
``` go
func ToDBName(name string) string
```
ToDBName convert string to db name





## Type Association
``` go
type Association struct {
    Error error
    // contains filtered or unexported fields
}
```
Association Mode contains some helper methods to handle relationship things easily.










### Association: Append
``` go
func (association *Association) Append(values ...interface{}) *Association
```
Append append new associations for many2many, has_many, replace current association for has_one, belongs_to




### Association: Clear
``` go
func (association *Association) Clear() *Association
```
Clear remove relationship between source & current associations, won't delete those associations




### Association: Count
``` go
func (association *Association) Count() int
```
Count return the count of current associations




### Association: Delete
``` go
func (association *Association) Delete(values ...interface{}) *Association
```
Delete remove relationship between source & passed arguments, but won't delete those arguments




### Association: Find
``` go
func (association *Association) Find(value interface{}) *Association
```
Find find out all related associations




### Association: Replace
``` go
func (association *Association) Replace(values ...interface{}) *Association
```
Replace replace current associations with new one




## Type Callback
``` go
type Callback struct {
    // contains filtered or unexported fields
}
```
Callback is a struct that contains all CURD callbacks


```go
Field `creates` contains callbacks will be call when creating object
Field `updates` contains callbacks will be call when updating object
Field `deletes` contains callbacks will be call when deleting object
Field `queries` contains callbacks will be call when querying object with query methods like Find, First, Related, Association...
Field `rowQueries` contains callbacks will be call when querying object with Row, Rows...
Field `processors` contains all callback processors, will be used to generate above callbacks in order
```










### Callback: Create
``` go
func (c *Callback) Create() *CallbackProcessor
```
Create could be used to register callbacks for creating object


```go
db.Callback().Create().After("gorm:create").Register("plugin:run_after_create", func(*Scope) {
  // business logic
  ...

  // set error if some thing wrong happened, will rollback the creating
  scope.Err(errors.New("error"))
})
```




### Callback: Delete
``` go
func (c *Callback) Delete() *CallbackProcessor
```
Delete could be used to register callbacks for deleting object, refer `Create` for usage




### Callback: Query
``` go
func (c *Callback) Query() *CallbackProcessor
```
Query could be used to register callbacks for querying objects with query methods like `Find`, `First`, `Related`, `Association`...
Refer `Create` for usage




### Callback: Row Query
``` go
func (c *Callback) RowQuery() *CallbackProcessor
```
RowQuery could be used to register callbacks for querying objects with `Row`, `Rows`, refer `Create` for usage




### Callback: Update
``` go
func (c *Callback) Update() *CallbackProcessor
```
Update could be used to register callbacks for updating object, refer `Create` for usage




## Type Callback Processor
``` go
type CallbackProcessor struct {
    // contains filtered or unexported fields
}
```
CallbackProcessor contains callback informations










### Callback Processor: After
``` go
func (cp *CallbackProcessor) After(callbackName string) *CallbackProcessor
```
After insert a new callback after callback `callbackName`, refer `Callbacks.Create`




### Callback Processor: Before
``` go
func (cp *CallbackProcessor) Before(callbackName string) *CallbackProcessor
```
Before insert a new callback before callback `callbackName`, refer `Callbacks.Create`




### Callback Processor: Get
``` go
func (cp *CallbackProcessor) Get(callbackName string) (callback func(scope *Scope))
```
Get registered callback


```go
db.Callback().Create().Get("gorm:create")
```




### Callback Processor: Register
``` go
func (cp *CallbackProcessor) Register(callbackName string, callback func(scope *Scope))
```
Register a new callback, refer `Callbacks.Create`




### Callback Processor: Remove
``` go
func (cp *CallbackProcessor) Remove(callbackName string)
```
Remove a registered callback


```go
db.Callback().Create().Remove("gorm:update_time_stamp_when_create")
```




### Callback Processor: Replace
``` go
func (cp *CallbackProcessor) Replace(callbackName string, callback func(scope *Scope))
```
Replace a registered callback with new callback


```go
    db.Callback().Create().Replace("gorm:update_time_stamp_when_create", func(*Scope) {
		   scope.SetColumn("Created", now)
		   scope.SetColumn("Updated", now)
    })
```




## Type DB
``` go
type DB struct {
    Value        interface{}
    Error        error
    RowsAffected int64
    // contains filtered or unexported fields
}
```
DB contains information for current db connection







### func Open
``` go
func Open(dialect string, args ...interface{}) (*DB, error)
```
Open initialize a new db connection, need to import driver first, e.g:


```go
import _ "github.com/go-sql-driver/mysql"
func main() {
  db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
}
```

GORM has wrapped some drivers, for easier to remember driver's import path, so you could import the mysql driver with


```go
import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"
```





### DB: Add Error
``` go
func (s *DB) AddError(err error) error
```
AddError add error to the db




### DB: Add Foreign Key
``` go
func (s *DB) AddForeignKey(field string, dest string, onDelete string, onUpdate string) *DB
```
AddForeignKey Add foreign key to the given scope, e.g:


```go
db.Model(&User{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
```




### DB: Add Index
``` go
func (s *DB) AddIndex(indexName string, columns ...string) *DB
```
AddIndex add index for columns with given name




### DB: Add Unique Index
``` go
func (s *DB) AddUniqueIndex(indexName string, columns ...string) *DB
```
AddUniqueIndex add unique index for columns with given name




### DB: Assign
``` go
func (s *DB) Assign(attrs ...interface{}) *DB
```
Assign assign result with argument regardless it is found or not with `FirstOrInit` <a href="https://jinzhu.github.io/gorm/curd.html#firstorinit">https://jinzhu.github.io/gorm/curd.html#firstorinit</a> or `FirstOrCreate` <a href="https://jinzhu.github.io/gorm/curd.html#firstorcreate">https://jinzhu.github.io/gorm/curd.html#firstorcreate</a>




### DB: Association
``` go
func (s *DB) Association(column string) *Association
```
Association start `Association Mode` to handler relations things easir in that mode, refer: <a href="https://jinzhu.github.io/gorm/associations.html#association-mode">https://jinzhu.github.io/gorm/associations.html#association-mode</a>




### DB: Attrs
``` go
func (s *DB) Attrs(attrs ...interface{}) *DB
```
Attrs initialize struct with argument if record not found with `FirstOrInit` <a href="https://jinzhu.github.io/gorm/curd.html#firstorinit">https://jinzhu.github.io/gorm/curd.html#firstorinit</a> or `FirstOrCreate` <a href="https://jinzhu.github.io/gorm/curd.html#firstorcreate">https://jinzhu.github.io/gorm/curd.html#firstorcreate</a>




### DB: Auto Migrate
``` go
func (s *DB) AutoMigrate(values ...interface{}) *DB
```
AutoMigrate run auto migration for given models, will only add missing fields, won't delete/change current data




### DB: Begin
``` go
func (s *DB) Begin() *DB
```
Begin begin a transaction




### DB: Callback
``` go
func (s *DB) Callback() *Callback
```
Callback return `Callbacks` container, you could add/change/delete callbacks with it


```go
db.Callback().Create().Register("update_created_at", updateCreated)
```

Refer <a href="https://jinzhu.github.io/gorm/development.html#callbacks">https://jinzhu.github.io/gorm/development.html#callbacks</a>




### DB: Close
``` go
func (s *DB) Close() error
```
Close close current db connection




### DB: Commit
``` go
func (s *DB) Commit() *DB
```
Commit commit a transaction




### DB: Common DB
``` go
func (s *DB) CommonDB() sqlCommon
```
CommonDB return the underlying `*sql.DB` or `*sql.Tx` instance, mainly intended to allow coexistence with legacy non-GORM code.




### DB: Count
``` go
func (s *DB) Count(value interface{}) *DB
```
Count get how many records for a model




### DB: Create
``` go
func (s *DB) Create(value interface{}) *DB
```
Create insert the value into database




### DB: Create Table
``` go
func (s *DB) CreateTable(models ...interface{}) *DB
```
CreateTable create table for models




### DB: DB
``` go
func (s *DB) DB() *sql.DB
```
DB get `*sql.DB` from current connection




### DB: Debug
``` go
func (s *DB) Debug() *DB
```
Debug start debug mode




### DB: Delete
``` go
func (s *DB) Delete(value interface{}, where ...interface{}) *DB
```
Delete delete value match given conditions, if the value has primary key, then will including the primary key as condition




### DB: Dialect
``` go
func (s *DB) Dialect() Dialect
```
Dialect get dialect




### DB: Drop Column
``` go
func (s *DB) DropColumn(column string) *DB
```
DropColumn drop a column




### DB: Drop Table
``` go
func (s *DB) DropTable(values ...interface{}) *DB
```
DropTable drop table for models




### DB: Drop Table If Exists
``` go
func (s *DB) DropTableIfExists(values ...interface{}) *DB
```
DropTableIfExists drop table if it is exist




### DB: Exec
``` go
func (s *DB) Exec(sql string, values ...interface{}) *DB
```
Exec execute raw sql




### DB: Find
``` go
func (s *DB) Find(out interface{}, where ...interface{}) *DB
```
Find find records that match given conditions




### DB: First
``` go
func (s *DB) First(out interface{}, where ...interface{}) *DB
```
First find first record that match given conditions, order by primary key




### DB: First Or Create
``` go
func (s *DB) FirstOrCreate(out interface{}, where ...interface{}) *DB
```
FirstOrCreate find first matched record or create a new one with given conditions (only works with struct, map conditions)
<a href="https://jinzhu.github.io/gorm/curd.html#firstorcreate">https://jinzhu.github.io/gorm/curd.html#firstorcreate</a>




### DB: First Or Init
``` go
func (s *DB) FirstOrInit(out interface{}, where ...interface{}) *DB
```
FirstOrInit find first matched record or initialize a new one with given conditions (only works with struct, map conditions)
<a href="https://jinzhu.github.io/gorm/curd.html#firstorinit">https://jinzhu.github.io/gorm/curd.html#firstorinit</a>




### DB: Get
``` go
func (s *DB) Get(name string) (value interface{}, ok bool)
```
Get get setting by name




### DB: Get Errors
``` go
func (s *DB) GetErrors() (errors []error)
```
GetErrors get happened errors from the db




### DB: Group
``` go
func (s *DB) Group(query string) *DB
```
Group specify the group method on the find




### DB: Has Table
``` go
func (s *DB) HasTable(value interface{}) bool
```
HasTable check has table or not




### DB: Having
``` go
func (s *DB) Having(query string, values ...interface{}) *DB
```
Having specify HAVING conditions for GROUP BY




### DB: Instant Set
``` go
func (s *DB) InstantSet(name string, value interface{}) *DB
```
InstantSet instant set setting, will affect current db




### DB: Joins
``` go
func (s *DB) Joins(query string, args ...interface{}) *DB
```
Joins specify Joins conditions


```go
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Find(&user)
```




### DB: Last
``` go
func (s *DB) Last(out interface{}, where ...interface{}) *DB
```
Last find last record that match given conditions, order by primary key




### DB: Limit
``` go
func (s *DB) Limit(limit interface{}) *DB
```
Limit specify the number of records to be retrieved




### DB: Log Mode
``` go
func (s *DB) LogMode(enable bool) *DB
```
LogMode set log mode, `true` for detailed logs, `false` for no log, default, will only print error logs




### DB: Model
``` go
func (s *DB) Model(value interface{}) *DB
```
Model specify the model you would like to run db operations


```go
// update all users's name to `hello`
db.Model(&User{}).Update("name", "hello")
// if user's primary key is non-blank, will use it as condition, then will only update the user's name to `hello`
db.Model(&user).Update("name", "hello")
```




### DB: Modify Column
``` go
func (s *DB) ModifyColumn(column string, typ string) *DB
```
ModifyColumn modify column to type




### DB: New
``` go
func (s *DB) New() *DB
```
New clone a new db connection without search conditions




### DB: New Record
``` go
func (s *DB) NewRecord(value interface{}) bool
```
NewRecord check if value's primary key is blank




### DB: New Scope
``` go
func (s *DB) NewScope(value interface{}) *Scope
```
NewScope create a scope for current operation




### DB: Not
``` go
func (s *DB) Not(query interface{}, args ...interface{}) *DB
```
Not filter records that don't match current conditions, similar to `Where`




### DB: Offset
``` go
func (s *DB) Offset(offset interface{}) *DB
```
Offset specify the number of records to skip before starting to return the records




### DB: Omit
``` go
func (s *DB) Omit(columns ...string) *DB
```
Omit specify fields that you want to ignore when saving to database for creating, updating




### DB: Or
``` go
func (s *DB) Or(query interface{}, args ...interface{}) *DB
```
Or filter records that match before conditions or this one, similar to `Where`




### DB: Order
``` go
func (s *DB) Order(value interface{}, reorder ...bool) *DB
```
Order specify order when retrieve records from database, set reorder to `true` to overwrite defined conditions


```go
db.Order("name DESC")
db.Order("name DESC", true) // reorder
db.Order(gorm.Expr("name = ? DESC", "first")) // sql expression
```




### DB: Pluck
``` go
func (s *DB) Pluck(column string, value interface{}) *DB
```
Pluck used to query single column from a model as a map


```go
var ages []int64
db.Find(&users).Pluck("age", &ages)
```




### DB: Preload
``` go
func (s *DB) Preload(column string, conditions ...interface{}) *DB
```
Preload preload associations with given conditions


```go
db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
```




### DB: Raw
``` go
func (s *DB) Raw(sql string, values ...interface{}) *DB
```
Raw use raw sql as conditions, won't run it unless invoked by other methods


```go
db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)
```




### DB: Record Not Found
``` go
func (s *DB) RecordNotFound() bool
```
RecordNotFound check if returning ErrRecordNotFound error




### DB: Related
``` go
func (s *DB) Related(value interface{}, foreignKeys ...string) *DB
```
Related get related associations




### DB: Remove Index
``` go
func (s *DB) RemoveIndex(indexName string) *DB
```
RemoveIndex remove index with name




### DB: Rollback
``` go
func (s *DB) Rollback() *DB
```
Rollback rollback a transaction




### DB: Row
``` go
func (s *DB) Row() *sql.Row
```
Row return `*sql.Row` with given conditions




### DB: Rows
``` go
func (s *DB) Rows() (*sql.Rows, error)
```
Rows return `*sql.Rows` with given conditions




### DB: Save
``` go
func (s *DB) Save(value interface{}) *DB
```
Save update value in database, if the value doesn't have primary key, will insert it




### DB: Scan
``` go
func (s *DB) Scan(dest interface{}) *DB
```
Scan scan value to a struct




### DB: Scan Rows
``` go
func (s *DB) ScanRows(rows *sql.Rows, result interface{}) error
```
ScanRows scan `*sql.Rows` to give struct




### DB: Scopes
``` go
func (s *DB) Scopes(funcs ...func(*DB) *DB) *DB
```
Scopes pass current database connection to arguments `func(*DB) *DB`, which could be used to add conditions dynamically


```go
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
    return db.Where("amount > ?", 1000)
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
    return func (db *gorm.DB) *gorm.DB {
        return db.Scopes(AmountGreaterThan1000).Where("status in (?)", status)
    }
}

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
```

Refer <a href="https://jinzhu.github.io/gorm/curd.html#scopes">https://jinzhu.github.io/gorm/curd.html#scopes</a>




### DB: Select
``` go
func (s *DB) Select(query interface{}, args ...interface{}) *DB
```
Select specify fields that you want to retrieve from database when querying, by default, will select all fields;
When creating/updating, specify fields that you want to save to database




### DB: Set
``` go
func (s *DB) Set(name string, value interface{}) *DB
```
Set set setting by name, which could be used in callbacks, will clone a new db, and update its setting




### DB: Set Join Table Handler
``` go
func (s *DB) SetJoinTableHandler(source interface{}, column string, handler JoinTableHandlerInterface)
```
SetJoinTableHandler set a model's join table handler for a relation




### DB: Set Logger
``` go
func (s *DB) SetLogger(log logger)
```
SetLogger replace default logger




### DB: Singular Table
``` go
func (s *DB) SingularTable(enable bool)
```
SingularTable use singular table by default




### DB: Table
``` go
func (s *DB) Table(name string) *DB
```
Table specify the table you would like to run db operations




### DB: Unscoped
``` go
func (s *DB) Unscoped() *DB
```
Unscoped return all record including deleted record, refer Soft Delete <a href="https://jinzhu.github.io/gorm/curd.html#soft-delete">https://jinzhu.github.io/gorm/curd.html#soft-delete</a>




### DB: Update
``` go
func (s *DB) Update(attrs ...interface{}) *DB
```
Update update attributes with callbacks, refer: <a href="https://jinzhu.github.io/gorm/curd.html#update">https://jinzhu.github.io/gorm/curd.html#update</a>




### DB: Update Column
``` go
func (s *DB) UpdateColumn(attrs ...interface{}) *DB
```
UpdateColumn update attributes without callbacks, refer: <a href="https://jinzhu.github.io/gorm/curd.html#update">https://jinzhu.github.io/gorm/curd.html#update</a>




### DB: Update Columns
``` go
func (s *DB) UpdateColumns(values interface{}) *DB
```
UpdateColumns update attributes without callbacks, refer: <a href="https://jinzhu.github.io/gorm/curd.html#update">https://jinzhu.github.io/gorm/curd.html#update</a>




### DB: Updates
``` go
func (s *DB) Updates(values interface{}, ignoreProtectedAttrs ...bool) *DB
```
Updates update attributes with callbacks, refer: <a href="https://jinzhu.github.io/gorm/curd.html#update">https://jinzhu.github.io/gorm/curd.html#update</a>




### DB: Where
``` go
func (s *DB) Where(query interface{}, args ...interface{}) *DB
```
Where return a new relation, filter records with given conditions, accepts `map`, `struct` or `string` as conditions, refer <a href="http://jinzhu.github.io/gorm/curd.html#query">http://jinzhu.github.io/gorm/curd.html#query</a>




## Type Default Foreign Key Namer
``` go
type DefaultForeignKeyNamer struct {
}
```
DefaultForeignKeyNamer contains the default foreign key name generator method










### Default Foreign Key Namer: Build Foreign Key Name
``` go
func (DefaultForeignKeyNamer) BuildForeignKeyName(tableName, field, dest string) string
```



## Type Dialect
``` go
type Dialect interface {
    // GetName get dialect's name
    GetName() string

    // SetDB set db for dialect
    SetDB(db *sql.DB)

    // BindVar return the placeholder for actual values in SQL statements, in many dbs it is "?", Postgres using $1
    BindVar(i int) string
    // Quote quotes field name to avoid SQL parsing exceptions by using a reserved word as a field name
    Quote(key string) string
    // DataTypeOf return data's sql type
    DataTypeOf(field *StructField) string

    // HasIndex check has index or not
    HasIndex(tableName string, indexName string) bool
    // HasForeignKey check has foreign key or not
    HasForeignKey(tableName string, foreignKeyName string) bool
    // RemoveIndex remove index
    RemoveIndex(tableName string, indexName string) error
    // HasTable check has table or not
    HasTable(tableName string) bool
    // HasColumn check has column or not
    HasColumn(tableName string, columnName string) bool

    // LimitAndOffsetSQL return generated SQL with Limit and Offset, as mssql has special case
    LimitAndOffsetSQL(limit, offset interface{}) string
    // SelectFromDummyTable return select values, for most dbs, `SELECT values` just works, mysql needs `SELECT value FROM DUAL`
    SelectFromDummyTable() string
    // LastInsertIdReturningSuffix most dbs support LastInsertId, but postgres needs to use `RETURNING`
    LastInsertIDReturningSuffix(tableName, columnName string) string

    // BuildForeignKeyName returns a foreign key name for the given table, field and reference
    BuildForeignKeyName(tableName, field, dest string) string

    // CurrentDatabase return current database name
    CurrentDatabase() string
}
```
Dialect interface contains behaviors that differ across SQL database










## Type Errors
``` go
type Errors struct {
    // contains filtered or unexported fields
}
```
Errors contains all happened errors










### Errors: Add
``` go
func (errs *Errors) Add(err error)
```
Add add an error




### Errors: Error
``` go
func (errs Errors) Error() string
```
Error format happened errors




### Errors: Get Errors
``` go
func (errs Errors) GetErrors() []error
```
GetErrors get all happened errors




## Type Field
``` go
type Field struct {
    *StructField
    IsBlank bool
    Field   reflect.Value
}
```
Field model field definition










### Field: Set
``` go
func (field *Field) Set(value interface{}) (err error)
```
Set set a value to the field




## Type Join Table Foreign Key
``` go
type JoinTableForeignKey struct {
    DBName            string
    AssociationDBName string
}
```
JoinTableForeignKey join table foreign key struct










## Type Join Table Handler
``` go
type JoinTableHandler struct {
    TableName   string          `sql:"-"`
    Source      JoinTableSource `sql:"-"`
    Destination JoinTableSource `sql:"-"`
}
```
JoinTableHandler default join table handler










### Join Table Handler: Add
``` go
func (s JoinTableHandler) Add(handler JoinTableHandlerInterface, db *DB, source interface{}, destination interface{}) error
```
Add create relationship in join table for source and destination




### Join Table Handler: Delete
``` go
func (s JoinTableHandler) Delete(handler JoinTableHandlerInterface, db *DB, sources ...interface{}) error
```
Delete delete relationship in join table for sources




### Join Table Handler: Destination Foreign Keys
``` go
func (s *JoinTableHandler) DestinationForeignKeys() []JoinTableForeignKey
```
DestinationForeignKeys return destination foreign keys




### Join Table Handler: Join With
``` go
func (s JoinTableHandler) JoinWith(handler JoinTableHandlerInterface, db *DB, source interface{}) *DB
```
JoinWith query with `Join` conditions




### Join Table Handler: Setup
``` go
func (s *JoinTableHandler) Setup(relationship *Relationship, tableName string, source reflect.Type, destination reflect.Type)
```
Setup initialize a default join table handler




### Join Table Handler: Source Foreign Keys
``` go
func (s *JoinTableHandler) SourceForeignKeys() []JoinTableForeignKey
```
SourceForeignKeys return source foreign keys




### Join Table Handler: Table
``` go
func (s JoinTableHandler) Table(db *DB) string
```
Table return join table's table name




## Type Join Table Handler Interface
``` go
type JoinTableHandlerInterface interface {
    // initialize join table handler
    Setup(relationship *Relationship, tableName string, source reflect.Type, destination reflect.Type)
    // Table return join table's table name
    Table(db *DB) string
    // Add create relationship in join table for source and destination
    Add(handler JoinTableHandlerInterface, db *DB, source interface{}, destination interface{}) error
    // Delete delete relationship in join table for sources
    Delete(handler JoinTableHandlerInterface, db *DB, sources ...interface{}) error
    // JoinWith query with `Join` conditions
    JoinWith(handler JoinTableHandlerInterface, db *DB, source interface{}) *DB
    // SourceForeignKeys return source foreign keys
    SourceForeignKeys() []JoinTableForeignKey
    // DestinationForeignKeys return destination foreign keys
    DestinationForeignKeys() []JoinTableForeignKey
}
```
JoinTableHandlerInterface is an interface for how to handle many2many relations










## Type Join Table Source
``` go
type JoinTableSource struct {
    ModelType   reflect.Type
    ForeignKeys []JoinTableForeignKey
}
```
JoinTableSource is a struct that contains model type and foreign keys










## Type Log Writer
``` go
type LogWriter interface {
    Println(v ...interface{})
}
```
LogWriter log writer interface










## Type Logger
``` go
type Logger struct {
    LogWriter
}
```
Logger default logger










### Logger: Print
``` go
func (logger Logger) Print(values ...interface{})
```
Print format & print log




## Type Model
``` go
type Model struct {
    ID        uint `gorm:"primary_key"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time `sql:"index"`
}
```
Model base model definition, including fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`, which could be embedded in your models


```go
type User struct {
  gorm.Model
}
```










## Type Model Struct
``` go
type ModelStruct struct {
    PrimaryFields []*StructField
    StructFields  []*StructField
    ModelType     reflect.Type
    // contains filtered or unexported fields
}
```
ModelStruct model definition










### Model Struct: Table Name
``` go
func (s *ModelStruct) TableName(db *DB) string
```
TableName get model's table name




## Type Relationship
``` go
type Relationship struct {
    Kind                         string
    PolymorphicType              string
    PolymorphicDBName            string
    ForeignFieldNames            []string
    ForeignDBNames               []string
    AssociationForeignFieldNames []string
    AssociationForeignDBNames    []string
    JoinTableHandler             JoinTableHandlerInterface
}
```
Relationship described the relationship between models










## Type Scope
``` go
type Scope struct {
    Search  *search
    Value   interface{}
    SQL     string
    SQLVars []interface{}
    // contains filtered or unexported fields
}
```
Scope contain current operation's information when you perform any operation on the database










### Scope: Add To Vars
``` go
func (scope *Scope) AddToVars(value interface{}) string
```
AddToVars add value as sql's vars, used to prevent SQL injection




### Scope: Begin
``` go
func (scope *Scope) Begin() *Scope
```
Begin start a transaction




### Scope: Call Method
``` go
func (scope *Scope) CallMethod(methodName string)
```
CallMethod call scope value's method, if it is a slice, will call its element's method one by one




### Scope: Combined Condition Sql
``` go
func (scope *Scope) CombinedConditionSql() string
```
CombinedConditionSql return combined condition sql




### Scope: Commit Or Rollback
``` go
func (scope *Scope) CommitOrRollback() *Scope
```
CommitOrRollback commit current transaction if no error happened, otherwise will rollback it




### Scope: DB
``` go
func (scope *Scope) DB() *DB
```
DB return scope's DB connection




### Scope: Dialect
``` go
func (scope *Scope) Dialect() Dialect
```
Dialect get dialect




### Scope: Err
``` go
func (scope *Scope) Err(err error) error
```
Err add error to Scope




### Scope: Exec
``` go
func (scope *Scope) Exec() *Scope
```
Exec perform generated SQL




### Scope: Field By Name
``` go
func (scope *Scope) FieldByName(name string) (field *Field, ok bool)
```
FieldByName find `gorm.Field` with field name or db name




### Scope: Fields
``` go
func (scope *Scope) Fields() []*Field
```
Fields get value's fields




### Scope: Get
``` go
func (scope *Scope) Get(name string) (interface{}, bool)
```
Get get setting by name




### Scope: Get Model Struct
``` go
func (scope *Scope) GetModelStruct() *ModelStruct
```
GetModelStruct get value's model struct, relationships based on struct and tag definition




### Scope: Get Struct Fields
``` go
func (scope *Scope) GetStructFields() (fields []*StructField)
```
GetStructFields get model's field structs




### Scope: Has Column
``` go
func (scope *Scope) HasColumn(column string) bool
```
HasColumn to check if has column




### Scope: Has Error
``` go
func (scope *Scope) HasError() bool
```
HasError check if there are any error




### Scope: Indirect Value
``` go
func (scope *Scope) IndirectValue() reflect.Value
```
IndirectValue return scope's reflect value's indirect value




### Scope: Instance Get
``` go
func (scope *Scope) InstanceGet(name string) (interface{}, bool)
```
InstanceGet get instance setting from current operation




### Scope: Instance ID
``` go
func (scope *Scope) InstanceID() string
```
InstanceID get InstanceID for scope




### Scope: Instance Set
``` go
func (scope *Scope) InstanceSet(name string, value interface{}) *Scope
```
InstanceSet set instance setting for current operation, but not for operations in callbacks, like saving associations callback




### Scope: Log
``` go
func (scope *Scope) Log(v ...interface{})
```
Log print log message




### Scope: New
``` go
func (scope *Scope) New(value interface{}) *Scope
```
New create a new Scope without search information




### Scope: New DB
``` go
func (scope *Scope) NewDB() *DB
```
NewDB create a new DB without search information




### Scope: Omit Attrs
``` go
func (scope *Scope) OmitAttrs() []string
```
OmitAttrs return omitted attributes




### Scope: Primary Field
``` go
func (scope *Scope) PrimaryField() *Field
```
PrimaryField return scope's main primary field, if defined more that one primary fields, will return the one having column name `id` or the first one




### Scope: Primary Fields
``` go
func (scope *Scope) PrimaryFields() (fields []*Field)
```
PrimaryFields return scope's primary fields




### Scope: Primary Key
``` go
func (scope *Scope) PrimaryKey() string
```
PrimaryKey get main primary field's db name




### Scope: Primary Key Value
``` go
func (scope *Scope) PrimaryKeyValue() interface{}
```
PrimaryKeyValue get the primary key's value




### Scope: Primary Key Zero
``` go
func (scope *Scope) PrimaryKeyZero() bool
```
PrimaryKeyZero check main primary field's value is blank or not




### Scope: Quote
``` go
func (scope *Scope) Quote(str string) string
```
Quote used to quote string to escape them for database




### Scope: Quoted Table Name
``` go
func (scope *Scope) QuotedTableName() (name string)
```
QuotedTableName return quoted table name




### Scope: Raw
``` go
func (scope *Scope) Raw(sql string) *Scope
```
Raw set raw sql




### Scope: SQ LD B
``` go
func (scope *Scope) SQLDB() sqlCommon
```
SQLDB return *sql.DB




### Scope: Select Attrs
``` go
func (scope *Scope) SelectAttrs() []string
```
SelectAttrs return selected attributes




### Scope: Set
``` go
func (scope *Scope) Set(name string, value interface{}) *Scope
```
Set set value by name




### Scope: Set Column
``` go
func (scope *Scope) SetColumn(column interface{}, value interface{}) error
```
SetColumn to set the column's value, column could be field or field's name/dbname




### Scope: Skip Left
``` go
func (scope *Scope) SkipLeft()
```
SkipLeft skip remaining callbacks




### Scope: Table Name
``` go
func (scope *Scope) TableName() string
```
TableName return table name




## Type Struct Field
``` go
type StructField struct {
    DBName          string
    Name            string
    Names           []string
    IsPrimaryKey    bool
    IsNormal        bool
    IsIgnored       bool
    IsScanner       bool
    HasDefaultValue bool
    Tag             reflect.StructTag
    TagSettings     map[string]string
    Struct          reflect.StructField
    IsForeignKey    bool
    Relationship    *Relationship
}
```
StructField model field's struct definition











