package validator

import (
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func PhoneNumberValidator(fl validator.FieldLevel) bool {
	phoneNumber := fl.Field().String()

	cleanedNumber := strings.ReplaceAll(phoneNumber, " ", "")
	cleanedNumber = strings.ReplaceAll(cleanedNumber, "-", "")

	if len(cleanedNumber) < 5 || len(cleanedNumber) > 15 {
		return false
	}

	match, _ := regexp.MatchString("^[0-9]+$", cleanedNumber)
	if !match {
		return false
	}

	return true
}
