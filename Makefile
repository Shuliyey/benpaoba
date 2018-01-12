all: deps build

.PHONY: deps
deps:
	go get -v github.com/urfave/cli
	#go get -d -v github.com/dustin/go-broadcast/...
	#go get -d -v github.com/manucorporat/stats/...

.PHONY: build
build: deps
	mkdir -p dist
	go build -o dist/benpaoba src/*.go
