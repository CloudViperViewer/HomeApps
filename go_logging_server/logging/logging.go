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
	"log"
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

// Used to writ the log to the required document and print to terminal
// Log information
func WriteLog(logIn Log) error {

	var level string = getLevel(logIn.Level)

	log.Println(level + "[" + logIn.Service + "] " + logIn.Message)

	// TODO: Implement persistence to file
	// 1. Writing to a rotating log file

	return nil
}

// Get Level
func getLevel(level int) string {

	//match integer to message level
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
