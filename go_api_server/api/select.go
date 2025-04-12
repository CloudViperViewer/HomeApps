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
 */

package api

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/CloudViperViewer/HomeApps/go_api_server/database"
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

	err = c.ShouldBindJSON(&selectQ)

	log.Println(selectQ)

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
