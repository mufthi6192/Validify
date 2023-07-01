package validation

import "fmt"

type SameValidation struct {
	Value        interface{}
	Field        string
	CompareValue interface{}
	CompareField string
}

func (sv *SameValidation) Validate() error {

	if sv.Value != sv.CompareValue {
		return fmt.Errorf(fmt.Sprintf("Failed ! %s must be same with %s", sv.Field, sv.CompareField))
	}

	return nil

}
