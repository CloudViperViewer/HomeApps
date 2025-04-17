/*
 * Holds functions related to health endpoint
 */

/*
* Package Components:
*
*
* Functions:
* - IsRunning: returns wheather the server is running
 */

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// returns if the server is running
func IsRunning(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "running"})
}
