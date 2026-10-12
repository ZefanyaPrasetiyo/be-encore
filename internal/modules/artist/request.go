
package artist

import "github.com/google/uuid"

type CreateArtistRequest struct {
	Name        string     `json:"name" validate:"required,max=100"`
	Description *string    `json:"description" validate:"omitempty,max=5000"`
	CreatedByID uuid.UUID  `json:"created_by" validate:"required"`
}

type UpdateArtistRequest struct {
	Name        *string `json:"name" validate:"omitempty,min=1,max=100"`
	Description *string `json:"description" validate:"omitempty,max=5000"`
}
