APP_NAME=kasir-api
BUILD_DIR=bin
VERSION := $(shell git describe --tags --always 2>/dev/null || echo "v0.0.0")

.PHONY: all build run run-build clean tidy tag

all: build

build:
	mkdir -p $(BUILD_DIR)
	# This finds the main.go file regardless of the subfolder name under cmd/
	go build -ldflags="-s -w -X main.version=$(VERSION)" \
		-o $(BUILD_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)/*.go

run:
	go run ./cmd/$(APP_NAME)

run-build:
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(BUILD_DIR)/$(APP_NAME)

tag:
	git tag $(VERSION)
	git push origin $(VERSION)

tidy:
	go mod tidy