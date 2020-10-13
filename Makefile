.PHONY: all build run fmt

all: build run fmt

Binary := "hash"

build:
	@go build -o ${Binary} .

fmt:
	@go fmt ./

run:
	@./hash