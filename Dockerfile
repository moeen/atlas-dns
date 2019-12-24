FROM golang:1.12-alpine

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN apk update && apk add --no-cache git

RUN go mod download
RUN go build -o dns_server

RUN chmod +x /app/dns_server

ENV GIN_MODE release

ENTRYPOINT ["/app/dns_server"]