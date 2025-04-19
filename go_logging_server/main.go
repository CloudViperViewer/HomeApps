package main

import (
	"github.com/CloudViperViewer/HomeApps/go_logging_server/api"
	"github.com/CloudViperViewer/HomeApps/go_logging_server/logging"
)

/*Entry point*/
func main() {

	//Setup structure
	logging.SetupLoggingFiles()

	//Start up server
	api.StartUpServer()
}
