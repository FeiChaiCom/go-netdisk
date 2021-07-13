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


FROM node as stage1

WORKDIR /root/web
COPY web .

RUN npm config set registry=http://registry.npm.taobao.org \
    && npm install \
    && npm run build 

FROM alpine

WORKDIR /app


COPY --from=stage0 /go/src/go-netdisk/ .
COPY --from=stage1 /root/static ./static
COPY --from=stage1 /root/templates ./templates
COPY ["./k8s/go/entrypoint", "./k8s/go/start", "/"]

RUN addgroup --system saas && adduser -S -G saas saas

RUN chmod +x /entrypoint /start \
    && chown -R saas /app

USER saas

EXPOSE 5000

ENTRYPOINT ["/entrypoint"]
