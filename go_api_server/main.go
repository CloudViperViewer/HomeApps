package main

import (
	"database/sql"
	"log"

	"github.com/CloudViperViewer/HomeApps/go_api_server/api"
	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
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
		Table: table,
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
		Fields:     []string{"BankID", "BankName"},
		PagingInfo: database.PagingInfo{StartIndex: 1, BatchSize: 10},
	}

	data, err := database.ExecuteSelectQuery(db, query)

	log.Println(data.GetRows())

	// router.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	// })

	api.StartUpServer()

}
