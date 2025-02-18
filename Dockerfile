FROM golang:1.23-alpine3.19 AS builder

USER root

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    ca-certificates \
    gcc \
    g++ \
    git \
    tzdata

RUN go install github.com/google/wire/cmd/wire@latest && go install github.com/swaggo/swag/cmd/swag@latest

WORKDIR /app

COPY go.* .

RUN go mod download && go mod verify 

COPY . .

RUN swag init -d ./internal/pkg/server -g server.go -o ./api

RUN wire ./application

RUN GOOS=linux go build -o app -buildvcs=false /app/cmd/server/main.go

FROM alpine:3.19

USER root

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    ca-certificates

WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/api ./api
COPY --from=builder /app/configs ./configs

CMD ["./app"]
