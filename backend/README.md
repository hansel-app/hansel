# Hansel Backend

## Getting Started

1. Download and install Go `>= 1.17` [here](https://golang.org/doc/install).
1. Install the protocol compiler plugins for Go using the following commands:
   ```sh
   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
   ```
1. Update your `PATH` so that the protocol compiler can find the plugins:
   ```sh
   $ export PATH="$PATH:$(go env GOPATH)/bin"
   ```
1. Run `make run` to start the server.

## Updating Protocol Buffer Definitions

After changing the definitions of the Protocol Buffer, run `make proto` to regenerate the respective Go files.
Note that `make build` and `make run` also regenerate the Go files automatically.

## Architecture

The Hansel backend makes use of the Ports & Adapters architecture.
Using the principle of dependency inversion, it allows for loosely-coupled components that can be easily swapped out by means of ports and adapters.
This not only enforces a strong separation of concerns, it also makes each component much more testable.
It is also overkill for this assignment and wholly unnecessary.

Do read this [article](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/) to get a better overview of this architecture.
