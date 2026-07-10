package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/docker/compose/v2/pkg/api"
	tccompose "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Example_compose() {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if err := compose.Down(context.Background(),
			tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
			log.Println(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = compose.Up(ctx, tccompose.Wait(true))
	if err != nil {
		log.Println(err)
		return
	}

	serviceNames := compose.Services()
	fmt.Println(serviceNames)

	// Output: [mysql nginx]
}

func Example_compose_waitForInvalidService() {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if err := compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
			log.Println(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = compose.
		WaitForService("non-existent-srv-1", wait.NewLogStrategy("started").WithStartupTimeout(10*time.Second).WithOccurrence(1)).
		Up(ctx, tccompose.Wait(true))

	if err == nil {
		log.Printf("Expected error to be thrown because service with wait strategy is not running: %s", err)
		return
	}

	fmt.Println(err.Error())

	// Output:
	// wait for services: no container found for service name non-existent-srv-1
}

func Example_compose_waitForLogStrategy() {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	if err != nil {
		log.Println(err)
		return
	}

	defer func() {
		if err := compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
			log.Println(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = compose.
		WaitForService("mysql", wait.NewLogStrategy("started").WithStartupTimeout(10*time.Second).WithOccurrence(1)).
		Up(ctx, tccompose.Wait(true))

	fmt.Println(err)

	// Output:
	// <nil>
}

func Example_composeWithStackFiles() {
	ctx := context.Background()

	dockerCompose, err := tccompose.NewDockerComposeWith(
		tccompose.WithStackFiles(filepath.Join("testdata", "docker-compose-stack.yml")),
		tccompose.StackIdentifier("nginx-compose"),
	)
	if err != nil {
		log.Println(err)
		return
	}

	err = dockerCompose.
		WaitForService("nginx", wait.NewHTTPStrategy("/").
			WithPort("80/tcp").
			WithStartupTimeout(20*time.Second),
		).
		Up(ctx, tccompose.Wait(true), tccompose.WithRecreate(api.RecreateNever))
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(err)

	// Output:
	// <nil>
}
