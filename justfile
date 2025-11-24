_default:
    @just --list

install:
    go mod download
    go get -t ./...

generate:
    go tool sqlc generate

setup: generate install

test:
    go test ./...
    go vet ./...

quality:
    golangci-lint fmt
    golangci-lint run
    govulncheck ./...
