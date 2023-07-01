package test

import (
	"Go-Validation/validator"
	"fmt"
	"testing"
)

func TestSameValidation(t *testing.T) {

	type data struct {
		Email        string `validation:"email"`
		ConfirmEmail string `validation:"email|same:Email"`
	}

	testData := data{
		Email:        "new@test.com",
		ConfirmEmail: "new@testa.com",
	}

	v := validator.NewValidator()
	err := v.Validate(testData)

	if err != nil {
		errMsg := v.GetFirstError()
		t.Fatalf(errMsg)
	}

	fmt.Println("OK")

}
