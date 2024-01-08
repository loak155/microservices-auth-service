include .env
DB_URL=postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

.PHONY: gen-proto
gen-proto:
	protoc -I. --go_out=. --go-grpc_out=. proto/*.proto

.PHONY: run-server
run-server:
	GO_ENV=dev go run .

.PHONY: new-migration
new-migration:
	migrate create -ext sql -dir migrations -seq $(name)

.PHONY: migrateup
migrateup:
	migrate -path migrations -database "$(DB_URL)" -verbose up

.PHONY: migrateup1
migrateup1:
	migrate -path migrations -database "$(DB_URL)" -verbose up 1

.PHONY: migratedown
migratedown:
	migrate -path migrations -database "$(DB_URL)" -verbose down

.PHONY: migratedown1
migratedown1:
	migrate -path migrations -database "$(DB_URL)" -verbose down 1
