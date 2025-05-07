.PHONY: build run test clean migrate-up migrate-down lint help db-create db-drop db-psql db-reset

# Variables
BINARY_NAME=daily-planner
DB_NAME=daily_planner
DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5434
MIGRATION_DIR=internal/repository/migrations

# Default target
all: build

# Build the application
build:
	@echo "Building application..."
	go build -o $(BINARY_NAME) cmd/api/main.go

# Run the application
run:
	@echo "Running application..."
	go run cmd/api/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	go clean

# Run database migrations up
migrate-up:
	@echo "Running database migrations up..."
	@if [ -z "$(DB_NAME)" ]; then \
		echo "Error: DB_NAME environment variable is not set"; \
		exit 1; \
	fi
	@if [ -d "$(MIGRATION_DIR)" ]; then \
		go run cmd/api/main.go migrate up; \
	else \
		echo "Error: Migration directory not found"; \
		exit 1; \
	fi

# Run database migrations down
migrate-down:
	@echo "Running database migrations down..."
	@if [ -z "$(DB_NAME)" ]; then \
		echo "Error: DB_NAME environment variable is not set"; \
		exit 1; \
	fi
	@if [ -d "$(MIGRATION_DIR)" ]; then \
		go run cmd/api/main.go migrate down; \
	else \
		echo "Error: Migration directory not found"; \
		exit 1; \
	fi

# Create database
db-create:
	@echo "Creating database $(DB_NAME)..."
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -c "CREATE DATABASE $(DB_NAME);"

# Drop database
db-drop:
	@echo "Dropping database $(DB_NAME)..."
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -c "DROP DATABASE IF EXISTS $(DB_NAME);"

# Connect to database using psql
db-psql:
	@echo "Connecting to database $(DB_NAME)..."
	PGPASSWORD=$(DB_PASSWORD) psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME)

# Reset database (drop and create)
db-reset: db-drop db-create migrate-up

# Run linter
lint:
	@echo "Running linter..."
	golangci-lint run

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  run          - Run the application"
	@echo "  test         - Run tests"
	@echo "  clean        - Clean build artifacts"
	@echo "  migrate-up   - Run database migrations up"
	@echo "  migrate-down - Run database migrations down"
	@echo "  db-create    - Create the database"
	@echo "  db-drop      - Drop the database"
	@echo "  db-psql      - Connect to database using psql"
	@echo "  db-reset     - Reset database (drop, create, and migrate)"
	@echo "  lint         - Run linter"
	@echo "  deps         - Install dependencies"
	@echo "  help         - Show this help message" 