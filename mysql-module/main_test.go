package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/wait"
)

func ExampleMySQLContainer() {
	ctx := context.Background()
	ctr, err := mysql.Run(ctx, "mysql:8.0",
		testcontainers.WithEnv(map[string]string{
			"MYSQL_ROOT_PASSWORD": "password",
			"MYSQL_DATABASE":      "database",
			"MYSQL_USER":          "username",
			"MYSQL_PASSWORD":      "password",
		}),
		testcontainers.WithWaitStrategy(
			wait.ForLog("port: 3306  MySQL Community Server - GPL").WithStartupTimeout(30*time.Second),
			wait.ForListeningPort("3306/tcp"),
		),
	)
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
