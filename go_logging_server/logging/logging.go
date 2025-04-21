/*
 * Holds methods and structs used to log calls
 */

/*
* Package Components:
*
* Structures
* - Log: struct used to define the log message
*
*
*Functions
* - WriteLog: Used to writ the log to the required document and print to terminal
* - getLevel: Used to convert the level to colourd string
 */

package logging

import (
	"fmt"
	"log"
	"os"

	"github.com/CloudViperViewer/HomeApps/utils"
)

// Structure to define log information
type Log struct {
	Level     int    `json:"level"`      //required 1: Debug, 2: INFO, 3: WARN, 4: ERROR, 5: FATAL
	Message   string `json:"message"`    //required
	Service   string `json:"service"`    //required
	RequestID int    `json:"request_id"` // required
	TimeStamp string `json:"timestamp"`
	Metadata  string `json:"metadata"` //optional
}

// Log levels constants
const (
	LevelDebug = 1
	LevelInfo  = 2
	LevelWarn  = 3
	LevelError = 4
	LevelFatal = 5
)

// Log format template
const logEntryTemplate = "\n--------------\n%s\n[%s]\n%s\n%s"

// Log message types
const (
	logTypeConsole = "console"
	logTypeFile    = "file"
)

// Used to writ the log to the required document and print to terminal
// Log information
func WriteLog(logIn Log) error {

	var level string = getLevel(logTypeConsole, logIn.Level)
	var err error

	log.Printf(logEntryTemplate, level, logIn.Service, logIn.TimeStamp, logIn.Message)

	// TODO: Implement persistence to file
	// 1. Writing to a rotating log file
	err = writeFile(logIn)

	return err
}

// Get Level
// - level: level of the message
func getLevel(messageType string, level int) string {

	//match integer to message level for console
	if messageType == logTypeConsole {
		switch level {
		case LevelDebug:
			return "\033[36m[DEBUG]\033[0m "
		case LevelInfo:
			return "\033[32m[INFO]\033[0m "
		case LevelWarn:
			return "\033[33m[WARN]\033[0m "
		case LevelError:
			return "\033[31m[ERROR]\033[0m "
		case LevelFatal:
			return "\033[35m[FATAL]\033[0m "
		default:
			return "\033[36m[DEBUG]\033[0m "
		}
	}

	//match integer to message level for file
	switch level {
	case LevelDebug:
		return "[DEBUG]"
	case LevelInfo:
		return "[INFO]"
	case LevelWarn:
		return "[WARN]"
	case LevelError:
		return "[ERROR]"
	case LevelFatal:
		return "[FATAL]"
	default:
		return "[DEBUG]"
	}

}

/* Writes the log message to a file
 * log struct containing the passed log details
 * Message to write to log
 */
func writeFile(logIn Log) error {

	var file *os.File
	var serviceIndex int
	var level string = getLevel(logTypeFile, logIn.Level)
	var err error

	//Find index for service
	mu.RLock()
	serviceIndex = utils.IndexOf(logIn.Service, serviceList)
	mu.RUnlock()

	if serviceIndex == -1 {
		return fmt.Errorf("failed to write to log could not find file for service %s", logIn.Service)
	}

	//Get file
	mu.RLock()
	file = files[serviceIndex].file
	mu.RUnlock()

	if file == nil {
		return fmt.Errorf("log file for service %s is nil", logIn.Service)
	}

	mu.Lock()
	_, err = file.WriteString(fmt.Sprintf(logEntryTemplate, level, logIn.Service, logIn.TimeStamp, logIn.Message))
	mu.Unlock()
	return err

}
