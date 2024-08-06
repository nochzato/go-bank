postgres:
	docker run --name go-bank-db -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres:16-alpine

createdb:
	docker exec -it go-bank-db createdb --username=postgres --owner=postgres go_bank

dropdb:
	docker exec -it go-bank-db dropdb go_bank

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/go_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/nochzato/go-bank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distibutor.go github.com/nochzato/go-bank/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=go_bank \
	proto/*.proto

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server mock proto redis
