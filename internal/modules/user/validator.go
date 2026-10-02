package user

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) ValidateCreate(req CreateUserRequest) error {
	return v.validate.Struct(req)
}

func (v *Validator) ValidateUpdate(req UpdateUserRequest) error {
	return v.validate.Struct(req)
}