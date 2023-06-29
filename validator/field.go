package validator

func (v *Validator) GetErrorFields() []string {

	return v.errorFields

}

func (v *Validator) GetFirstField() string {
	for _, value := range v.errorFields {
		return value
	}
	return ""
}
