lint:
    golangci-lint run

build:
    go build -o bin/echoservice main.go

run:
    go run main.go

test:
    go test ./...

clean:
    rm -rf bin/

rebuild:
    just clean
    just build
    just run
