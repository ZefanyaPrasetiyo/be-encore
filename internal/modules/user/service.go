package user

import (
	"encore-be/internal/models"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repository *Repository
	validator  *Validator
}

func NewService(repository *Repository, validator *Validator) *Service {
	return &Service{
		repository: repository,
		validator:  validator,
	}
}

func (s *Service) GetAll() ([]UserResponse, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	responses := make([]UserResponse, 0, len(users))

	for _, user := range users {
		responses = append(responses, ToResponse(user))
	}

	return responses, nil
}

func (s *Service) GetByID(id uuid.UUID) (*UserResponse, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return nil, err
	}

	response := ToResponse(*user)

	return &response, nil
}

func (s *Service) Create(req CreateUserRequest) (*UserResponse, error) {
	if err := s.validator.ValidateCreate(req); err != nil {
		return nil, err
	}

	existingUser, err := s.repository.FindByEmail(req.Email)

	if err == nil && existingUser.ID != uuid.Nil {
		return nil, errors.New("email sudah digunakan")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}

	password := string(passwordHash)

	user := &models.User{
		Email: 			req.Email,
		PasswordHash: 	&password,
		FullName: 		req.FullName,
		PhoneNumber: 	req.PhoneNumber,
		Role:			models.UserRole(req.Role),
	}

	err = s.repository.Create(user)
	if err != nil {
		return nil, err
	}
	response := ToResponse(*user)

	return &response, nil
}
