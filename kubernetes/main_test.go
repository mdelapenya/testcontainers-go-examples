package kubernetes

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/k3s"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
)

//go:embed scripts/k6.js
var k6Script string

//go:embed scripts/playwright.spec.js
var playwrightScript string

func TestKubernetesDeployment(t *testing.T) {
	// Create a shared network (like QuickPizza's quickpizza_network)
	nw, err := network.New(t.Context())
	testcontainers.CleanupNetwork(t, nw)
	require.NoError(t, err)

	// start a k3s container
	k3sAlias := "k3s"
	k3sContainer, err := k3s.Run(
		t.Context(),
		"rancher/k3s:v1.35.0-k3s1",
		k3s.WithManifest(filepath.Join("manifests", "quickpizza.yaml")),
		testcontainers.WithExposedPorts("3333/tcp"),
		network.WithNetwork([]string{k3sAlias}, nw),
	)
	testcontainers.CleanupContainer(t, k3sContainer)
	require.NoError(t, err)

	err = k3sContainer.CopyFileToContainer(
		t.Context(),
		"./manifests/quickpizza.yaml",
		"quickpizza.yaml",
		0x644,
	)
	require.NoError(t, err)

	t.Logf("🍕 Quickpizza manifest applied successfully")

	rc, stdout, err := k3sContainer.Exec(
		t.Context(),
		[]string{
			"kubectl",
			"wait",
			"pods",
			"--all",
			"--for=condition=Ready",
			"--timeout=90s",
		},
	)
	require.NoError(t, err)
	if rc != 0 {
		output := bytes.Buffer{}
		_, err = io.Copy(&output, stdout)
		require.NoError(t, err)
		t.Fatalf("pods not ready \n%s\n", output.String())
	}
	t.Logf("🍕 Quickpizza pods are ready")

	frontEndUrl := fmt.Sprintf("http://%s:3333", k3sAlias)

	// Test if the endpoint is actually reachable
	t.Logf("Frontend URL: %s", frontEndUrl)

	t.Run("k6 tests", func(t *testing.T) {
		// Run k6 tests using the official Grafana k6 Docker image
		// Following the pattern: docker run -i --network=quickpizza_network -e BASE_URL=http://quickpizza:3333 grafana/k6 run - <01.basic.js
		k6Container, err := testcontainers.GenericContainer(t.Context(), testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "grafana/k6:latest",
				Files: []testcontainers.ContainerFile{
					{
						Reader:            strings.NewReader(k6Script),
						ContainerFilePath: "/scripts/k6.js",
						FileMode:          0o644,
					},
				},
				Env: map[string]string{
					"FRONTEND_URL": frontEndUrl,
				},
				Cmd: []string{"run", "/scripts/k6.js"},
				// Wait for the container to exit
				WaitingFor: wait.ForExit(),
				// Connect to the same network as k3s
				Networks: []string{nw.Name},
			},
			Started: true,
		})
		testcontainers.CleanupContainer(t, k6Container)
		require.NoError(t, err)

		// Check the test results
		state, err := k6Container.State(t.Context())
		require.NoError(t, err)

		if state.ExitCode != 0 {
			logs := bytes.Buffer{}
			logReader, err := k6Container.Logs(t.Context())
			if err != nil {
				t.Logf("getting logs %v", err)
			} else {
				logs.ReadFrom(logReader)
			}

			t.Fatalf("k6 tests failed with exit code %d\n%s\n", state.ExitCode, logs.String())
		}

		t.Logf("✅ All k6 tests passed!")
	})

	t.Run("playwright tests", func(t *testing.T) {
		// Run Playwright tests using the official Playwright Docker image
		playwrightContainer, err := testcontainers.GenericContainer(t.Context(), testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "mcr.microsoft.com/playwright:v1.58.0-noble",
				Files: []testcontainers.ContainerFile{
					{
						Reader:            strings.NewReader(playwrightScript),
						ContainerFilePath: "/tests/playwright.spec.js",
						FileMode:          0o644,
					},
				},
				Env: map[string]string{
					"FRONTEND_URL": frontEndUrl,
				},
				// Install Playwright and run tests
				Cmd: []string{
					"sh", "-c",
					"cd /tests && npm init -y && npm install @playwright/test && npx playwright test playwright.spec.js",
				},
				// Wait for the container to exit
				WaitingFor: wait.ForExit(),
				// Connect to the same network as k3s
				Networks: []string{nw.Name},
			},
			Started: true,
		})
		testcontainers.CleanupContainer(t, playwrightContainer)
		require.NoError(t, err)

		// Check the test results
		state, err := playwrightContainer.State(t.Context())
		require.NoError(t, err)

		if state.ExitCode != 0 {
			logs := bytes.Buffer{}
			logReader, err := playwrightContainer.Logs(t.Context())
			if err != nil {
				t.Logf("getting logs %v", err)
			} else {
				logs.ReadFrom(logReader)
			}

			t.Fatalf("Playwright tests failed with exit code %d\n%s\n", state.ExitCode, logs.String())
		}

		t.Logf("✅ All Playwright tests passed!")
	})
}
