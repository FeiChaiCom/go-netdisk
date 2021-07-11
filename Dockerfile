FROM golang:1.16.5-alpine

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.io,direct"

RUN apk --no-cache add git ca-certificates

WORKDIR /go/src/github.com/gaomugong/go-netdisk/
COPY . .

RUN go env \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .

FROM alpine:latest
WORKDIR /go/src/github.com/gaomugong/go-netdisk/

COPY --from=0 /go/src/github.com/gaomugong/go-netdisk/ ./

EXPOSE 5000

ENTRYPOINT ./server -c .envs/docker.yaml
