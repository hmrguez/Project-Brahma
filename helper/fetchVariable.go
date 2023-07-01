package helper

import "reflect"

func FetchVariable(name string, variables map[string]interface{}) interface{} {
	// Get the value of the named variable using reflection
	value := reflect.ValueOf(variables[name])

	// Check if the value is valid and return it
	if value.IsValid() {
		return value.Interface()
	}

	// Return nil if the variable is not found
	return nil
}
