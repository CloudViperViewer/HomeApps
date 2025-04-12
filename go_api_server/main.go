package main

import (
	"database/sql"

	"github.com/CloudViperViewer/HomeApps/go_api_server/api"
	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	_ "github.com/go-sql-driver/mysql"
)

/*Function that sets up connection to the database*/

/*Main function for the api server*/
func main() {

	database.DatabaseInit()
	var db *sql.DB = database.GetDb()
	/*Defer won't execute till main returns*/
	defer db.Close()

	api.StartUpServer()

}
