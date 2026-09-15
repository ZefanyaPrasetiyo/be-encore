package models

import (
	"time"
	"github.com/google/uuid"
)

type EventStatus string

const (
	EventDraft     EventStatus = "draft"
	EventPublished EventStatus = "published"
	EventEnded     EventStatus = "ended"
	EventCancelled EventStatus = "cancelled"
)

type Event struct {
	ID          uuid.UUID   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	ArtistID    uuid.UUID   `gorm:"type:uuid;not null" json:"artist_id"`
	VenueID     uuid.UUID   `gorm:"type:uuid;not null" json:"venue_id"`
	Title       string      `gorm:"type:varchar(200);not null" json:"title"`
	Description string      `gorm:"type:text;not null" json:"description"`
	StartDate   time.Time   `gorm:"not null" json:"start_date"`
	Status      EventStatus `gorm:"type:event_status;default:'draft';not null" json:"status"`
	BannerURL   *string     `gorm:"type:varchar(255)" json:"banner_url,omitempty"`
	Artist Artist `gorm:"foreignKey:ArtistID" json:"artist"`
	Venue  Venue  `gorm:"foreignKey:VenueID" json:"venue"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}