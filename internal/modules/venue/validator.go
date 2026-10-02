package venue

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func NewValidator() *Validator{
	return &Validator{
		validate: validator.New(),
	}
}

func (v Validator) validCreateVenue(req CreateVenueRequest) error{
	return v.validate.Struct(req)
}

func (v Validator) validCreateCategory(req CreateVenueCategoryRequest) error{
	return v.validate.Struct(req)
}

func (v Validator) validUpdateVenue(req UpdateVenueRequest) error{
	return v.validate.Struct(req)
}

func (v Validator) validUpdateCategory(req UpdateVenueCategoryRequest) error{
	return v.validate.Struct(req)
}