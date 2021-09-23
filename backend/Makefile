BINARY_PATH=./bin/server/server.out
SERVER_PATH=./cmd/server/main.go
DB_CREATE_PATH=./cmd/database/create/main.go
DB_MIGRATE_PATH=./cmd/database/migrate/main.go
DB_SEED_PATH=./cmd/database/seed/main.go
DB_DROP_PATH=./cmd/database/drop/main.go
PROTO_DEF_PATH=../api-specs
PROTO_GO_PATH=./protobuf
PROTO_IMPORT_PREFIX=github.com/hansel-app/hansel/protobuf

createdb:
	@echo "Creating database..."
	@go run ${DB_CREATE_PATH}

migratedb:
	@echo "Migrating database..."
	@go run ${DB_MIGRATE_PATH}

seeddb:
	@echo "Seeding database..."
	@go run ${DB_SEED_PATH}

dropdb:
	@echo "Dropping database..."
	@go run ${DB_DROP_PATH}

setupdb: createdb migratedb seeddb

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

lint:
	@echo "Linting source files..."
	@golangci-lint run

lintfix:
	@echo "Linting and automatically fixing source files..."
	@golangci-lint run --fix

clean:
	@echo "Removing object files and cached files..."
	@go clean
	@rm ${BINARY_PATH}

proto:
	@echo "Generating Go code for Protocol Buffer definitions..."
	@rm -rf ${PROTO_GO_PATH}
	@mkdir ${PROTO_GO_PATH}
	@protoc --proto_path=${PROTO_DEF_PATH} --go_out=${PROTO_GO_PATH} --go-grpc_out=${PROTO_GO_PATH} $(wildcard ${PROTO_DEF_PATH}/*.proto) --go_opt=module=${PROTO_IMPORT_PREFIX} --go-grpc_opt=module=${PROTO_IMPORT_PREFIX}
