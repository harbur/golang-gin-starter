# Well documented Makefile following: https://www.thapaliya.com/en/writings/well-documented-makefiles/
.DEFAULT_GOAL:=help

# Go parameters
APP=golang-gin-starter
ALL=./...

-include .env

.EXPORT_ALL_VARIABLES:

.PHONY: all docs test install

all: docs test install ## Runs test and install
install: docs ## Installs app
	go install ./
test: docs  ## Runs tests
	go test ./...

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

run: ## Runs app in dev mode (listens at :8080)
	go run ./cmd/golang-gin-starter

deps: ## Checks dependencies
	go mod download
	go mod tidy
	go mod verify
docs: ## Generates swagger files
	swag init

changelog: ## Generates changelog
	conventional-changelog -r0 > CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "chore(changelog) update changelog"

setup: ## Installs dev dependencies
	go install github.com/swaggo/swag/cmd/swag@v1.16.4
