/*
 *Holds utilities to use across all go API's
 */

/*
* Package Components:
*
*
* Functions:
* - IsBodyEmpty: Checks if the api body is empty
 */

package apiutilities

import (
	"fmt"
	"io"
)

// Checks if the api body is empty
// - body of the api call
// - error from reading the body
func IsBodyEmpty(body []byte, err error) error {

	//checks if body empty
	if err == io.EOF || len(body) <= 0 {
		return fmt.Errorf("body cannot be empty")

	}

	//Checks if error occured or body empty
	if err != nil {
		return fmt.Errorf("failed to read request body: %s", err.Error())

	}

	return nil
}
