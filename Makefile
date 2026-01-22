APP_NAME=kasir-api
BUILD_DIR=bin

.PHONY: all build run clean tidy

all: build

build:
	mkdir -p $(BUILD_DIR)
	go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) main.go

run:
	go run main.go

run-build:
	./$(BUILD_DIR)/$(APP_NAME)

clean:
	rm -rf $(APP_NAME)

tidy:
	go mod tidy