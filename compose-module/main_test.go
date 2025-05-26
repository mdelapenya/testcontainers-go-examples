package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	tccompose "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
)

func Example_compose() {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := compose.Down(context.Background(),
			tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
			log.Fatal(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = compose.Up(ctx, tccompose.Wait(true))
	if err != nil {
		log.Fatal(err)
	}

	serviceNames := compose.Services()
	fmt.Println(serviceNames)

	// Output: [mysql nginx]
}

func Example_compose_waitForInvalidService() {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
			log.Fatal(err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = compose.
		WaitForService("non-existent-srv-1", wait.NewLogStrategy("started").WithStartupTimeout(10*time.Second).WithOccurrence(1)).
		Up(ctx, tccompose.Wait(true))

	if err == nil {
		log.Fatalf("Expected error to be thrown because service with wait strategy is not running: %s", err)
	}

	fmt.Println(err.Error())

	// Output:
	// wait for services: no container found for service name non-existent-srv-1
}

func Example_compose_waitForLogStrategy() {
	compose, err := tccompose.NewDockerCompose(filepath.Join("testdata", "docker-compose.yml"))
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := compose.Down(context.Background(), tccompose.RemoveOrphans(true), tccompose.RemoveImagesLocal); err != nil {
			log.Fatal(err)
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
