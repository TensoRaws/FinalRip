FROM golang:1.22-bookworm AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /build

COPY . .

RUN go mod download

RUN make worker

FROM lychee0/vs-pytorch:dev AS app

# worker app
WORKDIR /app

ENV TZ=Asia/Shanghai

ENV FINALRIP_EASYTIER_HOST 10.126.126.251

COPY --from=builder /build/worker/worker /app/
COPY --from=builder /build/conf/finalrip.yml /app/conf/

CMD ["/app/worker", "encode"]
