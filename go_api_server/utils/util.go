/*
 * Defines the utility functions for the entire application
 */

/*
* Package Components:

* Functions:
* - GetAllTags: gets the tags for a passed struct
* - JoinArray concatenates a slice of strings together with a delimiter
 */

package utils

import (
	"fmt"
	"reflect"
)

// Gets All Tags for the passed struct
//
//   - Input defined as any but should always be a struct
//   - String passed as tag name
//   - Return slice of strings
func GetAllTags(structure any, tagName string) []string {

	/*Define return variable*/
	var tagValues []string
	/*Get struct type*/
	var structType reflect.Type = reflect.TypeOf(structure)

	/*Loop over fields and get specific tags*/
	for i := 0; i < structType.NumField(); i++ {
		/*get tag for field*/
		field := structType.Field(i)
		tag := field.Tag.Get(tagName)

		/*add tag to slice*/
		tagValues = append(tagValues, tag)

	}

	/*Return values*/
	return tagValues

}

// Takes a list of string and concatenates them with a delimiter
//
//   - array list of times to concatenate
//   - delimiter string to use as seperator
//   - Returns the string from connected arrat
func JoinArray[T string | any](array []T, delimiter string) string {

	var connectedString string

	//Loops over list and concats to string
	for i := range array {

		connectedString = fmt.Sprintf("%s%v", connectedString, array[i])

		//if not end of arracy concat delimiter
		if i != len(array)-1 {
			connectedString = connectedString + delimiter
		}
	}

	return connectedString

}
