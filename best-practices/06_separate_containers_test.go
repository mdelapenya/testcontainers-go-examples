package bestpractices

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

// Each test starts its own container - slow!

func TestCreateUser(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), "mysql:8.0")
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	db := openDBSeparate(t, ctr)
	// Test create user...
	_ = db
}

func TestUpdateUser(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), "mysql:8.0")
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	db := openDBSeparate(t, ctr)
	// Test update user...
	_ = db
}

func TestDeleteUser(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), "mysql:8.0")
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	db := openDBSeparate(t, ctr)
	// Test delete user...
	_ = db
}

func openDBSeparate(t *testing.T, ctr *mysql.MySQLContainer) any {
	t.Helper()
	// Simplified for example.
	_ = ctr
	return nil
}
