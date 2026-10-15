package events

import (
	"time"

	"encore-be/internal/models"

	"github.com/google/uuid"
)

type EventArtistResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
}

type EventVenueResponse struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Address  string    `json:"address"`
	City     string    `json:"city"`
	Capacity int       `json:"capacity"`
}

type EventResponse struct {
	ID          uuid.UUID           `json:"id"`
	ArtistID    uuid.UUID           `json:"artist_id"`
	VenueID     uuid.UUID           `json:"venue_id"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	StartDate   time.Time           `json:"start_date"`
	Status      models.EventStatus  `json:"status"`
	BannerURL   *string             `json:"banner_url,omitempty"`
	Artist      EventArtistResponse `json:"artist"`
	Venue       EventVenueResponse  `json:"venue"`
	CreatedAt   time.Time           `json:"created_at"`
	UpdatedAt   time.Time           `json:"updated_at"`
}

func ToResponse(event models.Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		ArtistID:    event.ArtistID,
		VenueID:     event.VenueID,
		Title:       event.Title,
		Description: event.Description,
		StartDate:   event.StartDate,
		Status:      event.Status,
		BannerURL:   event.BannerURL,
		Artist: EventArtistResponse{
			ID:          event.Artist.ID,
			Name:        event.Artist.Name,
			Description: event.Artist.Description,
		},
		Venue: EventVenueResponse{
			ID:       event.Venue.ID,
			Name:     event.Venue.Name,
			Address:  event.Venue.Address,
			City:     event.Venue.City,
			Capacity: event.Venue.Capacity,
		},
		CreatedAt: event.CreatedAt,
		UpdatedAt: event.UpdatedAt,
	}
}

func ToResponses(events []models.Event) []EventResponse {
	responses := make([]EventResponse, 0, len(events))

	for _, event := range events {
		responses = append(responses, ToResponse(event))
	}

	return responses
}