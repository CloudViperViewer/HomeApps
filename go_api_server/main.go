package main

import (
	"database/sql"
	"net/http"

	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/*Function that sets up connection to the database*/

/*Main function for the api server*/
func main() {

	database.DatabaseInit()
	var db *sql.DB = database.GetDb()
	/*Defer won't execute till main returns*/
	defer db.Close()

	/*Table type testing*/
	table := tables.TableFactory("bank")

	var bankData tables.Bank

	database.ASelectQuery(db, table.GetDatabase(), table.GetTableName(), []any{bankData.BankID, bankData.BankName})

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":8080")
}
