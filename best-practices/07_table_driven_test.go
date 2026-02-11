package bestpractices

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

// Table-driven tests - poor IDE integration.

func TestUserOperationsTableDriven(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), "mysql:8.0")
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	db := openDBTable(t, ctr)

	tests := []struct {
		name string
		fn   func(t *testing.T, db *sql.DB)
	}{
		{"CreateUser", testCreateUserTable},
		{"UpdateUser", testUpdateUserTable},
		{"DeleteUser", testDeleteUserTable},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.fn(t, db)
		})
	}
}

func openDBTable(t *testing.T, ctr *mysql.MySQLContainer) *sql.DB {
	t.Helper()
	// Simplified for example.
	_ = ctr
	return nil
}

func testCreateUserTable(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}

func testUpdateUserTable(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}

func testDeleteUserTable(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}
