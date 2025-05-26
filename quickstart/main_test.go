package main

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func ExampleRun() {
	ctx := context.Background()
	moduleOpts := []testcontainers.ContainerCustomizer{
		testcontainers.WithExposedPorts("6379/tcp"),
		testcontainers.WithWaitStrategy(wait.ForLog("Ready to accept connections")),
	}

	redisC, err := testcontainers.Run(ctx, "redis:latest", moduleOpts...)
	defer func() {
		if err := testcontainers.TerminateContainer(redisC); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	fmt.Println(redisC.IsRunning())

	// Output:
	// true
}
