package test

import (
	"Go-Validation/validator"
	"fmt"
	"testing"
)

func TestArrayValidation(t *testing.T) {
	tests := []struct {
		Value        interface{}
		ShouldPass   bool
		ErrorMessage string
	}{
		// Valid arrays
		{[]int{1, 2, 3}, true, ""},
		{[]string{"apple", "banana", "orange"}, true, ""},
		{[3]int{1, 2, 3}, true, ""},

		// Invalid arrays
		{"not an array", false, "failed! Field must be an array"},
		{123, false, "failed! Field must be an array"},
		{map[string]int{"key": 123}, false, "failed! Field must be an array"},
	}

	v := validator.NewValidator()

	for _, test := range tests {
		testData := struct {
			Value interface{} `validation:"array"`
		}{Value: test.Value}

		err := v.Validate(testData)
		errMsg := v.GetErrorMessage("Value")

		if test.ShouldPass {
			if err != nil {
				t.Fatalf("Expected value '%v' to pass validation, but got error: %v\n", test.Value, errMsg)
			} else {
				fmt.Println("Validation passed")
			}
		} else {
			if err == nil {
				t.Fatalf("Expected value '%v' to fail validation, but got no error\n", test.Value)
			} else {
				if errMsg != test.ErrorMessage {
					t.Fatalf("Expected error message '%s', but got '%s'\n", test.ErrorMessage, errMsg)
				}
			}
		}
	}
}
