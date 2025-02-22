# Makefile

.PHONY: build-server build-client run-server run-client lint test

build-server:
	docker build -f Dockerfile.server -t word-of-wisdom-server .

build-client:
	docker build -f Dockerfile.client -t word-of-wisdom-client .

run-server:
	docker compose up server

run-client:
	docker compose up client

lint:
	golangci-lint run

test:
	go test ./...
