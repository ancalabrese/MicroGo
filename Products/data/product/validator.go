package product

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	validator.FieldError
}

//ValidationErrors is a collection of errors
type ValidationErrors []ValidationError

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return &Validation{validate}
}
func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag(),
	)
}
