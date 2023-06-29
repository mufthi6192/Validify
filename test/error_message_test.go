package test

import (
	"Go-Validation/validator"
	"fmt"
	"testing"
)

func TestErrorMessage(t *testing.T) {

	type data struct {
		Email      string `validation:"email"`
		NewEmail   string `validation:"email"`
		OlderEmail string `validation:"email"`
	}

	testData := data{
		Email:      "test@gmail.com",
		NewEmail:   "testnew",
		OlderEmail: "testold",
	}

	validation := validator.NewValidator()

	err := validation.Validate(testData)

	if err != nil {
		firstError := validation.GetFirstError()
		getError := validation.GetErrorMessage("OlderEmail")

		if firstError != "failed! NewEmail must be a valid email" {
			t.Fatalf("Expected error message '%s' to be 'failed! Email must be a valid email'", firstError)
		}

		if getError != "failed! OlderEmail must be a valid email" {
			t.Fatalf("Expected error message '%s' to be 'failed! Email must be a valid email'", getError)
		}
	} else {
		t.Fatalf("Validation should be error on TestErrorMessage")
	}

	fmt.Println("PASSED")

}
