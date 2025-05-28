---
title: Todo App + Auth + GORM + Testcontainers
keywords: [todo app, gorm, authentication, testcontainers, postgres]
description: A Todo application with authentication using GORM and Postgres.
---

# Todo App with Auth using GORM and Testcontainers

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/mdelapenya/testcontainers-go-examples/tree/main/gofiber-services) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github.com/mdelapenya/testcontainers-go-examples/tree/main/gofiber-services)

This project demonstrates a Todo application with authentication using GORM and Testcontainers.

The database is a Postgres instance created using the GoFiber's [Testcontainers Service module](https://github.com/gofiber/contrib/testcontainers).

## Prerequisites

Ensure you have the following installed and available in your `GOPATH`:

- Golang
- [Air](https://github.com/air-verse/air) for hot reloading

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/mdelapenya/testcontainers-go-examples.git
    cd testcontainers-go-examples/gofiber-services
    ```

2. Install dependencies:
    ```sh
    go get
    ```

## Running the Application

1. Start the application:
    ```sh
    air
    ```

## Environment Variables

Create a `.env` file in the root directory and add the following variables:

```shell
# PORT returns the server listening port
# default: 8000
PORT=

# DB returns the name of the sqlite database
# default: postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

# TOKENKEY returns the jwt token secret
TOKENKEY=

# TOKENEXP returns the jwt token expiration duration.
# Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
# default: 10h
TOKENEXP=
```
