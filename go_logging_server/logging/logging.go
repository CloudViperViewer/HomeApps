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
)

// Structure to define log information
type Log struct {
	Level     int    `json:"level"`      //required 1: Debug, 2: INFO, 3: WARN, 4: ERROR, 5: FATAL
	Message   string `json:"message"`    //required
	Service   string `json:"service"`    //required
	RequestID int    `json:"request_id"` // required
	Metadata  string `json:"metadata"`   //optional
}

// Used to writ the log to the required document and print to terminal
// Log information
func WriteLog(logIn Log) {

	var level string = getLevel(logIn.Level)

	log.Println(level + "[" + logIn.Service + "] " + logIn.Message)

}

// Get Level
func getLevel(level int) string {

	//match integer to message level
	switch level {
	default:
		return fmt.Sprintln("\033[DEBUG]\033[36m ")
	}
}
