package models

import (
	"time"
	"github.com/google/uuid"
)


type VenueCategory struct {
	ID 	  uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name  string    `gorm:"type:varchar(100);not null" json:"name"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type Venue struct {
	ID 		 uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Name     string    `gorm:"type:varchar(150);not null" json:"name"`
	Address  string    `gorm:"type:varchar(255);not null" json:"address"`
	City	 string    `gorm:"type:varchar(100);not null" json:"city"`
	Capacity int       `gorm:"type:int;not null" json:"capacity"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
