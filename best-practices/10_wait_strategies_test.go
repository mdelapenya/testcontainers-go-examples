package bestpractices

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestWaitStrategies(t *testing.T) {
	ctx := t.Context()

	t.Run("nginx", func(t *testing.T) {
		// ✅ Best: HTTP health check.
		webApp, err := testcontainers.Run(ctx, "nginx:latest",
			testcontainers.WithExposedPorts("80/tcp"),
			testcontainers.WithWaitStrategy(
				wait.ForHTTP("/").WithPort("80/tcp").WithStartupTimeout(30*time.Second),
			),
		)
		testcontainers.CleanupContainer(t, webApp)
		require.NoError(t, err)
	})

	t.Run("postgres", func(t *testing.T) {
		// ✅ Good: SQL ping check.
		dbCtr, err := postgres.Run(ctx, "postgres:16",
			testcontainers.WithWaitStrategy(
				wait.ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
					connURL := url.URL{
						Scheme:   "postgres",
						User:     url.UserPassword("postgres", "postgres"),
						Host:     net.JoinHostPort(host, port.Port()),
						Path:     "postgres",
						RawQuery: "sslmode=disable",
					}
					return connURL.String()
				}).WithStartupTimeout(60*time.Second),
			),
		)
		testcontainers.CleanupContainer(t, dbCtr)
		require.NoError(t, err)
	})

	// ❌ Avoid: Log-based checks (logs can change between versions).
	// wait.ForLog("Ready to accept connections")
}
