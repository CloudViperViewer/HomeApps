package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/CloudViperViewer/HomeApps/go_api_server/utils"
	_ "github.com/go-sql-driver/mysql"
)

// This is a temporary function for testing eventually will be made more generic
func SelectQuery(db *sql.DB, database string, table string, columns []any) {

	var row *sql.Rows
	var err error
	var columnNames []string
	var bankData tables.Bank
	columnNames = utils.GetAllTags(bankData, "db")

	query := fmt.Sprintf("Select %s FROM %s.%s WHERE ?", utils.JoinArray(columnNames, ", "), database, table)

	row, err = db.Query(query, QueryFilter("=", "bank_id", 1))

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
//   - string Operator used in the query "includes" "=" "<=" "<" ">=" ">" "is null" "is not null" "in" "not in"
//   - field to compare against
//   - value in for filter (not required for "is null" and "is not null")
func QueryFilter(operator string, field string, value any) string {

	var condition string

	switch operator {
	case "=":
		condition = fmt.Sprintf("%s = %v", field, value)
	}

	log.Println(condition)

	return condition
}
