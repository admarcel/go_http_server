FROM golang:1.10 as builder
ENV GOOS=linux
ENV CGO_ENABLED=0
RUN apt-get update
RUN apt-get install -y build-essential git make
WORKDIR /go/src/github.com/admarcel/go_http_server
COPY ./ .
RUN dep ensure
RUN make

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root
COPY --from=builder /go/src/github.com/admarcel/go_http_server/bin/go_http_server ./go_http_server
CMD ["/root/go_http_server"]
