package validationmsghandlers

import "github.com/go-playground/validator/v10"

type ProductMsgHandler struct {
}

func (pmh *ProductMsgHandler) ProductValidate(err error) string {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			switch fieldError.Field() {
				case "Title":
					switch fieldError.Tag() {
						case "required":
							return "Title is required."
						case "min":
							return "Title should have a minimum length of 2 characters."
					}
				case "Desc":
					switch fieldError.Tag() {
						case "required":
							return "Description is required."
						case "min":
							return "Description should have a minimum length of 20 characters."
					}
				case "Price":
					switch fieldError.Tag() {
						case "required":
							return "Price is required."
					}
			}
		}
	} 
		

	return "Validation failed."
}