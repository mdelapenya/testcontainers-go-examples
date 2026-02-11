---
marp: true
theme: default
paginate: true
backgroundColor: #fff
style: |
  section {
    font-size: 28px;
  }
  code {
    font-size: 18px;
  }
  pre {
    font-size: 16px;
  }
  h1 {
    color: #2d6df6;
  }
  h2 {
    color: #333;
  }
---

# Testcontainers-Go Best Practices

## Writing Fast, Reliable Integration Tests

---

## Why Best Practices Matter

- Slow tests = skipped tests
- Flaky tests erode confidence
- Modern Go + testcontainers-go = powerful combo

---

## Starting Point - Basic Container

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
    if err != nil { t.Fatal(err) }
    defer ctr.Terminate(ctx)
    // Use container...
}
```

---

## What Can We Improve?

❌ Lots of boilerplate for MySQL setup
❌ Using `defer` for cleanup (problematic in subtests)
❌ Manual error handling pattern
❌ Using `context.Background()`
❌ Using `wait.ForLog` is fragile

---

## Improvement #1 - Use Modules

Pre-configured modules for popular services (MySQL, Redis, Postgres, etc.)

```go
func TestUseModules(t *testing.T) {
    ctx := context.Background()

    // ✅ Use the MySQL module - sensible defaults, less code
    ctr, err := mysql.Run(ctx, "mysql:8.0")
    if err != nil { t.Fatal(err) }
    defer ctr.Terminate(ctx)

    // ✅ Module provides connection string helper
    connStr, err := ctr.ConnectionString(ctx)
    require.NoError(t, err)
}
```

---

## Modules - What's Still Wrong?

✅ Much less boilerplate!
❌ Still using `defer` for cleanup
❌ Inconsistent error handling
❌ Using `context.Background()` instead of test context

---

## Improvement #2 - Proper Cleanup

Use `testcontainers.CleanupContainer()` for safe cleanup

```go
func TestCleanupContainer(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")

    // ✅ CleanupContainer BEFORE error check
    // Handles partial container starts
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    // ✅ Module provides connection string helper
    connStr, err := ctr.ConnectionString(t.Context())
    require.NoError(t, err)
}
```

---

## Cleanup - Progress Check

✅ Much less boilerplate!
✅ Proper cleanup with `CleanupContainer`
✅ Using `t.Context()` for test lifecycle
❌ Each test starts its own container - slow!

---

## The Problem - One Container Per Test

```go
func TestCreateUser(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)
    // Test create user...
}

func TestUpdateUser(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)
    // Test update user...
}
```

**Issues:** 3 tests = 3 container starts = SLOW 🐢

---

## Table-Driven Tests?

```go
func TestUserOperationsTableDriven(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    tests := []struct {
        name string
        fn   func(t *testing.T, db *sql.DB)
    }{
        {"CreateUser", testCreateUser},
        {"UpdateUser", testUpdateUser},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) { tc.fn(t, db) })
    }
}
```

---

## Table-Driven - Trade-offs

✅ Single container for all tests - fast!
❌ Poor IDE integration - can't click to run a single test
❌ Harder to debug - which test case failed?
❌ Loop variable capture issues (pre-Go 1.22)

---

## Improvement #3 - Explicit Subtests

Best of both worlds - one container, explicit `t.Run()` calls:

```go
func TestUserOperationsExplicit(t *testing.T) {
    ctr, err := mysql.Run(t.Context(), "mysql:8.0")
    testcontainers.CleanupContainer(t, ctr)
    require.NoError(t, err)

    db := openDB(t, ctr)

    // ✅ Explicit subtests - IDE friendly, easy to debug
    t.Run("CreateUser", func(t *testing.T) { testCreateUser(t, db) })
    t.Run("UpdateUser", func(t *testing.T) { testUpdateUser(t, db) })
    t.Run("DeleteUser", func(t *testing.T) { testDeleteUser(t, db) })
}
```

---

## Explicit Subtests - Benefits

✅ Single container - fast!
✅ Click to run any subtest in IDE
✅ Clear test names in output
✅ Easy to add/remove tests
❌ Subtests share same database - data conflicts!

---

## Improvement #4 - Isolated Databases

Create a unique database for each subtest:

```go
t.Run("CreateUser", func(t *testing.T) {
    t.Parallel()
    db := openUniqueDB(t, ctr) // Creates unique DB for this subtest
    testCreateUser(t, db)
})

t.Run("UpdateUser", func(t *testing.T) {
    t.Parallel()
    db := openUniqueDB(t, ctr) // Creates unique DB for this subtest
    testUpdateUser(t, db)
})
```

---

## openUniqueDB Helper

```go
func openUniqueDB(t *testing.T, ctr *mysql.MySQLContainer) *sql.DB {
    t.Helper()
    ctx := t.Context()

    // Generate unique database name using test name
    dbName := strings.ReplaceAll(t.Name(), "/", "_")

    db, err := sql.Open("mysql", connStr)
    require.NoError(t, err)
    t.Cleanup(func() { db.Close() })

    // Create isolated database for this test
    _, err = db.ExecContext(ctx,
        fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
    require.NoError(t, err)

    return db
}
```

---

## Isolated Databases - Benefits

✅ Single container - fast!
✅ Subtests can run in parallel
✅ No data conflicts between tests
✅ Each test starts with clean state

---

## Bonus: Effective Wait Strategies

```go
// ✅ Best: HTTP health check
testcontainers.WithWaitStrategy(
    wait.ForHTTP("/").WithPort("80/tcp")
         .WithStartupTimeout(30*time.Second),
)

// ✅ Good: SQL ping check
testcontainers.WithWaitStrategy(
    wait.ForSQL("5432/tcp", "postgres", connFunc)
         .WithStartupTimeout(60*time.Second),
)

// ❌ Avoid: Log-based checks (logs change between versions)
wait.ForLog("Ready to accept connections")
```

---

## Reusable Containers

```go
// Singleton pattern with sync.OnceValues
var runPostgresOnce = sync.OnceValues(func() (*postgres.PostgresContainer, error) {
    return postgres.Run(context.Background(), "postgres:16",
        testcontainers.WithReuseByName("shared-postgres"),
    )
})

func TestReuse_A(t *testing.T) {
    db, err := runPostgresOnce()
    require.NoError(t, err)
    require.True(t, db.IsRunning())
}
```

---

## Performance: Pure Reuse vs sync.OnceValues

```log
         │      pure       │                 once                  │
         │     sec/op      │     sec/op      vs base               │
Reuse-16   127.312m ± 167%   4.574m ± 1346%  -96.41% (p=0.002 n=6)
```

**Key insight:** `sync.OnceValues` avoids Docker API calls after first run, making subsequent lookups ~30x faster

---

## Performance: Per-test vs Shared Container

```log
             │  per-test   │              shared               │
             │   sec/op    │   sec/op    vs base               │
Container-16   65.237 ± 2%   6.533 ± 4%  -89.99% (p=0.002 n=6)
```

**Key insight:** Sharing containers across subtests = 10x faster

---

## Key Takeaways

✅ Use modules for less boilerplate
✅ `CleanupContainer` before error check
✅ `t.Context()` for test lifecycle
✅ Explicit subtests over table-driven tests
✅ `t.Parallel()` for speed
✅ Functional wait strategies (avoid `ForLog`)

---

## Resources & Q&A

- 📚 [testcontainers-go documentation](https://golang.testcontainers.org/)
- 💻 This repo's examples
- ❓ Questions?
