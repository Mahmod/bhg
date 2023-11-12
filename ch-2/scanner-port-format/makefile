# Phony targets are always executed
.PHONY: all build clean clean-cross-platform help

# Name of the binary - current directory name
BINARY_NAME := $(notdir $(shell pwd))

# Default target
all: build

# Build the project
build:
	@echo "Building the project..."
	@go build -o $(BINARY_NAME) portformat.go || (echo "Build failed"; exit 1)

# Clean up the build
clean:
	@echo "Cleaning up..."
	@rm -f $(BINARY_NAME)

# Cross-platform clean command
clean-cross-platform:
	@echo "Performing cross-platform clean..."
ifeq ($(OS),Windows_NT)
	@del /f /q $(BINARY_NAME) 2>nul || (echo "Clean failed"; exit 1)
else
	@rm -f $(BINARY_NAME) || (echo "Clean failed"; exit 1)
endif

# Help command to display makefile usage
help:
	@echo "Usage: make [target]"
	@echo "Available targets:"
	@echo "  all                  : Build the project (default target)"
	@echo "  build                : Compile the project"
	@echo "  clean                : Remove the compiled binary"
	@echo "  clean-cross-platform : Remove the compiled binary (cross-platform)"
	@echo "  help                 : Display this help"
