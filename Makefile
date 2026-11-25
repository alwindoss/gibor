APP_NAME := gibor

.PHONY: dev build run clean docker docker-publish

dev:
	@echo "\nRunning the backend and frontend in dev mode"
	go tool task dev

build: clean
	@echo "\nBuilding $(APP_NAME)"
	go tool task build

clean:
	@echo "\nCleaning up the build directories"
	go tool task clean

docker:
	@echo "\nBuilding Docker image"
	go tool task build:docker

docker-publish:
	@echo "\nBuilding Docker image"
	go tool task publish:docker

