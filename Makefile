APP_NAME=kasir-api
BUILD_DIR=bin
VERSION := $(shell git describe --tags --always)

.PHONY: all build run run-build clean tidy tag

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w -X main.version=$(VERSION)" \
		-o $(BUILD_DIR)/$(APP_NAME) ./cmd/$(APP_NAME)

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