package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func HandleValidationErrors(err error) ([]ValidationErrorResponse, bool) {
	if err == nil {
		return make([]ValidationErrorResponse, 0), false
	}

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return make([]ValidationErrorResponse, 0), false
	}

	var errs []ValidationErrorResponse
	for _, err := range validationErrors {
		errs = append(errs, ValidationErrorResponse{
			Field:   err.Field(),
			Message: err.Tag(),
		})
	}

	return errs, true
}
