.PHONY: build run test clean docker-build docker-up docker-down install

# Build the application
build:
	@echo "Building FMR Backend (FIPS-enabled)..."
	@CGO_ENABLED=0 go build -o bin/server ./cmd/server

# Run the application locally
run:
	@echo "Running FMR Backend..."
	@go run ./cmd/server/main.go

# Run tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/

# Install dependencies
install:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker build -t fmr-backend-fips:latest .

# Start services with docker-compose
docker-up:
	@echo "Starting services..."
	@docker-compose up -d

# Stop services
docker-down:
	@echo "Stopping services..."
	@docker-compose down

# View logs
logs:
	@docker-compose logs -f backend

# Check FIPS mode
check-fips:
	@echo "Checking Go version and FIPS configuration..."
	@go version
	@echo "GODEBUG: $$GODEBUG"

# Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@golangci-lint run

help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application locally"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install dependencies"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-up    - Start services with docker-compose"
	@echo "  docker-down  - Stop services"
	@echo "  logs         - View backend logs"
	@echo "  check-fips   - Check FIPS configuration"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code"
