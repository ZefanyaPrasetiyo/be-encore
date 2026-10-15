package events

import (
	"errors"

	"encore-be/internal/models"

	"github.com/google/uuid"
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

func (s *Service) GetAllEvents() ([]models.Event, error) {
	return s.repository.findAllEvents()
}

func (s *Service) GetEventById(id uuid.UUID) (*models.Event, error) {
	return s.repository.findEventById(id)
}

func (s *Service) CreateEvent(req CreateEventRequest) (*models.Event, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	if _, err := s.repository.findArtistById(req.ArtistID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("artist tidak ditemukan")
		}
		return nil, err
	}

	if _, err := s.repository.findVenueById(req.VenueID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("venue tidak ditemukan")
		}
		return nil, err
	}

	status := req.Status
	if status == "" {
		status = models.EventDraft
	}

	event := &models.Event{
		ArtistID:    req.ArtistID,
		VenueID:     req.VenueID,
		Title:       req.Title,
		Description: req.Description,
		StartDate:   req.StartDate,
		Status:      status,
		BannerURL:   req.BannerURL,
	}

	if err := s.repository.createEvent(event); err != nil {
		return nil, err
	}

	return s.repository.findEventById(event.ID)
}

func (s *Service) UpdateEvent(id uuid.UUID, req UpdateEventRequest) (*models.Event, error) {
	if err := s.validator.Validate(req); err != nil {
		return nil, err
	}

	if req.ArtistID == nil &&
		req.VenueID == nil &&
		req.Title == nil &&
		req.Description == nil &&
		req.StartDate == nil &&
		req.Status == nil &&
		req.BannerURL == nil {
		return nil, errors.New("minimal satu field harus diisi")
	}

	event, err := s.repository.findEventById(id)
	if err != nil {
		return nil, err
	}

	if req.ArtistID != nil {
		if _, err := s.repository.findArtistById(*req.ArtistID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("artist tidak ditemukan")
			}
			return nil, err
		}
		event.ArtistID = *req.ArtistID
	}

	if req.VenueID != nil {
		if _, err := s.repository.findVenueById(*req.VenueID); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("venue tidak ditemukan")
			}
			return nil, err
		}
		event.VenueID = *req.VenueID
	}

	if req.Title != nil {
		event.Title = *req.Title
	}

	if req.Description != nil {
		event.Description = *req.Description
	}

	if req.StartDate != nil {
		event.StartDate = *req.StartDate
	}

	if req.Status != nil {
		event.Status = *req.Status
	}

	if req.BannerURL != nil {
		event.BannerURL = req.BannerURL
	}

	if err := s.repository.updateEvent(event); err != nil {
		return nil, err
	}

	return s.repository.findEventById(id)
}

func (s *Service) DeleteEvent(id uuid.UUID) error {
	event, err := s.repository.findEventById(id)
	if err != nil {
		return err
	}

	return s.repository.deleteEvent(event)
}