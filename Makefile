.PHONY: lint test scan

full-test: lint test

GIT_COMMIT ?= $(shell git rev-parse --verify HEAD)
GIT_VERSION ?= $(shell git describe --tags --always --dirty="-dev")
DATE ?= $(shell date -u '+%Y-%m-%d %H:%M UTC')
BUILDER ?= Makefile
VERSION_FLAGS := -X "main.version=$(GIT_VERSION)" -X "main.date=$(DATE)" -X "main.commit=$(GIT_COMMIT)" -X "main.builtBy=$(BUILDER)"

build:
	go build -ldflags='$(VERSION_FLAGS)' ./...

lint:
	go vet ./...
	golangci-lint run -E revive,gofmt ./...

test:
	go test -race -v -coverprofile="c.out" ./...
	go tool cover -func="c.out"

scan:
	gosec -no-fail -fmt sarif -out security.sarif ./...

run:
	go build -o api
	api

bench:
	go test -bench=. ./...