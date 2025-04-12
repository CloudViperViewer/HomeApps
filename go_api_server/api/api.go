/*
 * Sets up the api server listening for requests
 */

/*
* Package Components:
*
*
* Functions:
* StartUpServer: Starts up the api server
* setupEndPoints: Setup endpoints
* dbQuerySelect: Handles db select endpoint
 */

package api

import (
	"github.com/gin-gonic/gin"
)

// Starts up the api server
func StartUpServer() {

	/*Setup router*/
	router := gin.Default()

	//Initial endpoint setup
	setupEndPoints(router)

	/*Run router*/
	router.Run("0.0.0.0:8080")
}

// Setups up endpoints
func setupEndPoints(router *gin.Engine) {

	router.POST("/api/select", dbQuerySelect)
}
