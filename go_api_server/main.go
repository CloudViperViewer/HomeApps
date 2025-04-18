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
	"time"

	"github.com/CloudViperViewer/HomeApps/go_api_server/api"
	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	"github.com/CloudViperViewer/HomeApps/utils"
	_ "github.com/go-sql-driver/mysql"
)

/*Function that sets up connection to the database*/

/*Main function for the api server*/
func main() {

	//Check logging server has started: this has to remain as is because log server unreachable
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

	var url string = fmt.Sprintf("%s%s", utils.GetLogServerUrl(), "health")
	var err error
	var response *http.Response
	var maxRetries int = 10
	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	//execute health check
	for i := 0; i < maxRetries; i++ {

		response, err = client.Get(url)

		log.Printf("Connecting to log server....: Attempt: %v", i+1)

		//Check for running
		if err == nil && response != nil && response.StatusCode == http.StatusOK {
			log.Printf("Connected to log server")
			return nil
		}
		log.Println(err.Error())
	}

	return fmt.Errorf("log server not healthy")
}
