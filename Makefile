# Go parameters
GOCMD=go
GOVVVCMD=govvv
SRC=./cmd/golang-starter
APP=golang-starter
ALL=./...
GOINSTALL=$(GOVVVCMD) install -pkg github.com/harbur/golang-starter/pkgs/apis
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGENERATE=$(GOCMD) generate
DEP=dep
DEPENSURE=$(DEP) ensure
DOCS=$(SWAGGER) serve
# Swagger Commands
SWAGGERCMD=swagger
SWAGGERSERVE=$(SWAGGERCMD) serve
SWAGGERGEN=$(SWAGGERCMD) generate

-include .env

.EXPORT_ALL_VARIABLES:

.PHONY: all test install

all: test install
install:
	$(GOINSTALL) $(SRC)
test:
	./scripts/coverage.sh
run: install
	$(APP) --port 8080
deps:
	$(GOCMD) mod download
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor
	$(GOCMD) mod verify
generate:
	$(SWAGGERGEN) server
docs:
	$(SWAGGERSERVE) swagger.yaml
int:
	cd test/ && ./integration.sh

changelog:
	conventional-changelog -r0 > CHANGELOG.md
	git add CHANGELOG.md
	git commit -m "chore(changelog) update changelog"
