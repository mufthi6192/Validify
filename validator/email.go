package validator

import (
	"fmt"
	"regexp"
	"strings"
)

type EmailValidation struct {
	value interface{}
	field string
}

// Validate performs the email validation.
func (ev *EmailValidation) Validate() error {
	email := ev.value.(string)
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(emailRegex, email)
	if err != nil {
		return fmt.Errorf("failed! %s must be a valid email", ev.field)
	}
	if !match || strings.HasPrefix(email, "-") || strings.HasPrefix(email, ".") || strings.HasSuffix(email, ".") || strings.Contains(email, "..") ||
		strings.Contains(email, "@.") || strings.Contains(email, ".@") || strings.Contains(email, "@-") || strings.Contains(email, "-@") ||
		strings.Contains(email, "-.") || strings.Contains(email, ".-") || len(email) > 254 {
		return fmt.Errorf("failed! %s must be a valid email", ev.field)
	}

	return nil
}
