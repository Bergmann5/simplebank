postgres:
	docker run --name postres12 --network bank-network -p 8000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:8000/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:8000/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:8000/simple_bank?sslmode=disable" -verbose up1

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:8000/simple_bank?sslmode=disable" -verbose down1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postregs createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server