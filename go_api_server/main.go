package main

import (
	"database/sql"
	"log"
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
	table, err := tables.TableFactory(tables.BankTableKey)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	var query database.SelectQuery = database.SelectQuery{
		Table: &table,
		LogicExpression: database.LogicExpression{
			Operator: "AND",
			Filters: []database.Filter{
				{
					Operator: "=",
					Field:    "bank_id",
					Value:    []any{1},
				},
			},
		},
		Fields: []string{"BankID"},
	}

	data, err := database.ExecuteSelectQuery(query)

	println(data.GetRows())

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.Run(":8080")
}
