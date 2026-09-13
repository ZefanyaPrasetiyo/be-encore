package user

import (
	"gorm.io/gorm"
	"encore-be/internal/models"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service{
	return &Service{db: db}
}

func (s *Service) GetAllUsers() ([]models.User, error)  {
	var users []models.User

	err := s.db.Find(&users).Error
	return users, err
}