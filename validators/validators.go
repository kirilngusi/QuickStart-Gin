package validators

import (
	"github.com/go-playground/validator/v10"
	"strings"
)

func ValidateCoolTitle(filed validator.FieldLevel) bool {
	return strings.Contains(filed.Field().String(), "Cool")
}
