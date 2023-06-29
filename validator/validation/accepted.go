package validation

import (
	"fmt"
)

type AcceptedValidation struct {
	Value interface{}
	Field string
}

func (av *AcceptedValidation) Validate() error {
	value := av.Value

	switch value.(type) {
	case string:
		return stringValidation(value, av.Field)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return integerValidation(value, av.Field)
	case bool:
		return boolValidation(value, av.Field)
	default:
		return fmt.Errorf("failed! %s must be (Yes, on, 1, or true)", av.Field)
	}
}

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
