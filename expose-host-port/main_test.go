package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	"github.com/testcontainers/testcontainers-go"
	tcexec "github.com/testcontainers/testcontainers-go/exec"
)

const expectedOutput = "Hello, Gopher!"

func Example_exposeHostPortContainer() {
	// create an http server and start it in a goroutine
	// This server will listen on a random port and will
	// respond with the expected output: "Hello, Gopher!"
	server, port, err := createHttpServer()
	if err != nil {
		log.Fatalf("failed to create http server: %v", err)
	}
	defer func() {
		_ = server.Close()
	}()

	go func() {
		// for demo purposes, we ignore the error returned by ListenAndServe
		_ = server.ListenAndServe()
	}()

	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:           "alpine:3.17",
			HostAccessPorts: []int{port},
			Cmd:             []string{"top"},
		},
		Started: true,
	}

	ctx := context.Background()
	ctr, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		log.Printf("failed to create container: %v\n", err)
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

	code, reader, err := ctr.Exec(
		context.Background(),
		[]string{"wget", "-q", "-O", "-", fmt.Sprintf("http://%s:%d", testcontainers.HostInternal, port)},
		tcexec.Multiplexed(),
	)
	if err != nil {
		log.Printf("failed to execute command: %v\n", err)
		return
	}

	fmt.Println("Exit code:", code)

	// read the response
	bs, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("failed to read response: %v\n", err)
		return
	}

	fmt.Println(string(bs))

	// Output:
	// Exit code: 0
	// Hello, Gopher!
}

func createHttpServer() (*http.Server, int, error) {
	port, err := getFreePort()
	if err != nil {
		return nil, 0, err
	}

	server := &http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}

	server.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedOutput)
	})

	return server, port, nil
}

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
