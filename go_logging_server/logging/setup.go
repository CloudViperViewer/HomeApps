/*
 * Setups up the logging data on start up
 */

/*
* Package Components:
*
* Structs:
* - fileRecord: is the struct that defines the log file location and name
*
* Functions:
* - SetupLoggingFiles: Sets up the logging files structure on start up
* - createLoggingFiles: Creates the files for each logging record eg. database
* - createLogFile: creates the specific log file
 */

package logging

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/CloudViperViewer/HomeApps/utils"
)

type fileRecord struct {
	path         string
	filebaseName string
	file         *os.File
}

// dir constants
const (
	baseLogPath            = "./logs"
	logDatabasePath        = "./logs/database"
	logDatabaseApiFileName = "databaseApi"
)

// Files
var files []fileRecord = []fileRecord{
	{
		path:         logDatabasePath,
		filebaseName: logDatabaseApiFileName}}

// Sets up the logging files structure on start up
func SetupLoggingFiles() {
	var err error

	err = utils.CreateDirectory(baseLogPath)

	if err != nil {
		log.Fatal(err)
	}

	err = createLoggingFiles()

	if err != nil {
		log.Fatal(err)
	}
}

// Creates the files for each logging record eg. database
func createLoggingFiles() error {
	var err error

	//loops through files
	for _, file := range files {

		//create dir
		err = utils.CreateDirectory(file.path)
		if err != nil {
			return err
		}

		file.file, err = createLogFile(file.filebaseName, file.path)
	}

	return nil
}

// used to create the log file
// - file name: name of the log file may differ between services
func createLogFile(fileName string, path string) (*os.File, error) {

	fileName = fmt.Sprintf("%s/%s %s.md", path, fileName, time.Now().UTC().Format(time.DateOnly))

	//Create file
	file, err := os.Create(fileName)

	//check error
	if err != nil && !os.IsExist(err) {
		return nil, fmt.Errorf("failed to create log file: %s", fileName)
	}

	return file, nil
}
