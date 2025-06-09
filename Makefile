build:
	docker compose build 

network:
	docker network inspect primary >/dev/null 2>&1 || \
	docker network create primary

run:	network build
	docker compose up

migrate-up:
	goose -dir db/migrations postgres postgres://postgres:12345@localhost:5432/postgres?sslmode=disable up

migrate-down:
	goose -dir db/migrations postgres postgres://postgres:12345@localhost:5432/postgres?sslmode=disable down

generate-sql:
	sqlc generate

