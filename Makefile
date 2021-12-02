.PHONY: all swag clean check build-image run help

all: swag run

swag:
	@echo "swag tool required"
	swag init

clean:
	go clean

check:
	go fmt ./
	go vet ./

build-image:
	@echo "docker required"
	docker build . -t axiangcoding/gin-template:latest

run:
	go run ./

help:
	@echo "make - generate swagger docs, run application"
	@echo "make swag - generate swagger docs"
	@echo "make prepare - format codes"
	@echo "make build-image - build docker image"
	@echo "make run - run application"



