package validator

import (
	"fmt"

	"github.com/faxa0-0/billy/user_service/internal/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// Initialize the validator
func InitValidator() {
	validate = validator.New()
}

// Validate User
func ValidateUser(user *models.User) error {
	err := validate.Struct(user)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return fmt.Errorf("validation failed for field %s, condition: %s", e.Field(), e.ActualTag())
		}
	}
	return nil
}
