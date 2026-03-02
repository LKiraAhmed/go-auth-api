package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"auth-go/internal/database"
	"auth-go/internal/routes"
)

func main() {
	database.ConnectDatabase()

	app := fiber.New()

	routes.SetupRoutes(app, database.DB)

	log.Fatal(app.Listen("127.0.0.1:8080"))
}