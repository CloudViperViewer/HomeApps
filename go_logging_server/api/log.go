/*
 * Holds methods for the log endpoint
 */

/*
* Package Components:
*
*
* Functions:
* - log main logging endpoint of the log server
 */

package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	apiutilities "github.com/CloudViperViewer/HomeApps/api_utilities"
	"github.com/CloudViperViewer/HomeApps/go_logging_server/logging"
	"github.com/gin-gonic/gin"
)

// Main function exposed by logging end point
// - context for the api call
func handleLogRequest(c *gin.Context) {

	var body []byte
	var err error
	const maxBodySize = 1 << 20 //1MB
	var log logging.Log

	//read the body
	body, err = io.ReadAll(io.LimitReader(c.Request.Body, maxBodySize))

	//reset body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	//check if body empty
	if err = apiutilities.IsBodyEmpty(body, err); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	//Ensure content is json
	if !strings.HasPrefix(c.GetHeader("Content-Type"), "application/json") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "Content-Type must be application/json",
		})
		return
	}

	//parse json
	err = c.ShouldBindJSON(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "error in parsing json: " + err.Error(),
		})
		return
	}

	//check for missing data
	if err = confirmData(log); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	if err = logging.WriteLog(log); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": err.Error()})
		return
	}

	//return response
	c.JSON(http.StatusOK, gin.H{
		"success": true})

}

// Confirms the passed json has all required data
// - log passed by json
func confirmData(log logging.Log) error {

	var missingData []string

	//level missing or incorrect
	if log.Level < 1 || log.Level > 5 {
		missingData = append(missingData, "level missing or incorrect level")
	}

	//Check message
	if log.Message == "" {
		missingData = append(missingData, "message is missing")
	}

	//Check request id
	if log.RequestID <= 0 {
		missingData = append(missingData, "request Id cannot be empty or less than or equal to 0")
	}

	//check service
	if log.Service == "" {
		missingData = append(missingData, "service is missing")
	}

	//data is missing
	if len(missingData) > 0 {
		return fmt.Errorf("json missing required fields %s", strings.Join(missingData, ", "))
	}

	return nil

}
