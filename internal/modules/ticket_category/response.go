package ticket_category

import (
	"time"
	"github.com/google/uuid"
)

type BenefitResponse struct {
	ID				uuid.UUID	`json:"id"`
	Name			string		`json:"name"`
	Description		*string		`json:"description,omitempty"`
}

type CreateTicketCategoryResponse struct {
	ID			uuid.UUID			`json:"id"`
	EventID		uuid.UUID			`json:"event_id"`
	Name		string				`json:"name"`
	Price		float64				`json:"price"`
	Quota		int					`json:"quota"`
	Benefits	[]BenefitResponse 	`json:"benefits"`
	CreatedAt	time.Time			`json:"created_at"`
	UpdatedAt	time.Time			`json:"updated_at"`
}

type UpdateTicketCategoryResponse struct {
	ID			uuid.UUID			`json:"id"`
	EventID		uuid.UUID			`json:"event_id"`
	Name		string				`json:"name"`
	Price		float64				`json:"price"`
	Quota		int					`json:"quota"`
	Benefits	[]BenefitResponse 	`json:"benefits"`
	CreatedAt	time.Time			`json:"created_at"`
	UpdatedAt	time.Time			`json:"updated_at"`
}