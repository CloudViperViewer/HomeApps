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
 */

package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/CloudViperViewer/HomeApps/go_api_server/utils"
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
func LogicalExpression(logicalExpression LogicExpression) (string, []any) {

	var expression string
	var expressionList []string
	var values []any

	//Loop Over Nested Logical Expressions
	for i := range logicalExpression.LogicExpressions {
		subExpression, subValues := LogicalExpression(logicalExpression.LogicExpressions[i])

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
		subExpression, subValue := queryFilter(logicalExpression.Filters[i])
		expressionList = append(expressionList, subExpression)

		//if sub values slice is not null append to values list to flatten list
		if subValue != nil {
			values = append(values, subValue...)
		}
	}

	//Check if expression list is empty
	if len(expressionList) == 0 {
		return "", values
	}

	//joing expression list into one single expression
	switch logicalExpression.Operator {
	case "AND":
		expression = utils.JoinArray(expressionList, " AND ")
	case "OR":
		expression = utils.JoinArray(expressionList, " OR ")
	default:
		//If invalid operator default ot AND and log error
		log.Printf("Warning: Unrecognised logical operator %s, defaulting to AND", logicalExpression.Operator)
		expression = utils.JoinArray(expressionList, " AND ")
	}

	return expression, values

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
		fieldPtrs, err = utils.GetStructFieldPtrs(baseTable, selectQuery.Fields)

		//check field ptrs didn't error
		if err != nil {
			log.Println(err)
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
func queryFilter(filter Filter) (string, []any) {

	var condition string
	var value []any

	/*Create query string*/
	//includes" "=" "<=" "<" ">=" ">" "is null" "is not null" "in" "not in"
	switch filter.Operator {
	case "=":
		condition = fmt.Sprintf("%s = ?", filter.Field)
		value = filter.Value
	case "includes":
		condition = fmt.Sprintf("%s LIKE ?", filter.Field)
		//Check if value empty
		if len(filter.Value) == 0 {
			log.Printf("Warning: Empty value slice for operator %s", filter.Operator)
			return "", nil
		}
		//covert to string
		strVal, ok := filter.Value[0].(string)
		if !ok {
			log.Printf("Warning: Non-string value for Like operator: %v", filter.Value[0])
			return "", nil
		}
		value = append(value, "%"+strVal+"%")
	case "<=":
		condition = fmt.Sprintf("%s <= ?", filter.Field)
		value = filter.Value
	case "<":
		condition = fmt.Sprintf("%s < ?", filter.Field)
		value = filter.Value
	case ">=":
		condition = fmt.Sprintf("%s >= ?", filter.Field)
		value = filter.Value
	case ">":
		condition = fmt.Sprintf("%s > ?", filter.Field)
		value = filter.Value
	case "is null":
		condition = fmt.Sprintf("%s IS NULL", filter.Field)
	case "is not null":
		condition = fmt.Sprintf("%s IS NOT NULL", filter.Field)
	case "in":
		placeHolders := make([]string, len(filter.Value))
		for i := range filter.Value {
			placeHolders[i] = "?"
		}
		condition = fmt.Sprintf("%s IN (%s)", filter.Field, utils.JoinArray(placeHolders, ", "))
		value = filter.Value
	case "not in":
		placeHolders := make([]string, len(filter.Value))
		for i := range filter.Value {
			placeHolders[i] = "?"
		}
		condition = fmt.Sprintf("%s NOT IN (%s)", filter.Field, utils.JoinArray(placeHolders, ", "))
		value = filter.Value
	default:
		log.Printf("Warning: Unrecognised operator %s in filter", filter.Operator)
		condition = ""
		value = nil
	}

	return condition, value
}

// Function used to generate the query string
// - SelectQuery struct
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
		paging = fmt.Sprintf("OFFSET %v", selectQuery.PagingInfo.StartIndex-1)
	} else {
		paging = fmt.Sprintf("LIMIT %v OFFSET %v", selectQuery.PagingInfo.BatchSize, selectQuery.PagingInfo.StartIndex-1)
	}

	//Get db Fields
	fields, err = utils.GetTagList(data.GetBaseTableStruct(),
		selectQuery.Fields,
		"db")

	//Check fields returned
	if err != nil {
		return "", nil, err
	}

	//get logical Expression
	logicalExpression, values = LogicalExpression(selectQuery.LogicExpression)

	//Construct Query
	query = fmt.Sprintf("Select %s FROM `%s`.`%s` WHERE %s %s",
		utils.JoinArray(fields, ", "),
		data.GetDatabase(),
		data.GetTableName(),
		logicalExpression,
		paging)

	return query, values, nil
}
