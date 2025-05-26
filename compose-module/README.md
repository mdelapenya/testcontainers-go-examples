# Example repository for Testcontainers for Go: Docker compose module

This repository contains an example of how to use [Testcontainers for Go](https://golang.testcontainers.org/) with Docker Compose.

Please refer to the documentation of the Compose module [here](https://golang.testcontainers.org/features/docker_compose/).

## CI

This project contains a CI configuration for running the tests with Testcontainers for Go in different CI providers.

### GitHub Actions

This repository contains a GitHub Actions configuration for running the tests with Testcontainers for Go.

The configuration is in the [`.github/workflows/ci.yml` file](./.github/workflows/ci.yml), and it uses a matrix to run the tests with different versions of Go and Testcontainers for Go, including the main branch for both Testcontainers for Go and the Compose module.
