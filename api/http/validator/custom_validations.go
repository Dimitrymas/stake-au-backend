package validation

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

// Regular expressions for different password validations
var (
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
	digitRegex       = regexp.MustCompile(`\d`)
	specialCharRegex = regexp.MustCompile(`[\W_]`)
)

// ValidateLowercase Validate lowercase letter
func ValidateLowercase(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return lowercaseRegex.MatchString(password)
}

// ValidateUppercase checks for at least one uppercase letter
func ValidateUppercase(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return uppercaseRegex.MatchString(password)
}

// ValidateDigit checks for at least one digit
func ValidateDigit(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return digitRegex.MatchString(password)
}

// ValidateSpecialChar checks for at least one special character
func ValidateSpecialChar(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	return specialCharRegex.MatchString(password)
}
