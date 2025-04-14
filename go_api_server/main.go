/*
 * Contains functions used for main statup
 */

/*
* Package Components:
*
*
*Functions
* main: Main entry point of app
* checkLoggingServer: Checks if logging server reachable
*
 */

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/CloudViperViewer/HomeApps/go_api_server/api"
	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	_ "github.com/go-sql-driver/mysql"
)

/*Function that sets up connection to the database*/

/*Main function for the api server*/
func main() {

	//Check logging server has started
	err := checkLoggingServer()
	if err != nil {
		log.Fatalf("Error connecting to log server: %s", err.Error())
	}

	database.DatabaseInit()
	var db *sql.DB = database.GetDb()
	/*Defer won't execute till main returns*/
	defer db.Close()

	api.StartUpServer()

}

// Check if logging server reachable
func checkLoggingServer() error {

	var serverPort string = os.Getenv("LOG_SERVER_PORT")
	var url string = fmt.Sprintf("http://192.168.0.171:%s/health", serverPort)
	var err error
	var response *http.Response
	var maxRetries int = 10

	//execute health check
	for i := 0; i < maxRetries; i++ {

		response, err = http.Get(url)

		log.Printf("Connecting to log server....: Attempt: %v", i+1)

		//Check for running
		if response.StatusCode == http.StatusOK || err == nil {
			log.Printf("Connected to log server")
			return fmt.Errorf("log server not health")
		}
	}

	return err
}
