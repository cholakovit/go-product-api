package validationmessages

import "github.com/go-playground/validator/v10"

func UserMessageValidate(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
				case "FullName":
					switch fieldError.Tag() {
						case "required":
							return "Full name is required."
						case "min":
							return "Full name should have a minimum length of 2 characters."
					}
				case "Pass":
					switch fieldError.Tag() {
						case "required":
							return "Password is required."
						case "min":
							return "Password should have a minimum length of 2 characters."
					}
				case "Email":
					switch fieldError.Tag() {
						case "required":
							return "Email is required."
						case "min":
							return "Email should have a minimum length of 2 characters."
					}
				case "Role":
					switch fieldError.Tag() {
						case "required":
							return "Role is required."
						case "min":
							return "Role should have a minimum length of 2 characters."
					}
			}
		}
	} 
	return "Validation failed."
}