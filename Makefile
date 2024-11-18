# Define the Go binary name for server and cli apps
SERVER_BINARY_NAME=quiz-app-server
CLI_BINARY_NAME=quiz-app-cli

# Default OS and ARCH (used for local development)
GOOS ?= linux
GOARCH ?= amd64

# Build the server app for the target OS/ARCH
build-server:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/$(SERVER_BINARY_NAME)-$(GOOS)-$(GOARCH) ./server
	chmod +x ./bin/$(SERVER_BINARY_NAME)-$(GOOS)-$(GOARCH)

# Build the CLI app for the target OS/ARCH
build-cli:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./bin/$(CLI_BINARY_NAME)-$(GOOS)-$(GOARCH) ./cli
	chmod +x ./bin/$(CLI_BINARY_NAME)-$(GOOS)-$(GOARCH)

# Build both server and cli apps for the target OS/ARCH
build: build-server build-cli

# Build for Linux (amd64) for both server and cli
build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(SERVER_BINARY_NAME)-linux-amd64 ./server
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(CLI_BINARY_NAME)-linux-amd64 ./cli

# Build for macOS (amd64) for both server and cli
build-macos:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(SERVER_BINARY_NAME)-darwin-amd64 ./server
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(CLI_BINARY_NAME)-darwin-amd64 ./cli

# Build for Windows (amd64) for both server and cli
build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(SERVER_BINARY_NAME)-windows-amd64.exe ./server
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(CLI_BINARY_NAME)-windows-amd64.exe ./cli

# Clean the build directory
clean:
	rm -rf ./bin/*

run-server:
	$(eval LAST_BINARY := $(shell ls -t ./bin/$(SERVER_BINARY_NAME)-* | head -n 1))
	@echo "Running server: $(LAST_BINARY)"
	$(LAST_BINARY)

# Run the cli application
run-cli:
	$(eval LAST_BINARY := $(shell ls -t ./bin/$(CLI_BINARY_NAME)-* | head -n 1))
	@echo "Running CLI: $(LAST_BINARY)"
	$(LAST_BINARY) $(ARGS)

# Help target to show available commands
help:
	@echo "Available targets:"
	@echo "  make build-linux      - Build server and cli for Linux"
	@echo "  make build-macos      - Build server and cli for macOS"
	@echo "  make build-windows    - Build server and cli for Windows"
	@echo "  make build            - Build both server and cli for the current OS/ARCH"
	@echo "  make clean            - Clean build directory"
	@echo "  make run-server       - Run the server application locally"
	@echo "  make run-cli          - Run the cli application locally"
	@echo "  make help             - Show this help message"
