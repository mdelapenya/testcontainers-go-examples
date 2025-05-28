package routes

import (
	"github.com/gofiber/fiber/v3"
	"testcontainers-go-examples/gotodo/app/services"
	"testcontainers-go-examples/gotodo/utils/middleware"
)

// TodoRoutes contains all routes relative to /todo
func TodoRoutes(app fiber.Router) {
	r := app.Group("/todo").Use(middleware.Auth)

	r.Post("/create", services.CreateTodo)
	r.Get("/list", services.GetTodos)
	r.Get("/:todoID", services.GetTodo)
	r.Patch("/:todoID", services.UpdateTodoTitle)
	r.Patch("/:todoID/check", services.CheckTodo)
	r.Delete("/:todoID", services.DeleteTodo)
}
