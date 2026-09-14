package models

import (
	"time"
	"github.com/google/uuid"
)

type Artist struct {
	ID 	  		uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description *string   `gorm:"type:text" json:"description"`
	CreatedByID uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedBY 	User      `gorm:"foreignKey:CreatedByID" json:"created_by_user"`
	Staff       []User    `gorm:"many2many:artist_staffs;" json:"staff"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}