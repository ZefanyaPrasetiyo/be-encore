package artist

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

func NewService(repository *Repository, validator *Validator) *Service {
	return &Service{
		repository: repository,
		validator:  validator,
	}
}


func(s *Service) GetAllArtist() ([]models.Artist, error) {
	return s.repository.findAllArtist()
}

func(s *Service) GetArtistById(id uuid.UUID) (*models.Artist, error) {
	return s.repository.findArtistById(id)
}

func (s *Service) CreateArtist(req CreateArtistRequest) (*models.Artist, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	_, err := s.repository.findUserById(req.CreatedByID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user pembuat tidak ditemukan")
		}

		return nil, err
	}

	artist := &models.Artist{
		Name:        req.Name,
		Description: req.Description,
		CreatedByID: req.CreatedByID,
	}

	if err := s.repository.createArtist(artist); err != nil {
		return nil, err
	}

	return s.repository.findArtistById(artist.ID)
}

func (s *Service) UpdateArtist(id uuid.UUID, req UpdateArtistRequest) (*models.Artist, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	if req.Name == nil && req.Description == nil {
		return nil, errors.New("minimal satu field harus diisi")
	}

	artist, err := s.repository.findArtistById(id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		artist.Name = *req.Name
	}

	if req.Description != nil {
		artist.Description = req.Description
	}

	if err := s.repository.updateArtist(artist); err != nil {
		return nil, err
	}

	return s.repository.findArtistById(id)
}

func (s *Service) DeleteArtist(id uuid.UUID) error {
	artist, err := s.repository.findArtistById(id)
	if err != nil {
		return err
	}

	return s.repository.deleteArtist(artist)
}