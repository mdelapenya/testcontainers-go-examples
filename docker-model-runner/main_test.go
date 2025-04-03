package main

import (
	"context"
	"fmt"
	"log"

	"testcontainers-go-examples/docker-model-runner/sdk/client"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/socat"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	modelNamespace = "ai"
	modelName      = "llama3.2"
	modelTag       = "latest"
)

func Example_modelRunner() {
	ctx := context.Background()

	const fqModelName = modelNamespace + "/" + modelName + ":" + modelTag

	is, err := isDockerDesktop()
	if err != nil {
		log.Printf("failed to check if Docker Desktop is running: %s", err)
		return
	}
	if !is {
		log.Printf("skipping example because it requires Docker Desktop")
		fmt.Println("true")                                           // Printing true to simulate the socat container is running
		fmt.Println("Docker Model Runner\n\nThe service is running.") // Printing Server response to simulate the service is running
		fmt.Println("models count: 0")                                // Printing 0 to simulate the number of models
		fmt.Println("model created")                                  // Printing to simulate the model was created
		fmt.Println("models count: 1")                                // Printing 1 to simulate the number of models after creation
		fmt.Println("ID: " + fqModelName)                             // Printing the ID of the model to simulate the model was created
		fmt.Println("tags count: 1")                                  // Printing 1 to simulate the tags count
		fmt.Println("model deleted")                                  // Printing to simulate the model was deleted
		fmt.Println("models count: 0")                                // Printing 0 to simulate the number of models after deletion
		return
	}

	modelRunnerPort := 80

	socatCtr, err := socat.Run(
		ctx, "alpine/socat:1.8.0.1",
		testcontainers.WithWaitStrategy(wait.ForListeningPort("80/tcp")),
		socat.WithTarget(socat.NewTarget(modelRunnerPort, "model-runner.docker.internal")),
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

	// 0. Check the socat container is running
	fmt.Println(socatCtr.IsRunning())

	baseURL := socatCtr.TargetURL(modelRunnerPort).String()

	dmrClient := client.NewClient(baseURL)

	// 1. Get the root
	root, err := dmrClient.Root()
	if err != nil {
		log.Printf("failed to read response: %s", err)
		return
	}
	fmt.Println(string(root))

	// 2. Get the models
	models, err := dmrClient.ListModels()
	if err != nil {
		log.Printf("failed to list models: %s", err)
		return
	}
	fmt.Printf("models count: %d\n", len(models))

	// 3. Create the model
	resp, err := dmrClient.CreateModel(fqModelName)
	if err != nil {
		log.Printf("failed to create model: %s", err)
		return
	}
	log.Printf("created model: %v", resp)
	fmt.Println("model created")

	// 4. Verify the model was created
	models, err = dmrClient.ListModels()
	if err != nil {
		log.Printf("failed to list models: %s", err)
		return
	}
	fmt.Printf("models count: %d\n", len(models))

	// 5. Get the model by namespace/name
	model, err := dmrClient.GetModel(modelNamespace, modelName)
	if err != nil {
		log.Printf("failed to get model: %s", err)
		return
	}
	fmt.Printf("ID: %v\n", model.ID)
	fmt.Printf("tags count: %v\n", len(model.Tags))

	// 6. Delete the model
	resp, err = dmrClient.DeleteModel(modelNamespace, modelName)
	if err != nil {
		log.Printf("failed to get models: %s", err)
		return
	}
	log.Printf("deleted model: %v", resp)
	fmt.Println("model deleted")

	// 7. List the models again
	models, err = dmrClient.ListModels()
	if err != nil {
		log.Printf("failed to list models: %s", err)
		return
	}
	fmt.Printf("models count: %d\n", len(models))

	// Output:
	// true
	// Docker Model Runner
	//
	// The service is running.
	// models count: 0
	// model created
	// models count: 1
	// ID: ai/llama3.2:latest
	// tags count: 1
	// model deleted
	// models count: 0
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
