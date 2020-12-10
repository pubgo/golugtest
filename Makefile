Project=github.com/pubgo/golug
GOPath=$(shell go env GOPATH)
Version=$(shell git tag --sort=committerdate | tail -n 1)
GoROOT=$(shell go env GOROOT)
BuildTime=$(shell date "+%F %T")
CommitID=$(shell git rev-parse HEAD)
LDFLAGS=-ldflags " \
-X '${Project}/golug_version.GoROOT=${GoROOT}' \
-X '${Project}/golug_version.BuildTime=${BuildTime}' \
-X '${Project}/golug_version.GoPath=${GOPath}' \
-X '${Project}/golug_version.CommitID=${CommitID}' \
-X '${Project}/golug_version.Project=${Project}' \
-X '${Project}/golug_version.Version=${Version:-v0.0.1}' \
"

.PHONY: build
build:
	@go build ${LDFLAGS} -mod vendor -race -v -o main cmd/golug/main.go

build_hello_test:
	go build ${LDFLAGS} -mod vendor -v -o main  example/hello/main.go

.PHONY: install
install:
	@go install ${LDFLAGS} .

.PHONY: release
release:
	@go build ${LDFLAGS} -race -v -o main main.go

.PHONY: test
test:
	@go test -race -v ./... -cover

.PHONY: proto
proto: clear
	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/gogo/protobuf \
   --go_out=plugins=grpc:. \
   --golug_out=. \
	proto/hello/*

	protoc -I. \
   -I/usr/local/include \
   -I${GOPATH}/src \
   -I${GOPATH}/src/github.com/googleapis/googleapis \
   -I${GOPATH}/src/github.com/gogo/protobuf \
   --go_out=plugins=grpc:. \
   --golug_out=. \
	proto/login/*

.PHONY: clear
clear:
	rm -rf proto/*.go
	rm -rf proto/**/*.go

.PHONY: build-example
build-example:
	go build ${LDFLAGS} -mod vendor -v -o main example/main.go


.PHONY: ossync
ossync:
	cd ossync && go install ${LDFLAGS} -v .

.PHONY: tickrun
tickrun:
	cd tickrun && go install ${LDFLAGS} -v .