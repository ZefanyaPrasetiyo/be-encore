package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	
	"encore-be/internal/config"
	"encore-be/internal/models"
	"encore-be/internal/modules/user"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Gagal konek database: ", err)
	}

	enumQuery := `
		DO $$ 
		BEGIN 
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN 
				CREATE TYPE user_role AS ENUM ('admin', 'staff', 'buyer'); 
			END IF; 
		END $$;
	`
	db.Exec(enumQuery)

	err = db.AutoMigrate(&models.User{}, &models.Venue{}, &models.Artist{})
	if err != nil {
		log.Fatal("Gagal migrasi model: ", err)
	}
	
	config.SeedUsers(db)

	userService := user.NewService(db)
	userController := user.NewController(userService)

	app := fiber.New()

	api := app.Group("/api/v1")
	api.Get("/users", userController.GetAll)

	log.Println("Server Fiber berjalan di port 8080...")
	log.Fatal(app.Listen(":8080")) 
}