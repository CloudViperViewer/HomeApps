package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/CloudViperViewer/HomeApps/utils"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
Initiates the connection to the database
*/
func DatabaseInit() {
	var err error

	var connectionString string = getConnectionString()

	/*Attempts to connect to db 10 times*/
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", connectionString)

		if err == nil {
			err = db.Ping()
		}

		if err == nil {
			utils.LogInfo(utils.ServiceDatabaseApi, "", "Connected to database")
			return
		}

		/*5 second delay between attempts*/
		utils.LogInfo(utils.ServiceDatabaseApi, "", "Waiting for database to be ready... Retrying in 5 seconds")
		time.Sleep(5 * time.Second)
	}

	/*If all ten fail terminate*/
	if err != nil {
		utils.LogFatal(utils.ServiceDatabaseApi, "", "Failed to connect: %v", err)

	}
}

/*
Constructs the database connection string from the env variables
*/
func getConnectionString() string {

	/*Get env variables*/
	var db_host string = os.Getenv("DB_HOST")
	var db_name string = os.Getenv("DB_NAME")
	var db_user string = os.Getenv("DB_USER")
	var db_pass string = os.Getenv("DB_PASS")

	/*construct string*/
	return fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", db_user, db_pass, db_host, db_name)
}

/*Returns the db*/
func GetDb() *sql.DB {
	return db

}
