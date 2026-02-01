.PHONY: build run test lint clean

BINARY_NAME ?= api
CMD_PATH    := ./cmd/api

build:
	go build -o bin/$(BINARY_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)/main.go

test:
	go test ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

lint:
	golangci-lint run ./...

clean:
	rm -rf bin/ coverage.out coverage.html

docker-build:
	docker build -t $(BINARY_NAME):latest -f build/Dockerfile .

docker-run:
	docker run --rm -p 8080:8080 $(BINARY_NAME):latest
