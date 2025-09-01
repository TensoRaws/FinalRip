FROM golang:1.24-bookworm AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go mod download

RUN make server

FROM debian:bookworm AS app

WORKDIR /app

ENV TZ=Asia/Shanghai

ENV FINALRIP_EASYTIER_HOST 10.126.126.251

COPY --from=builder /build/server/server /app/
COPY --from=builder /build/conf/finalrip.yml /app/conf/

EXPOSE 8848

CMD ["/app/server", "server"]
