FROM golang:1.20-alpine3.16 as builder
WORKDIR /go/src/kubemanagement.com/server
COPY . .

RUN go env -w GO111MODULE=on \
   && go env -w GOPROXY=https://goproxy.cn,direct \
   && go env -w CGO_ENABLED=0 \
   && go env \
   && go mod tidy \
   && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="zj20162325@163.com"

WORKDIR /go/src/kubemanagement.com/server
COPY --from=0 /go/src/kubemanagement.com/server/config.yaml ./config.yaml
COPY --from=0 /go/src/kubemanagement.com/server/.kube/config ./.kube/config
COPY --from=0 /go/src/kubemanagement.com/server/server ./
EXPOSE 8082
ENTRYPOINT ./server