package user

import (
	"encore-be/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserResponse struct {
	ID 			uuid.UUID 	`json:"id"`
	Email		string		`json:"email"`		
	GoogleID	*string		`json:"google_id,omitempty"`
	FullName	string		`json:"full_name"`
	PhoneNumber	*string		`json:"phone_number,omitempty"`
	Role		string		`json:"role"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

func ToResponse(user models.User) UserResponse {
	return UserResponse{
		ID:          user.ID,
		Email:       user.Email,
		GoogleID:    user.GoogleID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		Role:        string(user.Role),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}