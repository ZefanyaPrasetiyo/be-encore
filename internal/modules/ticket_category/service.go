
package ticket_category

import (
	"encore-be/internal/models"

	"github.com/google/uuid"
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

func (s *Service) GetAllTicketCategory() ([]models.TicketCategory, error) {
	return s.repository.findAllTicketCategory()
}

func (s *Service) GetTicketCategoryById(id uuid.UUID) (*models.TicketCategory, error) {
	return s.repository.findTicketCategoryById(id)
}

func (s *Service) CreateTicketCategory(req CreateTicketCategoryRequest) (*models.TicketCategory, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	ticketCategory := &models.TicketCategory{
		Name:    req.Name,
		Price:   req.Price,
		Quota:   req.Quota,
		EventID: req.EventID,
	}

	if err := s.repository.createTicketCategory(ticketCategory); err != nil {
		return nil, err
	}

	return s.repository.findTicketCategoryById(ticketCategory.ID)
}

func (s *Service) UpdateTicketCategory(id uuid.UUID, req UpdateTicketCategoryRequest) (*models.TicketCategory, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	ticketCategory, err := s.repository.findTicketCategoryById(id)
	if err != nil {
		return nil, err
	}

	ticketCategory.Name = req.Name

	if req.Price != nil {
		ticketCategory.Price = *req.Price
	}

	if req.Quota != nil {
		ticketCategory.Quota = *req.Quota
	}

	if err := s.repository.updateTicketCategory(ticketCategory); err != nil {
		return nil, err
	}

	return s.repository.findTicketCategoryById(ticketCategory.ID)
}

func(s *Service) DeleteTicketCategory(id uuid.UUID) error{
	ticketCategory, err := s.repository.findTicketCategoryById(id)
	
	if err != nil {
		return err
	}

	return s.repository.DeleteTicketCategory(ticketCategory)
}