package user

import (
	"encore-be/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db,}
}

func(r *Repository) FindAll() ([]models.User, error) {
	var users []models.User

	err := r.db.Find(&users).Error

	return  users, err
}

func(r *Repository) FindById(id uuid.UUID) (*models.User,error ) {
	var user models.User

	err := r.db.First(&user, "id = ?").Error

	return &user, err
}

func(r *Repository) FindByEmail(email string)(*models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return &user, err
}

func(r *Repository) Create(user *models.User) error{
	return r.db.Create(user).Error
} 

func(r *Repository) Update(user *models.User) error{
	return r.db.Save(user).Error
}

func(r *Repository) Delete(user *models.User) error{
	return r.db.Delete(user).Error
}