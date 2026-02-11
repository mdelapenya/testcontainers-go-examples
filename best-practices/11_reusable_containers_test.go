package bestpractices

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const (
	postgresImage         = "postgres:16"
	postgresContainerName = "shared-postgres"
)

// runPostgres is a helper to start a Postgres container with reuse enabled.
func runPostgres(ctx context.Context) (*postgres.PostgresContainer, error) {
	return postgres.Run(ctx, postgresImage,
		testcontainers.WithReuseByName(postgresContainerName),
	)
}

// runPostgresOnce is a helper to start a Postgres container with reuse enabled, but only once.
var runPostgresOnce = sync.OnceValues(func() (*postgres.PostgresContainer, error) {
	return runPostgres(context.Background())
})

// runPostgresReuse is a test helper to start a Postgres container with reuse enabled.
func runPostgresReuse(t *testing.T) *postgres.PostgresContainer {
	db, err := runPostgres(t.Context())
	if err != nil {
		// If the create failed ensure it's fully cleaned up.
		testcontainers.CleanupContainer(t, db)
		t.Fatal(err)
	}

	return db
}

func TestReuse_A(t *testing.T) {
	db := runPostgresReuse(t)
	t.Log("container id", db.GetContainerID())
	require.True(t, db.IsRunning())
}

func TestReuse_B(t *testing.T) {
	db := runPostgresReuse(t)
	t.Log("container id", db.GetContainerID())
	require.True(t, db.IsRunning())
}

func TestOnce_A(t *testing.T) {
	db, err := runPostgresOnce()
	require.NoError(t, err)
	t.Log("container id", db.GetContainerID())
	require.True(t, db.IsRunning())
}

func TestOnce_B(t *testing.T) {
	db, err := runPostgresOnce()
	require.NoError(t, err)
	t.Log("container id", db.GetContainerID())
	require.True(t, db.IsRunning())
}
