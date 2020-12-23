APP ?= events-fanout
CMD ?= cmd
BIN ?= bin
TAG ?= latest

## build: build the application
.PHONY: build
build: clean
	@echo "Building"
	go build -o ${BIN}/${APP} ${CMD}/main.go

## build-linux: build the application for linux platform
.PHONY: build-linux
build-linux: clean
	@echo "Building for Linux"
	GOOS=linux GOARCH=amd64 go build -o ${BIN}/${APP}-linux-amd64 ${CMD}/main.go

## build-windows: build the application for Windows platform
.PHONY: build-windows
build-windows: clean
	@echo "Building for Windows"
	GOOS=windows GOARCH=amd64 go build -o ${BIN}/${APP}-windows-amd64 ${CMD}/main.go

## lint: lint the application
.PHONY: lint
lint:
	golint ./...

## run: run the application
.PHONY: run
run:
	go run ${CMD}/main.go

## test: test the application
.PHONY: test
test:
	go test -race ./...

## clean: cleans binary
.PHONY: clean
clean:
	@echo "Cleaning"
	go clean

## test: run tests with cache disabled
.PHONY: test
test:
	go test -v -count=1 -race ./...

## docker-build: builds docker image
.PHONY: docker-build
docker-build:
	docker build -t ${APP}:${TAG} .

## docker-run: run the docker image
.PHONY: docker-run
docker-run:
	docker run ${APP}:${TAG}

## help: Prints this help message
.PHONY: help
help:
	@echo "Usage: \n"
	sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'