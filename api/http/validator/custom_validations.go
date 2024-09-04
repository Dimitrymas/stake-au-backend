package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// ValidatePassword checks password complexity
func ValidatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[\W_]).{8,20}$`)
	return passwordRegex.MatchString(password)
}
