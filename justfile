_default:
    @just --list

install:
    go mod download
    go get -t ./...

generate:
    go generate ./ent

setup: generate install

test:
    go test ./...
    go vet ./...

quality:
    golangci-lint fmt
    golangci-lint run
    govulncheck ./...

runcli:
    go run ./cmd/cli
