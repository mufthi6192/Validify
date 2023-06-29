package validation

import (
	"fmt"
	"reflect"
)

type ArrayValidation struct {
	Value interface{}
	Field string
}

// Validate performs the array validation.
func (av *ArrayValidation) Validate() error {
	value := reflect.ValueOf(av.Value)
	if value.Kind() != reflect.Slice && value.Kind() != reflect.Array {
		return fmt.Errorf("failed! Field must be an array")
	}

	return nil
}
