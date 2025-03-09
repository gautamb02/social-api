package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ErrorMessage struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

// IsValid validate the given struct based on its rule
func IsValid(data any) ([]ErrorMessage, bool) {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	if err := validate.Struct(data); err != nil && !strings.Contains(err.Error(), "nil") {
		errs := translateError(err)
		return errs, false
	}
	return []ErrorMessage{}, true
}

func translateError(err error) (errs []ErrorMessage) {
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		errs = append(errs, ErrorMessage{
			Field: e.Field(),
			Error: e.Error(),
		})
	}
	return errs
}
