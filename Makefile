run:
	go run ./cmd/api

build:
	go build -o chirpy ./cmd/api

run-bin: build
	./chirpy

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