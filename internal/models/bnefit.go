package models

import (
	"time"
	"github.com/google/uuid"
)

type Benefit struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(150);not null" json:"name"`
	Description *string   `gorm:"type:text" json:"description,omitempty"`

	TicketCategories []TicketCategory `gorm:"many2many:ticket_category_benefits" json:"-"`
	
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}