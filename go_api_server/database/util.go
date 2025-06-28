/*
 * Defines the utility functions databate sql executions
 */

/*
* Package Components:
*
* Structs:
* - SelectQuery: Used to construct a select db query
* - PagingInfo: Determines how many rows to return and what index to start
* - Filter: Structure used in query filters and logical expressions
* - LogicExpression: Structure for complex queries with combined conditions
*
* Functions:
* - ExecuteSelectQuery: Runs a select query on the db
* - queryFilter: Builds filters for db queries
* - logicalExpression: Creates expressions combining conditions with AND/OR operators
* - generateSelectQueryString: Generates the query string
* - ExecuteInsertQuery: Runs a insert query in the db
* - confirmInsertData: Confirms insert call has required fields
* - generatesInsertSQLStatment: Generates the insert query
 */

package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/CloudViperViewer/HomeApps/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Used to construct a select db query
// - Table to query
// - Fields to return in query
// - Logical expression to filter query
// - page size of query cruicial for performance
type SelectQuery struct {
	Table           tables.Table
	Fields          []string
	LogicExpression LogicExpression
	PagingInfo      PagingInfo
}

// Determines how many rows to return and what index to start
//   - startIndex Offset for sql query
//   - batchSize Limit for sql queru
type PagingInfo struct {
	//starts at 1
	StartIndex int `json:"startIndex"`
	//size of batch -1 to return all
	BatchSize int `json:"batchSize"`
}

// Structure to be used in query filter and logical expressions
//   - string Operator used in the query "includes" "=" "<=" "<" ">=" ">" "is null" "is not null" "in" "not in"
//   - field to compare against
//   - value for the filter to compare against
type Filter struct {
	Operator string `json:"operator"`
	Field    string `json:"field"`
	Value    []any  `json:"value"`
}

// Structure to be used with a logical expression for more complex quieries
//   - Operator "AND", "OR"
//   - Filters
//   - Logic Expressions to further construct
type LogicExpression struct {
	Operator         string            `json:"operator"`
	Filters          []Filter          `json:"filters"`
	LogicExpressions []LogicExpression `json:"logicalExpressions"`
}

// Function will create a logical expression combining conditions with AND OR
//   - Logic expression for functions to evaluate
func LogicalExpression(logicalExpression LogicExpression, structure any) (string, []any, error) {

	var expression string
	var expressionList []string
	var values []any
	var err error

	//Loop Over Nested Logical Expressions
	for i := range logicalExpression.LogicExpressions {
		subExpression, subValues, err := LogicalExpression(logicalExpression.LogicExpressions[i], structure)

		if err != nil {
			return "", nil, err
		}

		/*Make sure sub Expression is not empty then add to expressionlist*/
		if subExpression != "" {
			expressionList = append(expressionList, "("+subExpression+")")
		}

		//if sub values slice is not null append to values list to flatten list
		if subValues != nil {
			values = append(values, subValues...)
		}
	}

	//Loop over logical expressions and create the slice of conditions
	for i := range logicalExpression.Filters {
		subExpression, subValue, err := queryFilter(logicalExpression.Filters[i], structure)
		expressionList = append(expressionList, subExpression)

		if err != nil {
			return "", nil, err
		}

		//if sub values slice is not null append to values list to flatten list
		if subValue != nil {
			values = append(values, subValue...)
		}
	}

	//Check if expression list is empty
	if len(expressionList) == 0 {
		return "", values, nil
	}

	//joing expression list into one single expression
	switch logicalExpression.Operator {
	case "AND":
		expression = strings.Join(expressionList, " AND ")
	case "OR":
		expression = strings.Join(expressionList, " OR ")
	default:
		//If invalid operator default ot AND and log error
		utils.LogWarn(utils.ServiceDatabaseApi, "", "Warning: Unrecognised logical operator %s, defaulting to AND", logicalExpression.Operator)
		expression = strings.Join(expressionList, " AND ")
	}

	return expression, values, err

}

// Generic select statment for db query
// - SelectQuery - select query struct holding the query instructions
func ExecuteSelectQuery(db *sql.DB, selectQuery SelectQuery) (tables.Table, error) {

	var err error
	var values []any
	var query string
	var rows *sql.Rows
	var data tables.Table = selectQuery.Table

	query, values, err = generateSelectQueryString(selectQuery, data)

	if err != nil {
		return nil, err
	}

	//Execute query
	rows, err = db.Query(query, values...)
	if err != nil {
		return nil, err
	}

	//Close the query after the function is completed
	defer rows.Close()

	for rows.Next() {
		var baseTable any = data.GetBaseTableStruct()
		var fieldPtrs []any

		//Get Field pointers
		if len(selectQuery.Fields) == 0 {
			fieldPtrs, err = utils.GetStructAllFieldPtrs(baseTable)
		} else {
			fieldPtrs, err = utils.GetStructFieldPtrs(baseTable, selectQuery.Fields)
		}

		//check field ptrs didn't error
		if err != nil {
			utils.LogError(utils.ServiceDatabaseApi, "", "error in getting field pointers: %s", err.Error())
			return nil, err
		}
		//Scan row
		if err = rows.Scan(fieldPtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row %s", err)

		}

		//Add to rows
		data.Append(baseTable)
	}

	return data, nil

}

// Function to build filter for db query
//   - Filter takes a filter struct
//   - returns a string with the expression and the corresponding value
func queryFilter(filter Filter, structure any) (string, []any, error) {

	var condition string
	var value []any
	var field []string
	var err error

	//Get db field
	field, err = utils.GetTagList(structure, []string{filter.Field}, "db")

	if err != nil {
		return "", nil, err
	}

	/*Create query string*/
	//includes" "=" "<=" "<" ">=" ">" "is null" "is not null" "in" "not in"
	switch filter.Operator {
	case "=":
		condition = fmt.Sprintf("%s = ?", field[0])
		value = filter.Value
	case "includes":
		condition = fmt.Sprintf("%s LIKE ?", field[0])
		//Check if value empty
		if len(filter.Value) == 0 {
			return "", nil, fmt.Errorf("empty value slice for operator %s", filter.Operator)
		}
		//covert to string
		strVal, ok := filter.Value[0].(string)
		if !ok {
			return "", nil, fmt.Errorf("non-string value for Like operator: %v", filter.Value[0])
		}
		value = append(value, "%"+strVal+"%")
	case "<=":
		condition = fmt.Sprintf("%s <= ?", field[0])
		value = filter.Value
	case "<":
		condition = fmt.Sprintf("%s < ?", field[0])
		value = filter.Value
	case ">=":
		condition = fmt.Sprintf("%s >= ?", field[0])
		value = filter.Value
	case ">":
		condition = fmt.Sprintf("%s > ?", field[0])
		value = filter.Value
	case "is null":
		condition = fmt.Sprintf("%s IS NULL", field[0])
	case "is not null":
		condition = fmt.Sprintf("%s IS NOT NULL", field[0])
	case "in":
		placeHolders := make([]string, len(filter.Value))
		for i := range filter.Value {
			placeHolders[i] = "?"
		}
		condition = fmt.Sprintf("%s IN (%s)", field[0], strings.Join(placeHolders, ", "))
		value = filter.Value
	case "not in":
		placeHolders := make([]string, len(filter.Value))
		for i := range filter.Value {
			placeHolders[i] = "?"
		}
		condition = fmt.Sprintf("%s NOT IN (%s)", field[0], strings.Join(placeHolders, ", "))
		value = filter.Value
	default:
		err = fmt.Errorf("unrecognised operator %s in filter", filter.Operator)
		condition = ""
		value = nil
	}

	return condition, value, err
}

// Function used to generate the query string
//   - SelectQuery struct
func generateSelectQueryString(selectQuery SelectQuery, data tables.Table) (string, []any, error) {

	var query string
	var paging string
	var fields []string
	var logicalExpression string
	var values []any
	var err error

	if selectQuery.PagingInfo == (PagingInfo{}) {
		return "", nil, fmt.Errorf("paging info cannot be empty")
	}

	if selectQuery.PagingInfo.BatchSize == -1 {
		paging = fmt.Sprintf("LIMIT 100000000 OFFSET %v", selectQuery.PagingInfo.StartIndex-1)
	} else {
		paging = fmt.Sprintf("LIMIT %v OFFSET %v", selectQuery.PagingInfo.BatchSize, selectQuery.PagingInfo.StartIndex-1)
	}

	//Get db Fields
	if len(selectQuery.Fields) == 0 {
		fields = append(fields, "*")
	} else {
		fields, err = utils.GetTagList(data.GetBaseTableStruct(),
			selectQuery.Fields,
			"db")
	}

	//Check fields returned
	if err != nil {
		return "", nil, err
	}

	//get logical Expression
	logicalExpression, values, err = LogicalExpression(selectQuery.LogicExpression, data.GetBaseTableStruct())

	if err != nil {
		return "", nil, err
	}

	//Construct Query
	query = fmt.Sprintf("Select %s FROM `%s`.`%s` WHERE %s %s",
		strings.Join(fields, ", "),
		data.GetDatabase(),
		data.GetTableName(),
		logicalExpression,
		paging)

	return query, values, nil
}

// Executes an insert onto the database table
//   - db database to insert into
//   - table data to insert
func ExecuteInsertQuery(db *sql.DB, table string, rows []map[string]any) error {

	var passedTable tables.Table
	var tableStruct any
	var err error
	var statement string
	var insertData []any
	utils.LogDebug(utils.ServiceDatabaseApi, "", "In Execute Insert Query")
	//get table strcut
	passedTable, err = tables.TableFactory(table)
	if err != nil {
		return err
	}
	utils.LogDebug(utils.ServiceDatabaseApi, "", "Table Factory successful")

	tableStruct = passedTable.GetBaseTableStruct()

	//check if req field missing
	err = confirmInsertData(tableStruct, rows)
	if err != nil {
		utils.LogError(utils.ServiceDatabaseApi, "", "error occured: %s", err.Error())
		return err
	}

	//Generate insert sql statement
	statement, insertData = generatesInsertSQLStatment(passedTable, rows)

	//Execute insert
	_, err = db.Exec(statement, insertData...)

	//check for error
	if err != nil {
		utils.LogError(utils.ServiceDatabaseApi, "", "error occured %s", err.Error())
		return err
	}

	return nil
}

// Checks if the passed data is complete and not missing required fields
//   - table struct
//   - data to check
//   - returns error if required data missing
func confirmInsertData(tableStruct any, rows []map[string]any) error {

	var jsonStrings []string = utils.GetAllTags(tableStruct, "json")
	utils.LogDebug(utils.ServiceDatabaseApi, "", "Data confirm started")
	utils.LogDebug(utils.ServiceDatabaseApi, "", "Data to confirm: %v", rows)

	//loop over data
	for _, row := range rows {
		utils.LogDebug(utils.ServiceDatabaseApi, "", "row %v", row)
		//loop over json
		for _, column := range jsonStrings {
			var columnName string = strings.TrimSuffix(column, ",omitempty")
			var currentVal any = row[columnName]
			utils.LogDebug(utils.ServiceDatabaseApi, "", "columnValue %v", currentVal)
			var isRequired []string
			// //get if field is required
			isRequired, _ = utils.GetTagList(tableStruct, []string{columnName}, "binding")
			utils.LogDebug(utils.ServiceDatabaseApi, "", "%s: %v", columnName, isRequired)
			if currentVal == nil && len(isRequired) == 1 && isRequired[0] == "required" {
				return fmt.Errorf("field is required: %s", columnName)
			}
		}
	}

	utils.LogDebug(utils.ServiceDatabaseApi, "", "Data Confirm Finished")
	return nil
}

// Generates insert sql statment
//   - table to insert into
func generatesInsertSQLStatment(table tables.Table, data []map[string]any) (string, []any) {
	var database string = table.GetDatabase()
	var tableName string = table.GetTableName()
	var jsonFields []string = utils.GetAllTags(table.GetBaseTableStruct(), "json")
	//get database column names
	var columns string = strings.Join(utils.GetAllTags(table.GetBaseTableStruct(), "db"), ", ")
	//No of rows to insert
	var rows []string
	//Data for rows
	var insertData []any

	//Loop over and create the no of rows for insert
	for _, row := range data {
		var wildCards []string
		//Loop over each column
		for _, column := range jsonFields {
			var columnName string = strings.TrimSuffix(column, ",omitempty")
			insertData = append(insertData, row[columnName])
			wildCards = append(wildCards, "?")

		}
		rows = append(rows, fmt.Sprintf("(%s)", strings.Join(wildCards, ", ")))
	}

	//generates sql script
	var statement string = fmt.Sprintf("INSERT INTO `%s`.`%s` (%s) VALUES %s", database, tableName, columns, strings.Join(rows, ", \n"))

	//Debug
	utils.LogDebug(utils.ServiceDatabaseApi, "", "%s", statement)

	return statement, insertData
}
