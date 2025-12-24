# Binary name
BINARY_NAME=veil-cli

# Build the project
all: build

build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) main.go
	@echo "Built successfully"

# Run the project
run:
	@go run main.go

# Clean up build files
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "Cleaned up"

# Tidy up Go modules
tidy:
	@echo "Tidying up Go modules..."
	@go mod tidy

.PHONY: all build run clean tidy

