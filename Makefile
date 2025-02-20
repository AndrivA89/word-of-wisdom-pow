# Makefile для сборки, тестирования и запуска приложения

.PHONY: build-server build-client run-server run-client test lint docker-build-server docker-build-client

build-server:
	@echo "Сборка сервера..."
	cd cmd/server && go build -o server .

build-client:
	@echo "Сборка клиента..."
	cd cmd/client && go build -o client .

run-server: build-server
	@echo "Запуск сервера..."
	cd cmd/server && ./server

run-client: build-client
	@echo "Запуск клиента..."
	cd cmd/client && ./client

test:
	@echo "Запуск тестов..."
	cd internal && go test -v ./...

lint:
	@echo "Запуск golangci-lint..."
	golangci-lint run

docker-build-server:
	@echo "Сборка Docker образа для сервера..."
	docker build -t wordofwisdom-server -f docker/Dockerfile.server .

docker-build-client:
	@echo "Сборка Docker образа для клиента..."
	docker build -t wordofwisdom-client -f docker/Dockerfile.client .
