package ticket_category

import (
	"time"

	"encore-be/internal/models"

	"github.com/google/uuid"
)

type BenefitResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
}

type TicketCategoryResponse struct {
	ID        uuid.UUID         `json:"id"`
	EventID   uuid.UUID         `json:"event_id"`
	Name      string            `json:"name"`
	Price     float64           `json:"price"`
	Quota     int               `json:"quota"`
	Benefits  []BenefitResponse `json:"benefits"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func ToResponse(category models.TicketCategory) TicketCategoryResponse {
	benefits := make([]BenefitResponse, 0, len(category.Benefits))

	for _, benefit := range category.Benefits {
		benefits = append(benefits, BenefitResponse{
			ID:          benefit.ID,
			Name:        benefit.Name,
			Description: benefit.Description,
		})
	}

	return TicketCategoryResponse{
		ID:        category.ID,
		EventID:   category.EventID,
		Name:      category.Name,
		Price:     category.Price,
		Quota:     category.Quota,
		Benefits:  benefits,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}