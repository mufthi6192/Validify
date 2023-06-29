package validator

type Validator struct {
	rules          map[string]string
	messages       map[string]string
	customMessages map[string]string
	errors         map[string]string
	errorFields    []string
}

type ValidationPart struct {
	Rule  string
	field string
	value interface{}
}

type ValidationRule interface {
	Validate() error
}
