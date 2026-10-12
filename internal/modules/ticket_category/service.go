package ticket_category

// import (
// 	"encore-be/internal/models"
// 	"github.com/google/uuid"
// )

// type Service struct {
// 	repository *Repository
// 	validator *Validator
// }

// func NewService(repository *Repository, validator *Validator) *Service {
// 	return &Service{
// 		repository: repository,
// 		validator: validator,
// 	}
// }

// func(s *Service) GetAllTicketCategory() ([]models.TicketCategory, error){
// 	return s.repository.findAllTicketCategory()
// }

// func(s *Service) GetTicketCategoryById(id uuid.UUID)(*models.TicketCategory, error){
// 	return s.repository.findTicketCategoryById(id)
// }

// func(s *Service) CreateTicketCategory(req CreateTicketCategoryRequest) (*models.TicketCategory, error){
// 	if err := s.validator.Validate(req); err != nil {
// 		return nil, err
// 	}

// 	_, err := s.repository.find
// }