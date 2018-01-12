all: deps build

.PHONY: deps
deps:
	go get -v github.com/urfave/cli

.PHONY: build
build: deps
	mkdir -p dist
	go build -o dist/benpaoba src/*.go

.PHONY: run
run:
	go run src/*.go
