package bestpractices

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestUseModules(t *testing.T) {
	ctx := context.Background()

	// ✅ Use the MySQL module - sensible defaults, less code.
	ctr, err := mysql.Run(ctx, "mysql:8.0")
	if err != nil {
		t.Fatal(err)
	}
	defer ctr.Terminate(ctx)

	// ✅ Module provides connection string helper.
	connStr, err := ctr.ConnectionString(ctx)
	require.NoError(t, err)
	// Use connStr.
	_ = connStr
}
