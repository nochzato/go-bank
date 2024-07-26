postgres:
	docker run --name go-bank-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:alpine

createdb:
	docker exec -it go-bank-db createdb --username=root --owner=root go_bank

dropdb:
	docker exec -it go-bank-db dropdb go_bank

migrateup:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/go_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/go_bank?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown
