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
		log.Printf("failed to start container: %v\n", err)
		return
	}
	defer func() {
		if ctr == nil {
			return
		}
		if err := ctr.Terminate(context.Background()); err != nil {
			log.Fatalf("failed to terminate container: %v", err)
		}
	}()

	fmt.Println(ctr.IsRunning())

	// Output: true
}
