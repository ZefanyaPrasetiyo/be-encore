package models

import (
	"time"
	"github.com/google/uuid"
)

type TicketCategory struct {
	ID 			uuid.UUID	`gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	EventID 	uuid.UUID	`gorm:"type:uuid;not null" json:"event_id"`
	Name 		string		`gorm:"type:varchar(100);not null" json:"name"`
	Price 		float64		`gorm:"type:numeric(15,2);not null" json:"price"`
	Quota 		int 		`gorm:"type:int;not null" json:"quota"`
	Event 		Event		`gorm:"foreignKey:EventID" json:"-"`
	CreatedAt 	time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt	time.Time	`gorm:"autoUpdateTime" json:"updated_at"`

}