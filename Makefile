postgres:
	docker run --name go-bank-db -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:alpine

createdb:
	docker exec -it go-bank-db createdb --username=postgres --owner=postgres go_bank

dropdb:
	docker exec -it go-bank-db dropdb go_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nochzato/go-bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock
