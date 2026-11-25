APP_NAME := gibor
SRC_DIR := ./cmd/$(APP_NAME)/... # If the Source files are in the root of the project then use .
BIN_DIR := bin
GO_CMD := go
BUILD_CMD := $(GO_CMD) build
BUILD_FLAGS := -o
BUILD_ENVS := CGO_ENABLED=0

.PHONY: dev build run clean

dev:
	@echo "\nRunning the dev backend and frontend $(APP_NAME)"
	go tool task dev

build: clean
	@echo "\nBuilding $(APP_NAME)"
	go tool task build

clean:
	@echo "\nDeleting $(BIN_DIR)"
	go tool task clean

docker:
	@echo "\nBuilding Docker image"
	go tool task build:docker

docker-publish:
	@echo "\nBuilding Docker image"
	go tool task publish:docker

