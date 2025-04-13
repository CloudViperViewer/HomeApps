
tags: [[Go Database API]] [[HomeApps]]
# HomeApps Go API Documentation

## 🧠 Overview
This documentation covers the structure and functionality of the Go-based backend for the MyBudgetApp project, grouped by Go package.

---

## 📦 `main` Package
**File**: `main.go`

### Functions
- **`main()`**: 
  - Initializes the database connection via `database.DatabaseInit()`.
  - Defers database close.
  - Starts the API server with `api.StartUpServer()`.

---

## 📦 `database` Package
**Files**: `database.go`, `util.go` (db-related), `select.go`

### Global Variables
- `db *sql.DB`: Shared database connection.

### Structs
- `SelectQuery`: Holds information for executing a SELECT query.
- `PagingInfo`: Controls pagination of query results.
- `Filter`: Encapsulates individual filter criteria.
- `LogicExpression`: Combines filters and nested logical expressions.

### Functions
- `DatabaseInit()`: Tries to connect to the DB up to 10 times.
- `getConnectionString()`: Builds DB connection string from environment vars.
- `GetDb()`: Returns the global database connection.
- `ExecuteSelectQuery(db, SelectQuery)`: Runs the query and maps results.
- `LogicalExpression()`: Recursively builds WHERE clause from logic expressions.
- `queryFilter()`: Builds SQL filter conditions.
- `generateSelectQueryString()`: Assembles SQL SELECT query string.

---

## 📦 `tables` Package
**Files**: `tables.go`, `accounts.go`, `bank.go`

### Interfaces
- `Table`: Interface that all table structs must implement.
  - Methods: `GetDatabase()`, `GetTableName()`, `GetBaseTableStruct()`, `Append(any)`, `GetRows()`

### Structs
#### Bank-related
- `Bank`: Model for `fin_ref_bank` table.
- `BankTable`: Implements `Table` for a slice of `Bank`.

#### Account-related
- `Account`: Model for `fin_accounts` table.
- `AccountTable`: Implements `Table` for a slice of `Account`.

### Functions
- `TableFactory(key string)`: Factory for producing `Table` implementations.

---

## 📦 `utils` Package
**File**: `util.go` (utils)

### Functions
- `GetStructAllFieldPtrs(any) ([]any, error)`: Gets pointers for all struct fields.
- `GetStructFieldPtrs(any, []string) ([]any, error)`: Gets field pointers by name.
- `GetTagList(any, []string, string) ([]string, error)`: Gets tag values for named fields.
- `GetAllTags(any, string) []string`: Gets all tags for a struct.
- `JoinArray[T](array []T, delimiter string) string`: Joins string elements with a delimiter.

---

## 📦 `api` Package
**Files**: `api.go`, `select.go`

### Structs
- `selectQuery`: Represents incoming API request for `/api/select`.

### Endpoints
- `POST /api/select`: Selects records from a table using filters and paging.

### Functions
- `StartUpServer()`: Initializes the API server using Gin.
- `setupEndPoints(router)`: Binds routes.
- `dbQuerySelect(c *gin.Context)`: Handler for select query.
- `confirmData(selectQuery)`: Validates query input.
- `queryDb(selectQuery)`: Converts to `SelectQuery`, runs DB query, and returns result.

---