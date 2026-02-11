package bestpractices

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestRawRun(t *testing.T) {
	ctx := context.Background()

	ctr, err := testcontainers.Run(ctx, "mysql:8.0",
		testcontainers.WithEnv(map[string]string{
			"MYSQL_ROOT_PASSWORD": "password",
			"MYSQL_DATABASE":      "testdb",
		}),
		testcontainers.WithExposedPorts("3306/tcp"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("port: 3306  MySQL Community Server"),
		),
	)
	if err != nil {
		t.Fatal(err)
	}
	defer ctr.Terminate(ctx)

	// Use container.
	host, _ := ctr.Host(ctx)
	port, _ := ctr.MappedPort(ctx, "3306/tcp")
	t.Logf("MySQL available at %s:%s", host, port.Port())
}
