package main

import (
	"log"

	"github.com/Rokiis1/trivia-api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
    database.ConnectDb()

    app := fiber.New()
    
    // setupRoutes(app)

    log.Fatal(app.Listen(":4000"))
}
