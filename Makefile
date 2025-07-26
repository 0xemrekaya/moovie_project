# Simplified Makefile for Moovie API

BINARY_NAME=moovie-api
MAIN_PATH=./cmd/api

.PHONY: all build run clean test docker-up docker-down

# Default command
all: build

# Build the application
build:
	@echo "Building application..."
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# Run the application locally
run:
	@echo "Running application..."
	go run $(MAIN_PATH)

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)
	go clean

# Start services with Docker Compose
docker-up:
	@echo "Starting Docker containers..."
	docker-compose up -d

# Stop services with Docker Compose
docker-down:
	@echo "Stopping Docker containers..."
	docker-compose down
