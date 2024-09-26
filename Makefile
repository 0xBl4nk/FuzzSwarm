# Makefile for FuzzSwarm

# Variáveis
BINARY_NAME = FuzzSwarm
SRC_DIR = ./src
MAIN_FILE = main.go

# Default target: build the project
all: build

# Build the Go binary
build:
	go build -o $(BINARY_NAME) $(MAIN_FILE)

# Run the project
run: build
	./$(BINARY_NAME)

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Lint the Go code using golint (you need to have golint installed)
lint:
	golint $(SRC_DIR)/*.go

# Format the Go code
fmt:
	go fmt $(SRC_DIR)/*.go

# Test the project (assuming there are test files)
test:
	go test ./...

# Run the fuzzer (assumption based on the presence of fuzzer.go)
fuzz:
	go run $(SRC_DIR)/fuzzer.go

# Help command
help:
	@echo "Usage:"
	@echo "  make         - Build the project"
	@echo "  make run     - Build and run the project"
	@echo "  make clean   - Clean build artifacts"
	@echo "  make lint    - Lint the source code"
	@echo "  make fmt     - Format the Go source code"
	@echo "  make test    - Run tests"
	@echo "  make fuzz    - Run the fuzzer"
