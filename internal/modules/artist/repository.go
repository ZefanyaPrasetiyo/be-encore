package artist

import (
	"encore-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository{
	return &Repository{db: db}
}

func(r *Repository) findAllArtist() ([]models.Artist, error) {
	var artist []models.Artist

	err := r.db.Preload("CreatedBY").Preload("Staff").Order("created_at DESC").Find(&artist).Error

	return artist,err
}

func (r *Repository) findArtistById(id uuid.UUID) (*models.Artist, error) {
	var artist models.Artist

	err := r.db.
		Preload("CreatedBY").
		Preload("Staff").
		First(&artist, "id = ?", id).Error

	if err != nil {
		return nil, err
	}

	return &artist, nil
}

func (r *Repository) findUserById(id uuid.UUID) (*models.User, error) {
	var user models.User

	err := r.db.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func(r *Repository) createArtist(artist *models.Artist) error {
	return r.db.Create(artist).Error
}

func(r *Repository) updateArtist(artist *models.Artist) error{
	return r.db.Save(artist).Error
}

func(r *Repository) deleteArtist(artist *models.Artist) error{
	return r.db.Delete(artist).Error
}
