//go:build dev

package config

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/contrib/testcontainers"
	"github.com/gofiber/fiber/v3"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

// ConfigureApp configures the fiber app, including the database connection string.
// The connection string is retrieved from the environment variable DB, or using
// tries to connect to a local postgres instance if the environment variable is not set.
func ConfigureApp(cfg fiber.Config) (*AppConfig, error) {
	// Define a context provider for the services startup.
	// This is useful to cancel the startup of the services if the context is canceled.
	// Default is context.Background().
	startupCtx, startupCancel := context.WithTimeout(context.Background(), 10*time.Second)
	cfg.ServicesStartupContextProvider = func() context.Context {
		return startupCtx
	}

	// Define a context provider for the services shutdown.
	// This is useful to cancel the shutdown of the services if the context is canceled.
	// Default is context.Background().
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	cfg.ServicesShutdownContextProvider = func() context.Context {
		return shutdownCtx
	}

	// Add the Postgres service to the app, including custom configuration.
	srv, err := setupPostgres(&cfg)
	if err != nil {
		startupCancel()
		shutdownCancel()
		return nil, fmt.Errorf("add postgres service: %w", err)
	}

	app := fiber.New(cfg)

	// Retrieve the Postgres service from the app, using the service key.
	postgresSrv := fiber.MustGetService[*testcontainers.ContainerService[*postgres.PostgresContainer]](app.State(), srv.Key())

	connString, err := postgresSrv.Container().ConnectionString(context.Background())
	if err != nil {
		startupCancel()
		shutdownCancel()
		return nil, fmt.Errorf("get postgres connection string: %w", err)
	}

	// Override the default database connection string with the one from the Testcontainers service.
	DB = connString

	return &AppConfig{
		App:            app,
		StartupCancel:  startupCancel,
		ShutdownCancel: shutdownCancel,
	}, nil
}

// setupPostgres adds a Postgres service to the app, including custom configuration to allow
// reusing the same container while developing locally.
func setupPostgres(cfg *fiber.Config) (*testcontainers.ContainerService[*postgres.PostgresContainer], error) {
	// Add the Postgres service to the app, including custom configuration.
	srv, err := testcontainers.AddService(cfg, testcontainers.NewModuleConfig(
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
		return nil, fmt.Errorf("add postgres service: %w", err)
	}

	return srv, nil
}
