GOLANGCI_VERSION := "v1.51.2"

install-linter:
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b /usr/local/bin
    golangci-lint --version


lint: install-linter
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
