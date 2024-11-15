package skelmi

import "github.com/go-playground/validator"

type DefaultValidator struct {
	validator *validator.Validate
}

func NewDefaultValidator() *DefaultValidator {
	return &DefaultValidator{
		validator: validator.New(),
	}
}

func (v *DefaultValidator) Validate(target interface{}) error {
	if err := v.validator.Struct(target); err != nil {
		return err
	}
	return nil
}
