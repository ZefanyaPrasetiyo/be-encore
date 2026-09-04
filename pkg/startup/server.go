package startup

import (
	  "github.com/gofiber/fiber/v3"
	"log"
)

func StartServer(app *fiber.App){
	log.Fatal(app.Listen(":5000"))
}