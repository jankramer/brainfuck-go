SHELL = /bin/bash

.PHONY: all test lint vet
all: test lint vet bin/brainfuck

test:
	go test -cover .

lint:
	golint .

vet:
	go vet .

bin/brainfuck: $(shell find . -name '*.go')
	go build -o $@

