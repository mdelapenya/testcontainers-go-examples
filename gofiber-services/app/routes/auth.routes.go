package routes

import (
	"github.com/gofiber/fiber/v3"

	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/app/services"
)

// AuthRoutes contains all the auth routes
func AuthRoutes(app fiber.Router) {
	r := app.Group("/auth")

	r.Post("/signup", services.Signup)
	r.Post("/login", services.Login)
}
