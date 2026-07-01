run:
	go run ./cmd/api

build:
	mkdir -p bin
	go build -o bin/chirpy ./cmd/api

run-bin: build
	./bin/chirpy

sqlc:
	sqlc generate

up:
	docker compose up -d

down:
	docker compose down

test:
	go test ./...

migrate-up:
	goose -dir sql/schema postgres "$(DB_URL)" up

migrate-down:
	goose -dir sql/schema postgres "$(DB_URL)" down

migrate-status:
	goose -dir sql/schema postgres "$(DB_URL)" status

migrate-reset:
	goose -dir sql/schema postgres "$(DB_URL)" reset