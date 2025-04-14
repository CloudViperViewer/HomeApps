/*
 * Sets up the api server listening for requests
 */

/*
* Package Components:
*
*
* This package implements API endpoints for database operations
* using the Gin framework for HTTP request handling.
*
*
* Functions:
* - StartUpServer: Starts up the api server
* - setupEndPoints: Setup endpoints
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
	router.Run("0.0.0.0:8090")
}

// Setups up endpoints
func setupEndPoints(router *gin.Engine) {

	//Health endpint
	router.GET("/health", IsRunning)
}
