# FROM golang:1.15
FROM golang:alpine3.13 AS build-env

# Set up dependencies
ENV PACKAGES bash curl make git libc-dev gcc linux-headers eudev-dev python3

WORKDIR /mempool-logger

COPY go.mod .
COPY go.sum .

COPY . .

RUN apk add --no-cache $PACKAGES && go build

FROM alpine:edge

RUN apk add --update ca-certificates

WORKDIR /mempool-logger

COPY --from=build-env /mempool-logger/minimal-mempool-logger /usr/bin/minimal-mempool-logger