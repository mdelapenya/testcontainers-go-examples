package bestpractices

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func TestCleanupContainer(t *testing.T) {
	ctr, err := mysql.Run(t.Context(), "mysql:8.0")
	// ✅ CleanupContainer BEFORE error check - handles partial container starts.
	testcontainers.CleanupContainer(t, ctr)
	require.NoError(t, err)

	// ✅ Module provides connection string helper.
	connStr, err := ctr.ConnectionString(t.Context())
	require.NoError(t, err)
	// Use connStr.
	_ = connStr
}
