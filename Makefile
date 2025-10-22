# Application name
APP_NAME=google_map2whatsapp
BINARY_NAME=bin/$(APP_NAME)
DOCKER_IMAGE=$(APP_NAME):latest
DOCKER_IMAGE_TEST=$(APP_NAME)_test:latest

# Directories
CACHE_DIR=cache

# Default target
all: setup build

# Setup environment
setup:
	@echo "Setting up environment..."
	@mkdir -p $(CACHE_DIR)
	@if [ ! -f .env ]; then \
		echo "Copying .env.example to .env"; \
		cp .env.example .env; \
		echo "Please update the .env file with your configuration"; \
	else \
		echo ".env file already exists, skipping..."; \
	fi

# Build the application
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(BINARY_NAME) .

# Run the application
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BINARY_NAME)

test:
	@echo "Running tests..."
	@go test -v ./...

lint:
	@echo "Running linter..."
	@golangci-lint run --config=./.golangci.yaml ./...

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

# Run application in Docker
docker-run: docker-build
	@echo "Running $(APP_NAME) in Docker..."
	docker run -it --rm \
		$(DOCKER_IMAGE) \
		./google_map2whatsapp

docker-build-test:
	@echo "Building Docker image..."
	docker build  --target dev -t $(DOCKER_IMAGE_TEST) .

# Run tests in Docker
docker-test: docker-build-test
	@echo "Running tests in Docker..."
	docker run --rm \
		-v $(PWD)/$(CACHE_DIR):/app/$(CACHE_DIR) \
		--env-file .env \
		$(DOCKER_IMAGE_TEST) \
		go test -v ./test/...

docker-lint: docker-build-test
	@echo "Running tests in Docker..."
	docker run --rm \
		-v $(PWD)/$(CACHE_DIR):/app/$(CACHE_DIR) \
		--env-file .env \
		$(DOCKER_IMAGE_TEST) \
		golangci-lint run --config=./.golangci.yaml ./...

.PHONY: all setup build run docker-build docker-run clean test