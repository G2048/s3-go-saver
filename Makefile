# Define variables
APPNAME = s3-cli
BINARY_NAME := ./build/$(APPNAME)
MAIN_GO = ./cmd/cli/main.go
# Use bash syntax
SHELL = /bin/bash

# Default target
# all: build

git_tag:
	git tag $(shell git cliff --bumped-version 2>/dev/null)

changelog:
	make git_tag
	git cliff --config detailed -o CHANGELOG.md

tidy: ## runs tidy to fix go.mod dependencies
	go mod tidy

build:
	make tidy
	GOOS=linux GOARCH=amd64 go build -v -buildvcs -ldflags '-s -w' -o $(BINARY_NAME) $(MAINGO)
	GOOS=windows GOARCH=amd64 go build -v -buildvcs -ldflags '-s -w' -o $(BINARY_NAME)-windows $(MAINGO)
	# GOOS=linux GOARCH=amd64 CGO_ENABLED=1 CC="zig cc -target x86_64-linux-gnu" go build -v -buildvcs -o $(BINARY_NAME)-zig $(MAIN_GO)
	printf "\n"
	make changelog

run:
	make -B build
	$(BINARY_NAME) -version

test:
	go test -v ./... -count=1
