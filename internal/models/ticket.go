package models

import (
	"time"
	"github.com/google/uuid"
)

type TicketStatus string

const (
	TicketPending 	TicketStatus = "pending"
	TicketPaid	  	TicketStatus = "paid"
	TicketUsed	  	TicketStatus = "used"
	TicketRefunded	TicketStatus = "refunded"
	TicketExpired 	TicketStatus = "expired"
)

type Ticket struct {
	ID 			uuid.UUID 		`gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	OrderID		uuid.UUID 		`gorm:"type:uuid;not null" json:"order_id"`
	CategoryID 	uuid.UUID 		`gorm:"type:uuuid;not null" json:"category_id"`
	TicketCode	string	  		`gorm:"type:varchar(50);uniqueIndex;not null" json:"ticket_code"`
	QRCode		*string	  		`gorm:"type:text" json:"qr_code,omitempty"`
	Status 		TicketStatus	`gorm:"type:ticket_status;default:'pending';not null" json:"status"`
	ScannedAt	*time.Time		`gorm:"scanned_at, omitempty"`
	Order    	Order         	`gorm:"foreignKey:OrderID" json:"-"`
	Category 	TicketCategory 	`gorm:"foreignKey:CategoryID" json:"category"`
}