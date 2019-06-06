
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/todo/todo.proto

.PHONY: build
build: proto

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todos-api main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker: build
	docker build . -t mingz2013/todos-api:latest

.PHONY: docker-push
docker-push: docker
	docker login
	docker push mingz2013/todos-api