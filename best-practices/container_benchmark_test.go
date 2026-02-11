package bestpractices

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/docker/go-connections/nat"
	mysqldriver "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/wait"
)

const numberOfTests = 10

var waitForMySQL = wait.ForSQL("3306/tcp", "mysql", func(host string, port nat.Port) string {
	cfg := mysqldriver.Config{
		User:   mysqlRootUser,
		Passwd: mysqlPassword,
		Net:    "tcp",
		Addr:   net.JoinHostPort(host, port.Port()),
	}
	return cfg.FormatDSN()
}).WithStartupTimeout(60 * time.Second)

// BenchmarkContainer benchmarks different container management strategies: per-test vs shared container.
func BenchmarkContainer(b *testing.B) {
	ctx := b.Context()

	b.Run("container=per-test", func(b *testing.B) {
		for b.Loop() {
			for j := range numberOfTests {
				ctr, err := mysql.Run(ctx, mysqlImage,
					mysql.WithUsername(mysqlRootUser),
					mysql.WithPassword(mysqlPassword),
					testcontainers.WithWaitStrategy(waitForMySQL),
				)
				require.NoError(b, err)

				connStr, err := ctr.ConnectionString(ctx)
				require.NoError(b, err)

				dbName := fmt.Sprintf("bench_per_%d", j)
				require.NoError(b, runSubtest(ctx, connStr, dbName))
				require.NoError(b, testcontainers.TerminateContainer(ctr))
			}
		}
	})

	b.Run("container=shared", func(b *testing.B) {
		for b.Loop() {
			ctr, err := mysql.Run(ctx, mysqlImage,
				mysql.WithUsername(mysqlRootUser),
				mysql.WithPassword(mysqlPassword),
				testcontainers.WithWaitStrategy(waitForMySQL),
			)
			testcontainers.CleanupContainer(b, ctr)
			require.NoError(b, err)

			connStr, err := ctr.ConnectionString(ctx)
			require.NoError(b, err)
			for j := range numberOfTests {
				dbName := fmt.Sprintf("bench_single_%d", j)
				require.NoError(b, runSubtest(ctx, connStr, dbName))
			}
			require.NoError(b, testcontainers.TerminateContainer(ctr))
		}
	})
}

// runSubtest simulates a test operation: create a database, run a query, drop the database.
func runSubtest(ctx context.Context, connStr, dbName string) error {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create isolated database
	if _, err := db.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE `%s`", dbName)); err != nil {
		return err
	}

	// Use the database and run a simple operation
	if _, err := db.ExecContext(ctx, fmt.Sprintf("USE `%s`", dbName)); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, "CREATE TABLE users (id INT PRIMARY KEY, name VARCHAR(100))"); err != nil {
		return err
	}

	if _, err := db.ExecContext(ctx, "INSERT INTO users (id, name) VALUES (1, 'test')"); err != nil {
		return err
	}

	var name string
	if err := db.QueryRowContext(ctx, "SELECT name FROM users WHERE id = 1").Scan(&name); err != nil {
		return err
	}

	// Cleanup
	if _, err := db.ExecContext(ctx, fmt.Sprintf("DROP DATABASE `%s`", dbName)); err != nil {
		return err
	}

	return nil
}
