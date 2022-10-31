package rest

import "log"

// map error code to error message
var errorMessage map[string]string

// Register a new error
func RegisterError(code string, message string) {
	_, ok := errorMessage[code]
	if ok {
		log.Fatalln("Error code has already been registered")
	}
	errorMessage[code] = message
}

// Get error message
func ErrorMessage(code string) string {
	return errorMessage[code]
}
