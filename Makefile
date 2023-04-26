# Makefile

# Set the name of the binary
BIN_NAME := myapp

# Set the Go command
GO_CMD := go

# Set the binary output directory
BIN_DIR := bin

# Set the path to the binary
BIN_PATH := $(BIN_DIR)/$(BIN_NAME)

# Default target, builds and runs the binary
.PHONY: all
all: build run

# Build the binary
.PHONY: build
build:
	$(GO_CMD) build -o $(BIN_PATH) .

# Create the bin directory if it doesn't exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Run the binary
.PHONY: run
run: $(BIN_DIR)
	$(BIN_PATH)

# Clean up the bin directory
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

