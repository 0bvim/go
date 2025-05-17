# This file is just to compile golang for "any" plataform

# Detect the host operating system
HOST_OS := $(shell uname -s)

# Default values for GOOS and GOARCH
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

# Set the Go binary name
BINARY_NAME := my_program

# Compilation flags
BUILD_FLAGS :=

ifeq ($(GOOS), windows)
    BINARY_NAME := my_program.exe
    BUILD_FLAGS += -ldflags "-H=windowsgui"
else
    BUILD_FLAGS += -ldflags "-s -w"
endif

# Default target
all: build

# Build target
build:
	@echo "Building for GOOS=$(GOOS), GOARCH=$(GOARCH) on Host OS: $(HOST_OS)"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(BUILD_FLAGS) -o $(BINARY_NAME) main.go

# Clean target
clean:
	@echo "Cleaning up..."
	rm -f $(BINARY_NAME)

# Phony targets
.PHONY: all build clean

