BINARY_PATH=./bin/server/server.out
SERVER_PATH=./cmd/server/main.go
PROTO_DEF_PATH=../api-specs
PROTO_GO_PATH=./protobuf

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
	@protoc --proto_path=${PROTO_DEF_PATH} --go_out=${PROTO_GO_PATH} $(wildcard ${PROTO_DEF_PATH}/*.proto)
