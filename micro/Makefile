GOPATH:=$(shell go env GOPATH)

.PHONY: build
build:
	cp ${GOPATH}/src/github.com/micro/micro/main.go .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o micro main.go plugin.go

.PHONY: docker
docker:
	docker build . -t micro-cors:latest