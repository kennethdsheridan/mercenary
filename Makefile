.PHONY: all clean

# Set the binary name and source directory
BINARY_NAME := mercenary
SOURCES_DIR := ./

# Determine the operating system
ifeq ($(shell uname -s),Linux)
	BINARY_EXT := linux
else
	BINARY_EXT := darwin
endif

# Set the output directory
OUTPUT_DIR := ./bin

# Set the output paths for Linux and macOS binaries
BINARY_LINUX := $(OUTPUT_DIR)/$(BINARY_NAME)_linux
BINARY_MACOS := $(OUTPUT_DIR)/$(BINARY_NAME)_macos

# Default target: build both Linux and macOS binaries
all: $(BINARY_LINUX) $(BINARY_MACOS)

# Build the Linux binary
$(BINARY_LINUX):
	GOOS=linux GOARCH=amd64 go build -trimpath -o $(BINARY_LINUX) $(SOURCES_DIR)

# Build the macOS binary
$(BINARY_MACOS):
	GOOS=darwin GOARCH=amd64 go build -trimpath -o $(BINARY_MACOS) $(SOURCES_DIR)

# Clean the generated binaries
clean:
	rm -rf $(OUTPUT_DIR)

