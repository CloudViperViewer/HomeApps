package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/CloudViperViewer/HomeApps/go_api_server/utils"
	_ "github.com/go-sql-driver/mysql"
)

// Structure to be used in query filter and logical expressions
//   - string Operator used in the query "includes" "=" "<=" "<" ">=" ">" "is null" "is not null" "in" "not in"
//   - field to compare against
//   - value for the filter to compare against
type Filter struct {
	Operator string
	Field    string
	value    any
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
func SelectQuery(db *sql.DB, database string, table string, columns []any) {

	var row *sql.Rows
	var err error
	var columnNames []string
	var bankData tables.Bank
	columnNames = utils.GetAllTags(bankData, "db")
	var filters string
	var values []any

	filters, values = LogicalExpression(LogicExpression{"AND", []Filter{{"=", "bank_id", 1}, {"=", "is_active", 1}}, nil})

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
func QueryFilter(filter Filter) string {

	var condition string

	/*Create query string*/
	switch filter.Operator {
	case "=":
		condition = fmt.Sprintf("%s = ?", filter.Field)
	}

	return condition
}

// Function will create a logical expression combining conditions with AND OR
//   - Logic expression for functions to evaluate
func LogicalExpression(LogicalExpression LogicExpression) (string, []any) {

	var expression string
	var expressionList []string
	var values []any

	//Loop over logical expressions and create the slice of conditions
	for i := range LogicalExpression.Filters {
		expressionList = append(expressionList, QueryFilter(LogicalExpression.Filters[i]))
		values = append(values, LogicalExpression.Filters[i].value)
	}

	//joing expression list into one single expression
	switch LogicalExpression.Operator {
	case "AND":
		expression = utils.JoinArray(expressionList, " AND ")
	}

	return expression, values

}
