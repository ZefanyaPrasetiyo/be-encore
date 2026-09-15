package config

import (
	"log"
	"gorm.io/gorm"
	"encore-be/internal/models"
)

func MigrateDB(db *gorm.DB) {
	enumQuery := `
	DO $$ 
	BEGIN 
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN 
			CREATE TYPE user_role AS ENUM ('admin', 'staff', 'buyer'); 
		END IF; 
		
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'event_status') THEN 
			CREATE TYPE event_status AS ENUM ('draft', 'published', 'ended', 'cancelled'); 
		END IF;
		
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN 
			CREATE TYPE order_status AS ENUM ('pending', 'paid', 'cancelled', 'refunded'); 
		END IF;
		
		IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'ticket_status') THEN 
			CREATE TYPE ticket_status AS ENUM ('pending', 'paid', 'used', 'expired'); 
		END IF;
	END $$;
	`
	db.Exec(enumQuery)

	err := db.AutoMigrate(&models.User{}, &models.Venue{},&models.VenueCategory{}, &models.Artist{}, &models.Event{}, &models.Order{}, &models.Ticket{}, &models.TicketCategory{})
	if err != nil {
		log.Fatal("Gagal migrasi model:", err)
	}

	log.Println("migrasi database berhasil")
}