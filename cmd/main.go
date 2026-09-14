package main

import (
	"flag"
	"log"

	"github.com/gofiber/fiber/v3"

	"encore-be/internal/config"
	"encore-be/internal/modules/user"
)

func main() {
	runSeed := flag.Bool("seed", false, "Jalankan database seeder")
	flag.Parse()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal("gagal konek database", err)
	}

	config.MigrateDB(db)

	if *runSeed {
		config.SeedUsers(db)
		log.Println("Seeder berhasil dijalankan. Menghentikan program...")
		return
	}

	userService := user.NewService(db)
	userController := user.NewController(userService)

	app := fiber.New()

	api := app.Group("/api/v1")
	api.Get("/users", userController.GetAll)

	log.Println("Server Fiber berjalan di port 8080...")
	log.Fatal(app.Listen(":8080")) 
}