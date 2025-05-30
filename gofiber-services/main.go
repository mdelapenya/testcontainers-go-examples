package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/app/dal"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/app/routes"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/config"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/config/database"
	"github.com/mdelapenya/testcontainers-go-examples/gofiber-services/utils"
)

func main() {
	cfg := fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	}

	appConfig, err := config.ConfigureApp(cfg)
	if err != nil {
		panic(err)
	}

	app := appConfig.App
	defer appConfig.StartupCancel()
	defer appConfig.ShutdownCancel()

	database.Connect(config.DB)
	if err := database.Migrate(&dal.User{}, &dal.Todo{}); err != nil {
		panic(err)
	}

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(fmt.Sprintf(":%v", config.PORT)); err != nil {
			log.Panic(err)
		}
	}()

	// Only handle signals if not running in development mode
	if os.Getenv("APP_ENV") != "dev" {
		quit := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

		<-quit // This blocks the main thread until an interrupt is received
		fmt.Println("Gracefully shutting down...")
		err = app.Shutdown()
		if err != nil {
			log.Panic(err)
		}
	} else {
		// In development mode, just block forever
		select {}
	}
}
