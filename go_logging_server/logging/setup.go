/*
 * Setups up the logging data on start up
 */

/*
* Package Components:
*
* Structs:
* - fileRecord: is the struct that defines the log file location and name
* 	- Close: closes the file
*
* Functions:
* - SetupLoggingFiles: Sets up the logging files structure on start up
* - initLogFilesData: Sets up the logging files structure on start up
* - createLoggingFiles: Creates the files for each logging record eg. database
* - createLogFile: creates the specific log file
* - CloseLoggingFiles: closes all logging files
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
var files []fileRecord

// Close log file
func (f *fileRecord) Close() error {
	if f.file != nil {
		return f.file.Close()
	}

	return nil
}

// Close all logging files
func CloseLoggingFiles() {
	for i := range files {
		if files[i].file != nil {
			files[i].Close()
			files[i].file = nil
		}
	}

}

// Sets up the logging files structure on start up
func SetupLoggingFiles() {
	var err error

	//Init log files data
	initLogFilesData()

	err = utils.CreateDirectory(baseLogPath)

	if err != nil {
		log.Fatal(err)
	}

	err = createLoggingFiles()

	if err != nil {
		log.Fatal(err)
	}
}

// Initalises the files slice that holds the details for the log files
func initLogFilesData() {
	files = []fileRecord{
		{
			path:         logDatabasePath,
			filebaseName: logDatabaseApiFileName}}
}

// Creates the files for each logging record eg. database
func createLoggingFiles() error {
	var err error

	//loops through files
	for i := range files {

		file := &files[i]

		//create dir
		err = utils.CreateDirectory(file.path)
		if err != nil {
			return err
		}

		//create file
		file.file, err = createLogFile(file.filebaseName, file.path)

		if err != nil {
			return err
		}
	}

	return nil
}

// used to create the log file
// - file name: name of the log file may differ between services
func createLogFile(fileName string, path string) (*os.File, error) {

	var file *os.File
	var err error

	fileName = fmt.Sprintf("%s/%s %s.md", path, fileName, time.Now().UTC().Format(time.DateOnly))

	//Create file
	file, err = os.Create(fileName)

	//check error
	if err != nil {

		//file exists try to open it
		if os.IsExist(err) {

			file, err = os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0644)

			if err != nil {
				return nil, fmt.Errorf("failed to create log file %s: %w", fileName, err)
			}

			return file, nil
		}
		return nil, fmt.Errorf("failed to create log file %s: %w", fileName, err)

	}

	return file, nil

}
