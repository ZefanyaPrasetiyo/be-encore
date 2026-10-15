package events

import (
	"time"

	"encore-be/internal/models"

	"github.com/google/uuid"
)

type CreateEventRequest struct {
	ArtistID    uuid.UUID          `json:"artist_id" validate:"required"`
	VenueID     uuid.UUID          `json:"venue_id" validate:"required"`
	Title       string             `json:"title" validate:"required,max=200"`
	Description string             `json:"description" validate:"required"`
	StartDate   time.Time          `json:"start_date" validate:"required"`
	Status      models.EventStatus `json:"status" validate:"omitempty,oneof=draft published ended cancelled"`
	BannerURL   *string            `json:"banner_url" validate:"omitempty,url,max=255"`
}

type UpdateEventRequest struct {
	ArtistID    *uuid.UUID          `json:"artist_id" validate:"omitempty"`
	VenueID     *uuid.UUID          `json:"venue_id" validate:"omitempty"`
	Title       *string             `json:"title" validate:"omitempty,min=1,max=200"`
	Description *string             `json:"description" validate:"omitempty"`
	StartDate   *time.Time          `json:"start_date" validate:"omitempty"`
	Status      *models.EventStatus `json:"status" validate:"omitempty,oneof=draft published ended cancelled"`
	BannerURL   *string             `json:"banner_url" validate:"omitempty,url,max=255"`
}
