# Testcontainers-Go Best Practices

A 10-minute presentation on writing fast, reliable integration tests with testcontainers-go.

---

## Presentation Outline

### Slide 1: Title Slide (30 sec)

- **Title:** "Testcontainers-Go Best Practices"
- Subtitle: Writing Fast, Reliable Integration Tests

---

### Slide 2: Why Best Practices Matter (1 min)

- Slow tests = skipped tests
- Flaky tests erode confidence
- Modern Go + testcontainers-go = powerful combo

---

### Slide 3: Starting Point - Basic Container (1 min)

Let's start with a basic MySQL test using `testcontainers.Run`:

```go
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
```

**What can we improve?**

- ❌ Lots of boilerplate for MySQL setup
- ❌ Using `defer` for cleanup (problematic in subtests)
- ❌ Manual error handling pattern
- ❌ Using `context.Background()`
- ❌ Using `wait.ForLog` is fragile

---

### Slide 4: Improvement #1 - Use Modules (1 min)

- Pre-configured modules for popular services (MySQL, Redis, Postgres, etc.)
- Less boilerplate, sensible defaults
- Built-in connection strings and helpers

```go
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
}
```

**What can we still improve?**

- ✅ Much less boilerplate!
- ❌ Still using `defer` for cleanup
- ❌ Inconsistent error handling
- ❌ Using `context.Background()` instead of test context

---

### Slide 5: Improvement #2 - Proper Cleanup (1.5 min)

- Use `testcontainers.CleanupContainer()` for safe cleanup
- Call cleanup BEFORE checking error (handles partial starts)
- Use `t.Context()` for automatic cancellation

```go
func TestCleanupContainer(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    // ✅ CleanupContainer BEFORE error check - handles partial container starts.
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    // ✅ Module provides connection string helper.
    connStr, err := ctr.ConnectionString(t.Context())
    require.NoError(t, err)
    // Use connStr.
}
```

**What can we still improve?**

- ✅ Much less boilerplate!
- ✅ Proper cleanup with `CleanupContainer`
- ✅ Using `t.Context()` for test lifecycle
- ❌ Each test starts its own container - slow!

---

### Slide 6: The Problem - One Container Per Test (1 min)

Each test function starts and stops its own container:

```go
func TestCreateUser(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    db := openDBSeparate(t, ctr)
    // Test create user...
}

func TestUpdateUser(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    db := openDBSeparate(t, ctr)
    // Test update user...
}

func TestDeleteUser(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    db := openDBSeparate(t, ctr)
    // Test delete user...
}
```

**Issues:**

- ❌ 3 tests = 3 container starts = SLOW
- ❌ Lots of duplicated setup code
- ❌ Even with `t.Parallel()`, still wasteful

---

### Slide 7: Table-Driven Tests? (1 min)

One container, multiple test cases via table-driven pattern:

```go
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
```

**Issues:**

- ✅ Single container for all tests - fast!
- ❌ Poor IDE integration - can't click to run a single test.
- ❌ Harder to debug - which test case failed?
- ❌ Loop variable capture issues (pre-Go 1.22).

---

### Slide 8: Improvement #3 - Explicit Subtests (1 min)

Best of both worlds - one container, explicit `t.Run()` calls:

```go
const mysqlImage = "mysql:8.0"

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
```

**Benefits:**

- ✅ Single container - fast!
- ✅ Click to run any subtest in IDE.
- ✅ Clear test names in output.
- ✅ Easy to add/remove tests.
- ❌ Subtests share same database - data conflicts!

---

### Slide 9: Improvement #4 - Isolated Databases per Subtest (1 min)

Create a unique database for each subtest to avoid data conflicts:

```go
const (
    mysqlPassword = "testpass"
    mysqlRootUser = "root"
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
```

**Benefits:**

- ✅ Single container - fast!
- ✅ Subtests can run in parallel.
- ✅ No data conflicts between tests.
- ✅ Each test starts with clean state.

---

### Slide 10: Improvement #5 - Parallel Tests (1.5 min)

- Add `t.Parallel()` for concurrent test execution
- Each test gets isolated container(s)
- Dramatically faster test suites

```go
func TestMySQL(t *testing.T) {
    t.Parallel() // ✅ Run this test in parallel with others.

    ctx := t.Context()

    ctr, err := mysql.Run(ctx, "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    connStr, err := ctr.ConnectionString(ctx)
    require.NoError(t, err)
    // Use connStr.
}

func TestPostgres(t *testing.T) {
    t.Parallel() // ✅ Runs concurrently with TestMySQL.

    ctx := t.Context()

    ctr, err := postgres.Run(ctx, "postgres:16")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    connStr, err := ctr.ConnectionString(ctx)
    require.NoError(t, err)
    // Use connStr.
}
```

**What can we still improve?**

- ✅ Tests run in parallel - much faster!
- ✅ One container shared across subtests
- ✅ IDE friendly, easy to debug

---

### Slide 11: Bonus - Effective Wait Strategies (1.5 min)

- Don't guess with `time.Sleep()`
- Prefer functional checks: `wait.ForHTTP()`, `wait.ForSQL()`, `wait.ForListeningPort()`
- Avoid `wait.ForLog()` — log messages change, functionality doesn't
- Set appropriate `WithStartupTimeout()`

```go
func TestWaitStrategies(t *testing.T) {
    ctx := t.Context()

    // ✅ Best: HTTP health check.
    webApp, err := testcontainers.Run(ctx, "nginx:latest",
        testcontainers.WithExposedPorts("80/tcp"),
        testcontainers.WithWaitStrategy(
            wait.ForHTTP("/").WithPort("80/tcp").WithStartupTimeout(30*time.Second),
        ),
    )
    testcontainers.CleanupContainer(t, webApp)
    require.NoError(t, err)

    // ✅ Good: SQL ping check.
    dbCtr, err := postgres.Run(ctx, "postgres:16",
        testcontainers.WithWaitStrategy(
            wait.ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
                connURL := url.URL{
                    Scheme:   "postgres",
                    User:     url.UserPassword("user", "pass"),
                    Host:     net.JoinHostPort(host, port.Port()),
                    Path:     "db",
                    RawQuery: "sslmode=disable",
                }
                return connURL.String()
            }).WithStartupTimeout(60*time.Second),
        ),
    )
    testcontainers.CleanupContainer(t, dbCtr)
    require.NoError(t, err)

    // ❌ Avoid: Log-based checks (logs can change between versions).
    // wait.ForLog("Ready to accept connections")
}
```

---

### Slide 12: Reusable Containers & Singleton Pattern (1 min)

- Share containers across tests where safe
- `testcontainers.WithReuseByName()` for local dev speed
- Trade-off: isolation vs speed
- Note: Ryuk waits 10s after test exit before cleanup (configurable via `TESTCONTAINERS_RYUK_CONNECTION_TIMEOUT`)

```go
const postgresImage = "postgres:16"

// Singleton pattern with sync.OnceValues for expensive containers.
var runPostgresOnce = sync.OnceValues(func() (*postgres.PostgresContainer, error) {
    return postgres.Run(context.Background(), postgresImage,
        testcontainers.WithReuseByName("shared-postgres"), // Reuse across test runs (dev speed).
    )
})

// Helper to start a Postgres container with reuse enabled.
func runPostgresReuse(t *testing.T) *postgres.PostgresContainer {
    db, err := postgres.Run(t.Context(), postgresImage,
        testcontainers.WithReuseByName("shared-postgres"),
    )
    if err != nil {
        // If the create failed ensure it's fully cleaned up.
        testcontainers.CleanupContainer(t, db)
        t.Fatal(err)
    }
    return db
}

func TestReuse_A(t *testing.T) {
    db := runPostgresReuse(t)
    require.True(t, db.IsRunning())
}

func TestOnce_A(t *testing.T) {
    db, err := runPostgresOnce()
    require.NoError(t, err)
    require.True(t, db.IsRunning())
}
```

**Pure reuse vs sync.OnceValues:**

```log
goos: darwin
goarch: arm64
pkg: testcontainers-go-examples/best-practices
cpu: Apple M3 Max
         │      pure       │                 once                  │
         │     sec/op      │     sec/op      vs base               │
Reuse-16   127.312m ± 167%   4.574m ± 1346%  -96.41% (p=0.002 n=6)
```

**Key insight:** `sync.OnceValues` avoids Docker API calls after first run, making subsequent lookups ~30x faster.

---

### Slide 13: Per-test vs Shared Container Performance (30 sec)

The following shows an example of running 5 sub tests sharing a MySQL container
vs 5 separate tests each starting and stopping their own container.

```log
goos: darwin
goarch: arm64
pkg: testcontainers-go-examples/best-practices
cpu: Apple M3 Max
             │  per-test   │              shared               │
             │   sec/op    │   sec/op    vs base               │
Container-16   65.237 ± 2%   6.533 ± 4%  -89.99% (p=0.002 n=6)
```

**Key insight:** Sharing containers across subtests is a must for performance.

---

### Slide 14: Key Takeaways (30 sec)

- ✅ Use modules for less boilerplate
- ✅ `CleanupContainer` before error check
- ✅ `t.Context()` for test lifecycle
- ✅ Explicit subtests over table-driven tests
- ✅ `t.Parallel()` for speed
- ✅ Functional wait strategies (avoid `ForLog`)

---

### Slide 15: Resources & Q&A (30 sec)

- [testcontainers-go documentation](https://golang.testcontainers.org/)
- This repo's examples
- Questions?

---

Total estimated time: ~11 minutes
