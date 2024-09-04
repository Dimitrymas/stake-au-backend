package validation

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
)

// HandleValidationError converts validation errors to a readable format
func HandleValidationError(err error) *map[string]string {
	if err == nil {
		return nil
	}

	validationErrors := make(map[string]string)

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			field := e.Field()
			var msg string

			switch e.Tag() {
			case "required":
				msg = fmt.Sprintf("Field '%s' is required", field)
			case "min":
				msg = fmt.Sprintf("Field '%s' must be at least %s characters long", field, e.Param())
			case "max":
				msg = fmt.Sprintf("Field '%s' must be no more than %s characters long", field, e.Param())
			case "lowercase":
				msg = fmt.Sprintf("Field '%s' must contain at least one lowercase letter", field)
			case "uppercase":
				msg = fmt.Sprintf("Field '%s' must contain at least one uppercase letter", field)
			case "digit":
				msg = fmt.Sprintf("Field '%s' must contain at least one digit", field)
			case "specialchar":
				msg = fmt.Sprintf("Field '%s' must contain at least one special character", field)
			default:
				msg = fmt.Sprintf("Field '%s' is invalid: %s", field, e.Tag())
			}

			validationErrors[field] = msg
		}
	}

	return &validationErrors
}
