package main

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

func ExampleElasticsearchContainer() {
	ctx := context.Background()

	ctr, err := elasticsearch.Run(ctx, "docker.elastic.co/elasticsearch/elasticsearch:8.14.2")
	defer func() {
		if err := testcontainers.TerminateContainer(ctr); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()
	if err != nil {
		log.Printf("failed to start container: %s", err)
		return
	}

	fmt.Println(ctr.IsRunning())

	// Output: true
}
