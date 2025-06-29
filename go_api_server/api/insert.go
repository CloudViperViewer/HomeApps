/*
 * Holds functions related to db insert end point
 */

/*
* Package Components:
*
* Structures
* - insert: represents the json struct to be passed in api call
*
* Functions:
* - dbInsert: Handles db insert endpoint
* - insertConfirmData: checks if required data is present in json
 */

package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	apiutilities "github.com/CloudViperViewer/HomeApps/api_utilities"
	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	"github.com/CloudViperViewer/HomeApps/utils"
	"github.com/gin-gonic/gin"
)

type insert struct {
	Table string           `json:"Table"`
	Rows  []map[string]any `json:"Data"`
}

/*Handles db insert endpoint
 * gin context
 */
func dbInsert(c *gin.Context) {
	var body []byte
	var err error
	const maxBodySize = 1 << 20 //1MB
	var insertQ insert

	//Read the body
	body, err = io.ReadAll(io.LimitReader(c.Request.Body, maxBodySize))

	//reset body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	//Check if body is empty
	if err = apiutilities.IsBodyEmpty(body, err); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		utils.LogDebug(utils.ServiceDatabaseApi, "", "base request in insert %s", err.Error())
		return

	}

	//Ensure content is json
	if !strings.HasPrefix(c.GetHeader("Content-Type"), "application/json") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "Content-Type must be application/json",
		})
		utils.LogDebug(utils.ServiceDatabaseApi, "", "base request in insert format not json")
		return
	}

	//parse data
	err = c.ShouldBindJSON(&insertQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "error in parsing json: " + err.Error(),
		})
		utils.LogDebug(utils.ServiceDatabaseApi, "", "base request in insert %s", err.Error())
		return
	}

	//confirm data is correct
	err = insertConfirmData(insertQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		utils.LogDebug(utils.ServiceDatabaseApi, "", "base request in insert %s", err.Error())
		return
	}

	//execute query
	err = database.ExecuteInsertQuery(database.GetDb(), insertQ.Table, insertQ.Rows)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "internal server error",
			"message": err.Error(),
		})
		utils.LogDebug(utils.ServiceDatabaseApi, "", "base request in insert %s", err.Error())
		return
	}

	//log successful insert
	utils.LogInfo(utils.ServiceDatabaseApi, "", "Insert Successful")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Insert successful",
		"rows":    len(insertQ.Rows),
	})

}

/*Confirm the data passed is complete
 * insert - data passed from from the json
 * return - error
 */
func insertConfirmData(insertQ insert) error {

	if insertQ.Table == "" {
		return fmt.Errorf("table is missing")
	}

	if len(insertQ.Rows) == 0 {
		return fmt.Errorf("no data to insert")
	}

	return nil
}
