ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk upgrade && \
    apk --update add git make build-base

WORKDIR /app

EXPOSE 9090

CMD go run ./cmd/go-api-server/api_server.go
