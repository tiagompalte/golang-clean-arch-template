wire:
	wire ./application

swagger:
	swag init -d ./internal/pkg/server -g server.go -o ./api

generate: wire swagger

migrate-up: 
	go run cmd/migrate/main.go up

migrate-down: 
	go run cmd/migrate/main.go down

# make migrate-create name=?
migrate-create: 
	go run cmd/migrate/main.go create $(name)

test-unit:
	go test ./... -cover

test-e2e:
	docker compose down --remove-orphans
	docker compose -f docker-compose.test.yml up --exit-code-from=app-test --build

docker-test-up-db:
	docker compose -f docker-compose.test.yml up mysql

run-server:
	go run cmd/server/main.go

docker-up-all:
	docker compose down --remove-orphans
	docker compose up --build

docker-up-db-cache:
	docker compose down --remove-orphans
	docker compose up --build mysql redis