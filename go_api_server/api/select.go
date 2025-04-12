/*
 * Holds functions related to db select end point
 */

/*
* Package Components:
*
*
* Functions:
* dbQuerySelect: Handles db select endpoint
 */

package api

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handles db select endpoint
func dbQuerySelect(c *gin.Context) {

	var body []byte
	var err error

	body, err = io.ReadAll(c.Request.Body)

	if err != nil || len(body) <= 0 {
		if err == io.EOF || len(body) <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "body cannot be empty",
			})

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "bad request",
				"message": "failed to read request body",
			})
			return
		}

		return

	}

	log.Println(body)
	log.Println(err)
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
