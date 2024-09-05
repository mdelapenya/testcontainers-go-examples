package main

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

func ExampleElasticsearchContainer() {
	ctx := context.Background()

	ctr, err := elasticsearch.Run(ctx, "docker.elastic.co/elasticsearch/elasticsearch:8.14.2")
	if err != nil {
		log.Fatalf("Could not start container: %s", err)
	}
	defer func() {
		if err := ctr.Terminate(context.Background()); err != nil {
			log.Fatalf("failed to terminate container: %s", err)
		}
	}()

	fmt.Println(ctr.IsRunning())

	// Output: true
}
