# Word of Wisdom TCP Server with PoW

## Task description
Design and implement “Word of Wisdom” tcp server.

- TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
- The choice of the POW algorithm should be explained.
- After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
- Docker file should be provided both for the server and for the client that solves the POW challenge.

## Features

- **Proof-of-Work (PoW)** challenge to mitigate DDoS attacks.
- Domain model for citations (quotes) with author parsing.
- Repository to load citations from a text file.
- Service layer for business logic.
- Dockerized server and client with Docker Compose.
- Comprehensive README, Makefile, and tests.
- Proper error handling and logging.
- Timeouts for PoW solving.
- golangci-lint configuration for code quality.

## Project Structure

```
word-of-wisdom-pow/
├── cmd/
│   ├── server/main.go             # Server entry point
│   └── client/main.go             # Client entry point
├── internal/
│   ├── application/
│   │   ├── citation_service.go    # Service layer for citations
│   │   └── deps.go                # Application's dependencies 
│   ├── domain/
│   │   └── citation.go            # Domain model forcCitation
│   ├── infrastructure/
│   │   ├── repository/
│   │   │   ├── repo.go            # Repository for loading citations
│   │   │   ├── repo_test.go       # Tests for repository
│   │   │   └── citations.txt      # Text file with citations 
│   │   └── tcp/
│   │       ├── server.go          # TCP server implementation
│   │       └── deps.go            # Server's dependencies
│   └── pow/
│       ├── proof_ow_work.go       # POW implementation
│       └── proof_ow_work_test.go  # Tests for POW
├── docker/
│   ├── server/Dockerfile          # Dockerfile for server
│   └── client/Dockerfile          # Dockerfile for client
├── docker-compose.yml             # Docker Compose configuration
├── Makefile                       # Build, test, deploy automation
├── .golangci.yml                  # golangci-lint configuration
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

- Docker Desktop (with Docker Compose)
- Go 1.21+ (for local development)
- [golangci-lint](https://golangci-lint.run/) (optional, for linting)

### Running with Docker Compose

1. Build and start the containers:
   ```bash
   docker compose up --build
2. The server will listen on port 9000. The client will connect, solve the PoW challenge, and receive a citation.

Running Locally
1. Ensure that the CITATIONS_FILE_PATH environment variable is set if needed:
   ```bash
   export CITATIONS_FILE_PATH="./internal/infrastructure/repository/citations.txt"
   ```
2. Run the server:
   ```bash
   go run cmd/server/main.go
   ```
3. In a separate terminal, run the client:
   ```bash
   go run cmd/client/main.go
   ```

Testing

Run tests using:
   ```bash
   make test
   ```

Linting

Run golangci-lint:
   ```bash
   make lint
   ```

Deployment

The project is Dockerized and can be deployed using Docker Compose. See the docker-compose.yml file for details
