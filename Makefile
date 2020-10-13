.PHONY: all build run fmt test

all: build run fmt

Binary := "hash"

build:
	@go build -o ${Binary} .

fmt:
	@go fmt ./

run:
	@./hash

test:
	@go test ./test -v