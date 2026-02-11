package bestpractices

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

const mysqlImage = "mysql:8.0"

// Explicit subtests - IDE friendly, easy to debug.

func TestUserOperationsExplicit(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), mysqlImage)
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	db := openDBExplicit(t, ctr)

	// ✅ Explicit subtests - IDE friendly, easy to debug.
	t.Run("CreateUser", func(t *testing.T) {
		testCreateUserExplicit(t, db)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		testUpdateUserExplicit(t, db)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		testDeleteUserExplicit(t, db)
	})
}

func openDBExplicit(t *testing.T, ctr *mysql.MySQLContainer) *sql.DB {
	t.Helper()
	// Simplified for example.
	_ = ctr
	return nil
}

func testCreateUserExplicit(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}

func testUpdateUserExplicit(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}

func testDeleteUserExplicit(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}
