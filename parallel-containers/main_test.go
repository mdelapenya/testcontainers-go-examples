package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/testcontainers/testcontainers-go"
)

func ExampleParallelContainers() {
	pReq := testcontainers.ParallelContainerRequest{
		{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "hello-world:latest",
				Env:   map[string]string{},
			},
			Started: true,
		},
		{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "hello-world:latest",
				Env:   map[string]string{},
			},
			Started: true,
		},
	}

	cs, err := testcontainers.ParallelContainers(context.Background(), pReq, testcontainers.ParallelContainersOptions{WorkersCount: runtime.NumCPU()})
	if err != nil {
		log.Fatalf("Could not start containers: %s", err)
	}

	for _, c := range cs {
		fmt.Println(c.IsRunning())
		defer func() {
			if err := c.Terminate(context.Background()); err != nil {
				log.Fatalf("failed to terminate container: %s", err)
			}
		}()
	}

	// Output:
	// true
	// true
}
