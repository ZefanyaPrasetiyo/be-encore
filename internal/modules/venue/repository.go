package venue

import (
	"encore-be/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{db: db,}
}

func(r *Repository) findAllVenueCategory() ([]models.VenueCategory, error) {
	var venueCategory []models.VenueCategory

	err := r.db.Find(&venueCategory).Error

	return venueCategory,err
}

func(r *Repository) findCategoryById(id uuid.UUID) (*models.VenueCategory, error) {
	var category models.VenueCategory

	err := r.db.First(&category, "id = ?", id).Error

	return &category, err
}

func(r *Repository) findAllVenue() ([]models.Venue, error) {
	var venue []models.Venue

	err := r.db.Preload("VenueCategory").Find(&venue).Error

	return venue, err
}

func(r *Repository) findVenueById(id uuid.UUID) (*models.Venue, error) {
	var venue models.Venue

	err := r.db.First(&venue, "id = ?", id).Error

	return &venue, err
}

func(r *Repository) CreateVenueCategory(venueCategory *models.VenueCategory) error{
	return r.db.Create(venueCategory). Error
}

func(r *Repository) CreateVenue(venue *models.Venue) error{
	return  r.db.Create(venue).Error
}

func(r *Repository) updateVenueCategory(venueCategory *models.VenueCategory) error{
	return r.db.Save(venueCategory).Error
}

func(r *Repository) updateVenue(venue *models.Venue) error{
	return  r.db.Save(venue).Error
}

func(r *Repository) deleteVenueCategory(venueCategory *models.VenueCategory) error{
	return  r.db.Delete(venueCategory).Error
}

func(r *Repository) deleteVenue(venue *models.Venue) error{
	return r.db.Delete(venue).Error
}
