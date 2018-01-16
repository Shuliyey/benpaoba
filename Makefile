all: deps build

.PHONY: deps
deps:
	go get -v github.com/urfave/cli
	go get -v github.com/cloudfoundry-attic/jibber_jabber

.PHONY: build
build: deps
	mkdir -p bin
	go build -o bin/benpaoba src/*.go

.PHONY: run
run:
	go run src/*.go
