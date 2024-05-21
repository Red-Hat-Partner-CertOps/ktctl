# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
TARGET=/usr/local/bin/ktctl
BUILD_DIR=./bin
SRC=./main.go

.PHONY: all build clean

all: build

build: clean
	@echo "Building binary..."
	@$(GOBUILD) -o $(TARGET) $(SRC)
	@echo "Build complete. Binary available at $(TARGET)"

clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Cleanup complete."
