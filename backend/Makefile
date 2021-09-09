BINARY_PATH=./bin/server/server.out
SERVER_PATH=./cmd/server/main.go
DB_CREATE_PATH=./cmd/database/create.go
DB_MIGRATE_PATH=./cmd/database/migrate.go
DB_DROP_PATH=./cmd/database/drop.go
PROTO_DEF_PATH=../api-specs
PROTO_GO_PATH=./protobuf

createdb:
	@echo "Creating database..."
	@go run ${DB_CREATE_PATH}

migratedb:
	@echo "Migrating database..."
	@go run ${DB_MIGRATE_PATH}

dropdb:
	@echo "Dropping database..."
	@go run ${DB_DROP_PATH}

build: proto
	@echo "Building server at ${BINARY_PATH}..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}

run: proto
	@echo "Running server..."
	@go build -o ${BINARY_PATH} ${SERVER_PATH}
	@eval ${BINARY_PATH}

fmt:
	@echo "Formatting source files..."
	@go fmt ./...

clean:
	@echo "Removing object files and cached files..."
	@go clean
	@rm ${BINARY_PATH}

proto:
	@echo "Generating Go code for Protocol Buffer definitions..."
	@rm -rf ${PROTO_GO_PATH}
	@mkdir ${PROTO_GO_PATH}
	@protoc --proto_path=${PROTO_DEF_PATH} --go_out=${PROTO_GO_PATH} --go-grpc_out=${PROTO_GO_PATH} $(wildcard ${PROTO_DEF_PATH}/*.proto)
