FROM golang:1.18-alpine as builder

ENV GO111MODULE=on

RUN apk add --no-cache git make

WORKDIR /app

COPY . .

RUN make build

FROM alpine

COPY --from=builder /app/bin/order-pack /order-pack

ENTRYPOINT [ "/order-pack"]