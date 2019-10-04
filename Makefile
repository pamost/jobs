GOCMD=go
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

all: gofmt test linter

gofmt:
	gofmt -w .

test:
	go test -v -cover ./...

linter:
	golangci-lint run --enable-all
