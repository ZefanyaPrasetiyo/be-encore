package artist

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

func (v *Validator) Validate(req any) error {
	return v.validate.Struct(req)
}


func(v Validator) validCreateArtist(req CreateArtistRequest) error{
	return v.validate.Struct(req)
}

func(v Validator) validUpdateArtist(req UpdateArtistRequest) error{
	return v.validate.Struct(req)
}



