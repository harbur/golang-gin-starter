# Go parameters
GOVVVCMD=govvv
APP=golang-gin-starter
ALL=./...

-include .env

.EXPORT_ALL_VARIABLES:

.PHONY: all docs test install

all: docs test install
install: docs
	govvv install -pkg github.com/harbur/golang-gin-starter/pkgs/store
test: docs
	./scripts/coverage.sh

generate:
	swag init

run:
	air

deps:
	go mod download
	go mod tidy
	go mod verify
docs:
	swag init
int:
	cd test/ && ./integration.sh

changelog:
	conventional-changelog -r0 > CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "chore(changelog) update changelog"

setup:
	go get -u github.com/ahmetb/govvv
	go get -u github.com/cosmtrek/air
	go get -u github.com/swaggo/swag/cmd/swag
