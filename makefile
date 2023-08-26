wire:
	wire ./application

swagger:
	swag init -d ./internal/pkg/server -g server.go -o ./api

generate: wire swagger

migrate-up: 
	migrate -path ./scripts/migrations -database "mysql://root:root@tcp(localhost:3306)/db_todo" -verbose up

tests:
	go run github.com/onsi/ginkgo/v2/ginkgo -r

run-server:
	go run cmd/server/main.go