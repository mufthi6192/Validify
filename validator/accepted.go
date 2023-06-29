package validator

import (
	"fmt"
)

type AcceptedValidation struct {
	value interface{}
	field string
}

// Validate performs the accepted validation.
func (av *AcceptedValidation) Validate() error {
	value := av.value

	switch value.(type) {
	case string:
		return stringValidation(value, av.field)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return integerValidation(value, av.field)
	case bool:
		return boolValidation(value, av.field)
	default:
		return fmt.Errorf("failed! %s must be (Yes, on, 1, or true)", av.field)
	}
}

// stringValidation performs validation for string values.
func stringValidation(value interface{}, field string) error {
	val := value.(string)
	param := []string{"yes", "on"}
	status := false

	for _, element := range param {
		if element == val {
			status = true
			break
		}
	}

	if !status {
		return fmt.Errorf("failed! %s must be (Yes, on, 1, or true)", field)
	}

	return nil
}

// integerValidation performs validation for integer values.
func integerValidation(value interface{}, field string) error {
	val := value.(int)
	param := []int{1}
	status := false

	for _, element := range param {
		if element == val {
			status = true
			break
		}
	}

	if !status {
		return fmt.Errorf("failed! %s must be (Yes, on, 1, or true)", field)
	}

	return nil
}

// boolValidation performs validation for boolean values.
func boolValidation(value interface{}, field string) error {
	val := value.(bool)
	param := []bool{true}
	status := false

	for _, element := range param {
		if element == val {
			status = true
			break
		}
	}

	if !status {
		return fmt.Errorf("failed! %s must be (Yes, on, 1, or true)", field)
	}

	return nil
}
