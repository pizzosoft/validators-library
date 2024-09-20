package validators

import (
	"errors"

	"github.com/go-playground/validator"
)

func ValidateStruct(obj interface{}) error {
	validate := validator.New()
	validate.RegisterValidation("is-good-password", ValidatePassword)
	err := validate.Struct(obj)
	if err != nil {

		validationErrors := err.(validator.ValidationErrors)

		validationError := validationErrors[0]

		switch validationError.Tag() {
		case "required":
			return errors.New(validationError.Field() + " is required")
		case "email":
			return errors.New(validationError.Field() + " is not a valid email")
		case "min":
			return errors.New(validationError.Field() + " must be at least " + validationError.Param() + " characters long")
		case "max":
			return errors.New(validationError.Field() + " must be at most " + validationError.Param() + " characters long")
		case "is-good-password":
			return errors.New(validationError.Field() + " must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one number and one special character")
		}
	}

	return nil
}

// validade if the password is strong
func ValidatePassword(password validator.FieldLevel) bool {
	validate := validator.New()
	err := validate.Var(password.Field().String(), "min=8,containsany=!@#$%^&*(),containsany=0123456789,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	return err == nil
}
