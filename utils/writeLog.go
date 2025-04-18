/*
 * Holds utility functions to write to log
 */

/*
* Package Components:

* Functions:
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
	"net/http"
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
	ServiceDatabase = "Mariadb"
)

// Creates a Debug Log
func LogDebug(service string, msg string, args ...any) {

	var serverUrl string = fmt.Sprintf("%s%s", GetLogServerUrl(), "log")
	var err error
	var response *http.Response
	var data map[string]any

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	data = map[string]any{
		"level":      levelDebug,
		"message":    fmt.Sprintf(msg, args...),
		"service":    service,
		"request_id": 2,
		"metadata":   ""}

	json, err := json.Marshal(data)

	if err != nil {
		log.Println("failed to write to log: " + err.Error())
	}

	response, err = client.Post(serverUrl, "application/json", bytes.NewReader(json))

	if response == nil || response.StatusCode != http.StatusOK || err != nil {
		if err != nil {
			log.Println("failed to write to log api failed: " + err.Error())
			return
		}

		log.Println("failed to write to log api failed")

	}

}

// Creates a Info Log
func LogInfo(service string, msg string, args ...any) {

	var serverUrl string = fmt.Sprintf("%s%s", GetLogServerUrl(), "log")
	var err error
	var response *http.Response
	var data map[string]any

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	data = map[string]any{
		"level":      levelInfo,
		"message":    fmt.Sprintf(msg, args...),
		"service":    service,
		"request_id": 2,
		"metadata":   ""}

	json, err := json.Marshal(data)

	if err != nil {
		log.Println("failed to write to log: " + err.Error())
	}

	response, err = client.Post(serverUrl, "application/json", bytes.NewReader(json))

	if response == nil || response.StatusCode != http.StatusOK || err != nil {
		if err != nil {
			log.Println("failed to write to log api failed: " + err.Error())
			return
		}

		log.Println("failed to write to log api failed")

	}

}

// Creates a Warn Log
func LogWarn(service string, msg string, args ...any) {

	var serverUrl string = fmt.Sprintf("%s%s", GetLogServerUrl(), "log")
	var err error
	var response *http.Response
	var data map[string]any

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	data = map[string]any{
		"level":      levelWarn,
		"message":    fmt.Sprintf(msg, args...),
		"service":    service,
		"request_id": 2,
		"metadata":   ""}

	json, err := json.Marshal(data)

	if err != nil {
		log.Println("failed to write to log: " + err.Error())
	}

	response, err = client.Post(serverUrl, "application/json", bytes.NewReader(json))

	if response == nil || response.StatusCode != http.StatusOK || err != nil {
		if err != nil {
			log.Println("failed to write to log api failed: " + err.Error())
			return
		}

		log.Println("failed to write to log api failed")

	}

}

// Creates a Warn Log
func LogError(service string, msg string, args ...any) {

	var serverUrl string = fmt.Sprintf("%s%s", GetLogServerUrl(), "log")
	var err error
	var response *http.Response
	var data map[string]any

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	data = map[string]any{
		"level":      levelError,
		"message":    fmt.Sprintf(msg, args...),
		"service":    service,
		"request_id": 2,
		"metadata":   ""}

	json, err := json.Marshal(data)

	if err != nil {
		log.Println("failed to write to log: " + err.Error())
	}

	response, err = client.Post(serverUrl, "application/json", bytes.NewReader(json))

	if response == nil || response.StatusCode != http.StatusOK || err != nil {
		if err != nil {
			log.Println("failed to write to log api failed: " + err.Error())
			return
		}

		log.Println("failed to write to log api failed")

	}

}

// Creates a Warn Log
func LogFatal(service string, msg string, args ...any) {

	var serverUrl string = fmt.Sprintf("%s%s", GetLogServerUrl(), "log")
	var err error
	var response *http.Response
	var data map[string]any

	var client = &http.Client{
		Timeout: 5 * time.Second,
	}

	data = map[string]any{
		"level":      levelFatal,
		"message":    fmt.Sprintf(msg, args...),
		"service":    service,
		"request_id": 2,
		"metadata":   ""}

	json, err := json.Marshal(data)

	if err != nil {
		log.Println("failed to write to log: " + err.Error())
	}

	response, err = client.Post(serverUrl, "application/json", bytes.NewReader(json))

	if response == nil || response.StatusCode != http.StatusOK || err != nil {
		if err != nil {
			log.Println("failed to write to log api failed: " + err.Error())
			return
		}

		log.Println("failed to write to log api failed")

	}

	os.Exit(1)

}
