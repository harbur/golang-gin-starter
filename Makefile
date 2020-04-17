# Go parameters
GOVVVCMD=govvv
APP=golang-gin-starter
ALL=./...

-include .env

.EXPORT_ALL_VARIABLES:

.PHONY: all test install

all: test install
install: docs
	govvv install
test:
	./scripts/coverage.sh
run: docs
	go run main.go
deps:
	go mod download
	go mod tidy
	go mod vendor
	go mod verify
docs:
	swag init
int:
	cd test/ && ./integration.sh

changelog:
	conventional-changelog -r0 > CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "chore(changelog) update changelog"
