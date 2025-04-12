package main

import (
	"database/sql"

	"github.com/CloudViperViewer/HomeApps/go_api_server/api"
	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	_ "github.com/go-sql-driver/mysql"
)

/*Function that sets up connection to the database*/

// main is the entry point of the API server. It initializes the database connection, defers its closure to ensure proper cleanup, and starts the API server.
func main() {

	database.DatabaseInit()
	var db *sql.DB = database.GetDb()
	/*Defer won't execute till main returns*/
	defer db.Close()

	api.StartUpServer()

}
