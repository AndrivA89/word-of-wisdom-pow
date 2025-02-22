.PHONY: build-server build-client run-server run-client test lint

build-server:
	@echo "Building server..."
	docker build -f docker/server/Dockerfile -t word-of-wisdom-server .

build-client:
	@echo "Building client..."
	docker build -f docker/client/Dockerfile -t word-of-wisdom-client .

run-server:
	@echo "Running server..."
	docker compose up server

run-client:
	@echo "Running client..."
	docker compose up client

test:
	@echo "Running tests..."
	go test ./...

lint:
	@echo "Running golangci-lint..."
	golangci-lint run