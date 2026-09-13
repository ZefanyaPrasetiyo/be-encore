package models


import (
	"time"
	"github.com/google/uuid"
)

type userRole string

const (
	RoleAdmin userRole = "admin"
	RoleStaff userRole = "staff"
	RoleBuyer userRole = "buyer"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	PasswordHash *string   `gorm:"type:varchar(255)" json:"-"`
	GoogleID     *string   `gorm:"type:varchar(255);uniqueIndex" json:"google_id,omitempty"`
	FullName     string    `gorm:"type:varchar(100);not null" json:"full_name"`
	PhoneNumber  *string   `gorm:"type:varchar(20)" json:"phone_number,omitempty"`
	Role         userRole  `gorm:"type:user_role;not null" json:"role"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}