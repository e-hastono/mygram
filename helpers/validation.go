package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidationErrorMessages(ve validator.ValidationErrors) (errorMessages []string) {
	for _, err := range ve {
		var errorMessage string
		switch err.Tag() {
		case "required":
			errorMessage = fmt.Sprintf("Field '%s' cannot be blank", err.Field())
		case "email":
			errorMessage = fmt.Sprintf("Field '%s' must be a valid email address", err.Field())
		case "len":
			errorMessage = fmt.Sprintf("Field '%s' must be exactly %v characters long", err.Field(), err.Param())
		case "min":
			errorMessage = fmt.Sprintf("Field '%s' must be a minimum of %v characters long", err.Field(), err.Param())
		case "gt":
			errorMessage = fmt.Sprintf("Field '%s' must be greater than %v", err.Field(), err.Param())
		case "lte":
			errorMessage = fmt.Sprintf("Field '%s' must be less than or equal to %v", err.Field(), err.Param())
		default:
			errorMessage = fmt.Sprintf("Field '%s': '%v' must satisfy '%s' '%v' criteria", err.Field(), err.Value(), err.Tag(), err.Param())
		}
		errorMessages = append(errorMessages, errorMessage)
	}

	return
}
