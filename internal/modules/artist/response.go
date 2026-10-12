package artist

import (
	"time"

	"encore-be/internal/models"

	"github.com/google/uuid"
)

type ArtistUserResponse struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
}

type ArtistResponse struct {
	ID          uuid.UUID            `json:"id"`
	Name        string               `json:"name"`
	Description *string              `json:"description,omitempty"`
	CreatedByID uuid.UUID            `json:"created_by"`
	CreatedBy   *ArtistUserResponse  `json:"created_by_user,omitempty"`
	Staff       []ArtistUserResponse `json:"staff"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
}

func ToResponse(artist models.Artist) ArtistResponse {
	response := ArtistResponse{
		ID:          artist.ID,
		Name:        artist.Name,
		Description: artist.Description,
		CreatedByID: artist.CreatedByID,
		Staff:       make([]ArtistUserResponse, 0, len(artist.Staff)),
		CreatedAt:   artist.CreatedAt,
		UpdatedAt:   artist.UpdatedAt,
	}

	if artist.CreatedBY.ID != uuid.Nil {
		response.CreatedBy = &ArtistUserResponse{
			ID:       artist.CreatedBY.ID,
			FullName: artist.CreatedBY.FullName,
			Email:    artist.CreatedBY.Email,
		}
	}

	for _, user := range artist.Staff {
		response.Staff = append(response.Staff, ArtistUserResponse{
			ID:       user.ID,
			FullName: user.FullName,
			Email:    user.Email,
		})
	}

	return response
}

func ToResponses(artists []models.Artist) []ArtistResponse {
	responses := make([]ArtistResponse, 0, len(artists))

	for _, artist := range artists {
		responses = append(responses, ToResponse(artist))
	}

	return responses
}
