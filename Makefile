BINARY_NAME=conf2env

.PHONY: all build clean test

all: build

build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) ./cmd/conf2env

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

test:
	@echo "Testing..."
	@go test -v ./...

install:
	@echo "Installing..."
	@go install ./cmd/conf2env

.PHONY: run
run: build
	@./$(BINARY_NAME)

.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	@go mod tidy 