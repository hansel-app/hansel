# Hansel Backend

## Getting Started

1. Download and install Go `>= 1.17` [here](https://golang.org/doc/install).
1. Install the protocol compiler plugins for Go using the following commands:
   ```sh
   $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```
1. Update your `PATH` so that the protocol compiler can find the plugins:
   ```sh
   $ export PATH="$PATH:$(go env GOPATH)/bin"
   ```
1. Make a copy of `.env.development` as `.env.development.local`.
   For development on your local system, we recommend that you connect to PostgreSQL via Unix-domain sockets so that there is no need for database server credentials, allowing you to leave `DB_USERNAME` and `DB_PASSWORD` empty.
   This can be done by setting the DB_HOSTNAME to the following locations:
   * Linux: `/var/run/postgresql`
   * macOS: `/tmp`

   Note that if you are connecting via `localhost`, you might need to set `DB_SSLMODE=disable`.

   For example:
   ```
   SERVER_PORT=8000
   SECRET_KEY=<output of `openssl rand -hex 64`>

   DB_HOSTNAME=/var/run/postgresql
   DB_PORT=5432
   DB_NAME=hansel_development
   DB_USERNAME=
   DB_PASSWORD=
   DB_SSLMODE=
   ```

   The `SECRET_KEY` is used to sign and verify the JSON Web Tokens.
   It is recommended to be sufficiently long and can be generated via `openssl rand -hex 64`.
1. Install PostgreSQL `>= 12` by following the instructions [here](https://www.postgresql.org/download/).
1. Create and migrate the database:
   ```sh
   $ make createdb
   $ make migratedb
   ```
1. Install `golangci-lint` by following the instructions [here](https://golangci-lint.run/usage/install/#local-installation).
1. Optionally, you can integrate the linter with your IDE if it is supported by following the instructions [here](https://golangci-lint.run/usage/integrations/).
   Otherwise, you will need to run `make fmt` and `make lintfix` to automatically format and fix any lint violations before you commit any changes, or add a pre-commit Git hook that does it for you automatically.
1. Run `make run` to start the server.

## Updating Protocol Buffer Definitions

After changing the definitions of the Protocol Buffer, run `make proto` to regenerate the respective Go files.
Note that `make build` and `make run` also regenerate the Go files automatically.

## Migrating Database

To apply new migrations to the database, run:
```sh
$ make migratedb
```

If the state of your database is out of sync, you can drop it and recreate it as such:
```sh
$ make dropdb
$ make createdb
$ make migratedb
```

## Architecture

The Hansel backend makes use of the Ports & Adapters architecture.
Using the principle of dependency inversion, it allows for loosely-coupled components that can be easily swapped out by means of ports and adapters.
This not only enforces a strong separation of concerns, it also makes each component much more testable.
It is also overkill for this assignment and wholly unnecessary.

Do read this [article](https://herbertograca.com/2017/11/16/explicit-architecture-01-ddd-hexagonal-onion-clean-cqrs-how-i-put-it-all-together/) to get a better overview of this architecture.
