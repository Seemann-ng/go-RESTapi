.PHONY: build
build:
	go build -v .\cmd\apiserver

.PHONY: test
test:
	go test -v -cover -race -timeout 30s ./...

.DEFAULT_GOAL := build
