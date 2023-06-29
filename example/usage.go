package example

import (
	"Go-Validation/validator"
	"fmt"
)

func GetSingleError() {

	type data struct {
		Email          string `validation:"email"`
		TermsOfService bool   `validation:"accepted"`
	}

	v := validator.NewValidator()

	userData := data{
		Email:          "mufthi@testing.com",
		TermsOfService: true,
	}

	//Recommended to use pointer
	// You can use v.Validate(userData) if you don't want to use pointer
	err := v.Validate(&userData)

	if err != nil {
		errMsg := v.GetFirstError()
		fmt.Println(errMsg)
	}

	//Validation passed
	fmt.Println("OK")

}
