package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/contrib/testcontainers"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

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

	// Define a context provider for the services startup.
	// This is useful to cancel the startup of the services if the context is canceled.
	// Default is context.Background().
	startupCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cfg.ServicesStartupContextProvider = func() context.Context {
		return startupCtx
	}

	// Define a context provider for the services shutdown.
	// This is useful to cancel the shutdown of the services if the context is canceled.
	// Default is context.Background().
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cfg.ServicesShutdownContextProvider = func() context.Context {
		return shutdownCtx
	}

	// Add the Postgres service to the app, including custom configuration.
	srv, err := testcontainers.AddService(&cfg, testcontainers.NewModuleConfig(
		"postgres-db",
		"postgres:16",
		postgres.Run,
		postgres.BasicWaitStrategies(),
		postgres.WithDatabase("todos"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		tc.WithReuseByName("postgres-db-todos"),
	))
	if err != nil {
		panic(err)
	}

	app := fiber.New(cfg)

	// Retrieve the Postgres service from the app, using the service key.
	postgresSrv := fiber.MustGetService[*testcontainers.ContainerService[*postgres.PostgresContainer]](app.State(), srv.Key())

	connString, err := postgresSrv.Container().ConnectionString(context.Background())
	if err != nil {
		panic(err)
	}

	// Override the default database connection string with the one from the Testcontainers service.
	config.DB = connString

	database.Connect(config.DB)
	if err := database.Migrate(&dal.User{}, &dal.Todo{}); err != nil {
		panic(err)
	}

	app.Use(logger.New())

	routes.AuthRoutes(app)
	routes.TodoRoutes(app)

	if err := app.Listen(fmt.Sprintf(":%v", config.PORT)); err != nil {
		panic(err)
	}
}
