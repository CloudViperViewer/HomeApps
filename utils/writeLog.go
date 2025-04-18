/*
 * Holds utility functions to write to log
 */

/*
* Package Components:

* Functions:
* - getRequestId: Generates a random number for the request
* - logMessageHandler: Handles standard log message actions
* - LogDebug: Writes a log message for debug
* - LogInfo: Writes a log message for info
* - LogWarn: Writes a log message for warn
* - LogError: Writes a log message for error
* - LogFatal: Writes a log message for fatal
 */

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Consts of log levels
const (
	levelDebug = 1
	levelInfo  = 2
	levelWarn  = 3
	levelError = 4
	levelFatal = 5
)

// Services constants
const (
	ServiceDatabaseApi = "DatabaseApi"
)

// rand number max
const randMax = 10000

// Generates a random number for the request
func getRequestId() int {

	return rand.Intn(randMax)
}

/*Handles standard log message actions
* - level: level of the log message
* - service: service log relates to
* - msg: message for the log
* - args: arguments for the message
 */
func logMessageHandler(level int, service string, metadata string, msg string, args ...any) {
	var serverUrl string = fmt.Sprintf("%s%s", GetLogServerUrl(), "log")
	var parsedUrl *url.URL
	var err error
	var response *http.Response
	var data map[string]any

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	//check url
	parsedUrl, err = url.Parse(serverUrl)
	if err != nil {
		log.Printf("invalid log server URL %q: %v", serverUrl, err)
		return
	}

	//setup data
	data = map[string]any{
		"level":      level,
		"message":    fmt.Sprintf(msg, args...),
		"service":    service,
		"request_id": getRequestId(),
		"metadata":   metadata}

	//conver to json
	json, err := json.Marshal(data)
	if err != nil {
		log.Println("failed to write to log: " + err.Error())
		return
	}

	//send response
	response, err = client.Post(parsedUrl.String(), "application/json", bytes.NewReader(json))

	//evaluate logging server response
	if response == nil || response.StatusCode != http.StatusOK || err != nil {
		if err != nil {
			log.Println("failed to write to log api failed: " + err.Error())
			//close body
			if response != nil && response.Body != nil {
				response.Body.Close()
			}
			return
		}

		log.Println("failed to write to log api failed")
		//close body
		if response != nil && response.Body != nil {
			response.Body.Close()
		}
	}

	//close body
	if response != nil && response.Body != nil {
		response.Body.Close()
	}
}

/* Creates a Debug Log
* - service: service log relates to
* - msg: message for the log
* - args: arguments for the message
 */
func LogDebug(service string, metadata string, msg string, args ...any) {

	//call handler
	logMessageHandler(levelDebug, service, metadata, msg, args...)
}

/* Creates a Info Log
* - service: service log relates to
* - msg: message for the log
* - args: arguments for the message
 */
func LogInfo(service string, metadata string, msg string, args ...any) {

	//call handler
	logMessageHandler(levelInfo, service, metadata, msg, args...)

}

/* Creates a Warn Log
* - service: service log relates to
* - msg: message for the log
* - args: arguments for the message
 */
func LogWarn(service string, metadata string, msg string, args ...any) {

	//call handler
	logMessageHandler(levelWarn, service, metadata, msg, args...)

}

/* Creates a Warn Log
* - service: service log relates to
* - msg: message for the log
* - args: arguments for the message
 */
func LogError(service string, metadata string, msg string, args ...any) {

	//call handler
	logMessageHandler(levelError, service, metadata, msg, args...)
}

/* Creates a Warn Log
* - service: service log relates to
* - msg: message for the log
* - args: arguments for the message
 */
func LogFatal(service string, metadata string, msg string, args ...any) {

	//call handler
	logMessageHandler(levelFatal, service, metadata, msg, args...)
	//exit application as fatal
	os.Exit(1)

}
