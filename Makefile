
IMG ?= sample-service:latest

run:
	go run cmd/sample-service/main.go

build:
	go build -o sample-service ./cmd/sample-service

docker-build:
	docker build -t $(IMG) -f Dockerfile .
