package test

import (
	"Go-Validation/validator"
	"fmt"
	"testing"
)

func TestActiveUrlValidation(t *testing.T) {
	tests := []struct {
		URL          string
		ShouldPass   bool
		ErrorMessage string
	}{
		// Valid URLs
		{"https://example.com", true, ""},
		{"http://subdomain.example.com", true, ""},
		{"https://www.example.co.uk", true, ""},

		// Invalid URLs
		{"example.com", false, "failed! Invalid URL: example.com"},
		{"http://", false, "failed! Invalid hostname: "},
		{"https://invalid.123", false, "failed! DNS lookup failed for hostname: invalid.123"},
		{"http://localhost", false, "failed! No valid A or AAAA record found for hostname: localhost"},
	}

	v := validator.NewValidator()

	for _, test := range tests {
		testData := struct {
			URL string `validation:"active_url"`
		}{URL: test.URL}

		err := v.Validate(testData)
		errMsg := v.GetErrorMessage("Email")

		if test.ShouldPass {
			if err != nil {
				t.Fatalf("Expected URL '%s' to pass validation, but got error: %v\n", test.URL, errMsg)
			} else {
				fmt.Println("Validation passed")
			}
		} else {
			if err == nil {
				t.Fatalf("Expected URL '%s' to fail validation, but got no error\n", test.URL)
			} else {
				fmt.Println("Validation PASSED")
			}
		}
	}
}
