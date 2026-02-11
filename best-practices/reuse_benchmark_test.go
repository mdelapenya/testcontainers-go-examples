package bestpractices

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

const reuseBenchmarkTests = 10

// BenchmarkReuse compares sync.OnceValues vs pure WithReuseByName for container reuse.
func BenchmarkReuse(b *testing.B) {
	// Pure reuse: each call goes through postgres.Run but reuses existing container.
	b.Run("reuse=pure", func(b *testing.B) {
		for b.Loop() {
			for range reuseBenchmarkTests {
				ctr, err := postgres.Run(b.Context(), postgresImage,
					testcontainers.WithReuseByName("bench-postgres-pure"),
				)
				require.NoError(b, err)
				require.True(b, ctr.IsRunning())
			}
		}
	})

	// OnceValues: container is started once and cached in memory.
	b.Run("reuse=once", func(b *testing.B) {
		var getPostgresOnce = sync.OnceValue(func() *postgres.PostgresContainer {
			ctr, err := postgres.Run(context.Background(), postgresImage,
				testcontainers.WithReuseByName("bench-postgres-once"),
			)
			require.NoError(b, err)
			return ctr
		})

		for b.Loop() {
			for range reuseBenchmarkTests {
				ctr := getPostgresOnce()
				require.True(b, ctr.IsRunning())
			}
		}
	})
}
