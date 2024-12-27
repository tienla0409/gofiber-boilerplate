DB_SOURCE=postgresql://root:password@localhost:5432/gofiber-boilerplate?sslmode=disable
MIGRATION_DIR=db/migrations

migrate_create:
	migrate create -ext sql -dir internal/db/migrations -seq $(name)

dc_up:
	docker-compose up -do

dc_down:
	docker-compose down

migrate_up:
	migrate -path $(MIGRATION_DIR) -database "$(DB_SOURCE)" -verbose up

migrate_down:
	migrate -path $(MIGRATION_DIR) -database "$(DB_SOURCE)" -verbose down

sqlc:
	sqlc generate

start:
	air

PHONY: migrate_create migrate_up migrate_down sqlc start