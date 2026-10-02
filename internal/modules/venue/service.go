package venue

import (
	"encore-be/internal/models"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	repository *Repository
	validator *Validator
}

func(s *Service) GetAllVenueCategory() ([]models.VenueCategory, error) {
	return s.repository.findAllVenueCategory()
}


func(s *Service) GetVenueCategoryById(id uuid.UUID)(*models.VenueCategory, error ){
	return s.repository.findCategoryById(id)
}


func(s * Service) CreateVenueCategory(req CreateVenueCategoryRequest) (*models.VenueCategory, error) {
	category := &models.VenueCategory{
		Name: req.Name,
	}

	err := s.repository.CreateVenueCategory(category)
	if err != nil{
		return nil, err
	}
	
	return category, nil
}

func (s *Service) UpdateVenueCategory(id uuid.UUID, req UpdateVenueCategoryRequest)(*models.VenueCategory, error) {
	category, err := s.repository.findCategoryById(id)
	if err != nil {
		return nil, err
	}

	category.Name = req.Name

	err = s.repository.updateVenueCategory(category)
	if err != nil {
		return nil, err
	}

	return  category, nil
}

func (s *Service) DeleteVenueCategory(id uuid.UUID) error{
	category, err := s.repository.findCategoryById(id)
	if err != nil {
		return  err
	}

	return s.repository.deleteVenueCategory(category)
}

func (s *Service) GetAllVenue() ([]models.Venue, error) {
	return s.repository.findAllVenue()
} 

func (s *Service) GetVenueById(id uuid.UUID) (*models.Venue, error){
	return s.repository.findVenueById(id)
}

func (s *Service) CreateVenue(req CreateVenueRequest) (*models.Venue, error) {
	_, err := s.repository.findCategoryById(req.VenueCategoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, errors.New("venue category tidak ditemukan")
		}
		return nil,err
	}

	venue := &models.Venue{
		Name: 				req.Name,
		Address: 			req.Address,
		City: 				req.City,
		Capacity: 			req.Capacity,
		VenueCategoryID: 	req.VenueCategoryID,				
	}

	err = s.repository.CreateVenue(venue)
	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (s *Service) updateVenue(id uuid.UUID, req UpdateVenueRequest) (*models.Venue, error) {
	venue, err := s.repository.findVenueById(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		venue.Name = req.Name
	}

	if req.Address != "" {
		venue.Address = req.Address
	}

	if req.City != "" {
		venue.City = req.City
	}

	if req.Capacity != 0 {
		venue.Capacity = req.Capacity
	}

	if req.VenueCategoryID != uuid.Nil {
		_, err := s.repository.findCategoryById(req.VenueCategoryID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("venue category tidak ditemukan")
			}

			return nil, err
		}
		venue.VenueCategoryID = req.VenueCategoryID
	}

	err = s.repository.updateVenue(venue)
	if err != nil {
		return nil, err
	}

	return venue, nil
}

func (s *Service) DeleteVenue(id uuid.UUID) error{
	venue, err := s.repository.findVenueById(id)
	if err != nil {
		return err
	}

	return s.repository.deleteVenue(venue)
}