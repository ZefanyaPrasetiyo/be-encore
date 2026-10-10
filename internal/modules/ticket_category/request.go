package ticket_category

import (
	"github.com/google/uuid"
)

type CreateTicketCategoryRequest struct {
	EventID   uuid.UUID         `json:"event_id" validate:"required"`
	Name      string            `json:"name" validate:"required"`
	Price     float64           `json:"price" validate:"gte=0"`
	Quota     int               `json:"quota" validate:"gte=0"`
	Benefits  []uuid.UUID 		`json:"benefits" validate:"omitempty, dive, required"`
}

type UpdateTicketCategoryRequest struct {
	Name		string			`json:"name" validate:"omitempty, max=100"`
	Price		*float64		`json:"price" validate:"omitempty, gte=0"`
	Quota		*int			`json:"quota" validate:"omitempty, gte=0"`
	Benefits	[]uuid.UUID		`json:"benefits" validate:"omitempty, dive, required"`
}