package venue

import (
	"encore-be/internal/models"
	"time"
	"github.com/google/uuid"
)

type VenueCategoryResponse struct {
	ID 	  		uuid.UUID 	`json:"id"`
	Name 		string    	`json:"name"`
	Venues 		*string		`json:"venues,omitempty"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}

type VenueResponse struct {
	ID 		 		uuid.UUID 		`json:"id"`
	Name     		string    		`json:"name"`
	Address  		string    		`json:"address"`
	City	 		string    		`json:"city"`
	VenueCategoryID	uuid.UUID 		`json:"venue_category_id"`
	VenueCategory	string			`json:"venue_category"`
	Capacity 		int       		`json:"capacity"`
	CreatedAt 		time.Time 		`json:"created_at"`
	UpdatedAt 		time.Time 		`json:"updated_at"`
}

func ToResponse(venue models.Venue) VenueResponse {
	return VenueResponse{
		ID: 				venue.ID,
		Name: 				venue.Name,
		Address: 			venue.Address,
		City: 				venue.City,
		VenueCategoryID: 	venue.VenueCategory.ID,
		VenueCategory: 		venue.VenueCategory.Name,
		Capacity: 			venue.Capacity,
		CreatedAt: 			venue.CreatedAt,
		UpdatedAt: 			venue.UpdatedAt,		
	}
}