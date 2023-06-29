package validator

import (
	"Go-Validation/validator/validation"
	"fmt"
)

func (vp *ValidationPart) Validate(v *Validator) error {

	basicValidation := []string{
		"email", "accepted", "active_url", "array",
	}
	isBasic := false

	for _, vb := range basicValidation {
		if vb == vp.Rule {
			isBasic = true
			break
		}
	}

	if isBasic {
		rule := createValidationBasic(vp.Rule, vp.value, vp.field)
		if rule == nil {
			return fmt.Errorf("failed! Rule '%s' not found", vp.Rule)
		}
		return rule.Validate()
	}

	return fmt.Errorf("failed! Rule not found")
}

func createValidationBasic(rule string, value interface{}, field string) ValidationRule {
	switch rule {
	case "email":
		return &validation.EmailValidation{Value: value, Field: field}
	case "accepted":
		return &validation.AcceptedValidation{Value: value, Field: field}
	case "active_url":
		return &validation.ActiveUrlValidation{Value: value, Field: field}
	case "array":
		return &validation.ArrayValidation{Value: value, Field: field}
	default:
		return nil
	}
}
