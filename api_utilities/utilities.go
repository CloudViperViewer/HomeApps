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

func IsBodyEmpty(body []byte, err error) error {

	//Checks if error occured or body empty
	if err != nil || len(body) <= 0 {
		if err == io.EOF || len(body) <= 0 {
			return fmt.Errorf("body cannot be empty")

		} else {
			return fmt.Errorf("failed to read request body: %s", err.Error())

		}

	}
	return nil
}
