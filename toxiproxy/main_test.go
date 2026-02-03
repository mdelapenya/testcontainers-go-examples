package toxiproxy

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	toxiclient "github.com/Shopify/toxiproxy/v2/client"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	tcredis "github.com/testcontainers/testcontainers-go/modules/redis"
	tctoxiproxy "github.com/testcontainers/testcontainers-go/modules/toxiproxy"
	"github.com/testcontainers/testcontainers-go/network"
)

func TestAddLatency(t *testing.T) {
	nw, err := network.New(t.Context())
	testcontainers.CleanupNetwork(t, nw)
	require.NoError(t, err)

	redisContainer, err := tcredis.Run(
		t.Context(),
		"redis:6-alpine",
		network.WithNetwork([]string{"redis"}, nw),
	)
	testcontainers.CleanupContainer(t, redisContainer)
	require.NoError(t, err)

	const proxyPort = "8666"

	// No need to create a proxy, as we are programmatically adding it below.
	toxiproxyContainer, err := tctoxiproxy.Run(
		t.Context(),
		"ghcr.io/shopify/toxiproxy:2.12.0",
		network.WithNetwork([]string{"toxiproxy"}, nw),
		// explicitly expose the ports that will be proxied using the programmatic API
		// of the toxiproxy client. Otherwise, the ports will not be exposed and the
		// toxiproxy client will not be able to connect to the proxy.
		testcontainers.WithExposedPorts(proxyPort+"/tcp"),
	)
	testcontainers.CleanupContainer(t, toxiproxyContainer)
	require.NoError(t, err)

	toxiURI, err := toxiproxyContainer.URI(t.Context())
	require.NoError(t, err)

	toxiproxyClient := toxiclient.NewClient(toxiURI)

	// Create the proxy using the network alias of the redis container,
	// as they run on the same network.
	proxy, err := toxiproxyClient.CreateProxy("redis", "0.0.0.0:"+proxyPort, "redis:6379")
	require.NoError(t, err)

	toxiproxyProxyPort, err := toxiproxyContainer.MappedPort(t.Context(), proxyPort+"/tcp")
	require.NoError(t, err)

	toxiproxyProxyHostIP, err := toxiproxyContainer.Host(t.Context())
	require.NoError(t, err)

	// Create a redis client that connects to the toxiproxy container.
	// We are defining a read timeout of 2 seconds, because we are adding
	// a latency toxic of 1 second to the request, +/- 200ms jitter.
	redisURI := fmt.Sprintf("redis://%s:%s?read_timeout=2s", toxiproxyProxyHostIP, toxiproxyProxyPort.Port())

	options, err := redis.ParseURL(redisURI)
	require.NoError(t, err)

	redisCli := redis.NewClient(options)
	t.Cleanup(func() {
		require.NoError(t, redisCli.FlushAll(context.TODO()).Err())
	})

	key := fmt.Sprintf("{user.%s}.favoritefood", uuid.NewString())
	value := "Cabbage Biscuits"
	ttl, err := time.ParseDuration("2h")
	require.NoError(t, err)

	err = redisCli.Set(t.Context(), key, value, ttl).Err()
	require.NoError(t, err)

	// Add Latency Toxic to the proxy
	const (
		latency = 1_000
		jitter  = 200
	)

	_, err = proxy.AddToxic("latency_down", "latency", "downstream", 1.0, toxiclient.Attributes{
		"latency": latency,
		"jitter":  jitter,
	})
	require.NoError(t, err)

	start := time.Now()
	savedValue, err := redisCli.Get(t.Context(), key).Result()
	require.NoError(t, err)

	duration := time.Since(start)

	log.Println("Duration:", duration)

	// The value is retrieved successfully
	require.Equal(t, value, savedValue)

	// Check that latency is within expected range (200ms-1200ms)
	// The latency toxic adds 1000ms (1000ms +/- 200ms jitter)
	minDuration := (latency - jitter) * time.Millisecond
	maxDuration := (latency + jitter) * time.Millisecond

	require.True(t, duration >= minDuration && duration <= maxDuration)
}

func TestConnectionCut(t *testing.T) {
	nw, err := network.New(t.Context())
	testcontainers.CleanupNetwork(t, nw)
	require.NoError(t, err)

	redisContainer, err := tcredis.Run(
		t.Context(),
		"redis:6-alpine",
		network.WithNetwork([]string{"redis"}, nw),
	)
	testcontainers.CleanupContainer(t, redisContainer)
	require.NoError(t, err)

	toxiproxyContainer, err := tctoxiproxy.Run(
		t.Context(),
		"ghcr.io/shopify/toxiproxy:2.12.0",
		// We create a proxy named "redis" that points to the redis container.
		tctoxiproxy.WithProxy("redis", "redis:6379"),
		network.WithNetwork([]string{"toxiproxy"}, nw),
	)
	testcontainers.CleanupContainer(t, toxiproxyContainer)
	require.NoError(t, err)

	proxiedRedisHost, proxiedRedisPort, err := toxiproxyContainer.ProxiedEndpoint(8666)
	require.NoError(t, err)

	toxiURI, err := toxiproxyContainer.URI(t.Context())
	require.NoError(t, err)

	toxiproxyClient := toxiclient.NewClient(toxiURI)

	// Retrieve the existing proxy
	proxies, err := toxiproxyClient.Proxies()
	require.NoError(t, err)

	proxy := proxies["redis"]

	// Create a redis client that connects to the toxiproxy container.
	// We are defining a read timeout of 2 seconds while testing proxy
	// enable/disable behavior (simulating a connection cut).
	redisURI := fmt.Sprintf("redis://%s:%s?read_timeout=2s", proxiedRedisHost, proxiedRedisPort)

	options, err := redis.ParseURL(redisURI)
	require.NoError(t, err)

	redisCli := redis.NewClient(options)
	t.Cleanup(func() {
		require.NoError(t, redisCli.FlushAll(context.TODO()).Err())
	})

	key := fmt.Sprintf("{user.%s}.favoritefood", uuid.NewString())
	value := "Cabbage Biscuits"
	ttl, err := time.ParseDuration("2h")
	require.NoError(t, err)

	err = redisCli.Set(t.Context(), key, value, ttl).Err()
	require.NoError(t, err)

	// Disable the proxy
	err = proxy.Disable()
	require.NoError(t, err)

	// Get data
	savedValue, err := redisCli.Get(t.Context(), key).Result()
	require.Error(t, err)

	// The value is not retrieved at all, so it's empty
	require.Empty(t, savedValue)

	// Re-enable the proxy
	err = proxy.Enable()
	require.NoError(t, err)

	savedValue, err = redisCli.Get(t.Context(), key).Result()
	require.NoError(t, err)

	// The value is retrieved successfully
	require.Equal(t, value, savedValue)
}
