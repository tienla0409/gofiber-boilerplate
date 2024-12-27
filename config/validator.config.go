package config

import "github.com/go-playground/validator/v10"

type StructValidator struct {
	Validator *validator.Validate
}

func NewValidator() *StructValidator {
	instance := validator.New()

	return &StructValidator{
		Validator: instance,
	}
}

func (v *StructValidator) Validate(out any) error {
	return v.Validator.Struct(out)
}

func (v *StructValidator) Engine() any {
	return v.Validate
}
