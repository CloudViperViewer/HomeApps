/*	------------------------------------------------------------------------------------------------*/
/*																									*/
/*																									*/
/*																									*/
/*																									*/
/*				Defines the utility functions databate sql executions								*/
/*																									*/
/*																									*/
/*																									*/
/*																									*/
/*																									*/
/*--------------------------------------------------------------------------------------------------*/

/*----------Structs---------------*/
//SelectQuery Used to construct a select db query
//PagingInfo Determines how many rows to return and what index to start
//Filter Structure to be used in query filter and logical expressions
//LogicExpression Structure to be used with a logical expression for more complex quieries

/*----------Functions---------------*/
//QueryFilter Function to build filter for db query
//LogicalExpression Function will create a logical expression combining conditions with AND OR

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
	Table           string
	Fields          []string
	LogicExpression LogicExpression
	PagingInfo      PagingInfo
}

// Determines how many rows to return and what index to start
//   - startIndex Offset for sql query
//   - batchSize Limit for sql queru
type PagingInfo struct {
	startIndex int
	batchSize  int
}

// Structure to be used in query filter and logical expressions
//   - string Operator used in the query "includes" "=" "<=" "<" ">=" ">" "is null" "is not null" "in" "not in"
//   - field to compare against
//   - value for the filter to compare against
type Filter struct {
	Operator string
	Field    string
	Value    []any
}

// Structure to be used with a logical expression for more complex quieries
//   - Operator "AND", "OR"
//   - Filters
//   - Logic Expressions to further construct
type LogicExpression struct {
	Operator         string
	Filters          []Filter
	LogicExpressions []LogicExpression
}

// This is a temporary function for testing eventually will be made more generic
func ASelectQuery(db *sql.DB, database string, table string, columns []any) {

	var row *sql.Rows
	var err error
	var columnNames []string
	var bankData tables.Bank
	columnNames = utils.GetAllTags(bankData, "db")
	var filters string
	var values []any

	var myExpress = LogicExpression{
		Operator: "AND",
		Filters: []Filter{
			{Operator: "in", Field: "bank_id", Value: []any{1, 2}},
		},
		LogicExpressions: nil,
	}

	filters, values = LogicalExpression(myExpress)
	log.Println(filters)
	log.Println(values)

	query := fmt.Sprintf("Select %s FROM %s.%s WHERE %s", utils.JoinArray(columnNames, ", "), database, table, filters)
	row, err = db.Query(query, values...)

	if err != nil {
		log.Fatal("Query failed ", err)
	} else {
		log.Println("Query successful")
	}

	defer row.Close()

	for row.Next() {
		var bankData tables.Bank

		if err = row.Scan(&bankData.BankID, &bankData.BankName, &bankData.DisplayOrder, &bankData.CreatedBy, &bankData.CreatedOn, &bankData.UpdatedBy, &bankData.UpdatedOn, &bankData.IsActive); err != nil {
			log.Fatal("Failed to scan row:", err)
		}

		log.Println(bankData.BankID, " ", bankData.BankName)

	}

	// Check for errors during iteration
	if err := row.Err(); err != nil {
		log.Fatal("Error iterating rows:", err)
	}
}

// Function to build filter for db query
//   - Filter takes a filter struct
//   - returns a string with the expression and the corresponding value
func QueryFilter(filter Filter) (string, []any) {

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
			log.Printf("Warning: Emty value sile for operator %s", filter.Operator)
			return "", nil
		}
		//covert to string
		strVal, ok := filter.Value[0].(string)
		if !ok {
			log.Printf("Warning: Non-string value for Like operator: %v", filter.Value[0])
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
		value = nil
	case "is not null":
		condition = fmt.Sprintf("%s IS NOT NULL", filter.Field)
		value = nil
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
		subExpression, subValue := QueryFilter(logicalExpression.Filters[i])
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
