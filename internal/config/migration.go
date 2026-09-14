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
		END $$;

		DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'event_status') THEN 
				CREATE TYPE event_status AS ENUM ('draft', 'published', 'ended'); 
			END IF; 
		END $$;
	`
	db.Exec(enumQuery)

	err := db.AutoMigrate(&models.User{}, &models.Venue{}, &models.Artist{}, &models.Event{})
	if err != nil {
		log.Fatal("Gagal migrasi model:", err)
	}

	log.Println("migrasi database berhasil")
}