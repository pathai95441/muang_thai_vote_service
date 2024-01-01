package application

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateOptional(fl validator.FieldLevel) bool {
	return true
}

func ValidateRequiredIfValue(fl validator.FieldLevel) bool {
	otherFieldName := strings.Split(fl.Param(), " ")[0]
	otherFieldCheck := strings.Split(fl.Param(), " ")[1]

	var otherFieldValue reflect.Value
	if fl.Parent().Kind() == reflect.Ptr {
		otherFieldValue = fl.Parent().Elem().FieldByName(otherFieldName)
	} else {
		otherFieldValue = fl.Parent().FieldByName(otherFieldName)
	}

	if otherFieldCheck == otherFieldValue.String() {
		return !fl.Field().IsZero()
	}

	return true
}
