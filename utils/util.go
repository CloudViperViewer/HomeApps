/*
 * Defines the utility functions for the entire application
 */

/*
* Package Components:

* Functions:
* - GetLogServerUrl: returns the url of the log server
* - GetStructAllFieldPtrs: Gets a slice of points for all fields in the struct
* - GetStructFieldPtrs: Gets Fields of the passed struct
* - GetTagList: Gets Tags for Specific fields in struct
* - GetAllTags: gets the tags for a passed struct
* - JoinArray_Deprecated: concatenates a slice of strings together with a delimiter
 */

package utils

import (
	"fmt"
	"os"
	"reflect"
)

// Gets the url for the logging server
func GetLogServerUrl() string {
	//Get host
	var serverHost string = os.Getenv("LOG_SERVER_HOST")
	if serverHost == "" {
		serverHost = "localhost"
	}

	//get port
	var serverPort string = os.Getenv("LOG_SERVER_PORT")

	if serverPort == "" {
		serverPort = "8090"
	}

	return fmt.Sprintf("http://%s:%s/", serverHost, serverPort)

}

// Gets a slice of points for all fields in the struct
func GetStructAllFieldPtrs(structure any) ([]any, error) {

	var ptrArray []any

	structFields := reflect.ValueOf(structure)

	// Keep unwrapping interface{} layers
	for structFields.Kind() == reflect.Interface {
		structFields = structFields.Elem()
	}

	//Check struct is the correct type
	if structFields.Kind() != reflect.Ptr || structFields.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected pointer to struct, got %T", structure)
	}

	if structFields.Kind() == reflect.Ptr {
		structFields = structFields.Elem()
	}

	//loop over fields
	for i := 0; i < structFields.NumField(); i++ {

		field := structFields.Field(i)

		if field.CanAddr() {
			ptrArray = append(ptrArray, field.Addr().Interface())
		}

	}

	return ptrArray, nil

}

// Gets Fields of the passed struct
// - Structure to get the fields from
// - List of fields you whish to get
func GetStructFieldPtrs(structure any, fields []string) ([]any, error) {

	var prtArray []any

	structFields := reflect.ValueOf(structure)

	// Keep unwrapping interface{} layers
	for structFields.Kind() == reflect.Interface {
		structFields = structFields.Elem()
	}

	//Chec struct is the correct type
	if structFields.Kind() != reflect.Ptr || structFields.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("expected pointer to struct, got %T", structure)
	}

	values := structFields.Elem()

	//Loop of fields to get
	for _, fieldName := range fields {

		field := values.FieldByName(fieldName)

		//Check if field is valid
		if !field.IsValid() {
			return nil, fmt.Errorf("no such filed %s in %T", fieldName, structure)

		}

		if !field.CanAddr() {
			return nil, fmt.Errorf("field %s is not addressable (unexported?)", fieldName)
		}

		prtArray = append(prtArray, field.Addr().Interface())

	}

	return prtArray, nil

}

// Gets Tags for Specific fields in struct
// - List struct to get fields from
// - List of fields wanted
// - Name of tag
func GetTagList(structure any, fields []string, tagName string) ([]string, error) {

	// Define variable to return
	var tagValues []string

	// Get struct type
	var structType reflect.Type = reflect.TypeOf(structure)

	// If it's a pointer, get the element type
	if structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}

	//Loop over fields
	for _, fieldName := range fields {

		//Get field compare see if field is in passed list of fields then get tag
		field, found := structType.FieldByName(fieldName)

		if !found {
			return nil, fmt.Errorf("error getting tag from field %s", fieldName)
		}

		tagValues = append(tagValues, field.Tag.Get(tagName))
	}

	return tagValues, nil
}

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

// Takes a list of string and concatenates them with a delimiter /Deprecated
//
//   - array list of times to concatenate
//   - delimiter string to use as seperator
//   - Returns the string from connected array
func JoinArray_Deprecated[T string | any](array []T, delimiter string) string {

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
