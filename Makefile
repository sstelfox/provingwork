COMPONENT_NAME = hashcash

CURRENT_GIT_REF = $(shell awk '{ print $$2 }' .git/HEAD)
VERSION = $(shell git describe --tags --always)
BUILD_TIME = $(shell date -u +'%Y%m%d-%H%M%S%z')
DIRTY = $(shell git diff --quiet --exit-code || echo '-dirty')

bin/$(COMPONENT_NAME): *.go .git/$(CURRENT_GIT_REF) Makefile
	godep go build -ldflags "-s -X main.Version=${VERSION}$(DIRTY) -X main.BuiltAt=${BUILD_TIME}" -o bin/$(COMPONENT_NAME) *.go

clean:
	rm -f ./bin/*

run:
	@make bin/$(COMPONENT_NAME)
	@./bin/$(COMPONENT_NAME)

setup:
	go get github.com/tools/godep
	go get golang.org/x/tools/cmd/cover
	godep get github.com/jteeuwen/go-bindata/...
	godep restore

test:
	godep go test -cover ./...

.DEFAULT_GOAL := bin/$(COMPONENT_NAME)
.PHONY: clean run setup test
