
GOPATH:=$(shell go env GOPATH)


.PHONY: build
build:

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o todos-web main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t mingz2013/todos-web:latest


.PHONY: docker-push
docker-push: docker
	docker login
	docker push mingz2013/todos-web