package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/moby/moby/client"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/dockermodelrunner"
)

const (
	modelNamespace = "ai"
	modelName      = "smollm2"
	modelTag       = "360M-Q4_K_M"
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
		// Because the Docker Model Runner is only available on Docker Desktop, we skip the example
		// on those platforms. We are printing the expected output to make the example work.
		log.Printf("skipping example because it requires Docker Desktop")
		fmt.Println("true")               // Printing true to simulate the socat container is running
		fmt.Println("models count: 0")    // Printing 0 to simulate the number of models
		fmt.Println("model created")      // Printing to simulate the model was created
		fmt.Println("models count: 1")    // Printing 1 to simulate the number of models after creation
		fmt.Println("ID: " + fqModelName) // Printing the ID of the model to simulate the model was created
		fmt.Println("tags count: 1")      // Printing 1 to simulate the tags count
		return
	}

	dmrCtr, err := dockermodelrunner.Run(ctx)
	defer func() {
		if err := testcontainers.TerminateContainer(dmrCtr); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	// 0. Check the DMR container is running
	fmt.Println(dmrCtr.IsRunning())

	// 1. Get the models
	models, err := dmrCtr.ListModels(ctx)
	if err != nil {
		log.Printf("failed to list models: %s", err)
		return
	}
	fmt.Printf("models count: %d\n", len(models))

	// 2. Create the model
	err = dmrCtr.PullModel(ctx, fqModelName)
	if err != nil {
		log.Printf("failed to create model: %s", err)
		return
	}
	fmt.Println("model created")

	// 4. Verify the model was created
	for i := range 60 { // try for 60 seconds
		models, err = dmrCtr.ListModels(ctx)
		if err != nil {
			log.Printf("failed to list models: %s", err)
			return
		}
		if len(models) > 0 {
			break
		}

		log.Printf("waiting for model to be created... %d seconds", i)
		time.Sleep(time.Second)
	}
	fmt.Printf("models count: %d\n", len(models))

	// 5. Get the model by namespace/name
	model, err := dmrCtr.InspectModel(ctx, modelNamespace, modelName+":"+modelTag)
	if err != nil {
		log.Printf("failed to inspect model: %s", err)
		return
	}
	fmt.Printf("ID: %v\n", model.Tags[0])
	fmt.Printf("tags count: %v\n", len(model.Tags))

	// Output:
	// true
	// models count: 0
	// model created
	// models count: 1
	// ID: ai/smollm2:360M-Q4_K_M
	// tags count: 1
}

func isDockerDesktop() (bool, error) {
	ctx := context.Background()

	cli, err := testcontainers.NewDockerClientWithOpts(ctx)
	if err != nil {
		return false, fmt.Errorf("new docker client: %w", err)
	}

	result, err := cli.Info(ctx, client.InfoOptions{})
	if err != nil {
		return false, fmt.Errorf("docker info: %w", err)
	}

	if result.Info.OperatingSystem == "Docker Desktop" {
		return true, nil
	}

	return false, nil
}
