# go-netdisk Makefile

the_edition?=inner_edition

ifdef HASTAG
	GITTAG=${HASTAG}
else
	GITTAG=$(shell git describe --always)
endif

PROJECT=go-netdisk
REGISTRY=registry.feichai.com/
BUILDTIME = $(shell date "+%Y-%m-%dT%T%z")
GITHASH=$(shell git rev-parse  HEAD)
VERSION=${GITTAG}-$(shell date +%y.%m.%d)

LDFLAGS=-ldflags "-X go-netdisk/pkg/version.PlatformName=${PROJECT} \
 -X go-netdisk/pkg/version.Version=${VERSION} \
 -X go-netdisk/pkg/version.BuildTime=${BUILDTIME} \
 -X go-netdisk/pkg/version.GitHash=${GITHASH} \
 -X go-netdisk/pkg/version.Tag=${GITTAG} \
 -X go-netdisk/pkg/version.Edition=${the_edition}"

# build path config
PACKAGE = ./build/${PROJECT}-${VERSION}
WORKSPACE=$(shell pwd)

# options
default: build

pre:
	@echo "git tag: ${GITTAG}"
	mkdir -p ${PACKAGE}
	go fmt ./...

proto:
	protoc --proto_path=. \
		--proto_path=./pkg/third_party/googleapis \
		--proto_path=./pkg/third_party/grpc-gateway \
		--micro_out=. \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true,register_func_suffix=Gw:. --swagger_out=logtostderr=true:. \
		./pkg/apis/proto/api.proto

clean:
	rm -rf ./build

build:pre
	mkdir -p ${PACKAGE}
	CGO_ENABLED=0 go build ${LDFLAGS} -o ${WORKSPACE}/${PACKAGE}/server ./main.go
	@echo "${WORKSPACE}/${PACKAGE}/server"

build-dev:pre
	mkdir -p ${PACKAGE}
	CGO_ENABLED=0 go build ${LDFLAGS} -o ./server ./main.go

build-linux:pre
	mkdir -p ${PACKAGE}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -v -x -o ${WORKSPACE}/${PACKAGE}/server ./main.go