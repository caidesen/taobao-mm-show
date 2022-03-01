FROM golang:1.17-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM alpine:3.15.0 AS scratch
WORKDIR /app

COPY --from=builder ./build/apiserver /app
COPY ./public /app/public
EXPOSE 3000

ENTRYPOINT ["/app/apiserver"]