
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/todo/todo.proto

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todos-srv main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t todos-srv:latest
