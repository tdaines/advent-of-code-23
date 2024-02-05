default: run

test:
    go test ./...

run:
    go run .

build:
    go build -v main.go
