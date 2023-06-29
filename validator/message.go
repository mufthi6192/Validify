package validator

func (v *Validator) GetFirstError() string {
	for _, value := range v.errors {
		return value
	}
	return ""
}

func (v *Validator) GetErrorMessage(field string) string {
	return v.errors[field]
}

func (v *Validator) GetErrors() map[string]string {
	return v.errors
}

func (v *Validator) SetCustomMessage(field string, message string) {
	getField := v.GetErrorFields()

	status := false
	for _, value := range getField {
		if value == field {
			status = true
		}
	}

	if status == false {
		panic("Failed ! Field not found")
	}
	v.customMessages[field] = message
}
