migrateup:
	@echo "Running migrations up..."
	@migrate -path=shared/sql/up -database="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"  -verbose up

migratedown:
	@echo "Running migrations down..."
	@migrate -path=shared/sql/down -database="postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"  -verbose down

protogen:
	@echo "Generating protobuf files..."
	@protoc --go_out=. --go-grpc_out=. shared/proto/*.proto

gorunmain:
	@echo "Running main.go..."
	@go run api/cmd/server/main.go


.PHONY: migrateup migratedown gorunmain protogen