FROM golang:1.16.5-alpine AS stage0

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.io,direct"
ENV APK_REP="mirrors.ustc.edu.cn"

RUN sed -i "s/dl-cdn.alpinelinux.org/${APK_REP}/g" /etc/apk/repositories \
    && apk --no-cache add git ca-certificates

WORKDIR /go/src/go-netdisk/

COPY . .

RUN go env \
    && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server .


FROM alpine:latest

RUN addgroup -S gogo && adduser -S -G gogo gogo

WORKDIR /app

COPY --from=stage0 /go/src/go-netdisk/ .
COPY ["./k8s/go/entrypoint", "./k8s/go/start", "/"]
RUN chmod +x /entrypoint /start && chown -R gogo /app

USER gogo

EXPOSE 5000

ENTRYPOINT ["/entrypoint"]
