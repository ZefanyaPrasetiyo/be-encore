package events

import(
	"github.com/go-playground/validator/v10"
)

type Validator struct{
	validate *validator.Validate
}

func NewValidator() *Validator{
	return &Validator{
		validate: validator.New(),
	}
}

func (v *Validator) Validate(req any) error {
	return v.validate.Struct(req)
}

func(v Validator) ValidCreateEvent(req CreateEventRequest) error{
	return v.validate.Struct(req)
}

func(v Validator) ValidUpdateEvent(req UpdateEventRequest) error{
	return v.validate.Struct(req)
}

