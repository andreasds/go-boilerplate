ARG GO_VERSION

FROM golang:${GO_VERSION}-alpine

RUN apk update && apk upgrade && \
    apk --update add git make build-base

WORKDIR /app

CMD make test
