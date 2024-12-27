package util

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

type ValidateError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

func GetValidateErrors(err error) []ValidateError {
	var ve validator.ValidationErrors
	out := make([]ValidateError, 0)

	if errors.As(err, &ve) {
		out = make([]ValidateError, len(ve))

		for i, formattedErr := range ve {
			out[i] = ValidateError{
				Field:  formattedErr.Field(),
				Reason: reasonByTag(formattedErr.Tag(), formattedErr.Param()),
			}
		}
	}

	return out
}

func reasonByTag(tag string, param string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "This field must be a valid email address"
	case "min":
		return fmt.Sprintf("This field must be at least %s characters", param)
	case "max":
		return fmt.Sprintf("This field must be at most %s characters", param)
	case "eqfield":
		return fmt.Sprintf("This field must be equal to %s", param)
	case "nefield":
		return fmt.Sprintf("This field must not be equal to %s", param)
	case "ltfield":
		return fmt.Sprintf("This field must be less than %s", param)
	case "ltefield":
		return fmt.Sprintf("This field must be less than or equal to %s", param)
	case "gtfield":
		return fmt.Sprintf("This field must be greater than %s", param)
	case "gtefield":
		return fmt.Sprintf("This field must be greater than or equal to %s", param)
	case "oneof":
		return fmt.Sprintf("This field must be one of %s", param)
	case "unique":
		return "This field must be unique"
	default:
		return "This field is invalid"
	}
}
