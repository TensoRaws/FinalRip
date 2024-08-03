FROM golang:1.22-bookworm AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go mod download

RUN make server

FROM alpine:3.20 AS app

WORKDIR /app

COPY --from=builder /build/server/server /app/
COPY --from=builder /build/conf/finalrip.yml /app/conf/

EXPOSE 8848

CMD ["/app/server", "server"]
