SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all
BINARY_NAME=chatgpt-tg

tidy:
	go mod tidy

build:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/${BINARY_NAME}/main.go
	chmod +x bin/*

default: tidy build