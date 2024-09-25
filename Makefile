SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all

BINARY_NAME=chatgpt-telegram
IMAGE_TAG=$(shell date +%Y%m%d%H%M%S)

tidy:
	go mod tidy

build:
	CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -o bin/${BINARY_NAME}-darwin-arm64 cmd/${BINARY_NAME}/main.go
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-amd64 cmd/${BINARY_NAME}/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o bin/${BINARY_NAME}-linux-arm64 cmd/${BINARY_NAME}/main.go
	CGO_ENABLED=0 GOARCH=arm GOOS=linux go build -o bin/${BINARY_NAME}-linux-arm cmd/${BINARY_NAME}/main.go
	chmod +x bin/*

test:
	go test ./...

docker:
	docker system prune -a
	docker build -t loqutus/$(BINARY_NAME):$(IMAGE_TAG) -f Dockerfile .
	docker push loqutus/$(BINARY_NAME):$(IMAGE_TAG)
	docker tag loqutus/$(BINARY_NAME):$(IMAGE_TAG) loqutus/$(BINARY_NAME):latest
	docker push loqutus/$(BINARY_NAME):latest

update:
	go get -u ./...
	go mod tidy

get:
	go get ./...

deploy:
	scp bin/chatgpt-telegram-linux-arm root@orangepi.govno.cloud:~/

default: tidy build