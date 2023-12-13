package reqvalidator

import (
	"fmt"
	"reflect"
	"strings"
)

// Validate check fieldswith tag "required"
func Validate(input interface{}) error {
	valueOfInput := reflect.ValueOf(input)

	if valueOfInput.Kind() == reflect.Ptr && valueOfInput.Elem().Kind() == reflect.Struct {
		elem := valueOfInput.Elem()
		return validateStruct(elem)
	}

	return fmt.Errorf("invalid input type for validation")
}

func validateStruct(elem reflect.Value) error {
	typeOfElem := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := typeOfElem.Field(i)
		value := elem.Field(i)
		required := field.Tag.Get("required")

		if required == "true" && isZero(value) {
			fieldName := strings.ToLower(field.Name)
			return fmt.Errorf("required field %s is empty", fieldName)
		}

		if value.Kind() == reflect.Struct {
			err := validateStruct(value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func isZero(value reflect.Value) bool {
	zeroValue := reflect.Zero(value.Type())
	return reflect.DeepEqual(value.Interface(), zeroValue.Interface())
}
