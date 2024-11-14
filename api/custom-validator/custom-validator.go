package customvalidator

import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func New(v *validator.Validate) *CustomValidator {
	cv := &CustomValidator{v}
	return cv
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
