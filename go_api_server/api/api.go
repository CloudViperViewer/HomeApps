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

// StartUpServer initializes the API server by creating a default Gin router, configuring API endpoints via setupEndPoints, and starting the server to listen on all interfaces at port 8080.
func StartUpServer() {

	/*Setup router*/
	router := gin.Default()

	//Initial endpoint setup
	setupEndPoints(router)

	/*Run router*/
	router.Run("0.0.0.0:8080")
}

// setupEndPoints registers the POST "/api/select" endpoint on the provided Gin router to handle database selection requests using dbQuerySelect.
func setupEndPoints(router *gin.Engine) {

	router.POST("/api/select", dbQuerySelect)
}
