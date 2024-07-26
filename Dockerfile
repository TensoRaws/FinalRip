FROM golang:1.22-alpine3.20 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go mod download

