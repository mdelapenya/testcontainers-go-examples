package bestpractices

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"testing"

	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func init() {
	// Suppress MySQL driver logs during connection retries in wait strategies.
	mysqldriver.SetLogger(log.New(nopWriter{}, "", 0))
}

type nopWriter struct{}

func (nopWriter) Write(p []byte) (int, error) { return len(p), nil }

// Isolated databases per subtest - no data conflicts.

const (
	mysqlPassword     = "testpass"
	mysqlRootPassword = "r00tPWd"
	mysqlRootUser     = "root"
)

func TestUserOperationsIsolated(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), mysqlImage,
		mysql.WithUsername(mysqlRootUser),
		mysql.WithPassword(mysqlPassword),
	)
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	// ✅ Each subtest gets its own isolated database.
	t.Run("CreateUser", func(t *testing.T) {
		t.Parallel()
		db := openUniqueDB(t, ctr) // Creates unique DB for this subtest.
		testCreateUserIsolated(t, db)
	})

	t.Run("UpdateUser", func(t *testing.T) {
		t.Parallel()
		db := openUniqueDB(t, ctr) // Creates unique DB for this subtest.
		testUpdateUserIsolated(t, db)
	})

	t.Run("DeleteUser", func(t *testing.T) {
		t.Parallel()
		db := openUniqueDB(t, ctr) // Creates unique DB for this subtest.
		testDeleteUserIsolated(t, db)
	})
}

// openUniqueDB creates a unique database for each test to avoid data conflicts.
func openUniqueDB(t *testing.T, ctr *mysql.MySQLContainer) *sql.DB {
	t.Helper()
	ctx := t.Context()

	// Generate unique database name using test name.
	dbName := strings.ReplaceAll(t.Name(), "/", "_")
	dbName = strings.ReplaceAll(dbName, "-", "_")

	connStr, err := ctr.ConnectionString(ctx)
	require.NoError(t, err)

	db, err := sql.Open("mysql", connStr)
	require.NoError(t, err)
	t.Cleanup(func() {
		require.NoError(t, db.Close())
	})

	// Create isolated database for this test.
	_, err = db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
	require.NoError(t, err)
	t.Cleanup(func() {
		// Don't use t.Context() in cleanup functions.
		_, err := db.ExecContext(context.Background(), fmt.Sprintf("DROP DATABASE IF EXISTS `%s`", dbName))
		require.NoError(t, err)
	})

	// Connect to the new database.
	_, err = db.ExecContext(ctx, fmt.Sprintf("USE `%s`", dbName))
	require.NoError(t, err)

	return db
}

func testCreateUserIsolated(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}

func testUpdateUserIsolated(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}

func testDeleteUserIsolated(t *testing.T, db *sql.DB) {
	t.Helper()
	// Test implementation.
	_ = db
}
