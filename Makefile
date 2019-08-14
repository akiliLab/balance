
.PHONY: install test build serve clean pack deploy ship

TAG?=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

export TAG

install:
	go get ./...

test: install
	go test ./...

proto-gen:
	protoc -I /usr/local/include -I. \
    	-I $(GOPATH)/src \
    	-I $(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--go_out=plugins=grpc:. \
    	--grpc-gateway_out=logtostderr=true:. \
		--swagger_out=logtostderr=true:. \
    	./proto/*.proto

docker-build:
	docker build -t akililab/balance:v0.0.1 .