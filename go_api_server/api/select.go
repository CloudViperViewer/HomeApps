/*
 * Holds functions related to db select end point
 */

/*
* Package Components:
*
* Structures
* - selectQuery represents the json struct to be passed in api call
*
* Functions:
* - dbQuerySelect: Handles db select endpoint
* - confirmData: confirms the passed data has all required fields
* - queryDb: executes db query
 */

package api

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
	"github.com/CloudViperViewer/HomeApps/go_api_server/tables"
	"github.com/CloudViperViewer/HomeApps/go_api_server/utils"
	"github.com/gin-gonic/gin"
)

// Rpresents the json struct to be passed in api call
type selectQuery struct {
	Table             string                   `json:"table"`
	Fields            []string                 `json:"fields"`
	LogicalExpression database.LogicExpression `json:"logicalExpression"`
	PagingInfo        database.PagingInfo      `json:"pagingInfo"`
}

// Handles db select endpoint
func dbQuerySelect(c *gin.Context) {

	var body []byte
	var err error
	const maxBodySize = 1 << 20 //1MB
	var selectQ selectQuery
	var data tables.Table

	//Read the body
	body, err = io.ReadAll(io.LimitReader(c.Request.Body, maxBodySize))

	//reset body
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	//Check if body is empty
	if err != nil || len(body) <= 0 {
		if err == io.EOF || len(body) <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "body cannot be empty",
			})
			return

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "failed to read request body: " + err.Error(),
			})
			return

		}

	}

	//Ensure content is json
	if c.GetHeader("Content-Type") != "application/json" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "Content-Type must be application/json",
		})
		return
	}

	//parse data
	err = c.ShouldBindJSON(&selectQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": "error in parsing json: " + err.Error(),
		})
		return
	}

	//Confirm data is there
	err = confirmData(selectQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	//Query db
	data, err = queryDb(selectQ)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "bad request",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ACCEPTED", "data": data.GetRows()})
}

// Confirms the passed meets requirments
func confirmData(selectQ selectQuery) error {

	var missingData []string

	//table required
	if selectQ.Table == "" {
		missingData = append(missingData, "table missing")
	}

	//startIndex missing
	if selectQ.PagingInfo.StartIndex == 0 {
		missingData = append(missingData, "start index cannot be 0 or empty")
	}

	//batchSize missing
	if selectQ.PagingInfo.BatchSize == 0 {
		missingData = append(missingData, "batch size must be greater than 0 or -1")
	}

	//data correct
	if missingData == nil {
		return nil
	}

	return fmt.Errorf("missing or incorrect data: %s", utils.JoinArray(missingData, ", "))

}

// Calls functions to query the db
func queryDb(selectQ selectQuery) (tables.Table, error) {

	var SQuery database.SelectQuery = database.SelectQuery{
		PagingInfo:      selectQ.PagingInfo,
		Fields:          selectQ.Fields,
		LogicExpression: selectQ.LogicalExpression,
	}
	var err error
	var data tables.Table

	//Get table
	SQuery.Table, err = tables.TableFactory(selectQ.Table)

	//Table factory failed
	if err != nil {
		return nil, err
	}

	data, err = database.ExecuteSelectQuery(database.GetDb(), SQuery)

	//Query failed
	if err != nil {
		return nil, err
	}

	return data, nil

}
