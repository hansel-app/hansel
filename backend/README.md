# Hansel Backend

## Getting Started

1. Download and install Go `>= 1.17` [here](https://golang.org/doc/install).
1. Install the protocol compiler plugins for Go using the following commands:
   ```sh
   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
   ```
1. Run `make run` to start the server.

## Updating Protocol Buffer Definitions

After changing the definitions of the Protocol Buffer, run `make proto` to regenerate the respective Go files.
Note that `make build` and `make run` also regenerate the Go files automatically.
