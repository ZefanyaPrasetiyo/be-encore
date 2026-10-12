package ticket_category

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

func (v Validator) validCreateTicketCategory(req CreateTicketCategoryRequest) error{
	return v.validate.Struct(req)
}

func (v Validator) validUpdateTicketCategory(req UpdateTicketCategoryRequest) error{
	return v.validate.Struct(req)
}



