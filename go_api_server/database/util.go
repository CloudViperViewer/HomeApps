package database

import (
	"database/sql"

	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/CloudViperViewer/HomeApps/go_api_server/utils"
	_ "github.com/go-sql-driver/mysql"
)

func SelectQuery(db *sql.DB, database string, table string, columns []any) {

	var row *sql.Rows
	var err error
	var columnNames []string
	var bankData tables.Bank
	columnNames = utils.GetAllTags(bankData, "db")

	row, err = db.Query("Select ? FROM ?.? WHERE `bank_id` = ?;", utils.JoinArray(columnNames, ", "), database, table, 1)

	// if err != nil {
	// 	log.Fatal("Query failed ", err)
	// } else {
	// 	log.Println("Query successful")
	// }

	// defer row.Close()

	// for row.Next() {
	// 	var bankData tables.Bank

	// 	if err = row.Scan(&bankData.BankID, &bankData.BankName); err != nil {
	// 		log.Fatal("Failed to scan row:", err)
	// 	}

	// 	log.Println(bankData.BankID, " ", bankData.BankName)

	// }

	// // Check for errors during iteration
	// if err := row.Err(); err != nil {
	// 	log.Fatal("Error iterating rows:", err)
	// }
}
