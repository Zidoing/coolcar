FROM golang:1.15-alpine AS builder

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY . /go/src/coolcar/server/

WORKDIR /go/src/coolcar/server

RUN go install ./gateway/...

FROM alpine:3.13

COPY --from=builder /go/bin/gateway /bin/gateway

EXPOSE 8080

ENTRYPOINT ["/bin/gateway"]