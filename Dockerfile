FROM golang:1.22-alpine3.19 AS builder

USER root

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    ca-certificates \
    gcc \
    g++ \
    git \
    tzdata

WORKDIR /app
COPY . .

RUN go mod download && go install github.com/swaggo/swag/cmd/swag@latest && swag init -d ./internal/pkg/server -g server.go -o ./api

RUN go install github.com/google/wire/cmd/wire@latest

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
