package kubernetes

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/k3s"
	"github.com/testcontainers/testcontainers-go/network"
	"github.com/testcontainers/testcontainers-go/wait"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
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
		testcontainers.WithAfterReadyCommand(
			// The PVC wait is the most critical piece since the StatefulSet won't start without a bound PVC,
			// and this can occasionally be slow in CI environments!
			testcontainers.NewRawCommand([]string{"kubectl", "wait", "pvc", "--all", "--for=jsonpath='{.status.phase}'=Bound", "--timeout=90s"}),
			testcontainers.NewRawCommand([]string{"kubectl", "wait", "statefulset", "--all", "--for=jsonpath='{.status.readyReplicas}'=1", "--timeout=90s"}),
			testcontainers.NewRawCommand([]string{"kubectl", "wait", "deployment", "--all", "--for=condition=Available", "--timeout=90s"}),
			testcontainers.NewRawCommand([]string{"kubectl", "wait", "pods", "--all", "--for=condition=Ready", "--timeout=90s"}),
		),
	)
	testcontainers.CleanupContainer(t, k3sContainer)
	require.NoError(t, err)

	t.Logf("🍕 Quickpizza manifest applied successfully")

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
				WaitingFor: wait.ForExit().WithExitTimeout(1 * time.Minute),
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
				_, err := logs.ReadFrom(logReader)
				require.NoError(t, err)
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
				WaitingFor: wait.ForExit().WithExitTimeout(1 * time.Minute),
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
				_, err := logs.ReadFrom(logReader)
				require.NoError(t, err)
			}

			t.Fatalf("Playwright tests failed with exit code %d\n%s\n", state.ExitCode, logs.String())
		}

		t.Logf("✅ All Playwright tests passed!")
	})

	t.Run("kubeconfig tests", func(t *testing.T) {
		// Get the kubeconfig from the K3s container
		kubeConfigYaml, err := k3sContainer.GetKubeConfig(t.Context())
		require.NoError(t, err)
		t.Logf("📝 Retrieved kubeconfig from K3s container")

		// Create a REST config from the kubeconfig
		restcfg, err := clientcmd.RESTConfigFromKubeConfig(kubeConfigYaml)
		require.NoError(t, err)

		// Create a Kubernetes clientset
		k8s, err := kubernetes.NewForConfig(restcfg)
		require.NoError(t, err)
		t.Logf("✅ Connected to Kubernetes API server")

		t.Run("all pods are running", func(t *testing.T) {
			pods, err := k8s.CoreV1().Pods("default").List(t.Context(), metav1.ListOptions{})
			require.NoError(t, err)
			require.NotEmpty(t, pods.Items, "Expected at least one pod to be running")

			runningPods := 0
			for _, pod := range pods.Items {
				if pod.Status.Phase == corev1.PodRunning {
					runningPods++
					t.Logf("  ✓ Pod %s is running", pod.Name)
				}
			}
			require.Positive(t, runningPods, "Expected at least one pod in Running state")
			t.Logf("🎯 Found %d running pods", runningPods)
		})

		t.Run("service exists and is of type LoadBalancer", func(t *testing.T) {
			service, err := k8s.CoreV1().Services("default").Get(t.Context(), "quickpizza-public-api", metav1.GetOptions{})
			require.NoError(t, err)
			require.Equal(t, corev1.ServiceTypeLoadBalancer, service.Spec.Type)
			require.Equal(t, int32(3333), service.Spec.Ports[0].Port)
			t.Logf("✅ Service quickpizza-public-api exists (type: %s, port: %d)", service.Spec.Type, service.Spec.Ports[0].Port)
		})

		t.Run("deployments exist and are ready", func(t *testing.T) {
			deployments, err := k8s.AppsV1().Deployments("default").List(t.Context(), metav1.ListOptions{
				LabelSelector: "app.k8s.io/name=quickpizza",
			})
			require.NoError(t, err)
			require.NotEmpty(t, deployments.Items, "Expected at least one deployment")

			for _, deployment := range deployments.Items {
				require.Equal(t, deployment.Status.ReadyReplicas, deployment.Status.Replicas,
					"Deployment %s: expected %d ready replicas, got %d",
					deployment.Name, deployment.Status.Replicas, deployment.Status.ReadyReplicas)
				t.Logf("  ✓ Deployment %s is ready (%d/%d replicas)",
					deployment.Name, deployment.Status.ReadyReplicas, deployment.Status.Replicas)
			}
			t.Logf("🎯 All %d deployments are ready", len(deployments.Items))
		})

		t.Run("statefulset exists and is ready", func(t *testing.T) {
			statefulSets, err := k8s.AppsV1().StatefulSets("default").List(t.Context(), metav1.ListOptions{
				LabelSelector: "app.kubernetes.io/component=database",
			})
			require.NoError(t, err)
			require.NotEmpty(t, statefulSets.Items, "Expected database StatefulSet to exist")

			for _, sts := range statefulSets.Items {
				require.Equal(t, sts.Status.ReadyReplicas, *sts.Spec.Replicas,
					"StatefulSet %s: expected %d ready replicas, got %d",
					sts.Name, *sts.Spec.Replicas, sts.Status.ReadyReplicas)
				t.Logf("  ✓ StatefulSet %s is ready (%d/%d replicas)",
					sts.Name, sts.Status.ReadyReplicas, *sts.Spec.Replicas)
			}
			t.Logf("🎯 All %d statefulsets are ready", len(statefulSets.Items))
		})

		t.Run("configmap exists and has data", func(t *testing.T) {
			configMap, err := k8s.CoreV1().ConfigMaps("default").Get(t.Context(), "quickpizza-env-common", metav1.GetOptions{})
			require.NoError(t, err)
			require.NotEmpty(t, configMap.Data, "Expected ConfigMap to have data")
			t.Logf("✅ ConfigMap quickpizza-env-common exists with %d keys", len(configMap.Data))
		})

		t.Run("secret exists and has data", func(t *testing.T) {
			secret, err := k8s.CoreV1().Secrets("default").Get(t.Context(), "quickpizza-db-credentials", metav1.GetOptions{})
			require.NoError(t, err)
			require.NotEmpty(t, secret.Data, "Expected Secret to have data")
			t.Logf("✅ Secret quickpizza-db-credentials exists")
		})

		t.Logf("✅ All Kubernetes resources verified successfully!")
	})
}
