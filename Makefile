wire:
	wire ./application

swagger:
	swag init -d ./internal/pkg/server -g server.go -o ./api

generate: wire swagger

migrate-up: 
	migrate -path ./scripts/migrations -database "mysql://root:root@tcp(localhost:3306)/db_todo" -verbose up

test-unit:
	go test ./... -cover

test-e2e:
	docker-compose down --remove-orphans
	docker-compose -f docker-compose.test.yml up --exit-code-from=app-test --build

run-server:
	go run cmd/server/main.go

docker-up:
	docker-compose down --remove-orphans
	docker-compose up --build