# Go parameters
GOVVVCMD=govvv
APP=golang-gin-starter
ALL=./...

-include .env

.EXPORT_ALL_VARIABLES:

.PHONY: all docs test install

all: test install
install: docs
	govvv install -pkg github.com/harbur/golang-gin-starter/pkgs/store
test:
	./scripts/coverage.sh
run: docs
	fresh
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

setup:
	go get -u github.com/ahmetb/govvv
	go get -u github.com/pilu/fresh
	go get -u github.com/swaggo/swag/cmd/swag
