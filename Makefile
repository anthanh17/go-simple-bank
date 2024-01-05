# Setup postgres database docker
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=abc123 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

# Migarte database
migrateup:
	migrate -path db/migration -database "postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable" -verbose down

# sqlc gen code golang
sqlc:
	sqlc generate

# Unit test
test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
