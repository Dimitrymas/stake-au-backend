package validation

import (
	"backend/api/pkg/constants"
	"errors"
	"fmt"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"strings"
)

// Global variables for validator and translator
var (
	validate      *validator.Validate
	trans         ut.Translator
	validStatuses = []string{
		constants.WithdrawalPending,
		constants.WithdrawalCanceled,
		constants.WithdrawalRejected,
		constants.WithdrawalRejectedNoRefund,
		constants.WithdrawalSuccess,
		constants.WithdrawalAll,
	}
)

// InitValidator initializes the validator and translator
func init() {
	validate = validator.New()

	russian := ru.New()
	uni := ut.New(russian, russian)

	var found bool
	trans, found = uni.GetTranslator("ru")
	if !found {
		log.Fatalf("translator not found")
	}

	err := validate.RegisterValidation("withdrawal_status", WithdrawalStatusValidation)
	if err != nil {
		fmt.Println("Ошибка при регистрации валидации:", err)
	}

	err = validate.RegisterTranslation(
		"required",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} обязательно", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)

	err = validate.RegisterTranslation(
		"min",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("min", "{0} должно быть больше {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("min", fe.Field(), fe.Param())
			return t
		},
	)

	err = validate.RegisterTranslation(
		"max",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("max", "{0} должно быть меньше {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("max", fe.Field(), fe.Param())
			return t
		},
	)

	err = validate.RegisterTranslation(
		"gt",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("gt", "{0} должно быть больше {1}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("gt", fe.Field(), fe.Param())
			return t
		},
	)
	// Регистрируем перевод для кастомного валидатора
	err = validate.RegisterTranslation(
		"withdrawal_status",
		trans, func(ut ut.Translator) error {
			return ut.Add("withdrawal_status", "Статус должен быть одним из: {0}", true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			validStatusesString := strings.Join(validStatuses, ", ")
			t, _ := ut.T("withdrawal_status", validStatusesString)
			return t
		},
	)
	if err != nil {
		log.Fatalf("error translating validator: %v", err)
	}
}

// WithdrawalStatusValidation is a custom validator for withdrawal status
func WithdrawalStatusValidation(fl validator.FieldLevel) bool {
	status := fl.Field().String()

	for _, validStatus := range validStatuses {
		if status == validStatus {
			return true
		}
	}
	return false
}

// ValidateStruct validates a given struct and returns translated error messages
func ValidateStruct(s interface{}) error {
	if err := validate.Struct(s); err != nil {
		return err
	}
	return nil
}

// TranslateError translates validation errors to Russian
func TranslateError(err error) map[string]string {
	if err == nil {
		return nil
	}

	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return map[string]string{"error": err.Error()}
	}

	if trans == nil {
		return map[string]string{"error": "translator not initialized"}
	}

	translatedErrors := errs.Translate(trans)
	if translatedErrors == nil {
		return map[string]string{"error": "translation failed"}
	}

	return translatedErrors
}
