package validator

import (
	"fmt"
	"reflect"
	"strings"
)

func NewValidator() *Validator {
	return &Validator{
		rules:          make(map[string]string),
		messages:       make(map[string]string),
		customMessages: make(map[string]string),
	}
}

func MakeStruct(data interface{}) error {
	validator := NewValidator()
	return validator.Validate(data)
}

func (v *Validator) Validate(data interface{}) error {
	v.errors = make(map[string]string)
	v.errorFields = make([]string, 0)

	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a struct")
	}

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		value := rv.Field(i)
		tag := field.Tag.Get("validation")

		if tag != "" {
			rules := strings.Split(tag, "|")
			var vp ValidationPart

			for _, rule := range rules {

				if strings.Contains(rule, ":") {
					keyword := strings.Split(rule, ":")[1]
					if !strings.Contains(keyword, ",") {
						compareVal, err := getFieldByKeyword(data, keyword)
						if err != nil {
							return err
						}
						vp = ValidationPart{
							Rule:  rule,
							field: field.Name,
							value: value.Interface(),
							ComparisonPart: ComparisonPart{
								field: keyword,
								value: compareVal,
							},
						}
					}
				} else {
					vp = ValidationPart{
						Rule:  rule,
						field: field.Name,
						value: value.Interface(),
					}
				}

				err := vp.Validate(v)
				if err != nil {
					if customMessage, ok := v.customMessages[field.Name]; ok {
						v.errors[field.Name] = customMessage
					} else {
						v.errors[field.Name] = err.Error()
					}
					v.errorFields = append(v.errorFields, field.Name)
				}
			}
		}
	}

	if len(v.errors) > 0 {
		return fmt.Errorf("validation failed")
	}

	return nil
}

func getFieldByKeyword(data interface{}, keyword string) (interface{}, error) {

	rt := reflect.TypeOf(data)
	rv := reflect.ValueOf(data)

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Name == keyword {
			value := rv.Field(i).Interface()
			return value, nil
		}
	}

	return nil, fmt.Errorf("field with keyword '%s' not found", keyword)
}
