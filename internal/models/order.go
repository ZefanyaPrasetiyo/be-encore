package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderPending 	OrderStatus = "pending"
	OrderPaid		OrderStatus = "paid"
	OrderCanceled	OrderStatus = "cancelled"
	OrderRefunded	OrderStatus = "refunded"
)


type Order struct {
	ID				uuid.UUID 	`gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	BuyerID			uuid.UUID 	`gorm:"type:uuid;not null" json:"buyer_id"`
	EventID			uuid.UUID 	`gorm:"type:uuid;not null" json:"event_id"`
	TotalAmount		float64		`gorm:"type:numeric(15,2);not null" json:"total_amount"`
	Status			OrderStatus	`gorm:"type:order_status;default:'pending';not null" json:"status"` 	
	PaymentMethod	*string		`gorm:"type:varchar(50)" json:"payment_method,omitempty"`
	BuyerUser		User		`gorm:"foreignKey:BuyerID" json:"buyer"`
	Event			Event		`gorm:"foreignKey:EventID" json:"event"`
	CreatedAt		time.Time	`gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt		time.Time	`gorm:"autoUpdateTime" json:"updated_at"`
}