package test

import (
	"Go-Validation/validator"
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {
	tests := []struct {
		Email        string
		ShouldPass   bool
		ErrorMessage string
	}{
		// Valid email addresses
		{"mufthi@example.com", true, ""},
		{"mufthi123@example.com", true, ""},
		{"mufthi.m@example.com", true, ""},
		{"mufthi-321@example.com", true, ""},
		{"mufthi@example.co", true, ""},
		{"mufthi@example.co.uk", true, ""},
		{"mufthi@subdomain.example.com", true, ""},
		{"mufthi@sub.sub.sub.example.com", true, ""},

		// Invalid email addresses
		{"mufthi", false, "failed! Email must be a valid email"},
		{"mufthi@example", false, "failed! Email must be a valid email"},
		{"mufthi@.com", false, "failed! Email must be a valid email"},
		{"mufthi@example..com", false, "failed! Email must be a valid email"},
		{"mufthi@-example.com", false, "failed! Email must be a valid email"},
		{"mufthi@example-.com", false, "failed! Email must be a valid email"},
		{"mufthi@[123.456.789.0", false, "failed! Email must be a valid email"},
		{"mufthi@[IPv6:2001:db8::1", false, "failed! Email must be a valid email"},

		// Additional test cases
		{"mufthi@example.com.", false, "failed! Email must be a valid email"},
		{"mufthi@example_com", false, "failed! Email must be a valid email"},
		{"mufthi@123.456.789", false, "failed! Email must be a valid email"},
		{"mufthi@123.456.789.256", false, "failed! Email must be a valid email"},
		{"mufthi@domain", false, "failed! Email must be a valid email"},
		{"mufthi@example.c", false, "failed! Email must be a valid email"},
		{"mufthi@example.12", false, "failed! Email must be a valid email"},
		{"mufthi@example.12a", false, "failed! Email must be a valid email"},
	}

	v := validator.NewValidator()

	for _, test := range tests {
		testData := struct {
			Email string `validation:"email"`
		}{Email: test.Email}

		err := v.Validate(testData)
		errMsg := v.GetErrorMessage("Email")

		if test.ShouldPass {
			if err != nil {
				t.Fatalf("Expected email '%s' to pass validation, but got error: %v\n", test.Email, errMsg)
			} else {
				fmt.Println("Validation passed")
			}
		} else {
			if err == nil {
				t.Fatalf("Expected email '%s' to fail validation, but got no error\n", test.Email)
			} else {
				errMsg := v.GetErrorMessage("Email")
				if errMsg != test.ErrorMessage {
					t.Fatalf("Expected error message '%s', but got '%s'\n", test.ErrorMessage, errMsg)
				}
			}
		}
	}

}
