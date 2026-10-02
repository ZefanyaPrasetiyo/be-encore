package venue

import(
	"github.com/google/uuid"
)

type CreateVenueCategoryRequest struct {
	Name		string	`json:"name" validate:"required, name"`
}

type UpdateVenueCategoryRequest struct {
	Name		string	`json:"name" validate:"required, name"`
}


type CreateVenueRequest struct {
	Name 			string 		`json:"name" validate:"required,name"`
	Address  		string 		`json:"address" validate:"required,max=255"`
	City	 		string 		`json:"city" validate:"required,max=100"`
	Capacity 		int    		`json:"capacity" validate:"required, min=1"`
	VenueCategoryID	uuid.UUID	`json:"venue_category_id" validate:"required"`
}

type UpdateVenueRequest struct {
	Name 				string 		`json:"name" validate:"omitempty,max=150"`
	Address  			string 		`json:"address" validate:"omitempty, max=255"`
	City	 			string 		`json:"city" validate:"omitempty, max=100"`
	Capacity 			int    		`json:"capacity" validate:"omitempty, min=1"`
	VenueCategoryID		uuid.UUID	`json:"venue_category_id" validate:"omitempty"`
}