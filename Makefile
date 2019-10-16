# Go parameters
GOCMD=go

deps:
	$(GOCMD) mod download
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor
	$(GOCMD) mod verify
