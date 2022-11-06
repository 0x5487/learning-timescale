SHELL := /bin/bash

APP := candlestick

VERSION := $(shell git describe --tags --always)
FULL_COMMIT := $(shell git rev-parse HEAD)
RELEASE_DATE := $(shell git show -s --format=%cI)
BUILD_DATE = $(shell date '+%FT%T%z')
LDFLAGS = -ldflags "-X $(APP)/internal/pkg/version.Version=$(VERSION) -X $(APP)/internal/pkg/version.FullCommit=$(FULL_COMMIT) -X $(APP)/internal/pkg/version.ReleaseDate=$(RELEASE_DATE) -X $(APP)/internal/pkg/version.BuildDate=$(BUILD_DATE)"

.PHONY: candlestick release clean testdata lint run

testdata:
	go build -o candlestick && ./candlestick testdata

candlestick:
	go build -o candlestick && ./candlestick

run:
	go run $(LDFLAGS) -race .

release: clean
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o ./bin/$(APP)-$(VERSION)

clean:
	rm -rf ./bin

lint:
	golangci-lint run ./... -v