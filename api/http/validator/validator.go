package validation

import (
	"github.com/go-playground/validator/v10"
)

// Global variable for validator
var validate *validator.Validate

// InitValidator initializes the validator
func init() {
	validate = validator.New()

	// Register custom validations here if needed
	registerCustomValidations()
}

// RegisterCustomValidations registers any custom validation rules
func registerCustomValidations() {
	err := validate.RegisterValidation("lowercase", ValidateLowercase)
	if err != nil {
		panic("failed to register lowercase validation")
	}

	err = validate.RegisterValidation("uppercase", ValidateUppercase)
	if err != nil {
		panic("failed to register uppercase validation")
	}

	err = validate.RegisterValidation("digit", ValidateDigit)
	if err != nil {
		panic("failed to register digit validation")
	}

	err = validate.RegisterValidation("specialchar", ValidateSpecialChar)
	if err != nil {
		panic("failed to register specialchar validation")
	}
}

// ValidateStruct validates a given struct and returns error messages
func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}
