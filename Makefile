all: deps build

.PHONY: deps
deps:
	go get -v github.com/urfave/cli
	go get -v github.com/cloudfoundry-attic/jibber_jabber

.PHONY: build
build: deps
	mkdir -p dist
	go build -o dist/benpaoba src/*.go

.PHONY: run
run:
	go run src/*.go
