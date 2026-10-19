package ticket_category

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

func(r *Repository)  findAllTicketCategory() ([]models.TicketCategory, error) {
	var ticketCategory []models.TicketCategory

	err := r.db.Find(&ticketCategory).Error

	return ticketCategory, err
}


func(r *Repository) findTicketCategoryById(id uuid.UUID) (*models.TicketCategory, error) {
	var ticketCategory models.TicketCategory

	err := r.db.First(&ticketCategory, "id = ?", id).Error

	return &ticketCategory, err
}

func(r *Repository) createTicketCategory(ticketCategory *models.TicketCategory) error{
	return r.db.Create(ticketCategory).Error

}

func(r *Repository) updateTicketCategory(ticketCategory *models.TicketCategory) error {
	return r.db.Save(ticketCategory).Error
}

func(r *Repository) DeleteTicketCategory(ticketCategory *models.TicketCategory) error {
	return r.db.Delete(ticketCategory).Error
}