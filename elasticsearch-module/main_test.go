package main

import (
	"context"
	"fmt"
	"log"

	"github.com/testcontainers/testcontainers-go/modules/elasticsearch"
)

func ExampleElasticsearch() {
	ctx := context.Background()

	ctr, err := elasticsearch.Run(ctx, "docker.elastic.co/elasticsearch/elasticsearch:8.14.2")
	if err != nil {
		log.Fatalf("Could not start container: %s", err)
	}
	defer ctr.Terminate(ctx)

	fmt.Println(ctr.IsRunning())

	// Output: true
}
