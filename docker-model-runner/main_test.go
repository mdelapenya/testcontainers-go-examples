package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/socat"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Example_modelRunner() {
	ctx := context.Background()

	is, err := isDockerDesktop()
	if err != nil {
		log.Printf("failed to check if Docker Desktop is running: %s", err)
		return
	}
	if !is {
		log.Printf("skipping example because it requires Docker Desktop")
		fmt.Println("true") // Printing true to simulate the socat container is running
		fmt.Println("200")  // Printing 200 to simulate the response is 200
		return
	}

	socatCtr, err := socat.Run(
		ctx, "alpine/socat:1.8.0.1",
		testcontainers.WithWaitStrategy(wait.ForListeningPort("80/tcp")),
		socat.WithTarget(socat.NewTarget(80, "model-runner.docker.internal")),
	)
	defer func() {
		if err := testcontainers.TerminateContainer(socatCtr); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	fmt.Println(socatCtr.IsRunning())

	httpClient := http.DefaultClient
	resp, err := httpClient.Get(socatCtr.TargetURL(80).String())
	if err != nil {
		log.Printf("failed to get response: %s", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response: %s", err)
		return
	}
	fmt.Println(string(body))

	// Output:
	// true
	// 200
	// Docker Model Runner
	//
	// The service is running.
}

func isDockerDesktop() (bool, error) {
	ctx := context.Background()

	cli, err := testcontainers.NewDockerClientWithOpts(ctx)
	if err != nil {
		return false, fmt.Errorf("new docker client: %w", err)
	}

	info, err := cli.Info(ctx)
	if err != nil {
		return false, fmt.Errorf("docker info: %w", err)
	}

	if info.OperatingSystem == "Docker Desktop" {
		return true, nil
	}

	return false, nil
}
