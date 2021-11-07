.PHONY: lint test scan

full-test: lint test

ifeq ($(OS),Windows_NT)
    RM = del //Q //F
    RRM = rmdir //Q //S
else
    RM = rm -f
    RRM = rm -f -r
endif

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