package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ValidateStruct(data interface{}) map[string]string {

	err := Validate.Struct(data)

	if err == nil {
		return nil
	}

	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {

		field := strings.ToLower(e.Field())

		switch e.Tag() {

		case "required":
			errors[field] = e.Field() + " is required"

		case "email":
			errors[field] = "Invalid email address"

		case "min":
			errors[field] = e.Field() + " is too short"

		case "max":
			errors[field] = e.Field() + " is too long"

		default:
			errors[field] = "Invalid value"
		}
	}

	return errors
}