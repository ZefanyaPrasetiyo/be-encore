package events

import (
	"encore-be/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) findAllEvents() ([]models.Event, error) {
	var events []models.Event

	err := r.db.
		Preload("Artist").
		Preload("Venue").
		Order("created_at DESC").
		Find(&events).Error

	return events, err
}

func (r *Repository) findEventById(id uuid.UUID) (*models.Event, error) {
	var event models.Event

	err := r.db.
		Preload("Artist").
		Preload("Venue").
		First(&event, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (r *Repository) findArtistById(id uuid.UUID) (*models.Artist, error) {
	var artist models.Artist

	err := r.db.First(&artist, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (r *Repository) findVenueById(id uuid.UUID) (*models.Venue, error) {
	var venue models.Venue

	err := r.db.First(&venue, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &venue, nil
}

func (r *Repository) createEvent(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *Repository) updateEvent(event *models.Event) error {
	return r.db.Save(event).Error
}

func (r *Repository) deleteEvent(event *models.Event) error {
	return r.db.Delete(event).Error
}