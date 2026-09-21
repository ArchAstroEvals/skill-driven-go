# harbor

Tiny records API in Go.

## Run

Run `go run .` then curl localhost:8080/health.

Create a record:

```sh
curl -X POST localhost:8080/records \
  -H "Authorization: Bearer tok-dev" \
  -d '{"name":"widget"}'
```

## Docker

Build `docker build -t harbor .` then run on port 8080.

## Test

Run `go test ./...` or `make test`.

## Layout

Server and mux, mutex store, validation, serializer, auth, query helpers, middleware.
