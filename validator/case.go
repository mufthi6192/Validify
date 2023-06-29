package validator

import "fmt"

func (vp *ValidationPart) Validate(v *Validator) error {
	rule := createValidationRule(vp.Rule, vp.value, vp.field)
	if rule == nil {
		return fmt.Errorf("failed! Rule '%s' not found", vp.Rule)
	}
	return rule.Validate()
}

func createValidationRule(rule string, value interface{}, field string) ValidationRule {
	switch rule {
	case "email":
		return &EmailValidation{value: value, field: field}
	case "accepted":
		return &AcceptedValidation{value: value, field: field}
	default:
		return nil
	}
}
