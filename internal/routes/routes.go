package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"       
	"auth-go/internal/handlers"
)
func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api")

	api.Post("/register", handlers.Register(db))
	api.Post("/login", handlers.Login(db))
	api.Get("/logout", handlers.Logout())
}