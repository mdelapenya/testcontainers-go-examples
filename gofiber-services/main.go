package main

import (
	"fmt"

	"testcontainers-go-examples/gotodo/app/dal"
	"testcontainers-go-examples/gotodo/app/routes"
	"testcontainers-go-examples/gotodo/config"
	"testcontainers-go-examples/gotodo/config/database"
	"testcontainers-go-examples/gotodo/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	database.Connect()
	database.Migrate(&dal.User{}, &dal.Todo{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	app.Listen(fmt.Sprintf(":%v", config.PORT))
}
